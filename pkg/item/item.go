package item

import (
	"encoding/json"

	"github.com/Covertness/ally/pkg/storage"
	"github.com/cockroachdb/apd"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Create create object in database
func Create(externalID string, price *apd.Decimal, meta map[string]interface{}) (*Item, error) {
	var metaByte []byte
	if meta != nil {
		metaByte, _ = json.Marshal(meta)
	}

	newItem := &Item{
		ExternalID: externalID,
		Price:      price,
		Meta:       postgres.Jsonb{RawMessage: metaByte},
	}

	err := storage.GetDB().Create(newItem).Error
	if err != nil {
		return nil, err
	}

	return newItem, nil
}
