package model

// Store defines what should be implemented by a model store
type Store interface {
	Save(model interface{}) (err error)
}
