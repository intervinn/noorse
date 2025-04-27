package noorse

var instance *Bot = nil

func Instance() *Bot {
	if instance == nil {
		instance = New()
	}
	return instance
}
