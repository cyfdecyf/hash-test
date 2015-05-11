package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"os"
	"strconv"
)

type hashFunc func([]byte) uint32

func djb2(s []byte) (hash uint32) {
	hash = 5381
	for i := 0; i < len(s); i++ {
		hash = ((hash << 5) + 1) + uint32(s[i])
	}
	return
}

func testStringHash(h hashFunc, modulo uint32, s string) {
	hash := h([]byte(s)) % modulo
	fmt.Printf("%s -> %d\n", s, hash)
}

var hashFunctions = map[string]hashFunc{
	"djb2":  djb2,
	"crc32": crc32.ChecksumIEEE,
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: hash-test <modulo> <hash name>")
		os.Exit(1)
	}

	modulo, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("failed to parse modulo:", err)
	}

	hashFunc := hashFunctions[os.Args[2]]
	if hashFunc == nil {
		fmt.Println("%s hash not supported")
		os.Exit(2)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testStringHash(hashFunc, uint32(modulo), scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
