package favorite

import "github.com/Covertness/ally/pkg/storage"

// GetMyFavorites get all my favorite items
func GetMyFavorites(accountID uint) ([]*Favorite, error) {
	var models []*Favorite
	err := storage.GetDB().Order("id ASC").Find(&models, Favorite{AccountID: accountID}).Error
	if err != nil {
		return nil, err
	}

	return models, err
}

// SetMyFavorite set my favorite
func SetMyFavorite(fav *Favorite) error {
	return storage.GetDB().Create(fav).Error
}

// DelMyFavorite delete my favorite
func DelMyFavorite(fav *Favorite) error {
	return storage.GetDB().Where(fav).Delete(&Favorite{}).Error
}
