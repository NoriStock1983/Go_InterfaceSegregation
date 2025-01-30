package usermst

import "InterfaceSegregation/domain/repositories"

type ISearchUsermstInteractor interface {
	SarchAllUsermst() string
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
