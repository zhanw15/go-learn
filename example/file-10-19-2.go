package main

import (
	"fmt"
	"strconv"
	"time"
)

type Producer interface {
	Product()
}

type Consumer interface {
	Consume()
}

type MessProducer struct {
	Mess chan <-string
}

type MessConsumer struct {
	Mess <-chan string
}

func NewStringProducer( ch chan string) Producer {
	p := new(MessProducer)
	p.Mess = ch
	return p
}

func NewStringConsumer(ch chan string) Consumer {
	c := new(MessConsumer)
	c.Mess = ch
	return c
}

func (c *MessConsumer) Consume() {
	t := time.NewTicker(3*time.Second)

	for range t.C {
		for i:=0; i<3; i++ {
			fmt.Println("consume message time:", time.Now())
			fmt.Println("message" + strconv.Itoa(i+1) + ":", <-c.Mess)
		}
	}
}

func (p *MessProducer) Product() {
	t := time.NewTicker(1*time.Second)

	for range t.C {
		p.Mess <- "product message time:" + time.Now().Format("2006/01/02_15:04:05")
	}
}

func main() {
	//开发生产者/消费者程序，按接口开发实现，消息只需要支持字符串即可
	//生产者每1秒钟生产一条消息，消息格式为：product message time:当前时间
	//消费者每3秒钟消费3条消息，并把消费的消息打印出来，消息格式为：consume message time:当前时间 message1或者message2或者message3:生产消息
	//要求：满足以下测试用例，消费者按消息格式每3秒钟输出3条消息
	messagechan := make(chan string,10)
	producer := NewStringProducer(messagechan)
	go producer.Product()
	consumer := NewStringConsumer(messagechan)
	consumer.Consume()
}