package dns

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/p2p/dnsdisc"
	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

type Resolver interface {
	LookupTXT(ctx context.Context, name string) ([]string, error)
}

type dnsDiscovery struct {
	domain string
	r Resolver
	c dnsdisc.Client
}

func NewDNSDiscovery(domain string, r Resolver, c dnsdisc.Client) discovery.Discoverer {
	return &dnsDiscovery{
		domain: domain,
		r: r,
		c: c,
	}
}

func (d *dnsDiscovery) FindPeers(ctx context.Context, ns string, opts ...discovery.Option) (<-chan peer.AddrInfo, error) {
	// Get options
	var options discovery.Options
	err := options.Apply(opts...)
	if err != nil {
		return nil, err
	}

	// @todo use options

	c := make(chan peer.AddrInfo)

	tree, err := d.c.SyncTree("ns")
	if err != nil {
		return nil, err
	}

	if tree == nil {
		return nil, errors.New("empty tree returned")
	}

	nodes := tree.Nodes()

	go func() {
		for _, n := range nodes {
			addr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d/udp/%d", n.IP(), n.TCP(), n.UDP()))
			if err != nil {
				continue
			}

			p, err := peer.AddrInfoFromP2pAddr(addr)
			if err != nil {
				continue
			}

			c <- *p
		}
	}()

	return c, nil
}
