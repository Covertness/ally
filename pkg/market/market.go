package market

import "github.com/Covertness/ally/pkg/storage"

// GetAll get all available markets
func GetAll() ([]*Market, error) {
	var models []*Market
	err := storage.GetDB().Order("id ASC").Find(&models, Market{Enable: true}).Error
	if err != nil {
		return nil, err
	}

	return models, err
}

// GetByName get the specified market by name
func GetByName(name string) (*Market, error) {
	var model Market
	err := storage.GetDB().Where(&Market{Name: name}).First(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, err
}
