import datetime
from django.apps import apps
from django.contrib.auth.decorators import login_required
from braces.views import LoginRequiredMixin
from django.contrib.auth.views import LoginView
from django.core.cache import cache
from django.core.exceptions import PermissionDenied
from django.urls import reverse
from django.db.models import Q
from django.http import HttpResponseRedirect
from django.shortcuts import render, redirect
from django.utils.decorators import method_decorator
from django.views.generic import DetailView, ListView, UpdateView

from aristotle_mdr import forms as MDRForms
from aristotle_mdr import models as MDR
from aristotle_mdr.views.utils import paginated_list, paginated_workgroup_list
from aristotle_mdr.utils import fetch_metadata_apps
from aristotle_mdr.utils import get_aristotle_url

import json


class FriendlyLoginView(LoginView):

    template_name='aristotle_mdr/friendly_login.html'
    redirect_field_name = 'next'
    redirect_authenticated_user = True

    def get_context_data(self, **kwargs):

        context = super().get_context_data(**kwargs)

        if self.request.GET.get('welcome', '') == 'true':
            context.update({'welcome': True})

        return context


@login_required
def home(request):
    from reversion.models import Revision

    recent = Revision.objects.filter(user=request.user).order_by('-date_created')[0:10]
    recentdata = []
    for rev in recent:
        revdata = {'revision': rev, 'versions': []}
        seen_ver_ids = []

        for ver in rev.version_set.all():

            seen = ver.object_id in seen_ver_ids
            add_version = None
            url = None

            if ver.format == 'json':
                object_data = json.loads(ver.serialized_data)

                try:
                    model = object_data[0]['model']
                except KeyError:
                    model = None

                if model:
                    add_version = (model != 'aristotle_mdr.status' and not seen)

                try:
                    name = object_data[0]['fields']['name']
                except KeyError:
                    name = None

                url = get_aristotle_url(object_data[0]['model'], object_data[0]['pk'], name)

            if add_version is None:
                # Fallback for if add_version could not be set, results in db query
                add_version = (not isinstance(ver.object, MDR.Status) and not seen)

            if add_version:

                if url:
                    revdata['versions'].append({'id': ver.object_id, 'text': str(ver), 'url': url})
                else:
                    # Fallback, results in db query
                    obj = ver.object
                    if hasattr(obj, 'get_absolute_url'):
                        revdata['versions'].append({'id': ver.object_id, 'text': str(ver), 'url': obj.get_absolute_url})

            seen_ver_ids.append(ver.object_id)

        recentdata.append(revdata)

    page = render(request, "aristotle_mdr/user/userHome.html", {"item": request.user, 'recentdata': recentdata})
    return page


@login_required
def roles(request):
    page = render(request, "aristotle_mdr/user/userRoles.html", {"item": request.user})
    return page


@login_required
def recent(request):
    from reversion.models import Revision
    from aristotle_mdr.views.utils import paginated_reversion_list
    items = Revision.objects.filter(user=request.user).order_by('-date_created')
    context = {}
    return paginated_reversion_list(request, items, "aristotle_mdr/user/recent.html", context)


@login_required
def inbox(request, folder=None):
    if folder is None:
        # By default show only unread
        folder='unread'
    folder=folder.lower()
    if folder == 'unread':
        notices = request.user.notifications.unread().all()
    elif folder == "all":
        notices = request.user.notifications.all()
    page = render(
        request,
        "aristotle_mdr/user/userInbox.html",
        {"item": request.user, "notifications": notices, 'folder': folder}
    )
    return page


@login_required
def admin_tools(request):
    if request.user.is_anonymous():
        return redirect(reverse('friendly_login') + '?next=%s' % request.path)
    elif not request.user.has_perm("aristotle_mdr.access_aristotle_dashboard"):
        raise PermissionDenied

    aristotle_apps = fetch_metadata_apps()

    from django.contrib.contenttypes.models import ContentType
    models = ContentType.objects.filter(app_label__in=aristotle_apps).all()
    model_stats = {}

    for m in models:
        if m.model_class() and issubclass(m.model_class(), MDR._concept) and not m.model.startswith("_"):
            # Only output subclasses of 11179 concept
            app_models = model_stats.get(m.app_label, {'app': None, 'models': []})
            if app_models['app'] is None:
                app_models['app'] = getattr(apps.get_app_config(m.app_label), 'verbose_name')
            app_models['models'].append(
                (
                    m.model_class(),
                    get_cached_object_count(m),
                    reverse("browse_concepts", args=[m.app_label, m.model])
                )
            )
            model_stats[m.app_label] = app_models

    page = render(
        request,
        "aristotle_mdr/user/userAdminTools.html",
        {"item": request.user, "models": model_stats}
    )
    return page


@login_required
def admin_stats(request):
    if request.user.is_anonymous():
        return redirect(reverse('friendly_login') + '?next=%s' % request.path)
    elif not request.user.has_perm("aristotle_mdr.access_aristotle_dashboard"):
        raise PermissionDenied

    aristotle_apps = fetch_metadata_apps()

    from django.contrib.contenttypes.models import ContentType
    models = ContentType.objects.filter(app_label__in=aristotle_apps).all()
    model_stats = {}

    # Get datetime objects for '7 days ago' and '30 days ago'
    t7 = datetime.date.today() - datetime.timedelta(days=7)
    t30 = datetime.date.today() - datetime.timedelta(days=30)
    mod_counts = []  # used to get the maximum count

    use_cache = True  # We still cache but its much, much shorter
    for m in models:
        if m.model_class() and issubclass(m.model_class(), MDR._concept) and not m.model.startswith("_"):
            # Only output subclasses of 11179 concept
            app_models = model_stats.get(m.app_label, {'app': None, 'models': []})
            if app_models['app'] is None:
                app_models['app'] = getattr(apps.get_app_config(m.app_label), 'verbose_name')
            if use_cache:
                total = get_cached_query_count(
                    qs=m.model_class().objects,
                    key=model_to_cache_key(m) + "__all_time",
                    ttl=60
                )
                t7_val = get_cached_query_count(
                    qs=m.model_class().objects.filter(created__gte=t7),
                    key=model_to_cache_key(m) + "__t7",
                    ttl=60
                )
                t30_val = get_cached_query_count(
                    qs=m.model_class().objects.filter(created__gte=t30),
                    key=model_to_cache_key(m) + "__t30",
                    ttl=60
                )
            else:
                total = m.model_class().objects.count()
                t7_val = m.model_class().objects.filter(created__gte=t7).count()
                t30_val = m.model_class().objects.filter(created__gte=t30).count()

            mod_counts.append(total)
            app_models['models'].append(
                (
                    m.model_class(),
                    {
                        'all_time': total,
                        't7': t7_val,
                        't30': t30_val
                    },
                    reverse("browse_concepts", args=[m.app_label, m.model])
                )
            )
            model_stats[m.app_label] = app_models

    page = render(
        request, "aristotle_mdr/user/userAdminStats.html",
        {"item": request.user, "model_stats": model_stats, 'model_max': max(mod_counts)}
    )
    return page


def get_cached_query_count(qs, key, ttl):
    count = cache.get(key, None)
    if not count:
        count = qs.count()
        cache.set(key, count, ttl)
    return count


def model_to_cache_key(model_type):
    return 'aristotle_adminpage_object_count_%s_%s' % (model_type.app_label, model_type.model)


def get_cached_object_count(model_type):
    CACHE_KEY = model_to_cache_key(model_type)
    query = model_type.model_class().objects
    return get_cached_query_count(query, CACHE_KEY, 60 * 60 * 12)  # Cache for 12 hours


class EditView(LoginRequiredMixin, UpdateView):

    template_name = "aristotle_mdr/user/userEdit.html"
    form_class = MDRForms.EditUserForm

    def get_object(self, querySet=None):
        return self.request.user

    def get_success_url(self):
        return reverse('aristotle:userHome')


@login_required
def favourites(request):
    items = request.user.profile.favourites.select_subclasses()
    context = {
        'help': request.GET.get("help", False),
        'favourite': request.GET.get("favourite", False),
        # "select_all_list_queryset_filter": 'favourited_by__user=user'  # no information leakage here.
    }
    return paginated_list(request, items, "aristotle_mdr/user/userFavourites.html", context)


@login_required
def registrar_tools(request):
    if not request.user.profile.is_registrar:
        raise PermissionDenied
    page = render(request, "aristotle_mdr/user/userRegistrarTools.html")
    return page


@login_required
def review_list(request):
    if not request.user.profile.is_registrar:
        raise PermissionDenied
    authorities = [i[0] for i in request.user.profile.registrarAuthorities.all().values_list('id')]

    # Registars can see items they have been asked to review
    q = Q(Q(registration_authority__id__in=authorities) & ~Q(status=MDR.REVIEW_STATES.cancelled))

    reviews = MDR.ReviewRequest.objects.visible(request.user).filter(q)
    return paginated_list(request, reviews, "aristotle_mdr/user/userReviewList.html", {'reviews': reviews})


@login_required
def my_review_list(request):
    # Users can see any items they have been asked to review
    q = Q(requester=request.user)
    reviews = MDR.ReviewRequest.objects.visible(request.user).filter(q)
    return paginated_list(request, reviews, "aristotle_mdr/user/my_review_list.html", {'reviews': reviews})


@login_required
def django_admin_wrapper(request, page_url):
    return render(request, "aristotle_mdr/user/admin.html", {'page_url': page_url})


class ReviewDetailsView(DetailView):
    pk_url_kwarg = 'review_id'
    template_name = "aristotle_mdr/user/request_review_details.html"
    context_object_name = "review"

    @method_decorator(login_required)
    def dispatch(self, *args, **kwargs):
        return super().dispatch(*args, **kwargs)

    def get_context_data(self, *args, **kwargs):
        # Call the base implementation first to get a context
        context = super().get_context_data(*args, **kwargs)
        context['next'] = self.request.GET.get('next', reverse('aristotle:userReadyForReview'))
        return context

    def get_queryset(self):
        return MDR.ReviewRequest.objects.visible(self.request.user)


class CreatedItemsListView(ListView):
    paginate_by = 25
    template_name = "aristotle_mdr/user/sandbox.html"

    @method_decorator(login_required)
    def dispatch(self, *args, **kwargs):
        return super().dispatch(*args, **kwargs)

    def get_queryset(self, *args, **kwargs):
        return MDR._concept.objects.filter(
            Q(
                submitter=self.request.user,
                statuses__isnull=True
            ) & Q(
                Q(review_requests__isnull=True) | Q(review_requests__status=MDR.REVIEW_STATES.cancelled)
            )
        )

    def get_context_data(self, *args, **kwargs):
        # Call the base implementation first to get a context
        context = super().get_context_data(*args, **kwargs)
        context['sort'] = self.request.GET.get('sort', 'name_asc')
        return context

    def get_ordering(self):
        from aristotle_mdr.views.utils import paginate_sort_opts
        self.order = self.request.GET.get('sort', 'name_asc')
        return paginate_sort_opts.get(self.order)


@login_required
def workgroups(request):
    text_filter = request.GET.get('filter', "")
    workgroups = request.user.profile.myWorkgroups
    if text_filter:
        workgroups = workgroups.filter(Q(name__icontains=text_filter) | Q(definition__icontains=text_filter))
    context = {'filter': text_filter}
    return paginated_workgroup_list(request, workgroups, "aristotle_mdr/user/userWorkgroups.html", context)


@login_required
def workgroup_archives(request):
    text_filter = request.GET.get('filter', None)
    workgroups = request.user.profile.workgroups.filter(archived=True)
    if text_filter:
        workgroups = workgroups.filter(Q(name__icontains=text_filter) | Q(definition__icontains=text_filter))
    context = {'filter': text_filter}
    return paginated_workgroup_list(request, workgroups, "aristotle_mdr/user/userWorkgroupArchives.html", context)
