package openssl_test

import (
	"testing"

	"github.com/golang-fips/openssl/v2"
)

func TestRand(t *testing.T) {
	_, err := openssl.NewRandReader().Read(make([]byte, 5))
	if err != nil {
		t.Fatal(err)
	}
}
