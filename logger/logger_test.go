package logger

import (
	"testing"
	"time"
)

func TestNewLogWriter(t *testing.T) {
	name := "container-x"
	logWriter := NewLogWriter(name)

	if name != logWriter.name {
		t.Errorf("Expected %s, but got %s", name, logWriter.name)
	}
}

func TestFormat(t *testing.T) {
	location, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		t.Fatal("could not load location 'Europe/Amsterdam'")
	}

	timestaps := []time.Time{
		time.Now().UTC(),
		time.Date(2019, time.January, 14, 3, 1, 1, 1, location),
	}
	for _, timestamp := range timestaps {
		formatted := formatTime(timestamp)

		if len(formatted) != 23 {
			t.Errorf("formatted date (%s) should be a fixed length 23 but was %d", formatted, len(formatted))
		}
	}
}

func TestLogWriter_format(t *testing.T) {
	logWriter := NewLogWriter("container-x")
	formattedLine := logWriter.format([]byte("test line"))

	if "container-x: test line" != formattedLine {
		t.Errorf("expected %s, but got %s", "container-x: test line", formattedLine)
	}
}
