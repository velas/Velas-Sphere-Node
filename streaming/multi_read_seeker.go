package main

import (
	"fmt"
	"io"
)

type multiReadSeeker struct {
	readSeekers  []io.ReadSeeker
	readSeeker   int
	globalOffset int64
	localOffset  int64
	whence       int
}

func (mr *multiReadSeeker) Read(p []byte) (n int, err error) {
	if len(mr.readSeekers) < 1 {
		return 0, io.EOF
	}

	n, err = mr.readSeekers[mr.readSeeker].Read(p)
	if n > 0 || err != io.EOF {
		if err == io.EOF && len(mr.readSeekers) > mr.readSeeker {
			// Don't return EOF yet. More readers remain.
			mr.readSeeker++
			err = nil
		}
		return
	}
	return 0, io.EOF
}

// TODO: implement
func (mr *multiReadSeeker) Seek(offset int64, whence int) (int64, error) {
	fmt.Println("whence", whence)

	mr.globalOffset += offset
	mr.whence = whence

	if mr.whence == 0 {
		mr.readSeeker = int(mr.globalOffset / (10 << 20))
		mr.localOffset = mr.globalOffset % (10 << 20)
	}

	if mr.whence == 2 {
		mr.readSeeker = len(mr.readSeekers) - int(mr.globalOffset/(10<<20)) - 1
		mr.localOffset = int64(len(mr.readSeekers)) - mr.globalOffset%(10<<20) - 1
	}

	seek, err := mr.readSeekers[mr.readSeeker].Seek(mr.localOffset, mr.whence)
	if err != nil {
		fmt.Println("err", err)
	}

	mr.whence = 0
	mr.globalOffset = seek + int64(mr.readSeeker*(10<<20))

	mr.readSeeker = int(mr.globalOffset / (10 << 20))
	mr.localOffset = mr.globalOffset % (10 << 20)

	return mr.globalOffset, err
}

func MultiReadSeeker(readers ...io.ReadSeeker) io.ReadSeeker {
	r := make([]io.ReadSeeker, len(readers))
	copy(r, readers)
	return &multiReadSeeker{r, 0, 0, 0, 0}
}
