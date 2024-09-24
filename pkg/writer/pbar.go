package writer

type ProgressBar interface {
	Add(n int)
	Close()
}
