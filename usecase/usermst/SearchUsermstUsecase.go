package usermst

import "InterfaceSegregation/domain/repositories"

type ISearchUsermstInteractor interface {
	SarchAllUsermst() string
	SearchUsercd(usercd string) string
}

type SearchUsermstInteractor struct {
	repositories.ISearchUsermstRepository
}

func NewSearchUsermstUsecase(searchusermst repositories.ISearchUsermstRepository) ISearchUsermstInteractor {
	return &SearchUsermstInteractor{
		searchusermst,
	}
}

func (s *SearchUsermstInteractor) SarchAllUsermst() string {
	return s.AllSearchUsermst()
}

func (s *SearchUsermstInteractor) SearchUsercd(usercd string) string {
	return s.SearchByUsercd(usercd)
}
