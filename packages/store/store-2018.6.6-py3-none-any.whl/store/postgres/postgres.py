from copy import copy

from pony.orm import Database, Required, db_session, select, desc, Json, count, Optional, Set, commit
from datetime import datetime

from store.base import BaseStore
from pony.orm import delete as db_delete


class PostgresStore(BaseStore):
    def __init__(self, data):
        db_type = data.get('type', 'postgres')
        db_user = data.get('user', 'dameng')
        db_password = data.get('password', 'hello')
        db_host = data.get('host', 'localhost')
        db_name = data.get('name', 'store')
        unique = data.get('unique', True)
        self.unique = unique

        table = data.get("table", "Store")
        if table[0].islower():
            table = table[0].upper() + table[1:]

        self.db = Database(db_type, user=db_user, password=db_password, host=db_host, database=db_name)
        body = dict(__doc__='docstring',
                    create_at=Required(datetime, sql_default='CURRENT_TIMESTAMP', default=lambda: datetime.utcnow()),
                    update_at=Required(datetime, sql_default='CURRENT_TIMESTAMP', default=lambda: datetime.utcnow()),
                    key=Required(str, index=True, unique=unique),
                    value=Required(Json, volatile=True, index=True),
                    pid=Optional(table, fk_name='pid'),
                    sub=Set(table, reverse='pid')
                    )

        self.Store = type(table, (self.db.Entity,), body)
        self.db.generate_mapping(create_tables=True, check_tables=True)

    @db_session
    def create(self, key, value, id_=None, pid=None):
        if self.unique:
            elem = select(e for e in self.Store if e.key == str(key)).order_by(lambda o: desc(o.create_at)).first()
            if elem is None:
                if id_:
                    elem = self.Store(key=str(key), value=value, id=id_, pid=pid)
                else:
                    elem = self.Store(key=str(key), value=value, pid=pid)
            else:
                elem.value = value
                elem.update_at = datetime.utcnow()
            commit()
            print('elem id:', elem.id)
            return elem.id
        else:
            if id_:
                elem = self.Store(key=str(key), value=value, id=id_, pid=pid)
            else:
                elem = self.Store(key=str(key), value=value, pid=pid)
            commit()
            print('elem id:', elem.id)
            return elem.id

    def _return_elem(self, elem):
        return {
            'id': copy(elem.id),
            'key': copy(elem.key),
            'value': copy(elem.value),
            'create_at': elem.create_at.strftime("%Y-%m-%d %H:%M:%S"),
            'update_at': elem.update_at.strftime("%Y-%m-%d %H:%M:%S"),
        }

    @db_session
    def read(self, key, pagenum=None, pagesize=None):
        if isinstance(key, str):
            if self.unique:
                elem = select(e for e in self.Store if e.key == key).order_by(lambda o: desc(o.create_at)).first()
                if elem:
                    return self._return_elem(elem)
            else:
                query = select(e for e in self.Store if e.key == key).order_by(lambda o: desc(o.create_at))
                elems = query.page(pagenum=pagenum, pagesize=pagesize) if pagenum and pagesize else query[:]
                if elems:
                    return [self._return_elem(elem) for elem in elems]
        elif isinstance(key, int):
            elem = select(e for e in self.Store if e.id == key).order_by(lambda o: desc(o.create_at)).first()
            if elem:
                return self._return_elem(elem)
        return None

    def _update_elem(self, elem, value, mode='update', pid=None):
        if mode == 'update':
            value_db = copy(elem.value)
            if isinstance(value_db, str) and isinstance(value, str):
                elem.value = value
                elem.update_at = datetime.utcnow()
            elif isinstance(value_db, dict) and isinstance(value, dict):
                value_db.update(value)
                elem.value = value_db
                elem.update_at = datetime.utcnow()
            elif isinstance(value_db, list) and isinstance(value, list):
                vv = [v for v in value if v not in value_db]
                value_db.extend(vv)
                elem.value = value_db
                elem.update_at = datetime.utcnow()
            elif type(value_db) != type(value):
                elem.value = value
                elem.update_at = datetime.utcnow()

        elif mode == 'replace':
            elem.value = value
            elem.update_at = datetime.utcnow()

        elif mode == 'delete':
            value_db = copy(elem.value)
            if isinstance(value_db, dict):
                if isinstance(value, dict):
                    pop_keys = value.keys()
                    for k in pop_keys:
                        value_db.pop(k)
                elif isinstance(value, list):
                    for k in value:
                        value_db.pop(k)
                elem.value = value_db
                elem.update_at = datetime.utcnow()
            elif isinstance(value_db, list) and isinstance(value, list):
                updated_db = [e for e in value_db if e not in value]
                elem.value = updated_db
                elem.update_at = datetime.utcnow()

        if pid is not None:
            elem.pid = pid
        commit()

    @db_session
    def update(self, key, value, mode='update', pid=None):
        if isinstance(key, str):
            if self.unique:
                elem = select(e for e in self.Store if e.key == key).order_by(lambda o: desc(o.create_at)).first()
                if elem:
                    self._update_elem(elem, value, mode, pid)
            else:
                elems = select(e for e in self.Store if e.key == key).order_by(lambda o: desc(o.create_at))[:]
                if elems:
                    for elem in elems:
                        self._update_elem(elem, value, mode, pid)
        elif isinstance(key, int):
            elem = select(e for e in self.Store if e.id == key).order_by(lambda o: desc(o.create_at)).first()
            if elem:
                self._update_elem(elem, value, mode, pid)

    @db_session
    def delete(self, key):
        if isinstance(key, str):
            db_delete(e for e in self.Store if e.key == key)
        elif isinstance(key, int):
            db_delete(e for e in self.Store if e.id == key)


    ##########
    # extended read
    @db_session
    def read_prefix(self, key, pagenum=None, pagesize=None):
        query = select(e for e in self.Store if e.key.startswith(key)).order_by(lambda o: desc(o.create_at))
        elems = query.page(pagenum=pagenum, pagesize=pagesize) if pagenum and pagesize else query[:]
        return [self._return_elem(elem) for elem in elems]

    @db_session
    def read_suffix(self, key, pagenum=None, pagesize=None):
        query = select(e for e in self.Store if e.key.endswith(key)).order_by(lambda o: desc(o.create_at))
        elems = query.page(pagenum=pagenum, pagesize=pagesize) if pagenum and pagesize else query[:]
        return [self._return_elem(elem) for elem in elems]

    @db_session
    def read_infix(self, key, pagenum=None, pagesize=None):
        query = select(e for e in self.Store if key in e.key).order_by(lambda o: desc(o.create_at))
        elems = query.page(pagenum=pagenum, pagesize=pagesize) if pagenum and pagesize else query[:]
        return [self._return_elem(elem) for elem in elems]

    @db_session
    def read_multi(self, keys, pagenum=None, pagesize=None):
        query = select(e for e in self.Store if e.key in keys).order_by(lambda o: desc(o.create_at))
        elems = query.page(pagenum=pagenum, pagesize=pagesize) if pagenum and pagesize else query[:]
        return [self._return_elem(elem) for elem in elems]


    @db_session
    def read_all(self, pagenum=None, pagesize=None):
        query = select(e for e in self.Store).order_by(lambda o: desc(o.create_at))
        elems = query.page(pagenum=pagenum, pagesize=pagesize) if pagenum and pagesize else query[:]
        return [self._return_elem(elem) for elem in elems]


    ########
    @db_session
    def delete_multi(self, keys):
        if isinstance(keys, list) and len(keys) > 0:
            if isinstance(keys[0], str):
                db_delete(e for e in self.Store if e.key in keys)
            elif isinstance(keys[0], int):
                db_delete(e for e in self.Store if e.id in keys)

    #########
    @db_session
    def count_key(self, key):
        return count(e for e in self.Store if e.key == key)

    @db_session
    def count_keys(self, keys):
        if isinstance(keys, list):
            return count(e for e in self.Store if e.key in keys)
        else:
            return self.count_key(keys)

    @db_session
    def count_all(self):
        return count(e for e in self.Store)


if __name__ == '__main__':
    import json
    st = PostgresStore({'name': 'new_store', "table": "test"})
    with db_session:
        id_1 = st.create(key='1', value='root', pid=None)
        print(id_1, '????')
        id_2 = st.create(key='2', value='e1', pid=id_1)
        id_3 = st.create(key='3', value='e2', pid=id_1)
        id_3 = st.create(key='4', value='e3', pid=id_2)
        es = st.read_all()
        print(json.dumps(es, indent=2))
        aa = select(e for e in st.Store)
        for a in aa:
            print('-'*40)
            if a.pid:
                print(a.pid.value, a.value)
                print(a.sub)
                print([e.value for e in a.pid.sub])
                print(a.pid)
            else:
                print('....', a.value)
                print(a.sub)
                print(a.pid)
        # st.delete('1')
    # print('haha')
