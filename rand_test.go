package roby

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandQueue(t *testing.T) {
	q := NewRandQueue(10, 5)
	assert.Equal(t, 5, len(q), "队列长度")

}
