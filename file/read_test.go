package file

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var(
	counter int //所有goroutine都要增加其值的变量
	wg sync.WaitGroup  //用来等待程序结束
	mutex sync.Mutex //用来定义一段代码临界区
)
func TestCreateFile(t *testing.T) {
	wg.Add(2) //表示要等待两个goroutine
	go inCounter(1)
	go inCounter(2)
	wg.Wait() //等待goroutine结束
	fmt.Printf("Final Counter: %d\\n",counter)
}

//incounter 使用互斥锁来同步并保证安全访问，增加counter的值
func inCounter(id int){
	defer wg.Done() //用来通知main 函数工作已完成
	for count:=0;count<2; count++{
		mutex.Lock() //同一时刻只允许一个goroutine 进入这一临界区
		value :=counter
		runtime.Gosched() //当前goroutine从线程退出，并放回到队列
		value++
		counter=value
		mutex.Unlock() //释放锁，允许正在等待的goroutine进入临界区
	}
}