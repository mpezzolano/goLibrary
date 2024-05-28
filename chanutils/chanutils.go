package chanutils

func CreateChannel(bufferSize int) chan interface{} {
	return make(chan interface{}, bufferSize)
}

func WriteToChannel(ch chan interface{}, data interface{}) {
	ch <- data
}

func ReadFromChannel(ch chan interface{}) interface{} {
	return <-ch
}
