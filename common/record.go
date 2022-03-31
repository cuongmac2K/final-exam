package common

import (
	"crypto/rand"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

var TimeFunc = time.Now

type Record struct {
	Otp       string    `json:"otp" gorm:"column:otp;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	Expiry    int       `json:"expiry" gorm:"column:expiry;"`
}

func NewRecord() *Record {
	return &Record{
		Otp:       EncodeToString(6),
		CreatedAt: time.Now(),
		Expiry:    180,
	}
}
func (r Record) GetOtp() string {
	return r.Otp
}
func (r Record) GetCreatedAt() time.Time {
	return r.CreatedAt
}
func (r Record) GetExpiry() int {
	return r.Expiry
}
func (r Record) CheckExpired() bool {
	now := TimeFunc()
	ExpiresAt := r.CreatedAt.Add(time.Second * time.Duration(r.Expiry))
	if now.After(ExpiresAt) {
		return true
	}
	return false
	//delta := time.Unix(now, 0).Sub(time.Unix(ExpiresAt, 0))
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (r *Record) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var record Record
	if err := json.Unmarshal(bytes, &record); err != nil {
		return err
	}

	*r = record
	return nil
}

// Value return json value, implement driver.Valuer interface
func (r *Record) Value() (driver.Value, error) {
	if r == nil {
		return nil, nil
	}
	return json.Marshal(r)
}
