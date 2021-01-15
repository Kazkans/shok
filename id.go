package main

import (
	"encoding/base32"
	"encoding/binary"
	"math/rand"
	"strings"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
var random = rand.New(source)

func getRandomID() string {
	bytes := make([]byte, 4)

	i := random.Uint32()
	binary.LittleEndian.PutUint32(bytes, i)
	id := base32.StdEncoding.EncodeToString(bytes)
	return strings.ToLower(strings.TrimSuffix(id, "="))
}
