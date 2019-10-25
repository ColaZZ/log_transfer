package log_transfer

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

var (
	esClient *elastic.Client
)

type LogMessage struct {
	App     string
	Topic   string
	Message string
}


func initES(addr string) (err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		fmt.Println("connect es error ", err)
		return
	}
	esClient = client
	return
}


func sendToES(topic string, data []byte) (err error) {

	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = string(data)

	_, err = esClient.Index().
		Index(topic).
		//Type(topic).
		//Id(fmt.Sprintf("%d", i)).
		BodyJson(msg).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
		return
	}
	return
}
