//go:build windows

package machineid

import (
	"crypto/rand"
	"fmt"
)

// machineID returns a random generated GUID instead of reading from registry.
// This will generate a new random ID every time it's called.
func machineID() (string, error) {
	// Tạo một mảng 16 byte ngẫu nhiên để tạo GUID
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// Định dạng đặc biệt cho GUID theo chuẩn RFC4122
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant bits
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4

	// Chuyển đổi thành chuỗi định dạng GUID
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16]), nil
}
