package broker

import (
	"github.com/micro/plugins/v5/broker/memory"
	"go-micro.dev/v5/broker"
)

var (
	// DefaultBroker is the default Broker.
	bk = memory.NewBroker()
)

func Init(opts ...broker.Option) error {
	return bk.Init(opts...)
}

func Connect() error {
	return bk.Connect()
}

func Disconnect() error {
	return bk.Disconnect()
}

func Publish(topic string, msg *broker.Message, opts ...broker.PublishOption) error {
	return bk.Publish(topic, msg, opts...)
}

func Subscribe(topic string, handler broker.Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error) {
	return bk.Subscribe(topic, handler, opts...)
}

// String returns the name of the Broker.
func String() string {
	return bk.String()
}
