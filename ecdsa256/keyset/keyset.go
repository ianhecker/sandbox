package keyset

import (
	"crypto/ecdsa"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Keyset struct {
	BaseKey    string
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    common.Address
}

func MakeKeyset(baseKey string) (keyset Keyset) {
	keyset.BaseKey = baseKey
	keyset.derivePrivateKey(keyset.BaseKey)
	keyset.derivePublicKey(keyset.PrivateKey)
	keyset.deriveAddress(keyset.PublicKey)
	return keyset
}

func (k *Keyset) derivePrivateKey(base string) {
	var err error
	k.PrivateKey, err = crypto.HexToECDSA(base)
	if err != nil {
		panic(err)
	}
}
func (k *Keyset) derivePublicKey(priv *ecdsa.PrivateKey) {
	var ok bool
	k.PublicKey, ok = priv.Public().(*ecdsa.PublicKey)
	if !ok {
		panic("couldnt cast to public key")
	}
}
func (k *Keyset) deriveAddress(pub *ecdsa.PublicKey) {
	k.Address = crypto.PubkeyToAddress(*pub)
}

func (k *Keyset) MarshalJSON() ([]byte, error) {
	type Alias Keyset
	return json.MarshalIndent(&struct{ *Alias }{Alias: (*Alias)(k)}, " ", " ")
}

func (k *Keyset) UnmarshalJSON(bytes []byte) error {
	type Alias Keyset
	tmp := &struct{ *Alias }{Alias: (*Alias)(k)}
	return json.Unmarshal(bytes, tmp)
}
