package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db: db,
	}, nil
}

func (p *PostgresStorage) Close() {
	closeDB(p.db)
}

func (s *PostgresStorage) AddSegment(segment Segment) error {
	_, err := s.db.Exec(`insert into segment (segment_name)
	values ($1)`, segment.Name)
	if err != nil {
		return errors.Wrap(err, "insert")
	}
	return nil
}

func (s *PostgresStorage) DeleteSegment(segment Segment) error {
	_, err := s.db.Exec(`delete from segment where segment.segment_name = $1`, segment.Name)
	if err != nil {
		return errors.Wrap(err, "insert")
	}
	return nil
}

func (s *PostgresStorage) GetSegmentsOfUser(user User) ([]Segment, error) {
	var segments []Segment

	err := s.db.Select(&segments,
		`select  segment.segment_name from users
left join segment_user on users.user_id = segment_user.user_id
left join segment on segment_user.segment_id = segment.segment_id
where users.user_id = $1`, user.Id)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}
	return segments, nil
}

func (s *PostgresStorage) AddSegmentToUser(user User, segments Segment) error {
	_, err := s.db.Exec(`insert into segment_user (user_id, segment_id) 
values ($1,$2)`, user.Id, segments.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) DeleteSegmentFromUser(user User, segment Segment) error {
	_, err := s.db.Exec(`delete from segment_user where segment_id=$1 and user_id=$2`, segment.Id, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) IsUserExists(id int) (bool, error) {
	var user User
	err := s.db.Get(&user, `select user_id from users where user_id=$1`, id)
	switch err {
	case nil:
		return true, nil
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, err
	}
}

func (s *PostgresStorage) IsSegmentExists(segment Segment) (bool, error) {
	var segmentFromDB Segment
	err := s.db.Get(&segmentFromDB, `select segment_name from segment where segment_name=$1`, segment.Name)
	switch err {
	case nil:
		return true, nil
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, err
	}
}

func (s *PostgresStorage) IsSegmentUser(segment Segment, user User) (int, bool, error) {
	var seg Segment
	err := s.db.Get(&seg, `select segment.segment_id from segment
join segment_user on segment.segment_id = segment_user.segment_id
where segment.segment_name=$1 and segment_user.user_id = $2
`, segment.Name, user.Id)
	switch err {
	case nil:
		return seg.Id, true, nil
	case sql.ErrNoRows:
		return 0, false, nil
	default:
		return 0, false, err
	}
}

func (s *PostgresStorage) GetSegmentID(segment Segment) (int, error) {
	var seg Segment
	err := s.db.Get(&seg, `select segment_id from segment where segment_name=$1`, segment.Name)
	if err != nil {
		return 0, err
	}
	return seg.Id, nil
}
