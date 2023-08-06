from golem.core.chat_session import ChatSession, Profile

class TestInterface():
    name = 'test'
    prefix = 'test'
    messages = []
    states = []

    @staticmethod
    def clear():
        TestInterface.messages = []
        TestInterface.states = []

    @staticmethod
    def fill_session_profile(session: ChatSession):
        if not session:
            raise ValueError("Session is None")
        session.profile = Profile(first_name='Test', last_name='')
        return session

    @staticmethod
    def post_message(session, response):
        TestInterface.messages.append(response)

    @staticmethod
    def send_settings(settings):
        pass

    @staticmethod
    def processing_start(session):
        pass

    @staticmethod
    def processing_end(session):
        pass

    @staticmethod
    def state_change(state):
        if not TestInterface.states or TestInterface.states[-1] != state:
            TestInterface.states.append(state)

    @staticmethod
    def parse_message(user_message, num_tries=1):
        return user_message
