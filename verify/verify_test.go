package verify

import (
	"github.com/xiaojun207/go-base-utils/assert"
	"testing"
)

func TestEmailFormat(t *testing.T) {
	ok := EmailFormat("abc@abc.com")
	assert.True(t, ok)
}

func TestMobileFormat(t *testing.T) {
	ok := MobileFormat("13412341234")
	assert.True(t, ok)
}

func TestPasswordFormat(t *testing.T) {
	ok := PasswordFormat("Abc123456")
	assert.True(t, ok)
}
