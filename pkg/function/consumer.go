package function

type FConsumerFunction func(v interface{})

type IConsumer interface {
	Accept(interface{})
}

type Consumer struct {
}

func (c *Consumer) Accept(v interface{}) {
	return
}

type IntConsumer struct {
	Consumer
}

func (c *IntConsumer) Accept(v interface{}) {

}
