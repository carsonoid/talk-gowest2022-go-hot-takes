package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// START OMIT
func IsEven(i int) bool {
	return i%2 == 0
}

func SendMessage(msg *Message) error {
	if msg == nil {
		return fmt.Errorf("message cannot be empty")
	}

	// ...

	return nil
}

// END OMIT

// START BAD OMIT
func TestIsEven(t *testing.T) {
	assert.True(t, IsEven(2))
	assert.False(t, IsEven(3))
}

func TestSendMessage(t *testing.T) {
	msg, err := getMessage()
	assert.NoError(t, err)

	err = msg.SetHeader("test", "test")
	assert.NoError(t, err)

	assert.AnError(SendMessage(msg))
}

// END BAD OMIT

// START FIXED1 OMIT
func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Error("expected 2 to be even")
	}

	if IsEven(3) {
		t.Error("expected 3 to be odd")
	}
}

// END FIXED1 OMIT

// START FIXED2 OMIT
func TestSendMessage(t *testing.T) {
	msg, err := getMessage()
	if err != nil {
		t.Error(err)
	}

	err = msg.SetHeader("test", "test")
	if err != nil {
		t.Error(err)
	}

	err = SendMessage(msg)
	if err == nil {
		t.Error("expected error:", err)
	}
}

// END FIXED2 OMIT
