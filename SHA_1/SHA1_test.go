package main

import (
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

// Test for comparing hash of own SHA-1 vs hash of library SHA-1
func TestHash(t *testing.T) {

	inValue := []byte("Own implementation of the SHA-1 hashing algorithm.")

	// own implementation of SHA-1
	ownHash := SHA_1(inValue)
	ownHashString := hex.EncodeToString(ownHash[:])

	// library implementation of SHA-1
	libHash := sha1.Sum(inValue)
	libHashString := hex.EncodeToString(libHash[:])

	// comparing hashes
	if ownHashString != libHashString {
		t.Errorf("Expected hash %s, the resulting hash %s", libHash, ownHash)
	}
}

// BenchmarkTest for library implementation of SHA-1
func BenchmarkSHA1LibraryTest(b *testing.B) {

	inValue := []byte("Own implementation of the SHA-1 hashing algorithm.")

	// benchmark function
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(inValue)
	}
}

// BenchmarkTest for own SHA-1 implementation
func BenchmarkSHA1OwnTest(b *testing.B) {

	inValue := []byte("Own implementation of the SHA-1 hashing algorithm.")

	// benchmark function
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SHA_1(inValue)
	}
}
