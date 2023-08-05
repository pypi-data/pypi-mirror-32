"""Exceptions raised by Fedora Messaging."""


class BaseException(Exception):
    """The base class for all exceptions raised by fedora_messaging."""


class PublishException(BaseException):
    """Base class for exceptions related to publishing."""


class PublishReturned(PublishException):
    """Raised when the broker rejects and returns the message to the publisher."""


class ConnectionException(BaseException):
    """Raised if a general connection error occurred."""

    def __init__(self, reason=None, **kwargs):
        super(ConnectionException, self).__init__(**kwargs)
        self.reason = reason


class ConsumeException(BaseException):
    """Base class for exceptions related to consuming."""


class Nack(ConsumeException):
    """
    Consumer callbacks should raise this to indicate they wish the message they
    are currently processing to be re-queued.
    """


class Drop(ConsumeException):
    """
    Consumer callbacks should raise this to indicate they wish the message they
    are currently processing to be dropped.
    """


class HaltConsumer(ConsumeException):
    """
    Consumer callbacks should raise this exception if they wish the consumer to
    be shut down.

    Args:
        exit_code (int): The exit code to use when halting.
    """

    def __init__(self, exit_code=0, reason=None, **kwargs):
        super(HaltConsumer, self).__init__(**kwargs)
        self.exit_code = exit_code
        self.reason = reason


class ValidationError(BaseException):
    """This error is raised when a message fails validation."""
