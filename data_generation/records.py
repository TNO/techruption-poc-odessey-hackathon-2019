# -*- coding: utf-8 -*-
"""
Created on Fri Apr 12 19:11:47 2019

@author: bosmanjw
"""

import numpy as np
import pandas as pd
# BN id 201.5 pp 52 https://zoek.officielebekendmakingen.nl/blg-784527.pdf
#https://www.prorail.nl/sites/default/files/phs-folder_gevaarlijke_stoffen_web.pdf

categories = np.array(["A", "B2", "B3", "C3", "D3", "D4"])
intensities = np.array([32680 ,1365 ,18120 ,2560 ,560 ,101 ,128550 ,10923 ,11820 ,1532 ,5100 ,152])
intensities = np.reshape(intensities,(6,2)).T[1]
probs = intensities / intensities.sum()
intensities = pd.DataFrame(data=np.atleast_2d(intensities),columns=categories)
substances={"A":["LPG", "propylene", "butadiene", "ethylene oxide"],
 "B2":["ammonia"],
 "B3":["chlorine"],
 "C3":["gasoline", "natural-gas condensate"],
 "D3":["acrylonitrile"],
 "D4":["hydrogen fluoride", "bromide"]
 }

#inhoud 14000-27000 liter (bron wiki)

def generate_records(number, maxIncrement = 10):
    records = []
    currId = 0
    for i in range(number):
        cat = np.random.choice(a=categories, p=probs)
        currId += np.random.randint(1, maxIncrement)
        cId = "{0:d}".format(currId)
        name = np.random.choice(substances[cat])
        vol = np.random.randint(14,27+1) * 1000
        vol = "{0:d}".format(vol)
        record = {"Container ID": cId ,"Substance category": cat, "Name": name, "Volume": vol, "Unit": "litre"}
        records.append(record)
    return records
