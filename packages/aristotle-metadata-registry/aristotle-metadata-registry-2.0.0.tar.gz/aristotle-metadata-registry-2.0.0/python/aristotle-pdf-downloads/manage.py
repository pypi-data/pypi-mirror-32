#!/usr/bin/env python
import os
import sys

# This is only needed for herokus collectstatic step to run :/

if __name__ == "__main__":
    from django.core.management import execute_from_command_line
    execute_from_command_line(sys.argv)
