
import ujson
import traceback

from .contexts import context
from .configuration import settings


class HttpStatus(Exception):
    status_code = None
    status_text = None
    info = None

    def __init__(self, message=None, reason=None, info=None, status_code=None, status_text=None):
        if info is not None:
            self.info = info

        if status_code is not None:
            self.status_code = status_code

        if status_text is not None:
            self.status_text = status_text

        if reason:
            context.response_headers.add_header('X-Reason', reason)

        super().__init__(message or self.status_text)

    @property
    def status(self):
        return '%s %s' % (self.status_code, self.status_text)

    def to_dict(self):
        return dict(
            description=self.info
        )

    def render(self):
        if context.response_content_type == 'application/json':
            return ujson.encode(self.to_dict())
        else:
            context.response_encoding = 'utf-8'
            context.response_content_type = 'text/plain'
            return "%s\n%s" % (self.status_text, self.info)


class HttpBadRequest(HttpStatus):
    status_code, status_text, info = 400, 'Bad Request', 'Bad request syntax or unsupported method'


class HttpUnauthorized(HttpStatus):
    status_code, status_text, info = 401, 'Unauthorized', 'No permission -- see authorization schemes'


class HttpForbidden(HttpStatus):
    status_code, status_text, info = 403, 'Forbidden', 'Request forbidden -- authorization will not help'


class HttpNotFound(HttpStatus):
    status_code, status_text, info = 404, 'Not Found', 'Nothing matches the given URI'


class HttpMethodNotAllowed(HttpStatus):
    status_code, status_text, info = 405, 'Method Not Allowed', 'Specified method is invalid for this resource'


class HttpConflict(HttpStatus):
    status_code, status_text, info = 409, 'Conflict', 'Request conflict'


class HttpGone(HttpStatus):
    status_code, status_text, info = 410, 'Gone', 'URI no longer exists and has been permanently removed'


class HttpPreconditionFailed(HttpStatus):
    status_code, status_text, info = 412, 'Precondition Failed', 'Request cannot be fulfilled'


class HttpRedirect(HttpStatus):
    """
    This is an abstract class for all redirects.
    """

    def __init__(self, location, *args, **kw):
        context.response_headers.add_header('Location', location)
        super().__init__(*args, **kw)


class HttpMovedPermanently(HttpRedirect):
    status_code, status_text, info = 301, 'Moved Permanently', 'Object moved permanently'


class HttpFound(HttpRedirect):
    status_code, status_text, info = 302, 'Found', 'Object moved temporarily'


class HttpNotModified(HttpStatus):
    status_code, status_text, info = 304, 'Not Modified', ''  # 304 is only header


class HttpInternalServerError(HttpStatus):
    status_code, status_text = 500, 'Internal Server Error'

    @property
    def info(self):
        if settings.debug:
            return traceback.format_exc()
        return 'Server got itself in trouble'


class HttpBadGatewayError(HttpStatus):
    status_code, status_text = 502, 'Bad Gateway'
