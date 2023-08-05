import time
from mock import patch

from django.utils.encoding import force_text

from oidc_provider.lib.utils.token import create_id_token

try:
    from urllib.parse import urlencode
except ImportError:
    from urllib import urlencode

from django.core.management import call_command
from django.test import TestCase, RequestFactory, override_settings
from django.core.urlresolvers import reverse
from django.utils import timezone

from oidc_provider.tests.app.utils import (
    create_fake_user,
    create_fake_client,
    create_fake_resource,
    create_fake_token,
    FAKE_RANDOM_STRING)
from oidc_provider.views import TokenIntrospectionView


class IntrospectionTestCase(TestCase):

    def setUp(self):
        call_command('creatersakey')
        self.factory = RequestFactory()
        self.user = create_fake_user()
        self.client = create_fake_client(response_type='id_token token')
        self.resource = create_fake_resource(allowed_clients=[self.client])
        self.scopes = ['openid', 'profile']
        self.token = create_fake_token(self.user, self.scopes, self.client)
        self.now = time.time()
        with patch('oidc_provider.lib.utils.token.time.time') as time_func:
            time_func.return_value = self.now
            self.token.id_token = create_id_token(self.user, self.client.client_id)
        self.token.save()

    def test_no_client_params_returns_inactive(self):
        response = self._make_request(client_id='')
        self._assert_inactive(response)

    def test_no_client_secret_returns_inactive(self):
        response = self._make_request(client_secret='')
        self._assert_inactive(response)

    def test_invalid_client_returns_inactive(self):
        response = self._make_request(client_id='invalid')
        self._assert_inactive(response)

    def test_token_not_found_returns_inactive(self):
        response = self._make_request(access_token='invalid')
        self._assert_inactive(response)

    def test_no_allowed_clients_returns_inactive(self):
        self.resource.allowed_clients.clear()
        self.resource.save()
        response = self._make_request()
        self._assert_inactive(response)

    def test_resource_inactive_returns_inactive(self):
        self.resource.active = False
        self.resource.save()
        response = self._make_request()
        self._assert_inactive(response)

    def test_token_expired_returns_inactive(self):
        self.token.expires_at = timezone.now() - timezone.timedelta(seconds=60)
        self.token.save()
        response = self._make_request()
        self._assert_inactive(response)

    def test_valid_request_returns_default_properties(self):
        response = self._make_request()
        self.assertEqual(response.status_code, 200)
        self.assertJSONEqual(force_text(response.content), {
            'active': True,
            'aud': self.resource.resource_id,
            'client_id': self.client.client_id,
            'sub': str(self.user.pk),
            'iat': int(self.now),
            'exp': int(self.now + 600),
            'iss': 'http://localhost:8000/openid',
        })

    @override_settings(
        OIDC_INTROSPECTION_PROCESSING_HOOK='oidc_provider.tests.app.utils.fake_introspection_processing_hook')
    def test_custom_introspection_hook_called_on_valid_request(self):
        response = self._make_request()
        self.assertEqual(response.status_code, 200)
        self.assertJSONEqual(force_text(response.content), {
            'active': True,
            'aud': self.resource.resource_id,
            'client_id': self.client.client_id,
            'sub': str(self.user.pk),
            'iat': int(self.now),
            'exp': int(self.now + 600),
            'iss': 'http://localhost:8000/openid',
            'test_introspection_processing_hook': FAKE_RANDOM_STRING
        })

    def _assert_inactive(self, response):
        self.assertEqual(response.status_code, 200)
        self.assertJSONEqual(force_text(response.content), {'active': False})

    def _make_request(self, **kwargs):
        url = reverse('oidc_provider:token-introspection')
        data = {
            'client_id': kwargs.get('client_id', self.resource.resource_id),
            'client_secret': kwargs.get('client_secret', self.resource.resource_secret),
            'token': kwargs.get('access_token', self.token.access_token),
        }

        request = self.factory.post(url, data=urlencode(data), content_type='application/x-www-form-urlencoded')

        return TokenIntrospectionView.as_view()(request)
