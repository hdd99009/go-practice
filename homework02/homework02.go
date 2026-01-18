package homework02

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func AddTen(x *int) {
	*x += 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func MultTwo(x *[]int) {
	for i := range *x {
		(*x)[i] = (*x)[i] * 2
	}
}

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func PrintNums() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			if i%2 == 1 {
				fmt.Println(i)
				time.Sleep(time.Second * 3)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			if i%2 == 0 {
				fmt.Println(i)
				time.Sleep(time.Second * 2)

			}
		}
	}()
	wg.Wait()
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

type TaskResult struct {
	ID   int
	Cost time.Duration
}

func Worker(wg *sync.WaitGroup, id int, task func(), results chan<- TaskResult) {
	defer wg.Done()
	start := time.Now()
	task()
	cost := time.Since(start)
	results <- TaskResult{
		id,
		cost,
	}
}

//对应的main函数
/*
package main

import (
	"fmt"
	"practice/homework02"
	"sync"
	"time"
)

func main() {
	// 模拟一组耗时不同的任务
	tasks := []func(){
		func() { time.Sleep(1 * time.Second); fmt.Println("任务A done") },
		func() { time.Sleep(2 * time.Second); fmt.Println("任务B done") },
		func() { time.Sleep(500 * time.Millisecond); fmt.Println("任务C done") },
	}
	results := make(chan homework02.TaskResult, len(tasks))
	var wg sync.WaitGroup
	wg.Add(3)
	for id, task := range tasks {
		go homework02.Worker(&wg, id, task, results)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	for res := range results {
		fmt.Printf("任务ID: %d, 耗时: %v\n", res.ID, res.Cost)
	}
}
*/

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return r.length*2 + r.width*2
}

func (r Circle) Area() float64 {
	return math.Pi * r.radius * r.radius
}

func (r Circle) Perimeter() float64 {
	return math.Pi * r.radius * 2
}

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个
// Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体
// 实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (s Employee) PrintInfo() {
	fmt.Printf("id,%d;Name,%s;Age,%d", s.EmployeeID, s.Person.Name, s.Person.Age)
}

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func MakeNum(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	for i := 1; i < 11; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
}

func PrintNum(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := range ch {
		time.Sleep(time.Second * 2)
		fmt.Println(i)
	}
}

/*
package main

import (
	"practice/homework02"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	wg.Add(1)
	wg2.Add(1)
	go homework02.MakeNum(&wg, ch)
	go homework02.PrintNum(&wg2, ch)
	go func() {
		wg.Wait()
		close(ch)
	}()
	wg2.Wait()
}
*/

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收
// 这些整数并打印。
// 考察点 ：通道的缓冲机制。
func Producer(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func Consumer(num int) {
	fmt.Println(num)
}

/*
package main

import (
	"practice/homework02"
	"sync"
)

func main() {
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	wg.Add(1)
	wg2.Add(1)

	go homework02.Producer(&wg, ch)

	go func() {
		defer wg2.Done()
		for num := range ch {
			homework02.Consumer(num)
		}
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()
	wg2.Wait()
}

*/

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
func AddOnethousand(count *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		*count++
	}
}

/*
package main

import (
	"fmt"
	"practice/homework02"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	count := 1
	for i := 0; i < 10; i++ {

		go homework02.AddOnethousand(&count, &mu, &wg)
	}
	wg.Wait()
	fmt.Println(count)
}
*/

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func AddOnethousand2(wg *sync.WaitGroup, count *int64) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(count, 1)
	}
}

/*
package main

import (
	"fmt"
	"practice/homework02"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	var count int64
	count = 1
	for i := 0; i < 10; i++ {

		go homework02.AddOnethousand2(&wg, &count)
	}
	wg.Wait()
	fmt.Println(count)
}

*/
