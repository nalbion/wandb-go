package wandb

type Table struct{}

type Loggable interface {
	Table | ~int | ~float64 | ~float32 | ~string
}

func Log[T Loggable](key string, value T) {
}
