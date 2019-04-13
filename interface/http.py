# -*- coding: utf-8 -*-
"""
Created on Sat Apr 13 12:27:30 2019

@author: bosmanjw
"""

from aiohttp import web
import aiohttp

class WebListener(object):
    def __init__(self, response_handle):
        self.app = web.Application()
        self.response_handle = response_handle
        self.app.add_routes([web.post('/', self.handler)])
        
    async def handler(self, request):
        data = await request.post()
        responseData = self.response_handle(dict(**data))
        return web.json_response(responseData)

    def start(self):
        web.run_app(self.app)

async def send_request(url, data):
    async with aiohttp.ClientSession() as session:
        async with session.post(url, data=data) as resp:
            status = resp.status
            body = await resp.text()
            return {"status":status, "body":body}

if __name__ == "__main__":
    def handle(arg):
        return arg
    
    w=WebListener(handle)

    w.start()
