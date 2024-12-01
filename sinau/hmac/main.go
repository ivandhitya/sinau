package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Fungsi untuk membuat HMAC
func generateHMAC(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// Fungsi untuk memverifikasi HMAC
func verifyHMAC(message, receivedHMAC, secret string) bool {
	expectedHMAC := generateHMAC(message, secret)
	return hmac.Equal([]byte(receivedHMAC), []byte(expectedHMAC))
}

func main() {
	secret := "mySecretKey"
	message := "This is a sensitive message."

	// Membuat HMAC untuk pesan
	hmacValue := generateHMAC(message, secret)
	fmt.Println("Generated HMAC:", hmacValue)

	// Verifikasi HMAC
	isValid := verifyHMAC(message, hmacValue, secret)
	fmt.Println("Is the HMAC valid?", isValid)
}
