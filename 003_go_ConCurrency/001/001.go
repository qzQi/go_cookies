package main

// https://livebook.manning.com/book/go-in-practice/chapter-3/1
import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

func ehco(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
func myFunc1() {
	go ehco(os.Stdin, os.Stdout)
	time.Sleep(10 * time.Second)
	fmt.Println("time out")
	os.Exit(0)
}

// You want to use a one-shot function in a way
// that doesn’t block the calling function,
// and you’d like to make sure that it runs.
func myFunc2() {
	fmt.Println("outside a goroutine")
	go func(str string) {
		fmt.Println("hello", str, "inside routine")
	}("qizhiyun")

	fmt.Println("again in outside")

	// omit this the routine may not be sched
	// runtime.Gosched()
	runtime.Gosched()
}

//

func compress(file string) error {
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(file + ".gz")
	if err != nil {
		return nil
	}
	gzout := gzip.NewWriter(out)
	// count :=
	_, err = io.Copy(gzout, in)
	return err
}

func myFunc3() {
	var wg sync.WaitGroup
	var i int = -1
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1)
		fileName := file
		go func(filename string) {
			compress(fileName)
			wg.Done()
		}(fileName)
		// 这里可以这样使用吗？不可以还是不能使用range里面的变量作为 参数
	}
	wg.Wait()
	fmt.Printf("compress %d file\n", i+1)
}

// use mutex
func myFunc4() {
	var lock_ sync.Mutex
	lock_.Lock()
	defer lock_.Unlock()

	var rwLock_ sync.RWMutex
	rwLock_.RLock()
	defer rwLock_.RUnlock()

	// rwlock's write lock the same as mutex
	rwLock_.Lock()
	defer rwLock_.Unlock()
}

// use channel
// out chan []byte : out是一个channel，限定方向就是
// out chan <- []byte
// out <-chan []byte 错误写法
func readStdin(out chan<- []byte) {
	for {
		// fmt.Println("is scheduling")
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}
func myFunc5() {
	time.Sleep(1 * time.Second)
	// the same
	sleep := time.After(1 * time.Second)
	<-sleep

	// done: a channel <- timr.Time
	done := time.After(30 * time.Second)
	// done:=make(chan <-time.Time)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)

		case <-done:
			fmt.Println("time out")
			os.Exit(0)
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("nothing typed in")
		}
	}
}

func testChannel() {
	// in:=make(<- chan int)
	// out:=make(chan <- int)
	// out<-2
	// var i int=<-in
	// // in<-i
	// 在声明的时候直接声明为单向的不容易，在使用的时候
	// 可以把它形参为单向
}
func useInChan(in <-chan int) {
	i := <-in
	fmt.Printf("chan type is %T,i is %d type is%T\n", in, i, i)
}
func useOutChan(out chan<- int) {
	fmt.Printf("out chan type is %T\n", out)
	out <- 2
}

// correct close the channel
func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(time.Millisecond * 500)
		}
	}
}
func closeChan() {
	// 其实没有缓冲的channel不会有内存泄漏，
	// 只会阻塞，这个就是为了演示
	msg := make(chan string)
	done := make(chan bool)
	until := time.After(5 * time.Second)
	go send(msg, done)
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			// 等待receive from msg
			time.Sleep(1 * time.Second)
			return
		}
	}
}

func printLine() {
	fmt.Println("+++++++++++")
}
func main() {
	// myFunc1()
	printLine()
	// myFunc2()

	// myFunc3()

	// myFunc5()

	chanInt := make(chan int)
	go useOutChan(chanInt)
	useInChan(chanInt)
}
