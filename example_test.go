package main

import (
	"fmt"
	"testing"
)

func TestExampleSuccess(t *testing.T) {
	result, err := example("hoge")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 1 {
		t.Fatal("failed test")
	}
}

func TestExampleFailed(t *testing.T) {
	result, err := example("fuga")
	if err == nil {
		t.Fatal("failed test")
	}
	if result != 0 {
		t.Fatal("failed test")
	}
}

func TestExample2(t *testing.T) {
	err := example2()
	if err != nil {
		t.Fatal("error1")
	} else {
		fmt.Print("error2")
	}
}
