// MIT License
//
// Copyright (c) 2018 Mark
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package logger

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(new(timestampedLogWriter))
}

type timestampedLogWriter struct {
}

func (timestampedLogWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf("%s %s", formatTime(time.Now().UTC()), string(bytes))
}

func formatTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05.000")
}

type LogWriter struct {
	log.Logger
	name string
}

func NewLogWriter(name string) *LogWriter {
	return &LogWriter{
		log.Logger{},
		name,
	}
}

func (w *LogWriter) Write(b []byte) (int, error) {
	log.Printf(w.format(b))
	return len(b), nil
}

func (w *LogWriter) format(b []byte) string {
	return fmt.Sprintf("%s: %s", w.name, string(b))
}
