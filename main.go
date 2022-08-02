package main // import "hello-rsa"

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"reflect"
)

func generatePrivateKEY() (string, *rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return "", key, err
	}

	keyBytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return "", key, err
	}

	keyBlock := pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes}
	keyString := string(pem.EncodeToMemory(&keyBlock))

	return keyString, key, err
}

func getDKIM(pemKey *rsa.PrivateKey) (string, crypto.PublicKey, string, error) {
	key := pemKey.Public()
	if reflect.TypeOf(key).String() != "*rsa.PublicKey" {
		return "", key, "", fmt.Errorf("not rsa")
	}

	keyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", key, "", err
	}

	keyBlock := pem.Block{Type: "PUBLIC KEY", Bytes: keyBytes}
	keyString := string(pem.EncodeToMemory(&keyBlock))

	dkim := "v=DKIM1;k=rsa;p=" + base64.StdEncoding.EncodeToString(keyBytes)

	return keyString, key, dkim, err
}

func main() {
	privString, priv, err := generatePrivateKEY()
	if err != nil {
		panic(err)
	}

	pubString, _, dkim, err := getDKIM(priv)
	if err != nil {
		panic(err)
	}

	fmt.Println(privString)
	fmt.Println(pubString)
	fmt.Println(dkim)
}
