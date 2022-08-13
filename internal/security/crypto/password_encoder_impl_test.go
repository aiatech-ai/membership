package crypto

import (
	"testing"

	"github.com/google/uuid"
)

func Test_passwordEncoder_Encode(t *testing.T) {

	tests := []struct {
		name        string
		rawPassword string
		numOfCheck  int
		wantErr     bool
	}{
		{name: "not same", rawPassword: "123456", numOfCheck: 5, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.numOfCheck; i++ {

				p := &passwordEncoder{}
				got1, err := p.Encode(tt.rawPassword)
				if (err != nil) != tt.wantErr {
					t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				got2, err := p.Encode(tt.rawPassword)
				if (err != nil) != tt.wantErr {
					t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got1 == got2 {
					t.Errorf("Encode() got1 = got2 = %v", got1)
				}
			}
		})
	}
}

func Test_passwordEncoder_Matches(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		want    bool
		wantLen int
	}{
		{name: "test matches", wantErr: false, want: true, wantLen: 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &passwordEncoder{}
			password := uuid.New().String()
			encodedPassword, err := p.Encode(password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantLen != len(encodedPassword) {
				t.Errorf("Encode() len = %v, wantLen %v", len(encodedPassword), tt.wantLen)
			}

			if got := p.Matches(password, encodedPassword); got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}
