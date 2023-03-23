package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

func signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		sig := <-ch
		fmt.Println("received signal:", sig)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			println("the program will exit")
			signal.Stop(ch)
			os.Exit(0)
			return
		default:
			println("the program will continue")
			continue
		}
	}
}

func main() {
	fmt.Println("PID:", os.Getpid())
	const topic = "quickstart-events"

	// 创建新的消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("fail to start consumer:", err)
		return
	}

	// 根据 topic 获取所有的分区列表
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("fail to get partition list:", err)
		return
	}
	fmt.Println("partition list:", partitionList)

	// 遍历所有的分区
	for partition := range partitionList {
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d: %v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s\n",
					msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}

	signalHandler()
}
