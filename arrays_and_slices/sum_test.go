package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1,2,3}

        got := Sum(numbers)
        want := 6
        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })
}

func TestSumAll(t *testing.T) {
    got := SumAll([]int{1, 2}, []int{0, 9})
    want := []int{3, 9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func TestSumAllTails(t *testing.T) {

    checksum := func(t testing.TB, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }

    t.Run("makes the sums of some slides", func (t *testing.T)  {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}
        checksum(t, got, want)
    })

    t.Run("safely sun empty slices", func (t *testing.T)  {
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}
        checksum(t, got, want)
    })
}
