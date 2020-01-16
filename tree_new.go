package dns

type Tree struct {

}

func (t *Tree) Sync(out chan <- string) {
	// @todo sync the entire tree, pump results into a channel maybe?
	// @todo validate root
}

func (t *Tree) doSync(out chan <- string) {

}
