package quark

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuark(t *testing.T) {

	c := New()

	res, err := c.Direct([]string{"c07bba0d73fe40e09c2c0e6ece1607cd"})
	assert.NoError(t, err)

	fmt.Printf("res: %v\n", res)

}
