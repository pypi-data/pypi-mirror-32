# -*- coding: utf-8 -*-
# Generated by Django 1.11.11 on 2018-03-19 14:53
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('djasana', '0013_adds_on_delete'),
    ]

    operations = [
        migrations.AlterField(
            model_name='user',
            name='email',
            field=models.EmailField(blank=True, max_length=254, null=True, verbose_name='email address'),
        ),
    ]
