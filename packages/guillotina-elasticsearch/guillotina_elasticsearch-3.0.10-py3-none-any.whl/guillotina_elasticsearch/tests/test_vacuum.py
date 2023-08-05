from guillotina.component import get_adapter
from guillotina.component import get_utility
from guillotina.interfaces import ICatalogUtility
from guillotina.utils import get_content_path
from guillotina_elasticsearch.commands.vacuum import Vacuum
from guillotina_elasticsearch.interfaces import IIndexManager
from guillotina_elasticsearch.tests.utils import add_content
from guillotina_elasticsearch.tests.utils import setup_txn_on_container

import asyncio
import os
import pytest
import random


DATABASE = os.environ.get('DATABASE', 'DUMMY')


@pytest.mark.skipif(DATABASE == 'DUMMY', reason='Not for dummy db')
async def test_adds_missing_elasticsearch_entry(es_requester):
    async with es_requester as requester:
        await add_content(requester)
        await asyncio.sleep(1)

        container, request, txn, tm = await setup_txn_on_container(requester)

        keys = await container.async_keys()
        key = random.choice(keys)
        ob = await container.async_get(key)
        search = get_utility(ICatalogUtility)
        await search.remove(container, [(
            ob._p_oid, ob.type_name, get_content_path(ob)
        )], request=request)

        await asyncio.sleep(1)

        vacuum = Vacuum(txn, tm, request, container)
        await vacuum.setup()
        await vacuum.check_missing()
        await vacuum.check_orphans()

        assert len(vacuum.missing) > 0

        await tm.abort(txn=txn)


@pytest.mark.skipif(DATABASE == 'DUMMY', reason='Not for dummy db')
async def test_removes_orphaned_es_entry(es_requester):
    async with es_requester as requester:
        container, request, txn, tm = await setup_txn_on_container(requester)
        search = get_utility(ICatalogUtility)
        im = get_adapter(container, IIndexManager)
        index_name = await im.get_index_name()
        await search.index(container, {
            'foobar': {
                'title': 'foobar',
                'type_name': 'Item'
            }
        })
        await search.refresh(container, index_name)
        await asyncio.sleep(1)

        vacuum = Vacuum(txn, tm, request, container)
        await vacuum.setup()
        await vacuum.check_missing()
        await vacuum.check_orphans()

        assert len(vacuum.orphaned) == 1

        await tm.abort(txn=txn)
