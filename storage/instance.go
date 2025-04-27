package storage

var instance *Storage = nil

func Instance() *Storage {
	if instance == nil {
		instance = New()
	}
	return instance
}
