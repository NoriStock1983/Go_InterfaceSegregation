package usermst

import "InterfaceSegregation/domain/repositories"

type ISearchUsermstInteractor interface {
	SarchAllUsermst() string
	SearchUsercd(usercd string) string
}

type SearchUsermstInteractor struct {
	repositories.ISearchUsermstRepository
}

// SearchUsermstInteractorのコンストラクタ
func NewSearchUsermstUsecase(searchusermst repositories.ISearchUsermstRepository) ISearchUsermstInteractor {
	return &SearchUsermstInteractor{
		searchusermst,
	}
}

// usermstの全件検索処理(ビジネスロジック)
func (s *SearchUsermstInteractor) SarchAllUsermst() string {
	return s.AllSearchUsermst()
}

// usermstのユーザーコード検索処理(ビジネスロジック)
func (s *SearchUsermstInteractor) SearchUsercd(usercd string) string {
	return s.SearchByUsercd(usercd)
}
