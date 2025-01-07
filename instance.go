package noorse

var instance *Bot = nil

func GetInstance() *Bot {
	if instance == nil {
		instance = New()
	}
	return instance
}
