# Techruption: Multi-Party All Night

shamir_secret_sharing.py contains classes:

- ShamirSecretSharingScheme, with attributes P (prime number, operations done over F_P), n (number of players), t (privacy threshold, or reconstruction threshold minus one), and method share_secret.

- Shares, with attributes ShamirSSS (a class as above), shares (the actual shares), and degree (equal to the above t), and methods reconstruct_secret and sum and multiplication.

## Usage example:

`Scheme = ShamirSecretSharingScheme(31, 5, 2)`
`Shares1 = Scheme.share_secret(4)`
`Shares2 = Scheme.share_secret(9)`
`(Shares1 + Shares2).reconstruct_secret()`
