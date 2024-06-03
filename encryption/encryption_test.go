package encryption

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "mysecretkey12345" // 16 bytes key for AES-128
	plaintext := "Hello, World!"

	// Test encryption
	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// Test decryption
	decryptedText, err := Decrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if decryptedText != plaintext {
		t.Errorf("Expected %s, but got %s", plaintext, decryptedText)
	}
}

func TestEncryptDecryptWithInvalidKey(t *testing.T) {
	key := "mysecretkey12345"        // 16 bytes key for AES-128
	invalidKey := "invalidkey123456" // 16 bytes invalid key
	plaintext := "Hello, World!"

	// Test encryption
	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// Test decryption with invalid key
	_, err = Decrypt(ciphertext, invalidKey)
	if err == nil {
		t.Fatal("Expected decryption to fail with invalid key, but it succeeded")
	}
}

func TestDecryptMalformedCiphertext(t *testing.T) {
	key := "mysecretkey12345" // 16 bytes key for AES-128
	malformedCiphertext := "malformedciphertext"

	// Test decryption with malformed ciphertext
	_, err := Decrypt(malformedCiphertext, key)
	if err == nil {
		t.Fatal("Expected decryption to fail with malformed ciphertext, but it succeeded")
	}
}
