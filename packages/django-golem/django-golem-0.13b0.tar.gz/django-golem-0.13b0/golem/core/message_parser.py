import emoji
import re
from django.conf import settings

ENTITY_EXTRACTORS = settings.GOLEM_CONFIG.get("ENTITY_EXTRACTORS", [])


def add_default_extractors():
    # compatibility for old chatbots
    if "WIT_TOKEN" in settings.GOLEM_CONFIG:
        print("Adding wit extractor from wit token")
        from golem.core.parsing import wit_extractor
        ENTITY_EXTRACTORS.append(wit_extractor.WitExtractor())


add_default_extractors()


def parse_text_message(text, num_tries=1):
    if len(ENTITY_EXTRACTORS) <= 0:
        print('No entity extractors configured!')
        return {'type': 'message', 'entities': {'_message_text': {'value': text}}}

    entities = {}

    for extractor in ENTITY_EXTRACTORS:
        append = extractor.extract_entities(text)
        for entity, values in append.items():
            entities.setdefault(entity, []).extend(values)

    print('Extracted entities:', entities)
    append = parse_additional_entities(text)

    for (entity, value) in re.findall(re.compile(r'/([^/]+)/([^/]+)/'), text):
        if not entity in append:
            append[entity] = []
        append[entity].append({'value': value})

    # add all new entity values
    for entity, values in append.items():
        if not entity in entities:
            entities[entity] = []
        entities[entity] += values

    entities['_message_text'] = [{'value': text}]

    parsed = {'entities': entities, 'type': 'message'}

    return parsed


def parse_additional_entities(text):
    # TODO custom nlp might receive them preprocessed (or not)
    entities = {}
    chars = {':)': ':slightly_smiling_face:', '(y)': ':thumbs_up_sign:', ':(': ':disappointed_face:',
             ':*': ':kissing_face:', ':O': ':face_with_open_mouth:', ':D': ':grinning_face:',
             '<3': ':heavy_black_heart:️', ':P': ':face_with_stuck-out_tongue:'}
    demojized = emoji.demojize(text)
    char_emojis = re.compile(
        r'(' + '|'.join(chars.keys()).replace('(', '\(').replace(')', '\)').replace('*', '\*') + r')')
    demojized = char_emojis.sub(lambda x: chars[x.group()], demojized)
    if demojized != text:
        match = re.compile(r':([a-zA-Z_0-9]+):')
        for emoji_name in re.findall(match, demojized):
            if 'emoji' not in entities:
                entities['emoji'] = []
            entities['emoji'].append({'value': emoji_name})
        # if re.match(match, demojized):
        #    entities['intent'].append({'value':'emoji'})
    return entities
