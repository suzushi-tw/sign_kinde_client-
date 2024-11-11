package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	// Replace these values with your actual credentials
	teamID := ""
	clientID := ""
	keyID := ""
	keyFilePath := "path/to/key.p8"

	// Read the private key file
	keyBytes, err := os.ReadFile(keyFilePath)
	if err != nil {
		fmt.Printf("Error reading key file: %v\n", err)
		return
	}

	// Parse the private key
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		fmt.Println("Error: Failed to decode PEM block")
		return
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("Error parsing private key: %v\n", err)
		return
	}

	// Create the claims
	now := time.Now()
	claims := jwt.MapClaims{
		"iss": teamID,
		"iat": now.Unix(),
		"exp": now.Add(180 * 24 * time.Hour).Unix(), // 180 days (maximum allowed)
		"aud": "https://appleid.apple.com",
		"sub": clientID,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = keyID

	// Sign the token
	signedToken, err := token.SignedString(privateKey.(*ecdsa.PrivateKey))
	if err != nil {
		fmt.Printf("Error signing token: %v\n", err)
		return
	}

	fmt.Println("Generated Client Secret:")
	fmt.Println(signedToken)
}