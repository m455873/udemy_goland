package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"
import "container/heap"


/*** This implementation of priority queue using heaps is from an example from the golang heap doc page, here:
     https://golang.org/pkg/container/heap/

***/

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
// End priority queue code



// Implement dijkstra alg, modified to find longest path rather than shortest path
func main() {
	// Read in the file, make table

	file, err := os.Open("pyramid.txt")
	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) //scanlines it the function to be used to split the file
	var results [][]int // the 2d slice for storing all rows
	
	for scanner.Scan() {
		wordlist := strings.Fields(scanner.Text())
		vals := []int{} // one row to be appended on end of results
		
		for _, v:= range wordlist {
			x, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error reading file!")
				fmt.Println(err)
			} else {
				vals = append(vals, x)
			}
		}
		results = append(results, vals)

	}
	for _, w := range results {
		fmt.Println(w)
	}
	
	// Do some traversal (nearest neighbors alg)
	var pathtotal int
	var rowno int
	var lastindex int

	for _,x := range results {
		if len(x) == 1 {
			fmt.Println(x[0])
			lastindex = 0
			pathtotal = x[0]
			G[
		} else if len(x) < 1 {
			fmt.Println("Error -- empty row")
		} else {
			var offset int
			if x[lastindex] < x[lastindex+1] {
				offset = 1
			}

			pathtotal += x[lastindex + offset]
			fmt.Println(x[lastindex+offset])
			lastindex += offset
		}
			
		rowno++
	}
	fmt.Printf ("All done,total:%d\n",pathtotal)

	// Try dijkstra
		// Dijkstra doesn't work without a priority queue, which we will have to implement (or borrow) to get it to work
		// So, we need to identify nodes such that we can stuff into priority queue and
		// when they come out we know where they came from (row/col).
		// Right now, order is implicit, we need explicit order. Because there are dups
		// numbers are unique.  We need to make each entry unique

		// option 1: hash, unique name like: "3_4": 45, where 3 is row, 4 is col, 45 is val
		// option 2:everything in hash, 3_4_45
		// optino 2: everything in slice, as source... but string 3_4_45
		// downside with string is unpacking data... probably just split on underscore
		// or use regex... dunno
		

	var D map[string]int
	var P map[string]int
	var Q map[string]int
	var G map[string]int
	D = make(map[string]int)
	P = make(map[string]int)
	Q = make(map[string]int)
	G = make(map[string]int)

	for _,y := range G {
		
			
	
	
	
	
	
	
}
