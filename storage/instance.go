package storage

var instance *Storage = nil

func GetInstance() *Storage {
	if instance == nil {
		instance = New()
	}
	return instance
}
