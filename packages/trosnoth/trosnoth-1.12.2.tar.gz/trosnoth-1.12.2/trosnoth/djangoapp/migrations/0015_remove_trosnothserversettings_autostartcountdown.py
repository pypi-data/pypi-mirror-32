# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('trosnoth', '0014_auto_20171215_1505'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='trosnothserversettings',
            name='autoStartCountDown',
        ),
    ]
