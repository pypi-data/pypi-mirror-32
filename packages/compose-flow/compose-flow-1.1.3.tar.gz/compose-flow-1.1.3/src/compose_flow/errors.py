class CommandError(Exception):
    """
    Raised when a problem with the command run is encountered
    """

class ErrorMessage(Exception):
    """
    Subclass to print out error message instead of entire stack trace
    """


class EnvError(ErrorMessage):
    """
    Error for when environment variables are not found
    """

class NoSuchConfig(Exception):
    """
    Raised when a requested config is not in the docker swarm
    """


class NoSuchProfile(Exception):
    """
    Raised when a requested profile is not listed in dc.yml
    """

class NotConnected(Exception):
    """
    Raised when not connected to a remote host
    """


class ProfileError(ErrorMessage):
    """
    Raised when there is a problem with a Profile
    """


class TagVersionError(Exception):
    """
    Raised when there is a problem running tag-version
    """
