package repositories

// 検索用のインターフェース
// 検索処理を行う場合にのみ必要な関数を定義する。
// 検索処理を行うためだけのインターフェースのため、登録、更新、削除処理は含まない。
type ISearchUsermstRepository interface {
	AllSearchUsermst() string
	CountByUsercd(usercd string) int
	SearchByUsercd(usercd string) string
}

// Insert用のインターフェース
// Insert処理を行う場合にのみ必要な関数を定義する。
// Insertを行う前に、対象データが存在するかどうか確認するためにCountByUsercd関数を定義する。
type IInsertUsermstRepository interface {
	CountByUsercd(usercd string) int
	InsertUsermst() bool
}

// Update用のインターフェース
// Update処理を行う場合にのみ必要な関数を定義する。
// Updateを行う前に、対象データが存在するかどうか確認するためにCountByUsercd関数を定義する。
// Update処理を実行する前に、対象データが他のユーザに更新されていないかなどを確認するために、SearchByUsercd関数を定義する。
type IUpdateUsermstRepository interface {
	CountByUsercd(usercd string) int
	SearchByUsercd() string
	UpdateUsermst() bool
}

// Delete用のインターフェース
// Delete処理を行う場合にのみ必要な関数を定義する。
// Deleteを行う前に、対象データが存在するかどうか確認するためにCountByUsercd関数を定義する。
// Delete処理を実行する前に、対象データが他のユーザに更新されていないかなどを確認するために、SearchByUsercd関数を定義する。
type IDeleteUsermstRepository interface {
	CountByUsercd(usercd string) int
	SearchByUsercd() string
	DeleteUsermst() bool
}
