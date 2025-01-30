package repositories

// 検索用のインターフェース
type ISearchUsermstRepository interface {
	AllSearchUsermst() string
	CountByUsercd(usercd string) int
	SearchByUsercd(usercd string) string
}

// Insert用のインターフェース
type IInsertUsermstRepository interface {
	CountByUsercd(usercd string) int
	InsertUsermst() bool
}

// Update用のインターフェース
type IUpdateUsermstRepository interface {
	CountByUsercd(usercd string) int
	SearchByUsercd() string
	UpdateUsermst() bool
}

// Delete用のインターフェース
type IDeleteUsermstRepository interface {
	CountByUsercd(usercd string) int
	SearchByUsercd() string
	DeleteUsermst() bool
}
