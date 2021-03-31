package kafka

import(
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

	if kafkaSender,err =sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config);err != nil{
		panic(err)
	}

	//TODO HANK 设置配置文件
	//初始化消费者信息
	if kafkaReceiver,err =sarama.NewConsumer([]string{"127.0.0.1:9092"}, &sarama.Config{});err != nil{
		panic(err)
	}
}

//生产者
func KafkaSender()sarama.SyncProducer{
	return kafkaSender
}

//消费者
func KafkaReceiver()sarama.Consumer{
	return kafkaReceiver
}
