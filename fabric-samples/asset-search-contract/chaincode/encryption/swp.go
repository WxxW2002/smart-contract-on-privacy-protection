package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"asset-search-contract/chaincode/models"
)

func EncryptAsset(asset *models.Asset) (*models.AssetPrivateDetails, error) {
	encryptedColor, err := encrypt(asset.Color)
	if err != nil {
		return nil, err
	}

	encryptedOwner, err := encrypt(asset.Owner)
	if err != nil {
		return nil, err
	}

	return &models.AssetPrivateDetails{
		ID:             asset.ID,
		EncryptedColor: encryptedColor,
		EncryptedOwner: encryptedOwner,
	}, nil
}

func encrypt(plaintext string) (string, error) {
	key := []byte("your_secret_key_here") // 请替换为你自己的密钥

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func BuildIndex(plaintext string) string {
	// 这里仅返回明文作为简单示例
	// TODO: 实现SWP索引生成逻辑
	return plaintext
}
