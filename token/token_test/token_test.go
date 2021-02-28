package token_test

import (
	"fmt"
	"home/token"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	err, generatedToken := token.GenerateToken("aqzhol@gmail.com")
	assert.NoError(t, err)

	os.Setenv("testToken", generatedToken)
}

func TestValidateToken(t *testing.T) {
	encodedToken := os.Getenv("testToken")
	fmt.Println(encodedToken)

	err, claims := token.ValidateToken(encodedToken)
	assert.NoError(t, err)

	assert.Equal(t, "aqzhol@gmail.com", claims.Username)
}
