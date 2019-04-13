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
totalIntensity = intensities.sum()
intensities = pd.DataFrame(data=np.atleast_2d(intensities),columns=categories)
substances={"A":["LPG", "propylene", "butadiene", "ethylene oxide"],
 "B2":["ammonia"],
 "B3":["chlorine"],
 "C3":["gasoline", "natural-gas condensate"],
 "D3":["acrylonitrile"],
 "D4":["hydrogen fluoride", "bromide"]
 }
weights = np.array(
        [intensities[cat][0]/len(substanceList) 
            for cat, substanceList in substances.items() 
                for substance in substanceList])
probs = weights / weights.sum()
#inhoud 14000-27000 liter (bron wiki)

def generate_container(containerID, minVolume=14, maxVolume=27):
    contentsList = []
    intensityList = []
    for cat, substanceList in substances.items():
        intensity = intensities[cat][0]/len(substanceList)
        for name in substanceList:            
            record = {"Substance category": cat, "Name": name, "Volume": "0", "Unit": "litre"}
            contentsList.append(record)
            intensityList.append(intensity)
    intensityList = np.array(intensityList)
    probs = intensityList / intensityList.sum()
    index = np.random.choice(range(len(probs)),p=probs)
    vol = np.random.randint(minVolume,maxVolume+1) * 1000
    vol = "{0:d}".format(vol)
    contentsList[index]["Volume"] = vol
    return {"Container ID": containerID, "Content": contentsList}

def generate_containers(number, maxIdIncrement = 10):
    containerList = []
    currId = 0
    for i in range(number):
        currId += np.random.randint(1, maxIdIncrement)
        containerId = "{0:d}".format(currId)
        nextContainer = generate_container(containerId)
        containerList.append(nextContainer)
    return containerList

def generate_ships(number, maxIncrement = 10, minContainers=100, maxContainers=1000):
    shipList = []
    currId = 0
    counts = np.random.randint(minContainers, maxContainers+1, size=number)
    counts = np.atleast_1d(counts)
    allContainersList = generate_containers(counts.sum())
    containerIds = [container["Container ID"] for container in allContainersList]
    np.random.permutation(containerIds)
    containerLists = np.split(containerIds, counts.cumsum()[:-1])
    
    for containerList in containerLists:
        currId += np.random.randint(1, maxIncrement)
        shipId = "{0:d}".format(currId)
        record = {"Ship ID": shipId , "Container list": containerList.tolist()}
        shipList.append(record)
    return shipList, allContainersList

def find_ship(shipList, containerId):
    return [ship for ship in shipList if containerId in ship["Container list"]]