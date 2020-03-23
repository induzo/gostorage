package mongo

import (
	"testing"
	"time"
)

func TestNewCli(t *testing.T) {
	tests := []struct {
		name      string
		mongoConf Config
		wantErr   bool
	}{
		{
			name: "Working connection",
			mongoConf: Config{
				Host:           "127.0.0.1",
				Port:           "27017",
				User:           "internal",
				Password:       "dev",
				ConnectTimeout: 100 * time.Millisecond,
			},
			wantErr: false,
		},
		{
			name: "Non existing db",
			mongoConf: Config{
				Host:     "none",
				Port:     "3306",
				User:     "internal",
				Password: "badtest",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCli(tt.mongoConf)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"NewCli() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
		})
	}
}
