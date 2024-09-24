package pbar

type ProgressBar interface {
	Add(n int)
	Close()
}
