package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetCacheHandle(t *testing.T) {
	var conn = GetCacheConnection()
	resp, err := conn.Do("PING")
	assert.Equal(t, "PONG", resp)
	assert.Equal(t, nil, err)
}
