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



def share_database(plaintextDatabase = plaintextDatabaseExample, n=11, r=4, substanceList = toySubstanceList, categoryList = toyCategoryList, maxVolume = toyMaxVoluem):
    P = sympy.nextprime( max(n+1, maxVolume) ) # Need enough numbers to encode all possible volumes, and at least n+1 due to Shamir secret sharing properties
    t = r-1
    enhancedDatabase = {
            'P' : P,
            'n' : n,
            'r' : r,
            't' : t,
            'Entries' : plaintextDatabase }
    databaseShares = enhancedDatabase # placeholder for now
    # WIP code
    # scheme = ShamirSecretSharingScheme(P, n, t)
    # databaseShares = [enhancedDatabase for _ in range(n)]
    # for share in databaseShares:
    #     enhancedDatabase['entries'[i in 0:len(plaintextDatabase)]['Volume']]

    return secretDatabase

if __name__ == "__main__":
    from toy_databases import plaintextDatabaseExample
    secretDatabase = share_database(plaintextDatabaseExample, 11, 4)
    print(secretDatabase)
