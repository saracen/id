package id

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

var ErrEmptyID = errors.New("empty id")

func ID() (string, error) {
	id, err := id()
	if err != nil {
		return "", err
	}
	if id == "" {
		return "", ErrEmptyID
	}

	return id, nil
}

func read(pathname string) func() (string, error) {
	return func() (string, error) {
		buf, err := os.ReadFile(pathname)

		if err != nil {
			return "", err
		}

		id := string(bytes.TrimSpace(buf))
		if len(id) == 0 {
			return "", ErrEmptyID
		}

		return id, nil
	}
}

func run(name string, arg ...string) func() (string, error) {
	return func() (string, error) {
		cmd := exec.Command(name, arg...)
		buf, err := cmd.Output()

		if err != nil {
			return "", err
		}

		id := string(bytes.TrimSpace(buf))
		if len(id) == 0 {
			return "", ErrEmptyID
		}

		return id, nil
	}
}

func fallback(fn ...func() (string, error)) (string, error) {
	var id string
	var err error
	for _, fn := range fn {
		id, err = fn()
		if err == nil {
			break
		}
	}

	return id, err
}
