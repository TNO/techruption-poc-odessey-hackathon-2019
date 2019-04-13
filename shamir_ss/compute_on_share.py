# Techruption: Multi-Party All Night

'''
    Code to accomplish two goals:
    - send shares of content of a container to relevant party;
    - or send shares of sum of volumes of a given substance in ship to relevant party.
'''

from shamir_secret_sharing import *

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
        if containerList[index]['Container ID'] == containerTarget:
            return index

def search_substance_index(substanceTarget, substanceList):
    '''
    Functions that looks for a substance with a given name
    in a list of substances, and returns its index.
    It is assumed that the substances in the list appear only once.
    '''
    for index in range(len(substanceList)):
        if substanceList[index]['Name'] == substanceTarget:
            return index


def mpc_compute(SSScheme, databaseShare, queryType, target):

    if queryType == 'Container content':
        targetIndex = search_container_index(target, databaseShare['Containers'])
        share = {
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
        raise ValueError('ERRRRRROR! Unknown query type')


if __name__ == "__main__":
    import share_database
    databaseShares, SSScheme, testShares = share_database.share_database()
    # print('Retrieving share of player 2 relating to content of container 13')
    # share = mpc_compute(SSScheme, databaseShares[2], 'Container content', '13')
    print('Retrieving sum of shares of player 2 relating to ammonia')
    share = mpc_compute(SSScheme, databaseShares[2], 'Substance amount', 'ammonia')
    print(share)
