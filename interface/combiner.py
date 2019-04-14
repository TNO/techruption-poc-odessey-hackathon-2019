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

class Combiner(object):
    def __init__(self, SSScheme, responseUrl):
        self.requestMap = {}
        self.SSScheme = SSScheme
        self.requestor = httpInterface.Requestor(responseUrl)
        self.listener = httpInterface.Listener(self.handle_request)

    def ensure_key_exists(self, key):
        if key not in self.requestMap.keys():
            self.requestMap[key] = []

    def has_sufficient_data(self, requestId, share):
        return self.requestMap[requestId] >= share['r']

    def request(self, requestId, share):
        #Create entry if it does not exist yet
        self.ensure_key_exists(requestId)
        
        #Append request data
        self.requestMap[requestId] += share
        
        if self.has_sufficient_data(requestId, share):
            print('Sufficient shares collected!')
            shares = self.requestMap[requestId]
            #Todo (Joost): queue task in worker pool
            value = reconstruct_value(self.SSScheme, shares)
            print('Combining completed!')
            data = {"id": requestId, "result": value}
            self.requestor.send_request(data)
    
    async def handle_request(self, data):
        requestId = data['id']
        share = data['share']
        return self.request(requestId, share)
    
    def start(self):
        self.listener.start()
                
if __name__ == "__main__":
    combiner = Combiner(config.SSScheme, config.responseUrl)
    combiner.start()
