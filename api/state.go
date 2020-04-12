package api

// State is interface of a simple State Manager
type State interface {
	Has(key string) bool

	Get(key string) (interface{}, error)
	GetInt(key string) (int, error)
	GetBool(key string) (bool, error)
	GetString(key string) (string, error)
	GetFloat(key string) (float32, error)

	Set(key string, value interface{}) error
}
