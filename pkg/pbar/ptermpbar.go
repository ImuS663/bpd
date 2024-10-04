package pbar

import (
	"github.com/pterm/pterm"
)

type PTermProgressBar struct {
	pb *pterm.ProgressbarPrinter
}

// NewPTermProgressBar creates a new pterm progress bar.
//
// name is the title of the progress bar.
//
// total is the total number of items to be processed.
//
// It returns a pointer to the new PTermProgressBar.
func NewPTermProgressBar(name string, total int64) *PTermProgressBar {
	pb, _ := pterm.DefaultProgressbar.WithTotal(int(total)).WithTitle(name).WithShowCount(false).WithMaxWidth(0).WithRemoveWhenDone().Start()

	return &PTermProgressBar{pb: pb}
}

// Add adds n to the progress bar's current value.
//
// n is the amount to add to the progress bar's current value.
func (pb *PTermProgressBar) Add(n int) {
	pb.pb.Add(int(n))
}

// Close stops the progress bar and resets its current value to 0.
func (pb *PTermProgressBar) Close() {
	pb.pb.WithCurrent(0)
	pb.pb.Stop()
}
