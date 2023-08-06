import time

from copy import deepcopy


# import sys
# import traceback

class Context:
    def __init__(self, dialog, entities, history, counter, max_depth=30, history_restart_minutes=30):
        self.counter = counter
        self.entities = entities
        self.history = history
        self.max_depth = max_depth
        self.dialog = dialog
        self.history_restart_minutes = history_restart_minutes

    def add_entities(self, new_entities):
        if new_entities is None:
            return
        inserted = {}
        # add all new entities
        for entity,values in new_entities.items():
            if entity not in self.entities:
                self.entities[entity] = []
            # allow also direct passing of {'entity' : 'value'}
            if not isinstance(values, dict) and not isinstance(values, list):
                values = {'value':values}
            if not isinstance(values, list):
                values = [values]
            
            inserted[entity] = deepcopy(values)
            # prepend each value to start of the list with 0 age 
            for value in values:
                value['counter'] = self.counter
                self.entities[entity] = [value] + self.entities[entity]
        self.debug()
        return inserted

    def add_state(self, state_name):
        timestamp = int(time.time())
        if len(self.history) > 20:
            self.history = self.history[-20:]

        state = {
            'name' : state_name,
            'timestamp' : timestamp
        }
        self.history.append(state)

    def clear(self, entities):
        for entity in entities:
            if entity in self.entities:
                del self.entities[entity]

    def get_min_entity_age(self, entities):
        ages = [self.get_age(entity)[1] for entity in entities]
        ages = filter(lambda x: x is not None, ages)
        return min(ages) if ages else None

    def get_history_state(self, index):
        return self.history[index] if len(self.history) > abs(index) else None

    def get_all(self, entity, max_age=None, limit=None, key='value', ignored_values=[]):
        values = []
        if entity not in self.entities:
            return values
        for value in self.entities[entity]:
            age = self.counter-value['counter']
            value['age'] = age
            # if I found a too old value, stop looking
            if max_age is not None and age > max_age:
                break
            v = value.get('value')
            if v in ignored_values:
                self.dialog.log.info('Skipping ignored entity value: {}={}'.format(entity, v))
                continue
            values.append(value.get(key) if key else value.copy())
            # if I already have enough values, stop looking
            if limit is not None and len(values) >= limit:
                break
        return values 
    
    def debug(self, max_age=5):
        self.dialog.log.info('-- HEAD of Context (max age {}): --'.format(max_age))
        for entity in self.entities:
            values = self.get_all_first(entity, key=None, max_age=max_age)
            if values:
                vs = [value.get('value') for value in values]
                self.dialog.log.info('{} (age {}): {}'.format(entity, values[0]['age'], vs if len(vs) > 1 else vs[0]))
        self.dialog.log.info('----------------------------------')

    def get(self, entity, max_age=None, key='value', ignored_values=[]):
        values = self.get_all(entity, max_age=max_age, limit=1, key=key, ignored_values=ignored_values)
        if not values:
            return None
        return values[0]

    def get_age(self, entity, max_age=None, key='value', ignored_values=[]):
        value = self.get_all(entity, max_age=max_age, limit=1, key=None, ignored_values=ignored_values)
        if not value:
            return (None, None)
        return value[0].get(key) if key else value[0], value[0]['age'] 

    def get_all_first(self, entity, max_age=None, key='value'):
        values = []
        if entity not in self.entities:
            return values
        found_age = None
        existing = []
        for value in self.entities[entity]:
            age = self.counter-value['counter']
            value['age'] = age
            # if I found a too old value, stop looking
            if max_age is not None and age > max_age:
                break
            if found_age is not None and age > found_age:
                break
            found_age = age
            if value.get('value') in existing:
                # skip duplicates
                continue
            existing.append(value.get('value'))
            values.append(value.get(key) if key else value)
        return values[::-1]

    def set(self, entity, value_dict):
        if not isinstance(value_dict, dict):
            raise ValueError('Use a dict to set a context value, e.g. {"value":"foo"}. Call multiple times to add more.')
        if entity not in self.entities:
            self.entities[entity] = []
        value_dict['counter'] = self.counter
        self.entities[entity] = [value_dict] + self.entities[entity][:self.max_depth-1]

    def set_value(self, entity, value):
        if entity not in self.entities:
            self.entities[entity] = []
        value_dict = {'counter': self.counter, 'value': value}
        self.entities[entity] = [value_dict] + self.entities[entity][:self.max_depth-1]

    def has_any(self, entities, max_age=None):
        for entity in entities:
            if self.get(entity, max_age=max_age, key=None):
                return True
        return False

    def has_all(self, entities, max_age=None):
        for entity in entities:
            if not self.get(entity, max_age=max_age, key=None):
                return False
        return True

    #def entity_value(self, entity):
    #    values = self.entity_values(entity)
    #    return values[0] if values else None

    #def __str__(self):
    #    res = "Message: %s\n" % (self.text)
    #    for entity in self.entities:
    #        for value in self.entities[entity]:
    #            res += "Entity %s: %s (%s)\n" % (entity, value['value'], value['confidence'])