package dns

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type root struct {
	root string
	lroot string
	seq uint
	sig []byte
}

func parseRoot(e string) (root, error) {
	var eroot, lroot, sig string
	var seq uint
	if _, err := fmt.Sscanf(e, RootPrefix+" r=%s l=%s seq=%d sig=%s", &eroot, &lroot, &seq, &sig); err != nil {
		return root{}, errors.New("invalid root")
	}
	if !isValidHash(eroot) || !isValidHash(lroot) {
		return root{}, errors.New("invalid hashes found")
	}
	sigb, err := b64format.DecodeString(sig)
	if err != nil || len(sigb) != crypto.SignatureLength {
		return root{}, errors.New("invalid signature")
	}
	return root{eroot, lroot, seq, sigb}, nil
}

func isValidHash(s string) bool {
	dlen := b32format.DecodedLen(len(s))
	if dlen < minHashLength || dlen > 32 || strings.ContainsAny(s, "\n\r") {
		return false
	}
	buf := make([]byte, 32)
	_, err := b32format.Decode(buf, []byte(s))
	return err == nil
}
