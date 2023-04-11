package iutils

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"golang.org/x/crypto/ssh"
)

var LogError func(err error) error = func(err error) error {
	if err != nil {
		Log(err)
	}

	return err
}

/*
	C-TODO
	4. PKI - 数字签名、CA、数字证书
	https://blog.csdn.net/ttyy1112/article/details/107083345
*/

/*
	ed25519
*/
func GenerateKeysByEd25519() (sshPublicKeyStr, privateKeyStr string, err error) {
	var ()

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		err = LogError(fmt.Errorf("E GenerateKeysByEd25519: %w", err))
		return
	}

	sshPublicKey, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		err = LogError(fmt.Errorf("E GenerateKeysByEd25519: %w", err))
		return
	}

	sshPublicKeyStr = string(ssh.MarshalAuthorizedKey(sshPublicKey))

	privateKeyDer, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		err = LogError(fmt.Errorf("E GenerateKeysByEd25519: %w", err))
		return
	}

	// !!!! PEM INSTEAD OF OPENSSH FORMAT
	privateKeyStr = string(
		pem.EncodeToMemory(
			&pem.Block{Bytes: privateKeyDer, Type: "PRIVATE KEY"},
		),
	)

	Log(
		publicKey,
		privateKey,
		sshPublicKey,
		sshPublicKeyStr,
		privateKeyDer,
		privateKeyStr,
	)

	return

}

/*
	rsa
*/
func EncodeRsaPublicKey(key *rsa.PublicKey) ([]byte, error) {
	bs, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(
		&pem.Block{
			Bytes: bs,
			Type:  "PUBLIC KEY",
		},
	), nil
}

func EncodeRsaPrivateKey(key *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(
		&pem.Block{
			Bytes: x509.MarshalPKCS1PrivateKey(key),
			Type:  "RSA PRIVATE KEY",
		},
	)
}

func GenerateKeysByRsa(bits int) (publicKeyStr, privateKeyStr string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		LogError(err)
		return
	}

	sshPublicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		LogError(err)
		return
	}

	publicKeyStr = string(
		ssh.MarshalAuthorizedKey(
			sshPublicKey,
		),
	)

	privateKeyStr = string(
		EncodeRsaPrivateKey(
			privateKey,
		),
	)

	return
}

/*
	ssh

	https://github.com/golang/crypto/blob/master/ssh/example_test.go
*/
func NewSshClient(user, host string, port int) (client *ssh.Client, err error) {
	var (
		addr         string
		clientConfig *ssh.ClientConfig
	)

	homePath, err := os.UserHomeDir()
	if err != nil {
		return
	}

	privateKeyBs, err := ioutil.ReadFile(path.Join(homePath, ".ssh", "id_rsa"))
	if err != nil {
		return
	}

	signer, err := ssh.ParsePrivateKey(privateKeyBs)
	if err != nil {
		return
	}

	addr = fmt.Sprintf("%s:%d", host, port)

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return
	}

	return
}
