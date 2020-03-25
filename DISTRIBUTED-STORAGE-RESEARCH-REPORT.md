# Velas Sphere Distributed Storage Research Report

This document describes how to store files securely and reliable in a trustless p2p environment with a simple economy model where requesters pay providers.

## Status

This research is completed!

## Requirements

1. Provable storage for limited time
2. Provable redundancy for ensuring files aren't lost
3. Only the owner can decrypt the files
4. Only the owner can get the files back
5. Providers don't accept raw files
6. Providers are motivated to perform tasks correctly
7. Requesters are motivated to pay

### Solution

Provider motivation is the following:
    
1. Each node pays a membership fee
2. Malicious nodes are banned via the smart contract

Requesters are motivated in this way:

1. They want a specific task to be performed

We can prove the file is stored on a remote node using a heuristic zero-knowledge algorithm based on challenges and Merkle trees.

1. A requester encrypts an arbitrary file using some asymmetric encryption algorithm (like AES)
2. The requester generates a certain amount of challenges based on desired storage time for the encrypted file, their Merkle tree, and constructs a path to the tree root for each of the challenges. 
3. The requester saves the challenges, their paths, and the file encryption key.
4. The requester sends a request to a provider to store the encrypted file providing the Merkle tree, time to store it, and his public key
5. The provider saves the file, tree, verifies it is encrypted using the entropy estimation and responds "OK"
6. Steps 1-5 can be repeated for redundancy
7. The requester doesn't need to store the file anymore
6. The requester polls the provider giving him each time a different challenge (using the set of saved ones)
7. The provider responds with the path of a to the hash(file+challenge) entry
8. Requester validates the path
9. Eventually, the requester can get his file back providing a signed using his private key message
10. Provider verifies the message, generates an invoice and sends the file back
11. Requester gets his file and sings the invoice