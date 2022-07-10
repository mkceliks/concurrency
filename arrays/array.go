package main

import (
	"fmt"
	"sort"
	"sync"
)

func sorting(x int, y int, z int, c chan int, wg *sync.WaitGroup) { // This is a sorting function comes from user input

	arr_func := make([]int, 3)
	arr_func[0] = x
	arr_func[1] = y
	arr_func[2] = z
	fmt.Printf("Unsorted Goroutines : [%d %d %d] \n", arr_func[0], arr_func[1], arr_func[2])
	sort.Ints(arr_func)
	c <- arr_func[0]
	c <- arr_func[1]
	c <- arr_func[2]
	fmt.Printf("Sorted Goroutines : [%d %d %d]\n", arr_func[0], arr_func[1], arr_func[2])
	defer wg.Done()
}

func main() {

	arr := make([]int, 12)
	for i := 0; i < 12; i++ { // taking numbers from user
		fmt.Printf("Enter a num : ")
		fmt.Scanln(&arr[i])
		fmt.Printf("\n")
	}
	c := make(chan int)                        // defining a channel 'c'
	var wg sync.WaitGroup                      // defining wg (waitGroup)
	wg.Add(4)                                  // we have 4 goroutines
	go sorting(arr[0], arr[1], arr[2], c, &wg) // taking the outcomes from sorting func with 1 goroutine
	arr[0] = <-c
	arr[1] = <-c
	arr[2] = <-c
	go sorting(arr[3], arr[4], arr[5], c, &wg) // taking the outcomes from sorting func with 1 goroutine
	arr[3] = <-c
	arr[4] = <-c
	arr[5] = <-c
	go sorting(arr[6], arr[7], arr[8], c, &wg) // taking the outcomes from sorting func with 1 goroutine
	arr[6] = <-c
	arr[7] = <-c
	arr[8] = <-c
	go sorting(arr[9], arr[10], arr[11], c, &wg) // taking the outcomes from sorting func with 1 goroutine
	arr[9] = <-c
	arr[10] = <-c
	arr[11] = <-c
	wg.Wait()
	sort.Ints(arr) // sorting the biggest array with 8 incomes

	//printing the sorted array.
	fmt.Printf("SORTED ARRAY[%d %d %d %d %d %d %d %d %d %d %d %d] \n", arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11])
}
