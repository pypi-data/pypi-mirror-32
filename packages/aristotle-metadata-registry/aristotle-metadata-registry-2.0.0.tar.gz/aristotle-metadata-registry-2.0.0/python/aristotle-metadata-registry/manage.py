#!/usr/bin/env python
import os
import sys

if __name__ == "__main__":
    if 'test' in sys.argv:
        os.environ['DJANGO_SETTINGS_MODULE'] = "aristotle_mdr.tests.settings.settings"
    elif 'schemamigration' in sys.argv:
        os.environ.setdefault("DJANGO_SETTINGS_MODULE", "aristotle_mdr.tests.settings.settings")
    else:
        os.environ.setdefault("DJANGO_SETTINGS_MODULE", "aristotle_mdr.required_settings")

    from django.core.management import execute_from_command_line

    execute_from_command_line(sys.argv)
