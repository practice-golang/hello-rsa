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

func getPEM() (string, *rsa.PrivateKey, error) {
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
	pub := pemKey.Public()
	if reflect.TypeOf(pub).String() != "*rsa.PublicKey" {
		return "", pub, "", fmt.Errorf("not rsa")
	}

	pubBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return "", pub, "", err
	}

	pubBlock := pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes}
	pubString := string(pem.EncodeToMemory(&pubBlock))

	dkim := "v=DKIM1;k=rsa;p=" + base64.StdEncoding.EncodeToString(pubBytes)

	return pubString, pub, dkim, err
}

func main() {
	privString, priv, err := getPEM()
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
