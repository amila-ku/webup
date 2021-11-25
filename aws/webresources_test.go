package aws

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeWebResources(t *testing.T) {
	err := MakeWebResources(context.TODO(), "abc.com")
	if err != nil {
		fmt.Println("failed")
	}
	assert.Nil(t, err)
}
