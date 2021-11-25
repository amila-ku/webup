package aws

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeBucket(t *testing.T) {
	err := MakeBucket(context.TODO(), "abc.com")
	if err != nil {
		fmt.Println("failed")
	}
	assert.Nil(t, err)
}
