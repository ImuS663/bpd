package pbar

import "testing"

func TestNewPTermProgressBar(t *testing.T) {
	pbar := NewPTermProgressBar("test", 10)

	if pbar == nil {
		t.Error("pbar is nil")
	}
}

func TestPTermProgressBarAdd(t *testing.T) {
	pbar := NewPTermProgressBar("test", 10)

	pbar.Add(5)

	if pbar.pb.Current != 5 {
		t.Error("expected 5, got", pbar.pb.Current)
	}
}

func TestPTermProgressBarClose(t *testing.T) {
	pbar := NewPTermProgressBar("test", 10)
	pbar.Close()

	if pbar.pb.Current != 0 {
		t.Error("expected 0, got", pbar.pb.Current)
	}
}
