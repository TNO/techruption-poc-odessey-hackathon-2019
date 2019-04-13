# Techruption: Multi-Party All Night

'''
    Code to recontruct secret-shared value of database.
    Input share can be of two type:
    - container conten, or
    - substance amount.
'''

from shamir_secret_sharing import *

queryTypes = [ 'Container content', 'Substance amount' ]

def reconstruct_value(SSScheme, shares, queryType):

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
            substanceShares = [int(shares[pID]['Share']['Content'][substanceIndex]['Volume']) for pID in range(SSScheme.n)]
            # print(len(substanceShares))
            # print(SSScheme.n)
            # print(SSScheme.t)
            substanceShares = Shares(SSScheme, substanceShares, SSScheme.t)
            substanceVolume = substanceShares.reconstruct_secret()
            container['Content'][substanceIndex]['Volume'] = substanceVolume
        return container

    if queryType == 'Substance amount' :
        # Initialize substance amount dictionary as first share
        substanceAmount = shares[0]['Share']
        # Then reconstruct each substance volume and fill it in
        substanceVolumeShares = [int(shares[pID]['Share']['Volume']) for pID in range(n)]
        substanceVolumeShares = Shares(SSScheme, substanceVolumeShares, SSScheme.t)
        substanceVolume = substanceVolumeShares.reconstruct_secret()
        substanceAmount['Volume'] = str(substanceVolume)
        return substanceAmount

    else:
        raise ValueError('ERRRRROR! Unknown query type')

if __name__ == "__main__":
    import share_database
    from compute_on_share import mpc_compute
    n = 11
    r = 4
    containerTarget = '13'
    substanceTarget = 'gasoline'
    print('TEMPORARY: Sharing database')
    databaseShares, SSScheme, testShares = share_database.share_database()

    # Option 0
    # containerShares = []
    # print('Retrieving shares relative to container ', containerTarget)
    # for pID in range(n):
    #     containerShare = mpc_compute(SSScheme, databaseShares[pID], 'Container content', containerTarget)
    #     containerShares.append(containerShare)
    # print('Reconstructing container content')
    # container = reconstruct_value(SSScheme, containerShares, 'Container content')
    # print(container)

    # Option 1
    substanceAmountShares = []
    print('Retrieving shares relative to substance ', substanceTarget)
    for pID in range(n):
        substanceAmountShare = mpc_compute(SSScheme, databaseShares[pID], 'Substance amount', substanceTarget)
        substanceAmountShares.append(substanceAmountShare)
    print('Reconstructing substance amount')
    substanceAmount = reconstruct_value(SSScheme, substanceAmountShares, 'Substance amount')
    print(substanceAmount)
