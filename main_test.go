package main

import (
	"encoding/base64"
	"encoding/csv"
	"os"
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

func BenchmarkOneFile(b *testing.B) {
	hashCsv, err := os.Open("hashes.csv")
	if err != nil {
		panic(err)
	}
	defer hashCsv.Close()

	csvLines, err := csv.NewReader(hashCsv).ReadAll()
	if err != nil {
		panic(err)
	}

	files := make([]File, 0, len(csvLines))

	for _, line := range csvLines {
		pdna, _ := base64.StdEncoding.DecodeString(line[1])
		if len(pdna) != 144 {
			panic("wrong hash length for file: " + line[0])
		}
		matchInfo := File{
			FileName:       line[0],
			PhotoDNA:       pdna,
			BestMatchIndex: -1,
			BestMatchScore: int(^uint(0) >> 1),
		}
		files = append(files, matchInfo)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hashForFile(len(files)/2, files)
	}
}
