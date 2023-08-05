from django.http import Http404
from django.conf import settings
from django.contrib.contenttypes.models import ContentType
from django.core.exceptions import PermissionDenied

from rest_framework import serializers, pagination, status, mixins
from rest_framework.views  import APIView
from rest_framework.response import Response
from rest_framework.reverse import reverse
from rest_framework.decorators import detail_route

from django.forms import model_to_dict
from aristotle_mdr import models, perms
from aristotle_mdr.forms.search import PermissionSearchQuerySet

from rest_framework import viewsets

from .utils import (
    DescriptionStubSerializerMixin,
    MultiSerializerViewSetMixin,
    api_excluded_fields,
    get_api_fields,
    aristotle_apps
)

class ConceptTypeSerializer(serializers.ModelSerializer):
    documentation = serializers.SerializerMethodField()
    fields = serializers.SerializerMethodField('get_extra_fields')
    class Meta:
        model = ContentType
        fields = ('name','app_label','model','documentation','fields')
    def get_documentation(self,instance):
        return instance.model_class().__doc__.strip()
    def get_extra_fields(self,instance):
        field_names = instance.model_class()._meta.get_fields()
        for field in get_api_fields(instance):
            if hasattr(field, 'help_text'):
                yield {'name':field.name, "help_text":field.help_text}
            else:
                yield {'name':field.name}

        for field in field_names:
            if field.name not in api_excluded_fields:
                f = field
                if f.auto_created == False: #, because the new get_field() API will find "reverse" relations), and:
                    if True: #f.is_relation and f.related_model is None: #, because the new get_field() API will find GenericForeignKey relations;
                        if hasattr(field, 'name'):
                            # yield field.name
                            if hasattr(field, 'help_text'):
                                yield {'name':field.name, "help_text":field.help_text}
                            else:
                                yield {'name':field.name}

        #return [field.field_name for field in field_names if field not in api_excluded_fields]
        #return [{'name':field,"help_text":""} for field in field_names if field not in api_excluded_fields]

class ConceptTypeViewSet(viewsets.ReadOnlyModelViewSet):
    """
    This viewset automatically provides `list` and `detail` actions.
    """
    queryset = ContentType.objects.filter(app_label__in=aristotle_apps).all()
    serializer_class = ConceptTypeSerializer

    def get_queryset(self):
        """
        Optionally restricts the returned purchases to a given user,
        by filtering against a `username` query parameter in the URL.
        """
        outputs = []
        for m in ContentType.objects.filter(app_label__in=aristotle_apps).all():
            if hasattr(m, 'model_class') and m.model_class():
                print(m.model_class().__class__, models._concept)
                if issubclass(m.model_class(), models._concept) and not m.model.startswith("_"):
                    outputs.append(m)
        return outputs
