package events

import (
	"github.com/gpabois/gostd/result"
)

type Subscriber = func(msg any) result.Result[bool]

type Subscribtion interface {
	// Pump the next message
	Pump(sub Subscriber) result.Result[bool]
}

type IEventService interface {
	Publish(name string, payload any) result.Result[bool]
	Subscribe(eventName string, subscriberName string) result.Result[Subscribtion]
}
