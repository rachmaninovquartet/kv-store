package interfaces

// KeyValueService defines the interface for key-value operations
type KeyValueService interface {
	SetKeyValue(key string, value interface{}, ttl *int) bool
	GetValue(key string) interface{}
	DeleteKey(key string) bool
	KeyExists(key string) bool
}
