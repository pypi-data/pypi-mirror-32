import json
import logging

import requests

from golem.core.parsing import date_utils
from golem.core.parsing.entity_extractor import EntityExtractor


class DucklingExtractor(EntityExtractor):

    def __init__(self, url, lang='en_US'):
        super().__init__()
        if not url:
            raise ValueError("Duckling URL must be set")
        self.duckling_url = url
        self.language = lang

    def extract_entities(self, text: str, max_retries=1):
        """
        Makes a duckling request for text entities.
        :param text: Text to be parsed by Duckling.
        :return: Json returned by Duckling. Empty on error.
        """
        payload = {
            'locale': self.language,
            'text': text
        }
        try:
            resp = requests.post(self.duckling_url + "/parse", data=payload)
            if resp.status_code == 200:
                jsn = resp.json()
                logging.debug('Duckling:', jsn)
                if jsn is not None:
                    return self.to_entities(jsn)
            else:
                resp.raise_for_status()
        except Exception:
            if max_retries > 0:
                return self.extract_entities(text, max_retries - 1)
            else:
                logging.exception('Exception @ Duckling')
        return {}

    def to_entities(self, jsn):
        """Converts duckling output to the correct format."""
        entities = {}
        for entity in jsn:
            key, value = entity['dim'], entity['value']
            if key == 'time': key = 'datetime'
            if key not in entities:
                entities[key] = []
            entities[key].append(value)
        return self._process_wit_entities(entities)

    def _process_wit_entities(self, entities: dict):

        entities = self._process_metadata(entities)

        if 'datetime' in entities:
            datetime = entities['datetime']
            duration = entities.get('duration', None)
            append = date_utils.process_datetime(datetime, duration)
            entities.update(append)

        return entities

    def _process_metadata(self, entities: dict):
        for entity, values in entities.items():
            for value in values:
                # parse string metadata from Wit into a dict
                metadata = value.get('metadata')
                if metadata and isinstance(metadata, str):
                    try:
                        value['metadata'] = json.loads(metadata)
                    except:
                        self.log.warning("Ignoring invalid metadata for entity {}: {}".format(
                            entity, metadata
                        ))
                        value['metadata'] = None
        return entities
