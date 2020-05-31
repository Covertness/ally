package timeline

import "github.com/Covertness/ally/pkg/storage"

// GetOrCreate get timeline, create if not exist
func GetOrCreate(timeline *Timeline) (*Timeline, error) {
	var t Timeline
	err := storage.GetDB().FirstOrCreate(&t, timeline).Error
	return &t, err
}

// GetByID get the specified timeline by id
func GetByID(id uint) (*Timeline, error) {
	var model Timeline
	err := storage.GetDB().First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return &model, err
}
