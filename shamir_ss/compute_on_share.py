# Techruption: Multi-Party All Night

'''
    Code to accomplish two goals:
    - send shares of content of a container to relevant party;
    - or send shares of sum of volumes of a given substance in ship to relevant party.
'''

from shamir_secret_sharing import *
import sys
sys.path.append('../')
import config
import interface.httpInterface as httpInterface

queryTypes = [ 'Container content', 'Substance amount' ]

toySubstanceList = ['ammonia', 'gasoline']
toyCategoryList = ['B2', 'C3']
toyMaxVolume = 30000

toyDatabase = [
        {'Container ID': '6', 
            'Content' : [{'Substance category': 'B2', 'Name': 'ammonia', 'Volume': '15000', 'Unit': 'litre'}, {'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '0', 'Unit': 'litre'}]}, 
        {'Container ID' : '13', 
            'Content' : [{'Substance category': 'B2', 'Name': 'ammonia', 'Volume': '0', 'Unit': 'litre'}, {'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '16000', 'Unit': 'litre'}]},
        {'Container ID' : '15', 
            'Content' : [{'Substance category': 'B2', 'Name': 'ammonia', 'Volume': '0', 'Unit': 'litre'}, {'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '20000', 'Unit': 'litre'}]} ]

# input: 
#     - either container identifier
#     - or ship and (optional?) substance name

def search_container_index(containerTarget, containerList):
    '''
    Functions that looks for a container with a given ID
    in a list of containers, and returns its index.
    It is assumed that the container IDs in the list are unique.
    '''
    for index in range(len(containerList)):
        if containerList[index]['Container ID'] == str(containerTarget):
            return index

    return -1

def search_substance_index(substanceTarget, substanceList):
    '''
    Functions that looks for a substance with a given name
    in a list of substances, and returns its index.
    It is assumed that the substances in the list appear only once.
    '''
    for index in range(len(substanceList)):
        if substanceList[index]['Name'] == substanceTarget:
            return index

    return -1


def mpc_compute(databaseShare, queryType, target, SSScheme=config.SSScheme):

    if queryType == 'Container content':
        targetIndex = search_container_index(target, databaseShare['Containers'])
        share = {
            'queryType' : queryType,
            'target' : target,
            'P' : databaseShare['P'],
            'n' : databaseShare['n'],
            'r' : databaseShare['r'],
            't' : databaseShare['t'],
            'pID' : databaseShare['pID'],
            'Share' : databaseShare['Containers'][targetIndex] } 

        return share

    if queryType == 'Substance amount':
        # Initialize share with 0 amount of substance
        share = {
            'queryType' : queryType,
            'target' : target,
            'P' : databaseShare['P'],
            'n' : databaseShare['n'],
            'r' : databaseShare['r'],
            't' : databaseShare['t'],
            'pID' : databaseShare['pID'],
            'Share' : {
                'Substance category' : 'TBD',
                'Name' : target,
                'Volume' : '0',
                'Unit' : 'litre' } }

        containers = databaseShare['Containers']
        for container in containers:
            substanceIndex = search_substance_index(target, container['Content'])
            share['Share']['Volume'] = str(int(share['Share']['Volume']) + int(container['Content'][substanceIndex]['Volume']) % SSScheme.P)

        return share

    else:
        raise ValueError('Unknown query type')

import share_database
import json
import asyncio


def start(param, share):
    # print('Retrieving share of player 2 relating to content of container 13')
    # share = mpc_compute(SSScheme, databaseShares[2], 'Container content', '13')
    print('Starting')
    targetUrl = ''
    requestor = httpInterface.Requestor(config.responseUrl)
    fh = open("/home/mpc/compute-initiator/fifo" + str(param), 'r')
    loop = asyncio.get_event_loop()
    while True:
        text = fh.readline()
        print("Received: ", text)
        data = json.loads(text)
	
        queryType = data['QueryType']
        # identifier = data['Identifier'] # Ensure this number is in database
        identifier = 1
        attribute = data['Attribute']
        clientReference = data['ClientReference']
        queryTypeList = ['Container content', 'Substance amount']
        share = mpc_compute(share, queryTypeList[queryType], identifier)
        jsonData = json.dumps(dict({"id": clientReference, "share": share}))
        print('Compute done: ', share)
        result = loop.run_until_complete(requestor.send_request(jsonData))
        #print('Result: ', json.loads(result['data']))

def main():
    databaseShares = share_database.share_database()
    fh0=open("share0", "wb")
    fh1=open("share1", "wb")
    fh2=open("share2", "wb")
    fh3=open("share3", "wb")
    import pickle
    pickle.dump(databaseShares[0], fh0)
    pickle.dump(databaseShares[1], fh1)
    pickle.dump(databaseShares[2], fh2)
    pickle.dump(databaseShares[3], fh3)
    fh0.close()
    fh1.close()
    fh2.close()
    fh3.close()
    

if __name__ == "__main__":
    #main()
    index = sys.argv[1]
    fh = open("share"+index,"rb")
    import pickle
    share = pickle.load(fh)
    fh.close()
    start(int(index), share)
