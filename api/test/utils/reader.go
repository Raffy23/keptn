package utils

import (
	"errors"
	"io"
)

type TestReader struct {
	data         []byte
	allowedLoops int
	loopCount    int
	pos          int
	throwError   bool
}

func NewTestReader(data []byte, allowedLoops int, throwError bool) *TestReader {
	return &TestReader{
		data:         data,
		allowedLoops: allowedLoops,
		throwError:   throwError,
	}
}

func (tr *TestReader) Read(p []byte) (n int, err error) {
	// check if we can reset
	if tr.pos >= len(tr.data) && tr.loopCount < tr.allowedLoops {
		tr.allowedLoops++
		tr.pos = 0
	}

	if tr.pos >= len(tr.data) {
		err := io.EOF

		if tr.throwError {
			err = errors.New("TestReader error")
		}

		// end of buffer
		return 0, err
	}

	// try to copy as much as possible in the buffer
	copied := copy(p, tr.data[tr.pos:])
	tr.pos += copied
	return copied, nil
}
