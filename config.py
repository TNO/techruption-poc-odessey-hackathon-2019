# Techruption: Multi-Party All Night
'''
    Configuration file:
    number of MPC nodes,
    reconstruction and privacy thresholds,
    upper bound on volume (input values),
    prime field size,
    secret-sharing scheme,
    urls.
'''

from shamir_secret_sharing import *
import sympy

n = 11
r = 4
t = r-1
maxVolume = 1e7
P = sympy.nextprime( max(n+1, maxVolume) ) # Need enough numbers to encode all possible volumes, and at least n+1 due to Shamir secret sharing properties
SSScheme = ShamirSecretSharingScheme(P,n,t)

responseUrl = "http://localhost:8080"
