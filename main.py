import logging
import sys
import pathlib
import asyncio

import aiohttp_jinja2
import jinja2
from aiohttp import web


PROJECT_ROOT = pathlib.Path(__file__).parent



@aiohttp_jinja2.template('index.html')
async def index(request):
    # async with request.app['db'].acquire() as conn:
    #     cursor = await conn.execute(db.question.select())
    #     records = await cursor.fetchall()
    #     questions = [dict(q) for q in records]
    return {'questions': 1111}

def setup_routes(app):
    app.router.add_get('/', index)

async def init_app():
    app = web.Application()
    # app['config'] = config

    setup_routes(app)

    # db_pool = await init_db(app)

    # setup_session(app, RedisStorage(redis_pool))
    aiohttp_jinja2.setup(
        app,
        loader=jinja2.FileSystemLoader(str(PROJECT_ROOT / 'templates'))
    )


    app.router.add_static('/static/', path=PROJECT_ROOT / 'static', name='static')

    return app

def main():
    # config = load_config(configpath)

    # loop = asyncio.get_event_loop()

    logging.basicConfig(level=logging.DEBUG)
    app = init_app()
    web.run_app(app, port=8080)

if __name__ == '__main__':
    main()