package main

// Worker pool pattern
// A worker pool is a collection of threads that are waiting for tasks to be allocated for concurrent execution. The worker pool pattern is used to limit the number of threads that are active at any given time, which can help to improve performance and reduce resource usage.
// In this example, we will create a worker pool that calculates the Fibonacci numbers for a given range of integers. We will use channels to communicate between the main goroutine and the worker goroutines.
// The main goroutine will send tasks to the worker goroutines through a channel, and the worker goroutines will send the results back to the main goroutine through another channel.
// We will use a buffered channel to limit the number of tasks that can be sent to the worker goroutines at any given time. This will help to prevent the worker goroutines from being overwhelmed with too many tasks and will allow them to process the tasks more efficiently.
func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := range 45 {
		tasks <- i
	}

	close(tasks)

	for i := range 45 {
		println(i, <-results)
	}

}

func worker(tasks <-chan int, results chan<- int) {
	for n := range tasks {
		results <- fibo(n)
	}
}

func fibo(pos int) int {
	if pos <= 1 {
		return pos
	}

	return fibo(pos-2) + fibo(pos-1)
}
