import logging

from django.core.exceptions import ImproperlyConfigured
from rest_framework.exceptions import PermissionDenied
from rest_framework.permissions import BasePermission, IsAuthenticated, SAFE_METHODS

from ...settings import oauth2_settings
from .authentication import OAuth2Authentication


log = logging.getLogger("oauth2_provider")


class TokenHasScope(BasePermission):
    """
    The request is authenticated as a user and the token used has the right scope
    """

    def has_permission(self, request, view):
        token = request.auth

        if not token:
            return False

        if hasattr(token, "scope"):  # OAuth 2
            required_scopes = self.get_scopes(request, view)
            log.debug("Required scopes to access resource: {0}".format(required_scopes))

            if token.is_valid(required_scopes):
                return True

            # Provide information about required scope?
            include_required_scope = (
                oauth2_settings.ERROR_RESPONSE_WITH_SCOPES and
                required_scopes and
                not token.is_expired() and
                not token.allow_scopes(required_scopes)
            )

            if include_required_scope:
                self.message = {
                    "detail": PermissionDenied.default_detail,
                    "required_scopes": list(required_scopes),
                }

            return False

        assert False, ("TokenHasScope requires the"
                       "`oauth2_provider.rest_framework.OAuth2Authentication` authentication "
                       "class to be used.")

    def get_scopes(self, request, view):
        try:
            return getattr(view, "required_scopes")
        except AttributeError:
            raise ImproperlyConfigured(
                "TokenHasScope requires the view to define the required_scopes attribute"
            )


class TokenHasReadWriteScope(TokenHasScope):
    """
    The request is authenticated as a user and the token used has the right scope
    """

    def get_scopes(self, request, view):
        try:
            required_scopes = super(TokenHasReadWriteScope, self).get_scopes(request, view)
        except ImproperlyConfigured:
            required_scopes = []

        # TODO: code duplication!! see dispatch in ReadWriteScopedResourceMixin
        if request.method.upper() in SAFE_METHODS:
            read_write_scope = oauth2_settings.READ_SCOPE
        else:
            read_write_scope = oauth2_settings.WRITE_SCOPE

        return required_scopes + [read_write_scope]


class TokenHasResourceScope(TokenHasScope):
    """
    The request is authenticated as a user and the token used has the right scope
    """

    def get_scopes(self, request, view):
        try:
            view_scopes = (
                super(TokenHasResourceScope, self).get_scopes(request, view)
            )
        except ImproperlyConfigured:
            view_scopes = []

        if request.method.upper() in SAFE_METHODS:
            scope_type = oauth2_settings.READ_SCOPE
        else:
            scope_type = oauth2_settings.WRITE_SCOPE

        required_scopes = [
            "{}:{}".format(scope, scope_type) for scope in view_scopes
        ]

        return required_scopes


class IsAuthenticatedOrTokenHasScope(BasePermission):
    """
    The user is authenticated using some backend or the token has the right scope
    This only returns True if the user is authenticated, but not using a token
    or using a token, and the token has the correct scope.

    This is usefull when combined with the DjangoModelPermissions to allow people browse
    the browsable api's if they log in using the a non token bassed middleware,
    and let them access the api's using a rest client with a token
    """
    def has_permission(self, request, view):
        is_authenticated = IsAuthenticated().has_permission(request, view)
        oauth2authenticated = False
        if is_authenticated:
            oauth2authenticated = isinstance(request.successful_authenticator, OAuth2Authentication)

        token_has_scope = TokenHasScope()
        return (is_authenticated and not oauth2authenticated) or token_has_scope.has_permission(request, view)
