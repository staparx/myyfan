package main

import (
	"math/rand"
	"time"
	"unsafe"
)

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type StringsPicker struct {
	ownRand *rand.Rand
}

func (p *StringsPicker) Pick(strArr ...string) string {
	return strArr[p.ownRand.Intn(len(strArr))]
}

var Picker = &StringsPicker{ownRand: rand.New(rand.NewSource(time.Now().Unix()))}
