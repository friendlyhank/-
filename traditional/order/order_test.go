package order

import (
	"fmt"
	_ "miaosha/common"
	"testing"
	"time"
)


var (
	total = 1000
	threads = 1000
)

type result struct{
}

func TestCreateOrder(t *testing.T){
	CreateOrder(100000,1)
}

func TestBenckPayOrder(t *testing.T){
	ch := make(chan *result,threads)
	start :=time.Now()
	for i :=0;i<threads;i++{
		go operate(i,total/threads,ch)
	}

	for i :=0;i < threads;i++{
		<-ch
	}

	d := time.Now().Sub(start)
	fmt.Printf("%f seconds total\n",d.Seconds())
}

func TestBenckPayOrder2(t *testing.T){
	for{
		ch := make(chan *result,threads)
		start := time.Now()
		for i := 0;i <threads;i++{
			go operate(i,total/threads,ch)
		}

		for i :=0;i < threads;i++{
			<-ch
		}
		d := time.Now().Sub(start)
		fmt.Printf("%f seconds total\n",d.Seconds())

		time.Sleep(1 *time.Second)
	}
}

func operate(id,count int,ch chan *result){
	r := &result{}
	for i := 0;i <count;i++{
		fmt.Printf("第%v个用户第%v次下单\n",id,i)
		CreateOrder(100000,1)
	}
	ch <- r
}



