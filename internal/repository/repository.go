package repository

type Repository interface {
	IsUserExists(id int) (bool, error)
	AddSegment(segment Segment) error
	DeleteSegment(segment Segment) error
	GetSegmentsOfUser(user User) ([]Segment, error)
	AddSegmentToUser(user User, segments Segment) error
	DeleteSegmentFromUser(user User, segment Segment) error
	IsSegmentExists(segment Segment) (bool, error)
	IsSegmentUser(segment Segment, user User) (int, bool, error)
	GetSegmentID(segment Segment) (int, error)
}
