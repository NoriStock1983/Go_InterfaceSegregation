package repositories

import "InterfaceSegregation/domain/repositories"

type UsermstRepository struct{}

// SearchUsermstRepositoryのコンストラクタ
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

// usermstのユーザーコード件数取得処理
func (countbyusercd *UsermstRepository) CountByUsercd(usercd string) int {
	return 1
}

// usermstのユーザーコード登録処理
func (insertusermst *UsermstRepository) InsertUsermst() bool {
	return true
}

// usermstのユーザーコード更新処理
func (updateusermst *UsermstRepository) UpdateUsermst() bool {
	return true
}

// usermstのユーザーコード削除処理
func (deleteusermst *UsermstRepository) DeleteUsermst() bool {
	return true
}
