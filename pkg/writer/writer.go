package writer

import (
	"io"

	"github.com/ImuS663/bpd/pkg/pbar"
)

type ProgressWriter struct {
	io.WriteCloser
	pb pbar.ProgressBar
}

func NewProgressWriter(writer io.WriteCloser, pb pbar.ProgressBar) *ProgressWriter {
	return &ProgressWriter{WriteCloser: writer, pb: pb}
}

func (w *ProgressWriter) Write(p []byte) (n int, err error) {
	if w.pb == nil {
		w.pb.Add(len(p))
	}

	return w.WriteCloser.Write(p)
}

func (w *ProgressWriter) Close() error {
	if w.pb != nil {
		w.pb.Close()
	}

	return w.WriteCloser.Close()
}
