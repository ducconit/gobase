package config

type Store interface {
	Load() error
	Save() error
	WasChanged() bool

	Get(key string) any
	GetString(key string) string
}
