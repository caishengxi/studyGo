package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

func add10(point *int) {
	*point += 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func multiplySlice(slice *[]int) {
	for i := range *slice {
		(*slice)[i] *= 2
	}
}

func testgoroutine1() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		if i%2 == 1 {

			fmt.Printf("奇数: %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func testgoroutine2() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {

			fmt.Printf("偶数: %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func task1() {

	defer wg.Done()
	start := time.Now()
	fmt.Printf("task 1 start: %s\n", start)
	sum := 0
	for i := 1; i <= 500000; i++ {
		sum += i
	}
	elasped := time.Since(start)
	fmt.Printf("Task 1 completed in %s\n", elasped)

}

func task2() {

	defer wg.Done()
	start := time.Now()
	fmt.Printf("task 2 start: %s\n", start)
	sum := 1000
	for i := 1; i <= 500000; i++ {
		sum += i
	}
	elasped := time.Since(start)
	fmt.Printf("Task 2 completed in %s\n", elasped)

}
func task3() {

	defer wg.Done()
	start := time.Now()
	fmt.Printf("task 3 start: %s\n", start)
	sum := 1
	for i := 1; i <= 50; i++ {
		sum += i
	}
	elasped := time.Since(start)
	fmt.Printf("Task 3 completed in %s\n", elasped)

}
func task4() {

	defer wg.Done()
	start := time.Now()
	fmt.Printf("task 4 start: %s\n", start)
	sum := 0
	for i := 1; i <= 500000; i++ {
		sum += i
	}
	elasped := time.Since(start)
	fmt.Printf("Task 4 completed in %s\n", elasped)

}
func task5() {

	defer wg.Done()
	start := time.Now()
	fmt.Printf("task 5 start: %s\n", start)
	sum := 0
	for i := 1; i <= 500000; i++ {
		sum += i
	}
	elasped := time.Since(start)
	fmt.Printf("Task 5 completed in %s\n", elasped)

}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}
func (c *Circle) Perimeter() float64 {
	return 2 * c.Radius * 3.14
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println(e)
}

func writeToChannel(ch chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("channel send:", i, "time:", time.Now())
		ch <- i

	}
	close(ch)
}

func readFromChannel(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received from channel:", num, "time:", time.Now())
	}
}
func writeToChannel01(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("channel send:", i, "time:", time.Now())
	}
	close(ch)
}

func readFromChannel01(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received from channel:", num, "time:", time.Now())
	}
}

func main() {
	fmt.Println("welcome to homework02")
	fmt.Println("第一题：传指针增加10")
	a := 29
	fmt.Println("原始值:", a)
	add10(&a)
	fmt.Println("增加10后:", a)

	fmt.Println("第二题：传切片指针，切片内的值乘以2")
	slice2 := []int{2, 3, 7, 9, 15, 4}
	fmt.Println("原始切片:", slice2)
	multiplySlice(&slice2)
	fmt.Println("乘以2后的切片:", slice2)

	fmt.Println("第三题：协程打印奇数和偶数")
	wg.Add(2)
	go testgoroutine1()
	go testgoroutine2()
	wg.Wait()

	fmt.Println("第四题：协程任务调度")
	wg.Add(5)
	go task1()
	go task2()
	go task3()
	go task4()
	go task5()
	wg.Wait()

	fmt.Println("第五题：接口实现")
	var s Shape
	s = &Rectangle{Width: 7, Height: 3}
	fmt.Printf("Rectangle Area: %.2f\n", s.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", s.Perimeter())

	s = &Circle{Radius: 5}
	fmt.Printf("Circle Area: %.2f\n", s.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", s.Perimeter())

	fmt.Println("第六题：结构体组合")
	emp := Employee{
		Person:     Person{Name: "才生息", Age: 35},
		EmployeeID: "012161",
	}
	emp.PrintInfo()

	fmt.Println("第七题：channel")
	chanel := make(chan int)

	go writeToChannel(chanel)
	go readFromChannel(chanel)

	time.Sleep(2 * time.Second)

	fmt.Println("第八题:channel 缓冲")
	chanel01 := make(chan int, 5)

	go writeToChannel01(chanel01)
	go readFromChannel01(chanel01)
	time.Sleep(2 * time.Second)

	fmt.Println("第九题 锁")
	var counter int

	startTime := time.Now()
	fmt.Println("starting counter:", counter, "time:", time.Now())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter, "cost", time.Since(startTime))

	fmt.Println("第十题 原子操作")
	var counter1 int64
	startTime1 := time.Now()
	fmt.Println("starting counter:", counter1, "time:", time.Now())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter1, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter with atomic:", counter1, "cost", time.Since(startTime1))

}
