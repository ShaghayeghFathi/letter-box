package letter_box

type ServerConfig struct {
	Name string
	URL  string
}

func NewServer(name string, url string) {
	switch name {
	case "emqx":
		{
			// setup emqx connection
		}

	default:
		// give validation error
	}
}
