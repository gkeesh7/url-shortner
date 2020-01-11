package randomutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode"
)

func TestRandomString(t *testing.T) {
	str := RandomString(7)
	isValid := true
	for _,r := range str {
		if !unicode.IsNumber(r) && !unicode.IsLetter(r) {
			isValid = false
		}
	}
	assert.Equal(t,7,len(str))
	assert.Equal(t,true,isValid)
}