package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	ivLength      = 16
	authTagLength = 16
)

// AESCrypto handles AES-256-GCM encryption/decryption
type AESCrypto struct {
	key []byte
}

// NewAESCrypto creates a new AESCrypto instance with the given Base64-encoded key
func NewAESCrypto(base64Key string) (*AESCrypto, error) {
	// Base64 디코딩
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		// URL-safe Base64 시도
		key, err = base64.URLEncoding.DecodeString(base64Key)
		if err != nil {
			return nil, fmt.Errorf("failed to decode key: %w", err)
		}
	}

	if len(key) != 32 {
		return nil, fmt.Errorf("invalid key length: expected 32 bytes, got %d", len(key))
	}

	return &AESCrypto{key: key}, nil
}

// Decrypt decrypts a Base64-encoded ciphertext using AES-256-GCM
// Returns the decrypted plaintext as raw bytes
func (c *AESCrypto) Decrypt(ciphertext string) ([]byte, error) {
	// Base64 디코딩
	combined, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	// 최소 길이 검증 (IV + AuthTag = 32 bytes)
	if len(combined) < ivLength+authTagLength {
		return nil, fmt.Errorf("ciphertext too short: %d bytes", len(combined))
	}

	// IV, Encrypted Data, Auth Tag 분리
	iv := combined[:ivLength]
	authTag := combined[len(combined)-authTagLength:]
	encrypted := combined[ivLength : len(combined)-authTagLength]

	// AES 블록 생성
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// GCM 모드 생성
	gcm, err := cipher.NewGCMWithNonceSize(block, ivLength)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// Auth Tag를 암호화 데이터 뒤에 붙여서 복호화
	ciphertextWithTag := append(encrypted, authTag...)

	// 복호화
	plaintext, err := gcm.Open(nil, iv, ciphertextWithTag, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}

	return plaintext, nil
}

// DecryptToString decrypts and converts to string
func (c *AESCrypto) DecryptToString(ciphertext string) (string, error) {
	plaintext, err := c.Decrypt(ciphertext)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// DecryptToQoo10Token decrypts and converts to Qoo10JP API token/key format
// Qoo10JP uses certification_key which is typically a plain string
func (c *AESCrypto) DecryptToQoo10Token(ciphertext string) (string, error) {
	plaintext, err := c.Decrypt(ciphertext)
	if err != nil {
		return "", err
	}

	// UTF-8 문자열로 변환
	tokenStr := string(plaintext)

	// 이미 HEX 형식인 경우 그대로 반환
	if isHexString(tokenStr) {
		return tokenStr, nil
	}

	// 바이너리 데이터인 경우 HEX로 변환
	if containsNonPrintable(plaintext) {
		return hex.EncodeToString(plaintext), nil
	}

	return strings.TrimSpace(tokenStr), nil
}

// DecryptToShopeeToken decrypts and converts to Shopee API token format
// - JWT tokens (starting with "eyJ"): returned as-is
// - Binary tokens: converted to HEX string (Node.js 'binary' encoding compatibility)
func (c *AESCrypto) DecryptToShopeeToken(ciphertext string) (string, error) {
	plaintext, err := c.Decrypt(ciphertext)
	if err != nil {
		return "", err
	}

	// UTF-8 문자열로 변환
	tokenStr := string(plaintext)

	// JWT인 경우 그대로 반환
	if strings.HasPrefix(tokenStr, "eyJ") {
		return tokenStr, nil
	}

	// Node.js의 'binary' 인코딩은 각 바이트를 그대로 문자로 취급
	// Go에서는 각 바이트를 HEX로 변환하면 됨
	hexToken := hex.EncodeToString(plaintext)

	return hexToken, nil
}

// isHexString checks if a string is a valid hex string
func isHexString(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// containsNonPrintable checks if bytes contain non-printable characters
func containsNonPrintable(data []byte) bool {
	for _, b := range data {
		if b < 32 || b > 126 {
			return true
		}
	}
	return false
}
