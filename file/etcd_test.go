package file

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestEtcD(t *testing.T) {
	fmt.Println( etcd.Get("name"))
	fmt.Println( etcd.Set("name\n:\"/```````", "zw7"))
	fmt.Println( etcd.Get("name\n:\"/```````"))
	//fmt.Println( etcd.Delete("name"))
	fmt.Println( etcd.Delete("zw7"))

	fmt.Println(strconv.Itoa(4))
}

func TestIO(t *testing.T) {

	for i:=0; i<100; i++ {
		name := strconv.Itoa(i)

		go func() {
			time.Sleep(time.Duration(rand.Intn(5000))* time.Millisecond)
			err := etcd.Set("name", name)
			if err!=nil {
				fmt.Println("Set Key Error:", err)
			}
			fmt.Println("Set name:", name)
		}()

		go func() {
			time.Sleep(time.Duration(rand.Intn(5000))* time.Millisecond)
			v, err := etcd.Get("name")
			if err!=nil {
				fmt.Println("Get Key Error:", err)
			}
			fmt.Println("Get name:", v)
		}()
	}

	time.Sleep(10*time.Second)
}