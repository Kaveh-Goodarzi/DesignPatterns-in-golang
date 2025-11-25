package main

import "fmt"

type Publisher interface {
	Register(subscriber Subscriber)
	NotifyAll(msg string)
}

type Subscriber interface {
	ReactToPublisherMsg(msg string)
}

type publisher struct {
	subscribers []Subscriber
}

func GetNewPublisher() publisher {
	return publisher{subscribers: make([]Subscriber, 0)}
}

func (p *publisher) Register(subscriber Subscriber) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p publisher) NotifyAll(msg string) {
	for _, subs := range p.subscribers {
		fmt.Println("Publisher notifying Subscriber with ID", subs.(subscriber).subscriberID)
		subs.ReactToPublisherMsg(msg)
	}
}

type subscriber struct {
	subscriberID string
}

func GetNewSubscriber(ID string) subscriber {
	return subscriber{subscriberID: ID}
}

func (s subscriber) ReactToPublisherMsg(msg string) {
	fmt.Printf("The Subscriber Received Message %s\n For Subscriber ID %s\n", msg, s.subscriberID)
}

func main() {
	pub := GetNewPublisher()

	subs := GetNewSubscriber("gEnIuS")

	pub.Register(subs)
	pub.NotifyAll("Hello everybody.")
}