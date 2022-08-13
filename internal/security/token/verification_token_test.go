package token

import (
	"strings"
	"testing"
)

func TestGenerateVerificationCode(t *testing.T) {
	verificationCode := GenerateVerificationCode()
	contains := strings.Contains(verificationCode, "-")
	if contains {
		t.Errorf("GenerateVerificationCode() contains illegal character")
	}

	if len(verificationCode) != 32 {
		t.Errorf("GenerateVerificationCode() lenght not matchs, got = %d, want=%d", len(verificationCode), 32)
	}
}
