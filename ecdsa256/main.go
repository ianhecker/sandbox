package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type derivation struct {
	BaseKey    string
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    common.Address
}

func main() {
	base := "355eb67a749754aa659a85ea66deb615bec9616a577078350c935e5a0e33faf5"
	priv, err := crypto.HexToECDSA(base)
	if err != nil {
		panic(err)
	}
	pub, ok := priv.Public().(*ecdsa.PublicKey)
	if !ok {
		panic("cant cast")
	}
	addr := crypto.PubkeyToAddress(*pub)

	d := derivation{base, priv, pub, addr}

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(d)
}
