# Techruption: Multi-Party All Night

'''
    Code to recontruct secret-shared value of database.
    Input share can be of two type:
    - container conten, or
    - substance amount.
'''

from shamir_secret_sharing import *
from compute_on_share import mpc_compute

queryTypes = [ 'Container content', 'Substance amount' ]

def reconstruct_value(Scheme, shares, queryType):

    if queryType == 'Container content' :
        # P = shares[0]['P']
        # n = shares[0]['n']
        # r = shares[0]['r']
        # t = shares[0]['t']
        # Initialize container dictionary as first share
        container = shares[0]['Share']
        # print(container)
        # Then reconstruct each substance volume and fill it in
        for substanceIndex in range(len(shares[0]['Share']['Content'])):
            substanceShares = [int(shares[pID]['Share']['Content'][substanceIndex]['Volume']) for pID in range(Scheme.n)]
            # print(len(substanceShares))
            # print(Scheme.n)
            # print(Scheme.t)
            substanceShares = Shares(Scheme, substanceShares, Scheme.t)
            substanceVolume = substanceShares.reconstruct_secret()
            container['Content'][substanceIndex]['Volume'] = substanceVolume
        return container

    if queryType == 'Substance amount' :
        pass

    else:
        return 'ERRRRROR!'

if __name__ == "__main__":
    import share_database
    n = 11
    r = 4
    containerTarget = '13'
    databaseShares, Scheme, testShares = share_database.share_database()
    containerShares = []
    for pID in range(n):
        containerShare = mpc_compute(databaseShares[pID], 'Container content', containerTarget)
        containerShares.append(containerShare)

    container = reconstruct_value(Scheme, containerShares, 'Container content')
    print(container)
