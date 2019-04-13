# Techruption: Multi-Party All Night

'''
    Code to secret-share a plaintext database,
    corresponding to several containers.
'''

from toy_databases import *
from shamir_secret_sharing import *
import sympy

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



def share_database(plaintextDatabase = toyDatabase, n=11, r=4, substanceList = toySubstanceList, categoryList = toyCategoryList, maxVolume = toyMaxVolume):
    P = sympy.nextprime( max(n+1, maxVolume) ) # Need enough numbers to encode all possible volumes, and at least n+1 due to Shamir secret sharing properties
    t = r-1
    enhancedDatabase = {
            'P' : P,
            'n' : n,
            'r' : r,
            't' : t,
            'Containers' : plaintextDatabase }

    # Initialize secret-sharing scheme
    SSScheme = ShamirSecretSharingScheme(P, n, t)
    # Initialize database shares as copies of the enhanced database
    databaseShares = [enhancedDatabase for _ in range(n)]

    # Share each 'Volume' entry of each substance of each container
    for containerIndex in range(len(plaintextDatabase)):
        container = plaintextDatabase[containerIndex]
        containerContent = container['Content']
        for substanceIndex in range(len(substanceList)):
            substance = containerContent[substanceIndex]
            substanceVolume = int(substance['Volume'])
            Shares = SSScheme.share_secret(substanceVolume)
            for pID in range(n):
                databaseShares[pID]['Containers'][containerIndex]['Content'][substanceIndex]['Volume'] = str(Shares.shares[pID])

    return databaseShares, Shares

# if __name__ == "__main__":
#     # from toy_databases import plaintextDatabaseExample
#     databaseShares = share_database(plaintextDatabaseExample, 11, 4)
#     print(databaseShares)
