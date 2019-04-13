# Techruption: Multi-Party All Night



# from shamir_secret_sharing import *
import sympy

n = 11
r = 4
t = r-1
maxVolume = 1e7
P = sympy.nextprime( max(n+1, maxVolume) ) # Need enough numbers to encode all possible volumes, and at least n+1 due to Shamir secret sharing properties
SSScheme = ShamirSecretSharingScheme(P,n,t)
