package account

import (
	"encoding/json"

	"github.com/Covertness/ally/pkg/storage"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Create create object in database
func Create(name string, meta map[string]interface{}) (*Account, error) {
	var metaByte []byte
	if meta != nil {
		metaByte, _ = json.Marshal(meta)
	}

	newAccount := &Account{
		Name: name,
		Meta: postgres.Jsonb{RawMessage: metaByte},
	}

	err := storage.GetDB().Create(newAccount).Error
	if err != nil {
		return nil, err
	}

	return newAccount, nil
}

// GetOrCreate get account by name, create if not exist
func GetOrCreate(name string) (*Account, error) {
	var myAccount Account
	err := storage.GetDB().FirstOrCreate(&myAccount, Account{Name: name}).Error
	return &myAccount, err
}
