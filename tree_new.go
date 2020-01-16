package dns

import ma "github.com/multiformats/go-multiaddr"

type Tree struct {
	r Resolver
}

func (t *Tree) Sync(out chan <- ma.Multiaddr) {
	// @todo sync the entire tree, pump results into a channel maybe?
	// @todo validate root

}

func (t *Tree) doSync(out chan <- ma.Multiaddr) {
	// @todo get record

	// @todo evaluate record type

	// @todo if leaf push into out
}

func (t *Tree) syncBranch() {

}
