package repositories

type ISearchUsermstRepository interface {
	AllSearchUsermst() string
}

type IInsertUsermstRepository interface {
	CountByUsercd(usercd string) int
	InsertUsermst() bool
}

type IUpdateUsermstRepository interface {
	CountByUsercd(usercd string) int
	SearchByUsercd() string
	UpdateUsermst() bool
}

type IDeleteUsermstRepository interface {
	CountByUsercd(usercd string) int
	SearchByUsercd() string
	DeleteUsermst() bool
}
