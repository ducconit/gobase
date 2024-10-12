package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

// Make sure FileStore implements the Store interface
var _ Store = (*FileStore)(nil)

type FileStoreConfig struct {
	Path string
	Type string
	Name string
}

type FileStore struct {
	*viper.Viper
	changed bool

	cfg *FileStoreConfig
}

func NewFileStore(cfg *FileStoreConfig) *FileStore {
	if cfg == nil {
		cfg = &FileStoreConfig{}
	}

	cfg.assignDefaults()

	v := viper.New()
	v.SetConfigType(cfg.Type)
	v.SetConfigName(cfg.Name)
	v.AddConfigPath(cfg.Path)

	return &FileStore{
		cfg:   cfg,
		Viper: v,
	}
}

func (s *FileStoreConfig) assignDefaults() {
	if s.Path != "" {
		ext := filepath.Ext(s.Path)

		if ext != "" {
			s.Name, _ = strings.CutSuffix(filepath.Base(s.Path), ext)
			s.Path = filepath.Dir(s.Path)
			s.Type = strings.TrimPrefix(ext, ".")
		}
	}

	if s.Name == "" {
		s.Name = "config"
	}

	if s.Type == "" {
		s.Type = "yml"
	}

	if s.Path == "" {
		// The project's executable directory
		s.Path = "."
	}
}

func (s *FileStore) Load() error {
	if err := s.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (s *FileStore) WasChanged() bool {
	return s.changed
}

func (s *FileStore) Save() error {
	return s.WriteConfig()
}

func (s *FileStore) GetAll() map[string]any {
	return s.AllSettings()
}

func (s *FileStore) GetArray(key string) []any {
	return s.Get(key).([]any)
}
