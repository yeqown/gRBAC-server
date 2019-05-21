package secret

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"time"
)

const (
	secretPath     = "secrets"
	secretFileName = "initialAdminPassword"
)

var (
	secretToken string
	token       string
)

// GetToken ...
// func GetToken() string {
// 	if token == "" {
// 		if err := Load
// 		token = generateSecret()
// 	}
// 	return token
// }

func generateSecret() string {
	byts := []byte(time.Now().String())
	m5 := md5.New()
	m5.Write(byts)
	return hex.EncodeToString(m5.Sum(nil))
}

// GetSecret ...
func GetSecret() string {
	if secretToken == "" {
		if err := loadSecretFile(); err != nil {
			panic(err)
		}
	}
	return secretToken
}

// InitSecretFile ... to generate
func InitSecretFile() {
	// write to file
	filename := path.Join(secretPath, secretFileName)
	if fd, err := os.Open(filename); os.IsExist(err) || fd != nil {
		println("cancel generate secret file, it exists!")
		return
	}

	// mkdir
	os.Mkdir(secretPath, 0755)
	fd, err := os.Create(filename)
	if err != nil {
		println(err.Error())
		return
	}
	defer fd.Close()

	hash := generateSecret()
	_, err = fd.Write([]byte(hash))
	if err != nil {
		println(err.Error())
		return
	}
}

// loadSecretFile ...
func loadSecretFile() error {
	filename := path.Join(secretPath, secretFileName)
	fd, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer fd.Close()
	secretToken, _ = bufio.NewReader(fd).ReadString(0)
	return nil
}
