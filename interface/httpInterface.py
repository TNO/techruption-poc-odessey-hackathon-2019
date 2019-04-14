# -*- coding: utf-8 -*-
"""
Created on Sat Apr 13 12:27:30 2019

@author: bosmanjw
"""

from aiohttp import web
import aiohttp

class Listener(object):
    def __init__(self, response_handle, port=8080):
        self.app = web.Application()
        self.port = port
        self.response_handle = response_handle
        self.app.add_routes([web.post('/', self.handler)])
        
    async def handler(self, request):
        data = await request.post()
        self.response_handle(dict(**data))
#        try:
#            responseData = await self.response_handle(dict(**data))
#        except:
#            responseData = {}
        return web.Response()

    def start(self):
        web.run_app(self.app, port=self.port)

class Requestor(object):
    def __init__(self, url):
        self.url = url
    
    async def send_request(self, data):
        async with aiohttp.ClientSession() as session:
            async with session.post(self.url, data=data) as resp:
                status = resp.status
                body = await resp.text()
                return {"status":status, "body":body}

    async def send_json(self, data):
        async with aiohttp.ClientSession() as session:
            async with session.post(self.url, json=data) as resp:
                status = resp.status
                body = await resp.text()
                return {"status":status, "body":body}

if __name__ == "__main__":
    def handle(arg):
        return arg
    
    w=Listener(handle)

    w.start()
