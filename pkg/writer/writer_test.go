package writer

import (
	"testing"
)

type mockProgressWriter struct {
	total int
}

func (pb *mockProgressWriter) Add(n int) {
	pb.total += n
}

func (pb *mockProgressWriter) Close() {
	pb.total = 0
}

type mockWriter struct {
	total int
}

func (mw *mockWriter) Write(p []byte) (n int, err error) {
	mw.total += len(p)
	return len(p), nil
}

func (mw *mockWriter) Close() error {
	mw.total = 0
	return nil
}

func TestNewProgressWriter(t *testing.T) {
	pb := mockProgressWriter{}
	mw := mockWriter{}
	writer := NewProgressWriter(&mw, &pb)

	if writer == nil {
		t.Error("writer is nil")
	}
}

func TestProgressWriterWrite(t *testing.T) {
	pb := mockProgressWriter{}
	mw := mockWriter{}
	writer := NewProgressWriter(&mw, &pb)
	writer.Write([]byte("test"))

	if mw.total != 4 {
		t.Error("writer total expected 4, got", pb.total)
	}

	if pb.total != 4 {
		t.Error("progress bar expected 4, got", pb.total)
	}

}

func TestProgressWriterClose(t *testing.T) {
	pb := mockProgressWriter{}
	mw := mockWriter{}
	writer := NewProgressWriter(&mw, &pb)
	writer.Close()

	if mw.total != 0 {
		t.Error("writer total expected 0, got", pb.total)
	}

	if pb.total != 0 {
		t.Error("progress bar expected 0, got", pb.total)
	}
}
