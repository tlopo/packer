package ntlmssp

import (
	"bytes"
	"testing"
)

// test cases from http://davenport.sourceforge.net/ntlm.html

var username = "user"
var password = "SecREt01"
var target = "DOMAIN"
var challenge = []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}

func TestCalculateNTLMv2Response(t *testing.T) {
	NTLMv2Hash := getNtlmV2Hash(password, username, target)
	ClientChallenge := []byte{0xff, 0xff, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44}
	Time := []byte{0x00, 0x90, 0xd3, 0x36, 0xb7, 0x34, 0xc3, 0x01}
	targetInfo := []byte{0x02, 0x00, 0x0c, 0x00, 0x44, 0x00, 0x4f, 0x00, 0x4d, 0x00, 0x41, 0x00, 0x49, 0x00, 0x4e, 0x00, 0x01, 0x00, 0x0c, 0x00, 0x53, 0x00, 0x45, 0x00, 0x52, 0x00, 0x56, 0x00, 0x45, 0x00, 0x52, 0x00, 0x04, 0x00, 0x14, 0x00, 0x64, 0x00, 0x6f, 0x00, 0x6d, 0x00, 0x61, 0x00, 0x69, 0x00, 0x6e, 0x00, 0x2e, 0x00, 0x63, 0x00, 0x6f, 0x00, 0x6d, 0x00, 0x03, 0x00, 0x22, 0x00, 0x73, 0x00, 0x65, 0x00, 0x72, 0x00, 0x76, 0x00, 0x65, 0x00, 0x72, 0x00, 0x2e, 0x00, 0x64, 0x00, 0x6f, 0x00, 0x6d, 0x00, 0x61, 0x00, 0x69, 0x00, 0x6e, 0x00, 0x2e, 0x00, 0x63, 0x00, 0x6f, 0x00, 0x6d, 0x00, 0x00, 0x00, 0x00, 0x00}

	v := computeNtlmV2Response(NTLMv2Hash, challenge, ClientChallenge, Time, targetInfo)

	if expected := []byte{
		0xcb, 0xab, 0xbc, 0xa7, 0x13, 0xeb, 0x79, 0x5d, 0x04, 0xc9, 0x7a, 0xbc, 0x01, 0xee, 0x49, 0x83,
		0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x90, 0xd3, 0x36, 0xb7, 0x34, 0xc3, 0x01,
		0xff, 0xff, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x0c, 0x00,
		0x44, 0x00, 0x4f, 0x00, 0x4d, 0x00, 0x41, 0x00, 0x49, 0x00, 0x4e, 0x00, 0x01, 0x00, 0x0c, 0x00,
		0x53, 0x00, 0x45, 0x00, 0x52, 0x00, 0x56, 0x00, 0x45, 0x00, 0x52, 0x00, 0x04, 0x00, 0x14, 0x00,
		0x64, 0x00, 0x6f, 0x00, 0x6d, 0x00, 0x61, 0x00, 0x69, 0x00, 0x6e, 0x00, 0x2e, 0x00, 0x63, 0x00,
		0x6f, 0x00, 0x6d, 0x00, 0x03, 0x00, 0x22, 0x00, 0x73, 0x00, 0x65, 0x00, 0x72, 0x00, 0x76, 0x00,
		0x65, 0x00, 0x72, 0x00, 0x2e, 0x00, 0x64, 0x00, 0x6f, 0x00, 0x6d, 0x00, 0x61, 0x00, 0x69, 0x00,
		0x6e, 0x00, 0x2e, 0x00, 0x63, 0x00, 0x6f, 0x00, 0x6d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
	}; !bytes.Equal(v, expected) {
		t.Fatalf("expected %x, got %x", expected, v)
	}
}

func TestCalculateLMv2Response(t *testing.T) {
	NTLMv2Hash := getNtlmV2Hash(password, username, target)
	ClientChallenge := []byte{0xff, 0xff, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44}

	v := computeLmV2Response(NTLMv2Hash, challenge, ClientChallenge)

	if expected := []byte{
		0xd6, 0xe6, 0x15, 0x2e, 0xa2, 0x5d, 0x03, 0xb7, 0xc6, 0xba, 0x66, 0x29, 0xc2, 0xd6, 0xaa, 0xf0, 0xff, 0xff, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44,
	}; !bytes.Equal(v, expected) {
		t.Fatalf("expected %x, got %x", expected, v)
	}
}

func TestToUnicode(t *testing.T) {
	v := toUnicode(password)
	if expected := []byte{0x53, 0x00, 0x65, 0x00, 0x63, 0x00, 0x52, 0x00, 0x45, 0x00, 0x74, 0x00, 0x30, 0x00, 0x31, 0x00}; !bytes.Equal(v, expected) {
		t.Fatalf("expected %v, got %v", expected, v)
	}
}

func TestNTLMhash(t *testing.T) {
	v := getNtlmHash(password)
	if expected := []byte{0xcd, 0x06, 0xca, 0x7c, 0x7e, 0x10, 0xc9, 0x9b, 0x1d, 0x33, 0xb7, 0x48, 0x5a, 0x2e, 0xd8, 0x08}; !bytes.Equal(v, expected) {
		t.Fatalf("expected %v, got %v", expected, v)
	}
}

func TestNTLMv2Hash(t *testing.T) {
	v := getNtlmV2Hash(password, username, target)
	if expected := []byte{0x04, 0xb8, 0xe0, 0xba, 0x74, 0x28, 0x9c, 0xc5, 0x40, 0x82, 0x6b, 0xab, 0x1d, 0xee, 0x63, 0xae}; !bytes.Equal(v, expected) {
		t.Fatalf("expected %v, got %v", expected, v)
	}
}
