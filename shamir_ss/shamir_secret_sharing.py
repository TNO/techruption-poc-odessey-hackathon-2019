# Shamir Secret Sharing Library
# Techruption: Multi-Party All Night
import sympy as sp
import gmpy2
import secrets


class ShamirSecretSharingScheme(object):

    def __init__(self,P,n,t):
        self.P = P
        self.n = n
        self.t = t
        # Vandermonde matrix for evaluation of polynomials at points [1,..,n]
        self.Vm = [[pow(i+1,j,P) for j in range(t+1)] for i in range(n)]

    def share_secret(self,s):
        #Sample random polynomial of degree t with constant coeffiecient
        secret_poly = [s] + [secrets.randbelow(self.P) for _ in range(self.t)]
        # Create an array of all the shares
        shares = [ sum([self.Vm[ind][i]*secret_poly[i] % self.P for i in range(self.t+1)]) % self.P for ind in range(self.n)]
        Sharing = Shares(self,shares,self.t)
        return Sharing

class Shares(object):

    def __init__(self,ShamirSSS,shares,degree):
        self.scheme = ShamirSSS
        self.shares = shares
        self.degree = degree                       # The degree of the polynomial used for sharing the secret, i.e. at least degree+1 shares are required to reconstruct.

    def reconstruct_secret(self):
        if len(self.shares)< self.degree+1:
            raise ValueError('Too little shares to reconstruct.')
        # We precompute some values so that the Langrage interpolation.
        # We assume that we can always use shares f(1), ... f(self.degree+1)
        invert_tmp = [gmpy2.invert(i,self.scheme.P) for i in range(1,self.degree+1)]
        lagrange_interpol = [ mult_list([ sign(j-i)*j*invert_tmp[abs(j-i)-1] % self.scheme.P   for j in range(1,self.degree+2) if i!=j ],self.scheme.P)    for i in range(1,self.degree+2) ]
        # Implicit assumption is that share[i+1] is the evaluation of the secret polynomial at i
        # For now, we will use the first self.degree+1 shares of to reconstruct. However, any other subset of size at least t+1
        # will suffice. This choice implies that the scheme is not robust yet!
        secret = int(sum(lagrange_interpol[i]*self.shares[i] % self.scheme.P for i in range(self.degree+1)) % self.scheme.P)
        return secret

    def __add__(self,other):
        if (self.scheme != other.scheme):
            raise ValueError("Different secret sharing schemes have been used, i.e. shares are incompatible.")

        shares= [(self.shares[i]+other.shares[i]) % self.scheme.P for i in range(len(self.shares))]
        degree = max(self.degree,other.degree)
        return Shares(self.scheme,shares,degree)

    def __mul__(self,other):
        if (self.scheme != other.scheme):
            raise ValueError("Different secret sharing schemes have been used, i.e. shares are incompatible.")
        shares= [(self.shares[i]*other.shares[i]) % self.scheme.P for i in range(len(self.shares))]
        degree = self.degree + other.degree
        return Shares(self.scheme,shares,degree)

def sign(a):
    return 2*(a>=0)-1

def mult_list(L,modulus):
    out=1
    for l in L:
        out=out*l % modulus
    return out

# Debugging

#P = 5               #sp.nextprime(2**5)              # prime field size
#n = 3               # Number of players
#t = 1               # Reconstructrion threshold, at least t+1 players are needed to reconstruct the secret

#ShamirSSS = ShamirSecretSharingScheme(P,n,t)
#a=ShamirSSS.share_secret(2)
#print(a.shares)
#print(a.reconstruct_secret())
#print('\n')
