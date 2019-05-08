# -*- coding: utf-8 -*-
"""
Created on Sat Apr 13 16:12:34 2019

@author: bosmanjw
"""

import httpInterface
import sys
sys.path.append('../')
sys.path.append('../shamir_ss')
from shamir_ss.reconstruct_value import reconstruct_value
import config
import asyncio
import json

class Combiner(object):
    def __init__(self, SSScheme, responseUrl):
        self.requestMap = {}
        self.SSScheme = SSScheme
        self.requestor = httpInterface.Requestor(responseUrl)
        self.listener = httpInterface.Listener(self.handle_request)
        self.loop = asyncio.get_event_loop()

    def ensure_key_exists(self, key):
        if key not in self.requestMap.keys():
            self.requestMap[key] = []

    def has_sufficient_data(self, requestId, share):
        return len(self.requestMap[requestId]) >= share['r']

    def request(self, requestId, share):
        print(type(share))
        #Create entry if it does not exist yet
        self.ensure_key_exists(requestId)
        
        #Append request data
        self.requestMap[requestId].append(share)
        
        if self.has_sufficient_data(requestId, share):
            print('Sufficient shares collected!')
            shares = self.requestMap[requestId]
            #Todo (Joost): queue task in worker pool
            value = reconstruct_value(shares, SSScheme=self.SSScheme)
            data = {"id": requestId, "result": json.dumps(value)}
            print('Combining completed!', data)
            asyncio.ensure_future(self.requestor.send_request(data))

    def handle_request(self, data):
        body = data
        requestId = body['id']
        share = body['share']
        return self.request(requestId, share)
    
    def start(self):
        self.listener.start()
                
if __name__ == "__main__":
    combiner = Combiner(config.SSScheme, config.phoneInterfaceUrl)
    combiner.start()
