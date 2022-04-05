package fibonacciTestify

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestFibonacciFinder_GetNumber(t *testing.T) {
    ff := New(10)

    assert.Equalf(t, uint(55), ff.GetNumber(10), "10th number is 55. Got %d", ff.GetNumber(10))
    assert.Equalf(t, uint(0), ff.GetNumber(0), "0th number is 1. Got %d", ff.GetNumber(0))
    assert.Equalf(t, uint(1), ff.GetNumber(1), "1st number is 1. Got %d", ff.GetNumber(1))
    assert.Equalf(t, uint(1), ff.GetNumber(2), "2nd number is 1. Got %d", ff.GetNumber(2))
    assert.NotEqual(t, uint(1), ff.GetNumber(3), "3rd number is not 1")
}
