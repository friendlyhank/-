package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var(
	kafkaSender sarama.SyncProducer
	kafkaReceiver sarama.Consumer
)

func Init() {
	var err error
	//初始化生产者信息
	config := sarama.NewConfig()
	config.Producer.RequiredAcks =sarama.WaitForAll //发送完数据需要leader和follow都确认
	config.Producer.Partitioner  = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true // 成功交付的消息将在success channel返回

	kafkaSender,err =sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config)
	if err != nil{
		panic(err)
	}

	//TODO HANK 设置配置文件
	//初始化消费者信息
	kafkaReceiver,err =sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil{
		panic(err)
	}

	partitionList, err := kafkaReceiver.Partitions("createorder") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}

	if len(partitionList) == 0{
		fmt.Printf("没有设置分区信息")
		return
	}

	// 针对每个分区创建一个对应的分区消费者
	pc, err := kafkaReceiver.ConsumePartition("createorder", int32(partitionList[0]), sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("failed to start consumer for partition %d,err:%v\n", partitionList[0], err)
		return
	}

	//这里不能关闭，否则回报错
	//defer pc.Close()

	go func() {
		for{
			msg :=<- pc.Messages()
			if msg == nil{
				fmt.Println(msg)
				continue
			}
			fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
		}
	}()
}

//消费者
func KafkaReceiver()sarama.Consumer{
	return kafkaReceiver
}
