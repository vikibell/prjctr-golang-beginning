package sensor

type Sensor[T any] interface {
	ReadData() T
}
