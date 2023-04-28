package cmq

import (
	"github.com/ShaghayeghFathi/letter-box/internal/config"
	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
)

// NatsStreaming is the struct for nats streaming
type NatsStreaming struct {
	stanConn      stan.Conn
	subscriptions []stan.Subscription
}

func newNatsStreaming(stanConn stan.Conn) *NatsStreaming {
	return &NatsStreaming{stanConn: stanConn}
}

// QueueSubscribe is a function for subscribing to a subject
// it will subscribe to the subject and will call the handler function when a message is received
func (n *NatsStreaming) QueueSubscribe(subject, qg string, handler func(subject string, data []byte)) {
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
func (n *NatsStreaming) UnsubscribeAll() {
	for _, subscription := range n.subscriptions {
		err := subscription.Unsubscribe()
		if err != nil {
			log.WithError(err).Error("error in unsubscribing from nats streaming subject")
		}
	}
}

func Connect(nc config.NatsStreaming) *NatsStreaming {
	opts := []stan.Option{
		stan.NatsURL(nc.Address),
		//stan.Pings(config.C.NatsStreaming.PingInterval, config.C.NatsStreaming.PingMaxOut),
		stan.SetConnectionLostHandler(func(conn stan.Conn, err error) {
			log.Fatalf("nats streaming connection lost acceptHandler: %s", err.Error())
		}),
	}

	stanConn, err := stan.Connect(
		nc.ClusterID,
		nc.ClientID+strconv.Itoa(rand.Int()),
		opts...,
	)
	if err != nil {
		log.Fatalf("error in stan connection %v", err)
	}
	ns := newNatsStreaming(stanConn)

	log.Info("using NatsStreaming as CMQ")

	return ns
}
