package repositories

import "InterfaceSegregation/domain/repositories"

type UsermstRepository struct{}

func NewSearchUsermstRepository() repositories.ISearchUsermstRepository {
	return &UsermstRepository{}
}

// usermstの全件検索処理
func (searchall *UsermstRepository) AllSearchUsermst() string {
	return "AllSearchUsermst"
}

// usermstのユーザーコード検索処理
func (searchbyusercd *UsermstRepository) SearchByUsercd(userced string) string {
	return userced
}

func (countbyusercd *UsermstRepository) CountByUsercd(usercd string) int {
	return 1
}

func (insertusermst *UsermstRepository) InsertUsermst() bool {
	return true
}

func (updateusermst *UsermstRepository) UpdateUsermst() bool {
	return true
}

func (deleteusermst *UsermstRepository) DeleteUsermst() bool {
	return true
}
