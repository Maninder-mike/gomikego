package main

import (
    "testing"

    "github.com/stretchr/testify/assert"

)

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
    assert.Equal(t, 123, 123, "they should be equal")

    assert.NotEqual(t, 123, 456, "they should not be equal")

}

