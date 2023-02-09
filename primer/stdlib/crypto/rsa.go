package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
)

func main() {
	fmt.Println("GenRsaKey")
	pubKey, prvKey := GenRsaKey()
	fmt.Println(string(pubKey))
	fmt.Println(string(prvKey))

	data := "plaintext"

	fmt.Println("\nencrypt and decrypt")
	ciphertext := RsaEncrypt(pubKey, []byte(data))
	fmt.Println("RsaEncrypt:", hex.EncodeToString(ciphertext))
	plaintext := RsaDecrypt(prvKey, ciphertext)
	fmt.Println("RsaDecrypt:", string(plaintext))

	fmt.Println("\nsign and verify")
	signature := RsaSignWithSha256(prvKey, []byte(data))
	fmt.Println("RsaSignWithSha256: ", hex.EncodeToString(signature))
	fmt.Println("RsaVerifyWithSha256:", RsaVerifyWithSha256(pubKey, []byte(data), signature))
}

// GenRsaKey 生成 RSA 公钥和私钥
func GenRsaKey() (pubKeyBytes, prvKeyBytes []byte) {
	// GenerateKey generates an RSA keypair of the given bit size using the
	// random source random (for example, crypto/rand.Reader).
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	// MarshalPKCS1PrivateKey converts an RSA private key to PKCS #1, ASN.1 DER form.
	derBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derBytes,
	}
	prvKeyBytes = pem.EncodeToMemory(block)

	publicKey := &privateKey.PublicKey
	// MarshalPKIXPublicKey converts a public key to PKIX, ASN.1 DER form.
	derBytes, err = x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	}
	pubKeyBytes = pem.EncodeToMemory(block)

	return
}

// RsaEncrypt 使用 RSA 公钥加密
func RsaEncrypt(pubKeyBytes, plaintext []byte) []byte {
	// Decode will find the next PEM formatted block (certificate, private key
	// etc) in the input. It returns that block and the remainder of the input. If
	// no PEM data is found, p is nil and the whole of the input is returned in
	// rest.
	block, _ := pem.Decode(pubKeyBytes)
	if block == nil {
		panic(errors.New("pem.Decode: block == nil"))
	}
	// ParsePKIXPublicKey parses a public key in PKIX, ASN.1 DER form.
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	pub := pubInterface.(*rsa.PublicKey)
	// EncryptPKCS1v15 encrypts the given message with RSA and the padding
	// scheme from PKCS #1 v1.5.  The message must be no longer than the
	// length of the public modulus minus 11 bytes.
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plaintext)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// RsaDecrypt 使用 RSA 私钥解密
func RsaDecrypt(pubKeyBytes, ciphertext []byte) []byte {
	// Decode will find the next PEM formatted block (certificate, private key
	// etc) in the input. It returns that block and the remainder of the input. If
	// no PEM data is found, p is nil and the whole of the input is returned in
	// rest.
	block, _ := pem.Decode(pubKeyBytes)
	if block == nil {
		panic(errors.New("pem.Decode: block == nil"))
	}
	// ParsePKCS1PrivateKey parses an RSA private key in PKCS #1, ASN.1 DER form.
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS #1 v1.5.
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		panic(err)
	}
	return plaintext
}

// RsaSignWithSha256 使用 SHA-256 签名
func RsaSignWithSha256(prvKeyBytes, data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	hashed := hash.Sum(nil)

	block, _ := pem.Decode(prvKeyBytes)
	if block == nil {
		panic(errors.New("pem.Decode: block == nil"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Errorf("x509.ParsePKCS1PrivateKey: %w", err))
	}

	// SignPKCS1v15 calculates the signature of hashed using
	// RSASSA-PKCS1-V1_5-SIGN from RSA PKCS #1 v1.5.  Note that hashed must
	// be the result of hashing the input message using the given hash
	// function. If hash is zero, hashed is signed directly. This isn't
	// advisable except for interoperability.
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		panic(fmt.Errorf("rsa.SignPKCS1v15: %w", err))
	}
	return signature
}

// RsaVerifyWithSha256 使用 SHA-256 验证签名
func RsaVerifyWithSha256(pubKeyBytes, data, signature []byte) bool {
	block, _ := pem.Decode(pubKeyBytes)
	if block == nil {
		panic(errors.New("pem.Decode: block == nil"))
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(fmt.Errorf("x509.ParsePKIXPublicKey: %w", err))
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signature)
	if err != nil {
		panic(fmt.Errorf("rsa.VerifyPKCS1v15: %w", err))
	}
	return true
}
