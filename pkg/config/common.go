package config

import (
	"encoding/json"
	"os"

	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Covertness/ally/pkg/storage"
)

func getenv(key, defaultValue string) string {
	if len(os.Getenv(key)) == 0 {
		return defaultValue
	}
	return os.Getenv(key)
}

func getconfig(key string) ([]byte, error) {
	var cfg Config
	err := storage.GetDB().First(&cfg, &Config{Key: key}).Error
	if err != nil {
		return nil, err
	}

	return cfg.Value.RawMessage, nil
}

func setconfig(key string, value []byte) error {
	var cfg Config
	err := storage.GetDB().FirstOrCreate(&cfg, Config{Key: key}).Error
	if err != nil {
		return err
	}

	return storage.GetDB().Model(&cfg).Updates(map[string]interface{}{
		"value": postgres.Jsonb{RawMessage: value},
	}).Error
}

func encodeValue(c map[string]string) ([]byte, error) {
	return json.Marshal(c)
}

func decodeValue(v []byte) (map[string]string, error) {
	var content map[string]string
	err := json.Unmarshal(v, &content)
	if err != nil {
		return nil, err
	}
	return content, nil
}
