package main

import (
	"encoding/base64"
	"testing"
)

func TestDistance(t *testing.T) {
	h1, _ := base64.StdEncoding.DecodeString("EX0BcwIkBf8MBQf/CgUA/CABAe1lHwhkCdYDFhY7IRggKTUkGRUlGi8MIBrgFQcZDLwHDCMsHQw+M0AWRjA6GChJLB7bGwUUAswHCD0dEw4eSh0rMSUPJxguCyngBQgMAtUSBhcTHAsIGR0UCBUjDREJJwfdABgJAXCLAAUL5AAEB+MABQjsAA8B8QBkAJMA")
	h2, _ := base64.StdEncoding.DecodeString("AngArQQbAf8XEAL/GwgC/xcIA+ZqAQGXBeYEEQ00DRomIisgNg0qKzAWDRvOBAMcBd4NARMgHAFDIUgVMTI4DScrFALeEQoGBNIHAhQYDAQvIxghLyUQHCgpCgbZFQoEBMoUABMSHgMZGC8GIBwiDikjFQ/SCx8CAGWaAAkH3wENC9gBER7rBRkI3QRmAYsA")

	distance := calcDistance(h1, h2)
	if distance != 28315 {
		t.Error("wrong distance", distance)
	}
}

func BenchmarkDistance(b *testing.B) {
    h1, _ := base64.StdEncoding.DecodeString("EX0BcwIkBf8MBQf/CgUA/CABAe1lHwhkCdYDFhY7IRggKTUkGRUlGi8MIBrgFQcZDLwHDCMsHQw+M0AWRjA6GChJLB7bGwUUAswHCD0dEw4eSh0rMSUPJxguCyngBQgMAtUSBhcTHAsIGR0UCBUjDREJJwfdABgJAXCLAAUL5AAEB+MABQjsAA8B8QBkAJMA")
	h2, _ := base64.StdEncoding.DecodeString("AngArQQbAf8XEAL/GwgC/xcIA+ZqAQGXBeYEEQ00DRomIisgNg0qKzAWDRvOBAMcBd4NARMgHAFDIUgVMTI4DScrFALeEQoGBNIHAhQYDAQvIxghLyUQHCgpCgbZFQoEBMoUABMSHgMZGC8GIBwiDikjFQ/SCx8CAGWaAAkH3wENC9gBER7rBRkI3QRmAYsA")

	b.ResetTimer()

    for i := 0; i < b.N; i++ {
        calcDistance(h1, h2)
    }
}
