package main

import (
	"sandbox/ecdsa256/keyset"
	"sandbox/ecdsa256/write"
)

func main() {
	baseKey := "b8a55f78560a345a5c006300a146c4b2f479874450e32389d1f2761e1efb5efb"

	k := keyset.MakeKeyset(baseKey)
	b, err := k.MarshalJSON()
	if err != nil {
		panic(err)
	}
	err = write.WriteFile(b, baseKey+".json")
	if err != nil {
		panic(err)
	}
}
