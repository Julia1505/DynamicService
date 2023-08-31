package repository

type Segment struct {
	Id   int    `db:"segment_id"`
	Name string `db:"segment_name"`
}

type User struct {
	Id int `db:"user_id"`
}

type UserWithSegments struct {
	ID   int    `db:"user_id"`
	Name string `db:"segment_name"`
}
