package main

/*
	生产者消费者模型--处理模拟订单:

	在实际的开发中，生产者消费者模式应用也非常的广泛，例如：在电商网站中，订单处理，就是非常典型的生产者消费者模式。
	当很多用户单击下订单按钮后，订单生产的数据全部放到缓冲区（队列）中，然后消费者将队列中的数据取出来发送者仓库管理等系统。
	通过生产者消费者模式，将订单系统与仓库管理系统隔离开，且用户可以随时下单（生产数据）。
	如果订单系统直接调用仓库系统，那么用户单击下订单按钮后，要等到仓库系统的结果返回。这样速度会很慢。

下面代码模拟一个订单处理的过程:

*/
import "fmt"

//定义商品信息结构体
type orderInfo struct {
	id int //字段 | 成员变量
}

//生产者 只写channel
func producer(out chan<- orderInfo) {

	for i := 0; i < 10; i++ { // 循环生成10份订单
		order := orderInfo{id: i + 1}
		out <- order // 写入channel
	}
	close(out) // 写完channel,关闭channel
}

//消费者 只读channel
func consumer(in <-chan orderInfo) {

	for order := range in { // 从channel中取出订单
		fmt.Println("订单id为: ", order.id) //模拟订单处理
	}

}

func main() {
	// 定义一个双向 channel， 指定数据类型为OrderInfo
	ch := make(chan orderInfo)

	// 创建新go程，传递只写channel 传递的是引用
	go producer(ch)

	// 主go程，传只读channel; 与 producer 传递的是同一个 channel
	consumer(ch)

}

/*
OrderInfo为订单信息，这里为了简单只定义了一个订单编号属性，然后生产者模拟10个订单，消费者对产生的订单进行处理。
输出结果:

	订单id为:  1
	订单id为:  2
	订单id为:  3
	订单id为:  4
	订单id为:  5
	订单id为:  6
	订单id为:  7
	订单id为:  8
	订单id为:  9
	订单id为:  10

*/
