package cmq

type CMQ interface {
	QueueSubscribe(subject, qg string, handler func(subject string, data []byte))
	UnsubscribeAll()
}
