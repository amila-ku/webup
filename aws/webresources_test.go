package aws

import (
	"testing"
	"context"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestMakeWebResources(t *testing.T){
	err := MakeWebResources(context.TODO(), "abc.com")
	if err != nil {
		fmt.Println("failed")
	}
	assert.Nil(t, err)
}