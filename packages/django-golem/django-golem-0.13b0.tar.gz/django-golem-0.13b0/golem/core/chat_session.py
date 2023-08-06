from golem.core.interfaces.all import create_from_prefix


class ChatSession:

    def __init__(self, interface, unique_id: str, meta=None, profile=None, is_logged=True, is_admin=False, is_test=False):
        """
        :param interface: Chat interface object
        :param unique_id: A string that uniquely identifies the user within the interface
        :param profile: Profile object
        :param meta: Any interface-specific data. type: dict
        """
        if interface is None:
            raise ValueError("Interface must be set")
        elif not isinstance(unique_id, str):
            raise ValueError("UID must be a string")
        elif not isinstance(profile, Profile) and profile is not None:
            raise ValueError("Profile must be instance of class Profile")
        self.interface = interface
        self.chat_id = self.interface.prefix + "_" + unique_id
        self.profile = profile or Profile()
        self.meta = meta or {}
        self.is_logged = is_logged
        self.is_admin = is_admin
        self.is_test = is_test

    def to_json(self):
        return {
            "interface": self.interface.prefix,
            "unique_id": self.chat_id.split('_', maxsplit=1)[1],
            "profile": self.profile.to_json(),
            "meta": self.meta,
            "is_logged": self.is_logged,
            "is_admin": self.is_admin
        }

    @staticmethod
    def from_json(data: dict):
        interface = create_from_prefix(data['interface'])
        return ChatSession(
            interface=interface,
            unique_id=data['unique_id'],
            profile=Profile.from_json(data['profile']),
            meta=data['meta'],
            is_logged=data['is_logged'],
            is_admin=data['is_admin']
        )


class Profile:
    def __init__(self, profile_id=None, first_name=None, last_name=None):
        self.profile_id = profile_id
        self.first_name = first_name
        self.last_name = last_name

    def to_json(self):
        return {
            "profile_id": self.profile_id,
            "first_name": self.first_name,
            "last_name": self.last_name
        }

    @staticmethod
    def from_json(data: dict):
        return Profile(
            profile_id=data['profile_id'],
            first_name=data['first_name'],
            last_name=data['last_name']
        )
