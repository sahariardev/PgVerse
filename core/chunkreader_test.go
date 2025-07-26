package core

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func newChunkReaderFromString(s string, minBuffSize int) *ChunkReader {
	return &ChunkReader{
		reader:     strings.NewReader(s),
		minBufSize: minBuffSize,
		buffer:     make([]byte, minBuffSize),
	}
}

func TestChunkReaderBasicRead(t *testing.T) {
	r := newChunkReaderFromString("abcdefg", 4)

	data, err := r.Next(3)
	if err != nil {
		t.Fatalf("Next returned error: %v", err)
	}

	if string(data) != "abc" {
		t.Fatalf("Next returned %s, expected %s", string(data), "ada")
	}

	data, err = r.Next(4)

	if err != nil {
		t.Fatalf("Next returned error: %v", err)
	}

	if string(data) != "defg" {
		t.Fatalf("Next returned %s, expected %s", string(data), "defg")
	}
}

func TestChunkReaderReadAcrossBufferResize(t *testing.T) {
	r := newChunkReaderFromString("abcdefg", 4)

	data, err := r.Next(7)

	if err != nil {
		t.Fatalf("Next returned error: %v", err)
	}

	if string(data) != "abcdefg" {
		t.Fatalf("Next returned %s, expected %s", string(data), "abcdef")
	}
}

func TestChunkReaderExactBufferSize(t *testing.T) {
	r := newChunkReaderFromString("abcdefg", 7)

	data, err := r.Next(7)

	if err != nil {
		t.Fatalf("Next returned error: %v", err)
	}

	if string(data) != "abcdefg" {
		t.Fatalf("Next returned %s, expected %s", string(data), "abcdef")
	}
}

func TestChunkReaderMultipleSmallRead(t *testing.T) {
	r := newChunkReaderFromString("abcdefg", 2)

	for _, expected := range []string{"a", "b", "c", "d", "e", "f", "g"} {

		data, err := r.Next(1)
		if err != nil {
			t.Fatalf("Next returned error: %v", err)
		}

		if string(data) != expected {
			t.Fatalf("Next returned %s, expected %s", string(data), expected)
		}
	}
}

func TestChunkReaderEndOfFile(t *testing.T) {
	r := newChunkReaderFromString("abc", 2)
	_, err := r.Next(10)
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatalf("Next returned error: %v", err)
	}
}
