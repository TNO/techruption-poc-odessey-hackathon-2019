# Techruption: Multi-Party All Night
'''
    Some toy databases.
    Done for quick testing of the python code.
    Some are json objects (good),
    some are python lists (bad, not nicely portable).
'''

################################################################################

# Containers jsons
# OUTDATED, SEE JOOST SCRIPT

container1 = {
        'id' : '134',}



container2 = {
        'id' : '23',
        'pepper' : 50,
        'ammonia' : 18,
        'plastic ducks' : 45 }

################################################################################

# Ship bills database (indicates containers per ship)
# SUBOPTIMAL, SEE JOOST SCRIPT

ship_content_bills = [
        { 'ship_id' : 'Aurora 2',
        'bill' : [ '134', '23' ] },
        { 'ship_id' : 'HMS Victory',
            'bill' : [ 'Cannon load', 'Sail load'] } ]

################################################################################

# List of possible goods
# OUTDATED, SEE JOOST SCRIPT

goodsList = [ 'chlorine', 'ammonia', 'sand', 'plastic ducks', 'pepper' ]

################################################################################

'''
    Plaintext database.
    Contains a list of of container entries,
    generated with records.generate_records(20).
'''

plaintextDatabaseExample = [{'Container ID': '6', 'Substance category': 'B2', 'Name': 'ammonia', 'Volume': '15000', 'Unit': 'litre'}, {'Container ID': '13', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '16000', 'Unit': 'litre'}, {'Container ID': '19', 'Substance category': 'A', 'Name': 'LPG', 'Volume': '20000', 'Unit': 'litre'}, {'Container ID': '26', 'Substance category': 'A', 'Name': 'propylene', 'Volume': '21000', 'Unit': 'litre'}, {'Container ID': '35', 'Substance category': 'C3', 'Name': 'natural-gas condensate', 'Volume': '26000', 'Unit': 'litre'}, {'Container ID': '37', 'Substance category': 'C3', 'Name': 'natural-gas condensate', 'Volume': '14000', 'Unit': 'litre'}, {'Container ID': '43', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '24000', 'Unit': 'litre'}, {'Container ID': '46', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '20000', 'Unit': 'litre'}, {'Container ID': '50', 'Substance category': 'C3', 'Name': 'natural-gas condensate', 'Volume': '21000', 'Unit': 'litre'}, {'Container ID': '52', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '16000', 'Unit': 'litre'}, {'Container ID': '60', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '14000', 'Unit': 'litre'}, {'Container ID': '68', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '22000', 'Unit': 'litre'}, {'Container ID': '72', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '14000', 'Unit': 'litre'}, {'Container ID': '77', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '20000', 'Unit': 'litre'}, {'Container ID': '86', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '18000', 'Unit': 'litre'}, {'Container ID': '87', 'Substance category': 'D3', 'Name': 'acrylonitrile', 'Volume': '20000', 'Unit': 'litre'}, {'Container ID': '91', 'Substance category': 'C3', 'Name': 'gasoline', 'Volume': '24000', 'Unit': 'litre'}, {'Container ID': '93', 'Substance category': 'B2', 'Name': 'ammonia', 'Volume': '26000', 'Unit': 'litre'}, {'Container ID': '95', 'Substance category': 'C3', 'Name': 'natural-gas condensate', 'Volume': '21000', 'Unit': 'litre'}, {'Container ID': '97', 'Substance category': 'D4', 'Name': 'bromide', 'Volume': '26000', 'Unit': 'litre'}] 

################################################################################

'''
    Database share.
    Obtained from the plaintext database, with modifications as follows:
    - Add meta data:
      - Number n of parties
      - party index pId (in range(n), i.e. from 0 to n-1)
      - Reconstruction threshold r
      - Privacy threshold t (t = r-1)
      - Prime number P expressing size of finite field.
    - Replace 'Volume' field of each entry with a share in finite field F_P.
'''
