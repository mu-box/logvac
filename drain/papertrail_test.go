package drain_test

import (
	"bytes"
	"regexp"
	"testing"

	logvac "github.com/mu-box/logvac/core"
	"github.com/mu-box/logvac/drain"
)

// Test adding drain.
func TestPTrailInit(t *testing.T) {
	trailTest := &drain.Papertrail{}
	trailTest.Close()
	trailTest.Init()
}

// Test writing and reading data, as well as closing.
func TestPTrailPublish(t *testing.T) {
	var b WriteCloseBuffer
	trailTest := &drain.Papertrail{Conn: &b}
	if trailTest.Conn == nil {
		t.Fatal("Failed to create a thing")
	}

	msg := logvac.Message{Content: "This is a message\n"}

	trailTest.Publish(msg)
	if match, err := regexp.Match(msg.Content, b.Bytes()); !match || err != nil {
		t.Fatalf("Failed to publish (%s) - Got '%s'; Expected '%s'", err.Error(), b.String(), msg.Content)
	}
	trailTest.Close()
}

type WriteCloseBuffer struct {
	bytes.Buffer
}

func (cb WriteCloseBuffer) Close() error {
	return nil
}
