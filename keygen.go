package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func CreateEncryptionKeys() {
	CreateRSAKey()
	CreateTLSCerts()

}

func CreateTLSCerts() {
	random := rand.Reader

	var key rsa.PrivateKey
	loadGobKey("private.key", &key)

	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000)
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "jan.newmarch.name",
			Organization: []string{"Jan Newmarch"},
		},
		NotBefore: now,
		NotAfter:  then,

		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,

		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{"jan.newmarch.name", "localhost"},
	}
	theBytes, err := x509.CreateCertificate(random, &template,
		&template, &key.PublicKey, &key)
	checkError(err)

	certPEMFile, err := os.Create("public.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: theBytes})
	certPEMFile.Close()

	keyPEMFile, err := os.Create("private.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&key)})
	keyPEMFile.Close()
}

func CreateRSAKey() {
	reader := rand.Reader
	bitSize := 4092
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)
	log.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	log.Println("Private key exponent", key.D.String())

	publicKey := key.PublicKey
	log.Println("Public key modulus", publicKey.N.String())
	log.Println("Public key exponent", publicKey.E)

	saveGobKey("private.key", key)
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func loadGobKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}
