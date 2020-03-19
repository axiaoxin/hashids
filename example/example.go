package main

import (
	"log"

	"github.com/axiaoxin/hashids"
)

func main() {
	salt := "my-salt:appid:region:uin"
	minLen := 8
	prefix := ""
	h, err := hashids.New(salt, minLen, prefix)
	if err != nil {
		log.Fatal(err)
	}
	var id int64 = 1
	strID, err := h.Encode(id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("int64 id %d encode to %s", id, strID)

	int64ID, err := h.Decode(strID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("string id %s decode to %d", strID, int64ID)
}
