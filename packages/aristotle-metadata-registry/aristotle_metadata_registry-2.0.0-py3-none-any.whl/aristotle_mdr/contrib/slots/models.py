"""
Aristotle MDR 11179 Slots models
================================

These are based on the Slots definition in ISO/IEC 11179 Part 3 - 7.2.2.4
"""

from django.db import models

from model_utils.models import TimeStampedModel

from aristotle_mdr import models as MDR
from aristotle_mdr.fields import ConceptForeignKey


class Slot(TimeStampedModel):
    # on save confirm the concept and model are correct, otherwise reject
    # on save confirm the cardinality
    name = models.CharField(max_length=256)  # Or some other sane length
    type = models.CharField(max_length=256, blank=True)  # Or some other sane length
    concept = ConceptForeignKey(MDR._concept, related_name='slots')
    value = models.TextField()

    def __str__(self):
        return u"{0} - {1}".format(self.name, self.value)


def concepts_with_similar_slots(user, name=None, _type=None, value=None, slot=None):
    assert(slot is not None or _type is not None or name is not None)

    if slot is not None:
        name = slot.name
        _type = slot.type
        value = slot.value

    slots = MDR._concept.objects.visible(user)

    if name:
        slots = slots.filter(slots__name=name)

    if _type:
        slots = slots.filter(slots__type=_type)

    if value is not None:
        slots = slots.filter(slots__value=value)

    if slot is not None:
        slots = slots.exclude(id=slot.concept.id)

    return slots.distinct()
