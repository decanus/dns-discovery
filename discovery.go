package dns

import (
	"context"
	"errors"
	"strings"

	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
)

const RootPrefix = ""

type Resolver interface {
	LookupTXT(ctx context.Context, name string) ([]string, error)
}

type dnsDiscovery struct {
	domain string
	r Resolver
}

func NewDNSDiscovery() discovery.Discoverer {
	return &dnsDiscovery{}
}

func (d *dnsDiscovery) FindPeers(ctx context.Context, ns string, opts ...discovery.Option) (<-chan peer.AddrInfo, error) {
	// Get options
	var options discovery.Options
	err := options.Apply(opts...)
	if err != nil {
		return nil, err
	}

	root, err := d.resolveRoot(ctx)
	if err != nil {
		return nil, err
	}

	c := make(chan peer.AddrInfo)

	// @todo start resolving rest of the root

	return c, nil
}

func (d *dnsDiscovery) resolveRoot(ctx context.Context) (string, error) {
	txts, err := d.r.LookupTXT(ctx, d.domain)
	if err != nil {
		return "", err
	}

	for _, txt := range txts {
		if strings.HasPrefix(txt, RootPrefix) {

		}
	}

	return "", errors.New("new root record found")
}
