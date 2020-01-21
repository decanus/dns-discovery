# dns-discovery

DNS based discovery for libp2p based on [EIP-1459](https://eips.ethereum.org/EIPS/eip-1459), the implementation is heavily inspired by the code found in [go-ethereum](https://github.com/ethereum/go-ethereum/tree/master/p2p/dnsdisc).
Currently the discovery function wraps around go-ethereum. 

## TODOs
 - Work on changes to go-ethereum that allow for EIP-1459 with various network address types. Abstraction of encoding decoding.
