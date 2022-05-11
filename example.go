package main

import (
	"errors"
)

func example(code string) (int, error) {
	if code == "hoge" {
		return 1, nil
	}
	return 0, errors.New("code must be hoge")
}

func example2() error {
	n := 1
	if n == 1 {
		return errors.New("error")
	}
	return nil
}
