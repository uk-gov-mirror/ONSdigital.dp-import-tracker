package main

import (
	"context"
	"flag"
	"time"

	"github.com/ONSdigital/dp-import/events"
	"github.com/ONSdigital/dp-kafka/kafka"
)

var instanceID = flag.String("instance", "ac280d98-7211-4b04-9497-40f199396cc3", "")
var jobID = flag.String("job", "cd8759b7-a2f7-49f1-9b56-245d848d50fc", "")
var url = flag.String("url", "https://s3-eu-west-1.amazonaws.com/dp-frontend-florence-file-uploads/159-coicopcomb-inc-geo_cutcsv", "")

var topic = flag.String("topic", "input-file-available", "")
var kafkaHost = flag.String("kafka", "localhost:9092", "")

func main() {

	flag.Parse()
	ctx := context.Background()

	var brokers []string
	brokers = append(brokers, *kafkaHost)

	producer, _ := kafka.NewProducer(ctx, brokers, *topic, int(2000000), kafka.CreateProducerChannels())

	fileEvent := events.InputFileAvailable{
		InstanceID: *instanceID,
		JobID:      *jobID,
		URL:        *url,
	}

	bytes, error := events.InputFileAvailableSchema.Marshal(fileEvent)
	if error != nil {
		panic(error)
	}
	producer.Channels().Output <- bytes

	time.Sleep(time.Duration(time.Second))

	producer.Close(ctx)
}
