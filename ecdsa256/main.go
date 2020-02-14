package main

import (
	"fmt"
	"sandbox/ecdsa256/keyset"
)

func main() {
	baseKey := "355eb67a749754aa659a85ea66deb615bec9616a577078350c935e5a0e33faf5"

	k := keyset.MakeKeyset(baseKey)
	b, _ := k.MarshalJSON()
	fmt.Printf("marshal: %s\n", string(b))

	// var k2 keyset.Keyset
	// k2.UnmarshalJSON(b)
	// fmt.Printf("unmarshal: %+v\n", k2)
}
