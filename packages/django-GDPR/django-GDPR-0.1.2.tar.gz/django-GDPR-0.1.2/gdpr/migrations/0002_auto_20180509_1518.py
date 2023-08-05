# -*- coding: utf-8 -*-
# Generated by Django 1.11 on 2018-05-09 13:18
from __future__ import unicode_literals

import datetime
from django.db import migrations, models
from django.utils.timezone import utc


class Migration(migrations.Migration):

    dependencies = [
        ('gdpr', '0001_initial'),
    ]

    operations = [
        migrations.RenameField(
            model_name='legalreasonrelatedobject',
            old_name='content_type',
            new_name='object_content_type',
        ),
        migrations.AddField(
            model_name='legalreason',
            name='issued_at',
            field=models.DateTimeField(default=datetime.datetime(2018, 5, 9, 13, 18, 7, 317147, tzinfo=utc), verbose_name='issued at'),
            preserve_default=False,
        ),
        migrations.AlterField(
            model_name='legalreason',
            name='purpose_slug',
            field=models.CharField(choices=[('precheck', 'Scoring'), ('application-created', 'Scoring')], max_length=100, verbose_name='purpose'),
        ),
    ]
