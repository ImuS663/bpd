package pbar

import (
	"github.com/pterm/pterm"
)

type PTermProgressBar struct {
	pb *pterm.ProgressbarPrinter
}

func NewPTermProgressBar(name string, total int64) ProgressBar {
	pb, _ := pterm.DefaultProgressbar.WithTotal(int(total)).WithTitle(name).WithShowCount(false).Start()

	return &PTermProgressBar{pb: pb}
}

func (pb *PTermProgressBar) Add(n int) {
	pb.pb.Add(int(n))
}

func (pb *PTermProgressBar) Close() {
	pb.pb.WithCurrent(0)
	pb.pb.Stop()
}
