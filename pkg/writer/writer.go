package writer

import (
	"io"

	"github.com/ImuS663/bpd/pkg/pbar"
)

type ProgressWriter struct {
	io.WriteCloser
	pb pbar.ProgressBar
}

// NewProgressWriter takes a io.WriteCloser and a pbar.ProgressBar and returns a new ProgressWriter.
//
// The ProgressWriter is a io.WriteCloser that wraps the given io.WriteCloser and also reports the progress of the write operation to the given pbar.ProgressBar.
//
// The ProgressWriter is useful for reporting the progress of a large write operation to the user.
func NewProgressWriter(writer io.WriteCloser, pb pbar.ProgressBar) *ProgressWriter {
	return &ProgressWriter{WriteCloser: writer, pb: pb}
}

// Write implements the io.Writer interface. It writes the given bytes to the underlying
// io.Writer and updates the given pbar.ProgressBar with the number of bytes written.
func (w *ProgressWriter) Write(p []byte) (n int, err error) {
	if w.pb != nil {
		w.pb.Add(len(p))
	}

	return w.WriteCloser.Write(p)
}

// Close implements the io.Closer interface. It closes the underlying io.WriteCloser and
// also stops the given pbar.ProgressBar.
func (w *ProgressWriter) Close() error {
	if w.pb != nil {
		w.pb.Close()
	}

	return w.WriteCloser.Close()
}
