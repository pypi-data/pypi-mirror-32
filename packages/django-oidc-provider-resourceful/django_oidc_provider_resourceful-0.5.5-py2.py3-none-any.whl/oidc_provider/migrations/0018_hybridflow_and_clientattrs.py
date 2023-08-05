# -*- coding: utf-8 -*-
# Generated by Django 1.9.7 on 2016-09-12 14:08
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('oidc_provider', '0017_auto_20160811_1954'),
    ]

    operations = [
        migrations.AddField(
            model_name='client',
            name='contact_email',
            field=models.CharField(blank=True, default='', max_length=255, verbose_name='Contact Email'),
        ),
        migrations.AddField(
            model_name='client',
            name='logo',
            field=models.FileField(
                blank=True, default='', upload_to='oidc_provider/clients', verbose_name='Logo Image'),
        ),
        migrations.AddField(
            model_name='client',
            name='terms_url',
            field=models.CharField(
                blank=True,
                default='',
                help_text='External reference to the privacy policy of the client.',
                max_length=255,
                verbose_name='Terms URL'),
        ),
        migrations.AddField(
            model_name='client',
            name='website_url',
            field=models.CharField(blank=True, default='', max_length=255, verbose_name='Website URL'),
        ),
        migrations.AlterField(
            model_name='client',
            name='jwt_alg',
            field=models.CharField(
                choices=[('HS256', 'HS256'), ('RS256', 'RS256')],
                default='RS256',
                help_text='Algorithm used to encode ID Tokens.',
                max_length=10,
                verbose_name='JWT Algorithm'),
        ),
        migrations.AlterField(
            model_name='client',
            name='response_type',
            field=models.CharField(
                choices=[
                    ('code', 'code (Authorization Code Flow)'), ('id_token', 'id_token (Implicit Flow)'),
                    ('id_token token', 'id_token token (Implicit Flow)'), ('code token', 'code token (Hybrid Flow)'),
                    ('code id_token', 'code id_token (Hybrid Flow)'),
                    ('code id_token token', 'code id_token token (Hybrid Flow)')],
                max_length=30,
                verbose_name='Response Type'),
        ),
    ]
