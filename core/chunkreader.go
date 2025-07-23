package core

import "io"

type ChunkReader struct {
	reader       io.Reader
	buffer       []byte
	minBufSize   int
	readPointer  int
	writePointer int
}

func (r *ChunkReader) Nex(n int) ([]byte, error) {
	if r.readPointer+n <= r.writePointer {
		buf := r.buffer[r.readPointer : r.readPointer+n]
		r.readPointer += n
		return buf, nil
	}

	if len(r.buffer) < n {
		r.copyBufferContent(r.newBuffer(n))
	}

	minDataNeedToPullFromIO := n - (r.writePointer - r.readPointer)

	if len(r.buffer)-r.writePointer < minDataNeedToPullFromIO {
		r.copyBufferContent(r.newBuffer(minDataNeedToPullFromIO))
	}

	if err := r.readAtLEast(minDataNeedToPullFromIO); err != nil {
		return nil, err
	}

	buf := r.buffer[r.readPointer : r.readPointer+n]
	r.readPointer += n

	return buf, nil
}

func (r *ChunkReader) readAtLEast(min int) error {
	n, err := io.ReadAtLeast(r.reader, r.buffer[r.writePointer:], min)
	r.writePointer += n
	return err
}

func (r *ChunkReader) newBuffer(length int) []byte {
	if length < r.minBufSize {
		length = r.minBufSize
	}

	return make([]byte, length)
}

func (r *ChunkReader) copyBufferContent(dest []byte) {
	r.writePointer = copy(dest, r.buffer[r.readPointer:r.writePointer])
	r.readPointer = 0
	r.buffer = dest
}
