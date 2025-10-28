package common

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql/driver"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
)

type Phone string

var encryptionKey = []byte("passphrasewhichneedstobe32bytes!")

// Scan implements the sql.Scanner interface
func (p *Phone) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var encrypted string
	switch v := value.(type) {
	case string:
		encrypted = v
	case []byte:
		encrypted = string(v)
	default:
		return errors.New("invalid type for MobileNumber")
	}

	// 解密逻辑
	decrypted, err := decrypt(encrypted)
	if err != nil {
		return err
	}

	*p = Phone(decrypted)
	return nil
}

// Value implements the driver.Valuer interface
func (p Phone) Value() (driver.Value, error) {
	// 加密逻辑
	return encrypt(string(p))
}

// 加密函数
func encrypt(plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}
	block, err := aes.NewCipher(encryptionKey)
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

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// 解密函数
func decrypt(ciphertext string) (string, error) {
	if len(ciphertext) == 0 {
		return "", nil
	}
	decoded, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	if len(decoded) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := decoded[:aes.BlockSize]
	decrypted := decoded[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, decrypted)

	return string(decrypted), nil
}

type Param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Items []string

func (t *Items) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Items) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type Swipers []string
type Videos []string
type Params []Param
type Keyword []string
type ImageUrl []string
type List []int

func (t *List) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t List) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *ImageUrl) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t ImageUrl) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *Swipers) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Swipers) Value() (driver.Value, error) {
	return json.Marshal(t)
}
func (t *Params) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Params) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *Videos) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Videos) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *Keyword) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Keyword) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type IdentityJson struct {
	Data []struct {
		Label string
		Name  string
		Value string
	} `json:"data"`
}

func (t *IdentityJson) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t IdentityJson) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type CoordRes struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func (c CoordRes) GormDataType() string {
	return "geometry"
}

// GormValue - 修复坐标顺序问题
func (c CoordRes) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if len(c.Coordinates) == 0 {
		return clause.Expr{SQL: "NULL"}
	}

	// 正确顺序：经度在前(X)，纬度在后(Y)
	return clause.Expr{
		SQL:  "ST_GeomFromText(?, 4326)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", c.Coordinates[1], c.Coordinates[0])},
	}
}

// Scan - 完全重写，使用安全可靠的WKB解析
func (c *CoordRes) Scan(value interface{}) error {
	c.Type = "Point"
	c.Coordinates = nil

	if value == nil {
		return nil
	}

	wkb, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid type: %T, expected []byte", value)
	}

	// 回退到SQL函数解析（兼容性）
	return parseViaSQL(c, wkb)
}

// 通过SQL函数解析
func parseViaSQL(c *CoordRes, wkb []byte) error {
	if len(wkb) == 0 {
		return fmt.Errorf("wkb is empty")
	}
	hexStr := hex.EncodeToString(wkb)
	query := fmt.Sprintf("SELECT ST_AsGeoJSON(0x%v) as coord", hexStr)

	var jsonStr string
	if err := facades.Orm().Query().Raw(query).Pluck("coord", &jsonStr); err != nil {
		return fmt.Errorf("SQL parse failed: %w", err)
	}

	if jsonStr == "" {
		return nil
	}

	// 临时结构体用于解析
	var tmp struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	if err := json.Unmarshal([]byte(jsonStr), &tmp); err != nil {
		return fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	c.Type = tmp.Type
	c.Coordinates = tmp.Coordinates
	return nil
}
