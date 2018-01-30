package server

import "fmt"

//IncrementCounter is a shit
func (w *Worker) IncrementCounter() chan bool {
	w.counter++
	fmt.Println("Counter's value is %d", w.counter)
	w.counterChan <- true
	fmt.Println("Channel sent...")
	return w.counterChan
}
