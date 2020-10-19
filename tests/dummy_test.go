package main_test

import "testing"

func TestSum(t *testing.T) {
    thisIsTrue := true
    if thisIsTrue != true {
       t.Errorf("True is not true")
    }
}