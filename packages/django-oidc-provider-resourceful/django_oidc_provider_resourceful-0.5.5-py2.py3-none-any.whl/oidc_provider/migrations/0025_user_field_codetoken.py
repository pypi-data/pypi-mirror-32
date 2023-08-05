# Generated by Django 2.0.3 on 2018-04-13 19:34

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('oidc_provider', '0024_auto_20180327_1959'),
    ]

    operations = [
        migrations.AddField(
            model_name='client',
            name='_scope',
            field=models.TextField(blank=True, default='', help_text='Specifies the authorized scope values for the client app.', verbose_name='Scopes'),
        ),
        migrations.AlterField(
            model_name='token',
            name='user',
            field=models.ForeignKey(null=True, on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL, verbose_name='User'),
        ),
    ]
