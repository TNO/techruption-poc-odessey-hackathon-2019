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


def mpc_compute(databaseShare, queryType, target):

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
        pass

    else:
        return 'ERRRRRROR!'


if __name__ == "__main__":
    import share_database
    databaseShares, Scheme, testShares = share_database.share_database()
    share = mpc_compute(databaseShares[2], 'Container content', '13')
    print(share)


