package cmq

import (
	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
)

type CMQ interface {
	QueueSubscribe(subject, qg string, handler func(subject string, data []byte))
	UnsubscribeAll()
}

// natsStreaming is the struct for nats streaming
type natsStreaming struct {
	stanConn      stan.Conn
	subscriptions []stan.Subscription
}

func newNatsStreaming(stanConn stan.Conn) *natsStreaming {
	return &natsStreaming{stanConn: stanConn}
}

// QueueSubscribe is a function for subscribing to a subject
// it will subscribe to the subject and will call the handler function when a message is received
func (n *natsStreaming) QueueSubscribe(subject, qg string, handler func(subject string, data []byte)) {
	subscription, err := n.stanConn.QueueSubscribe(subject, qg, func(msg *stan.Msg) {
		handler(subject, msg.Data)
	})
	if err != nil {
		log.WithFields(log.Fields{
			"err":     err,
			"subject": subject,
		}).Fatal("error in subscribing to nats streaming")
	}

	n.subscriptions = append(n.subscriptions, subscription)
}

// UnsubscribeAll is a function for unsubscribing from all subjects
func (n *natsStreaming) UnsubscribeAll() {
	for _, subscription := range n.subscriptions {
		err := subscription.Unsubscribe()
		if err != nil {
			log.WithError(err).Error("error in unsubscribing from nats streaming subject")
		}
	}
}
