package redis

import (
	"bytes"
	"testing"
)

func TestGet(t *testing.T) {
	redPool, _ := NewConnPool(Config{
		Host:                 "127.0.0.1",
		Port:                 "6379",
		Password:             "dev",
		MaxActiveConnections: 256,
	})
	if _, err := redPool.Get().Do("SET", "123", ""); err != nil {
		t.Errorf("Do(SET(): %v", err)
		return
	}

	if _, err := redPool.Get().Do("SET", "12", "-"); err != nil {
		t.Errorf("Do(SET(): %v", err)
		return
	}

	tests := []struct {
		name     string
		key      string
		expected []byte
		wantErr  bool
	}{
		{
			name:     "exist with empty val",
			key:      "123",
			expected: []byte(""),
			wantErr:  false,
		},
		{
			name:     "exists",
			key:      "12",
			expected: []byte("-"),
			wantErr:  false,
		},
		{
			name:     "empty",
			key:      "1234",
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(redPool, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists error %v unexpected", err)
				return
			}
			if err == nil {
				if !bytes.Equal(got.([]byte), tt.expected) {
					t.Errorf("Get() shd exists but didnt")
					return
				}
			}
		})
	}
}
