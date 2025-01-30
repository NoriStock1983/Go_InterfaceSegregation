package usermst

import "InterfaceSegregation/usecase/usermst"

type SearchUsermstController struct {
	SearchUsermst usermst.ISearchUsermstInteractor
}

// SearchUsermstControllerのコンストラクタ
func NewUsermstController(s usermst.ISearchUsermstInteractor) SearchUsermstController {
	return SearchUsermstController{
		SearchUsermst: s,
	}
}

// usermstの全件検索処理
func (s *SearchUsermstController) SarchAllUsermst() string {

	return s.SearchUsermst.SarchAllUsermst()
}

// usermstのユーザーコード検索処理
func (s *SearchUsermstController) SearchByUsercd() string {
	return s.SearchUsermst.SearchUsercd("usercd")
}
