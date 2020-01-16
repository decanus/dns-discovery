package dns

import (
	"context"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	ma "github.com/multiformats/go-multiaddr"
)

var (
	b64format = base64.RawURLEncoding
	b32format = base32.StdEncoding.WithPadding(base32.NoPadding)
)

type Tree struct {
	name string
	r Resolver
}

func (t *Tree) Sync(ctx context.Context, out chan <- ma.Multiaddr) {
	// @todo sync the entire tree, pump results into a channel maybe?
	// @todo validate root

}

func (t *Tree) doSync(ctx context.Context, out chan <- ma.Multiaddr, name string) {
	root, err := t.resolveRoot(ctx, name)
	if err != nil {
		// @todo
	}

	// @todo get record

	// @todo evaluate record type

	// @todo if leaf push into out
}

func (t *Tree) resolveRoot(ctx context.Context, name string) (root, error) {
	txts, err := t.r.LookupTXT(ctx, name)
	if err != nil {
		return root{}, err
	}

	for _, txt := range txts {
		if strings.HasPrefix(txt, RootPrefix) {
			return parseRoot(txt)
		}
	}

	return root{}, errors.New("new root record found")
}

func (t *Tree) syncBranch() {

}