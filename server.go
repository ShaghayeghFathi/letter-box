package letter_box

import (
	"github.com/ShaghayeghFathi/letter-box/internal/cmq"
	"github.com/ShaghayeghFathi/letter-box/internal/config"
)

type ServerConfig struct {
	Name string
	URL  string
}

type Server struct {
	//Emq EMQ
	NatsStreaming *cmq.CMQ
}

func NewServer(name string, url string) {
	switch name {
	case "emqx":
		{
			// setup emqx connection
		}
	case "nats":
		{
			c := config.NatsStreaming{
				Address:   "",
				ClusterID: "",
				ClientID:  "",
			}
			cmq.Connect(c)

		}

	default:
		// give validation error
	}
}
