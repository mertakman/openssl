package openssl_test

import (
	"bytes"
	"crypto"
	"testing"

	"github.com/golang-fips/openssl/v2"
)

type tls1prfTest struct {
	hash   crypto.Hash
	secret []byte
	label  []byte
	seed   []byte
	out    []byte
}

var tls1prfTests = []tls1prfTest{
	// TLS 1.0/1.1 test generated with OpenSSL and cross-validated
	// with Windows CNG.
	{
		crypto.MD5SHA1,
		[]byte{
			0x9b, 0xbe, 0x43, 0x6b, 0xa9, 0x40, 0xf0, 0x17,
			0xb1, 0x76, 0x52, 0x84, 0x9a, 0x71, 0xdb, 0x35,
		},
		[]byte{
			0x74, 0x65, 0x73, 0x74, 0x20, 0x6c, 0x61, 0x62,
			0x65, 0x6c},
		[]byte{
			0xa0, 0xba, 0x9f, 0x93, 0x6c, 0xda, 0x31, 0x18,
			0x27, 0xa6, 0xf7, 0x96, 0xff, 0xd5, 0x19, 0x8c,
		},
		[]byte{
			0x66, 0x17, 0x40, 0xe6, 0xf9, 0x8b, 0xc9, 0x01,
		},
	},
	// Tests from https://mailarchive.ietf.org/arch/msg/tls/fzVCzk-z3FShgGJ6DOXqM1ydxms/
	{
		crypto.SHA256,
		[]byte{
			0x9b, 0xbe, 0x43, 0x6b, 0xa9, 0x40, 0xf0, 0x17,
			0xb1, 0x76, 0x52, 0x84, 0x9a, 0x71, 0xdb, 0x35,
		},
		[]byte{
			0x74, 0x65, 0x73, 0x74, 0x20, 0x6c, 0x61, 0x62,
			0x65, 0x6c},
		[]byte{
			0xa0, 0xba, 0x9f, 0x93, 0x6c, 0xda, 0x31, 0x18,
			0x27, 0xa6, 0xf7, 0x96, 0xff, 0xd5, 0x19, 0x8c,
		},
		[]byte{
			0xe3, 0xf2, 0x29, 0xba, 0x72, 0x7b, 0xe1, 0x7b,
			0x8d, 0x12, 0x26, 0x20, 0x55, 0x7c, 0xd4, 0x53,
			0xc2, 0xaa, 0xb2, 0x1d, 0x07, 0xc3, 0xd4, 0x95,
			0x32, 0x9b, 0x52, 0xd4, 0xe6, 0x1e, 0xdb, 0x5a,
			0x6b, 0x30, 0x17, 0x91, 0xe9, 0x0d, 0x35, 0xc9,
			0xc9, 0xa4, 0x6b, 0x4e, 0x14, 0xba, 0xf9, 0xaf,
			0x0f, 0xa0, 0x22, 0xf7, 0x07, 0x7d, 0xef, 0x17,
			0xab, 0xfd, 0x37, 0x97, 0xc0, 0x56, 0x4b, 0xab,
			0x4f, 0xbc, 0x91, 0x66, 0x6e, 0x9d, 0xef, 0x9b,
			0x97, 0xfc, 0xe3, 0x4f, 0x79, 0x67, 0x89, 0xba,
			0xa4, 0x80, 0x82, 0xd1, 0x22, 0xee, 0x42, 0xc5,
			0xa7, 0x2e, 0x5a, 0x51, 0x10, 0xff, 0xf7, 0x01,
			0x87, 0x34, 0x7b, 0x66,
		},
	},
	{
		crypto.SHA384,
		[]byte{
			0xb8, 0x0b, 0x73, 0x3d, 0x6c, 0xee, 0xfc, 0xdc,
			0x71, 0x56, 0x6e, 0xa4, 0x8e, 0x55, 0x67, 0xdf,
		},
		[]byte{
			0x74, 0x65, 0x73, 0x74, 0x20, 0x6c, 0x61, 0x62,
			0x65, 0x6c},
		[]byte{
			0xcd, 0x66, 0x5c, 0xf6, 0xa8, 0x44, 0x7d, 0xd6,
			0xff, 0x8b, 0x27, 0x55, 0x5e, 0xdb, 0x74, 0x65,
		},
		[]byte{
			0x7b, 0x0c, 0x18, 0xe9, 0xce, 0xd4, 0x10, 0xed,
			0x18, 0x04, 0xf2, 0xcf, 0xa3, 0x4a, 0x33, 0x6a,
			0x1c, 0x14, 0xdf, 0xfb, 0x49, 0x00, 0xbb, 0x5f,
			0xd7, 0x94, 0x21, 0x07, 0xe8, 0x1c, 0x83, 0xcd,
			0xe9, 0xca, 0x0f, 0xaa, 0x60, 0xbe, 0x9f, 0xe3,
			0x4f, 0x82, 0xb1, 0x23, 0x3c, 0x91, 0x46, 0xa0,
			0xe5, 0x34, 0xcb, 0x40, 0x0f, 0xed, 0x27, 0x00,
			0x88, 0x4f, 0x9d, 0xc2, 0x36, 0xf8, 0x0e, 0xdd,
			0x8b, 0xfa, 0x96, 0x11, 0x44, 0xc9, 0xe8, 0xd7,
			0x92, 0xec, 0xa7, 0x22, 0xa7, 0xb3, 0x2f, 0xc3,
			0xd4, 0x16, 0xd4, 0x73, 0xeb, 0xc2, 0xc5, 0xfd,
			0x4a, 0xbf, 0xda, 0xd0, 0x5d, 0x91, 0x84, 0x25,
			0x9b, 0x5b, 0xf8, 0xcd, 0x4d, 0x90, 0xfa, 0x0d,
			0x31, 0xe2, 0xde, 0xc4, 0x79, 0xe4, 0xf1, 0xa2,
			0x60, 0x66, 0xf2, 0xee, 0xa9, 0xa6, 0x92, 0x36,
			0xa3, 0xe5, 0x26, 0x55, 0xc9, 0xe9, 0xae, 0xe6,
			0x91, 0xc8, 0xf3, 0xa2, 0x68, 0x54, 0x30, 0x8d,
			0x5e, 0xaa, 0x3b, 0xe8, 0x5e, 0x09, 0x90, 0x70,
			0x3d, 0x73, 0xe5, 0x6f,
		},
	},
	{
		crypto.SHA512,
		[]byte{
			0xb0, 0x32, 0x35, 0x23, 0xc1, 0x85, 0x35, 0x99,
			0x58, 0x4d, 0x88, 0x56, 0x8b, 0xbb, 0x05, 0xeb,
		},
		[]byte{
			0x74, 0x65, 0x73, 0x74, 0x20, 0x6c, 0x61, 0x62,
			0x65, 0x6c,
		},
		[]byte{
			0xd4, 0x64, 0x0e, 0x12, 0xe4, 0xbc, 0xdb, 0xfb,
			0x43, 0x7f, 0x03, 0xe6, 0xae, 0x41, 0x8e, 0xe5,
		},
		[]byte{
			0x12, 0x61, 0xf5, 0x88, 0xc7, 0x98, 0xc5, 0xc2,
			0x01, 0xff, 0x03, 0x6e, 0x7a, 0x9c, 0xb5, 0xed,
			0xcd, 0x7f, 0xe3, 0xf9, 0x4c, 0x66, 0x9a, 0x12,
			0x2a, 0x46, 0x38, 0xd7, 0xd5, 0x08, 0xb2, 0x83,
			0x04, 0x2d, 0xf6, 0x78, 0x98, 0x75, 0xc7, 0x14,
			0x7e, 0x90, 0x6d, 0x86, 0x8b, 0xc7, 0x5c, 0x45,
			0xe2, 0x0e, 0xb4, 0x0c, 0x1c, 0xf4, 0xa1, 0x71,
			0x3b, 0x27, 0x37, 0x1f, 0x68, 0x43, 0x25, 0x92,
			0xf7, 0xdc, 0x8e, 0xa8, 0xef, 0x22, 0x3e, 0x12,
			0xea, 0x85, 0x07, 0x84, 0x13, 0x11, 0xbf, 0x68,
			0x65, 0x3d, 0x0c, 0xfc, 0x40, 0x56, 0xd8, 0x11,
			0xf0, 0x25, 0xc4, 0x5d, 0xdf, 0xa6, 0xe6, 0xfe,
			0xc7, 0x02, 0xf0, 0x54, 0xb4, 0x09, 0xd6, 0xf2,
			0x8d, 0xd0, 0xa3, 0x23, 0x3e, 0x49, 0x8d, 0xa4,
			0x1a, 0x3e, 0x75, 0xc5, 0x63, 0x0e, 0xed, 0xbe,
			0x22, 0xfe, 0x25, 0x4e, 0x33, 0xa1, 0xb0, 0xe9,
			0xf6, 0xb9, 0x82, 0x66, 0x75, 0xbe, 0xc7, 0xd0,
			0x1a, 0x84, 0x56, 0x58, 0xdc, 0x9c, 0x39, 0x75,
			0x45, 0x40, 0x1d, 0x40, 0xb9, 0xf4, 0x6c, 0x7a,
			0x40, 0x0e, 0xe1, 0xb8, 0xf8, 0x1c, 0xa0, 0xa6,
			0x0d, 0x1a, 0x39, 0x7a, 0x10, 0x28, 0xbf, 0xf5,
			0xd2, 0xef, 0x50, 0x66, 0x12, 0x68, 0x42, 0xfb,
			0x8d, 0xa4, 0x19, 0x76, 0x32, 0xbd, 0xb5, 0x4f,
			0xf6, 0x63, 0x3f, 0x86, 0xbb, 0xc8, 0x36, 0xe6,
			0x40, 0xd4, 0xd8, 0x98,
		},
	},
}

func TestTLS1PRF(t *testing.T) {
	if !openssl.SupportsTLS1PRF() {
		t.Skip("TLS PRF is not supported")
	}
	for _, tt := range tls1prfTests {
		tt := tt
		t.Run(tt.hash.String(), func(t *testing.T) {
			if !openssl.SupportsHash(tt.hash) {
				t.Skip("skipping: hash not supported")
			}
			out, err := openssl.TLS1PRF(tt.secret, tt.label, tt.seed, len(tt.out), cryptoToHash(tt.hash))
			if err != nil {
				t.Fatalf("error deriving TLS 1.2 PRF: %v.", err)
			}
			if !bytes.Equal(out, tt.out) {
				t.Errorf("incorrect key output: have %v, need %v.", out, tt.out)
			}
		})
	}
}
