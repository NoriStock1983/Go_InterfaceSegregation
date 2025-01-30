package usermst

import "InterfaceSegregation/usecase/usermst"

type SearchUsermstController struct {
	SearchUsermst usermst.ISearchUsermstInteractor
}

func NewUsermstController(s usermst.ISearchUsermstInteractor) SearchUsermstController {
	return SearchUsermstController{
		SearchUsermst: s,
	}
}

func (s *SearchUsermstController) SarchAllUsermst() string {

	return s.SearchUsermst.SarchAllUsermst()
}

func (s *SearchUsermstController) SearchByUsercd() string {
	return s.SearchUsermst.SearchUsercd("usercd")
}
