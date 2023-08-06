"""CAS authentication backend"""

from six.moves.urllib_parse import urlencode, urljoin
from six.moves.urllib.request import urlopen

import requests
from django.conf import settings
from .utils import cas_response_callbacks
from django.contrib.auth import get_user_model

__all__ = ['CASBackend']


def _verify_cas1(ticket, service, cas_url=None):
    """Verifies CAS 1.0 authentication ticket.

    Returns username on success and None on failure.
    """

    params = {'ticket': ticket, 'service': service}
    url = (
        urljoin(settings.CAS_SERVER_URL, 'validate') + '?' + urlencode(params))
    page = urlopen(url)
    try:
        verified = page.readline().strip()
        if verified == 'yes':
            return page.readline().strip()
        else:
            return None
    finally:
        page.close()


def _verify_cas2(ticket, service, cas_url=None, suffix='proxyValidate'):
    return _internal_verify_cas(ticket, service, cas_url, suffix)


def _verify_cas3(ticket, service, cas_url=None, suffix='p3/proxyValidate'):
    return _internal_verify_cas(ticket, service, cas_url, suffix)


def _internal_verify_cas(ticket, service, cas_url, suffix):
    """Verifies CAS 2.0 and 3.0  XML-based authentication ticket.

    Returns username on success and None on failure.
    """

    try:
        from xml.etree import ElementTree
    except ImportError:
        from elementtree import ElementTree

    if settings.CAS_PROXY_CALLBACK:
        params = {
            'ticket': ticket,
            'service': service,
            'pgtUrl': settings.CAS_PROXY_CALLBACK
        }
    else:
        params = {'ticket': ticket, 'service': service}

    url = (urljoin(cas_url, suffix) + '?' + urlencode(params))

    response = requests.get(url, verify=False).text
    tree = ElementTree.fromstring(response)

    if tree[0].tag.endswith('authenticationSuccess'):
        if settings.CAS_RESPONSE_CALLBACKS:
            cas_response_callbacks(tree)
        return tree[0][0].text
    else:
        return None


def verify_proxy_ticket(ticket, service):
    """Verifies CAS 2.0+ XML-based proxy ticket.

    Returns username on success and None on failure.
    """

    try:
        from xml.etree import ElementTree
    except ImportError:
        from elementtree import ElementTree

    params = {'ticket': ticket, 'service': service}

    url = (urljoin(settings.CAS_SERVER_URL, 'proxyValidate') + '?' +
           urlencode(params))

    page = urlopen(url)

    try:
        response = page.read()
        tree = ElementTree.fromstring(response)
        if tree[0].tag.endswith('authenticationSuccess'):
            username = tree[0][0].text
            proxies = []
            if len(tree[0]) > 1:
                for element in tree[0][1]:
                    proxies.append(element.text)
            return {"username": username, "proxies": proxies}
        else:
            return None
    finally:
        page.close()


_PROTOCOLS = {'1': _verify_cas1, '2': _verify_cas2, '3': _verify_cas3}

if settings.CAS_VERSION not in _PROTOCOLS:
    raise ValueError('Unsupported CAS_VERSION %r' % settings.CAS_VERSION)

_verify = _PROTOCOLS[settings.CAS_VERSION]


class CASBackend(object):
    """CAS authentication backend"""

    supports_object_permissions = False
    supports_inactive_user = False

    def authenticate(self, ticket, service, cas_url=None):
        """Verifies CAS ticket and gets or create User object
            NB: Use of PT to identify proxy
        """

        email = _verify(ticket, service, cas_url)
        user_model = get_user_model()
        if not email:
            return None
        try:
            if settings.CAS_USERNAME_FORMAT:
                email = settings.CAS_USERNAME_FORMAT(email)
            user = user_model.objects.get(email=email)
        except user_model.DoesNotExist:
            if settings.CAS_USER_CREATION:
                # user will have an "unusable" password
                user = user_model.objects.create_user(email, email, '')
                user.save()
            else:
                return None
        return user

    def get_user(self, user_id):
        """Retrieve the user's entry in the User model if it exists"""

        user_model = get_user_model()
        try:
            return user_model.objects.get(pk=user_id)
        except user_model.DoesNotExist:
            return None
