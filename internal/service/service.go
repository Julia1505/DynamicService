package service

import (
	"UserSegmentationService/internal/repository"
	"github.com/pkg/errors"
)

var (
	ErrSegmentExists = errors.New("Segment already exists")
	ErrNoSegment     = errors.New("Segment doesn't exists")
	ErrNoUser        = errors.New("No such user")
)

type DynamicSegment interface {
	CreateSegment(segment repository.Segment) error
	DeleteSegment(segment repository.Segment) error
	ChangeUserSegments(user repository.User, add []repository.Segment, delete []repository.Segment) error
	ShowUserSegments(user repository.User) ([]repository.Segment, error)
}

type Service struct {
	DynamicSegment
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		NewDynamicSegmentService(repo),
	}
}

type DynamicSegmentService struct {
	repo repository.Repository
}

func NewDynamicSegmentService(repo repository.Repository) *DynamicSegmentService {
	return &DynamicSegmentService{repo: repo}
}

func (s *DynamicSegmentService) CreateSegment(segment repository.Segment) error {
	is, err := s.repo.IsSegmentExists(segment)
	if err != nil {
		return err
	}

	if is {
		return ErrSegmentExists
	}

	err = s.repo.AddSegment(segment)
	if err != nil {
		return err
	}
	return nil
}

func (s *DynamicSegmentService) DeleteSegment(segment repository.Segment) error {
	is, err := s.repo.IsSegmentExists(segment)
	if err != nil {
		return err
	}

	if !is {
		return ErrNoSegment
	}

	err = s.repo.DeleteSegment(segment)
	if err != nil {
		return err
	}
	return nil
}
func (s *DynamicSegmentService) ChangeUserSegments(user repository.User, add []repository.Segment, delete []repository.Segment) error {
	is, err := s.repo.IsUserExists(user.Id)
	if err != nil {
		return err
	}

	if !is {
		return ErrNoUser
	}

	segmentsToAdd := make([]repository.Segment, 0, len(add))
	for _, seg := range add {
		is, err = s.repo.IsSegmentExists(seg)
		if err != nil {
			return err
		}
		if is {
			segmentsToAdd = append(segmentsToAdd, seg)
		}
	}

	for i := range segmentsToAdd {
		_, is, err := s.repo.IsSegmentUser(segmentsToAdd[i], user)
		if err != nil {
			return err
		}

		if is {
			continue
		}
		segmentsToAdd[i].Id, _ = s.repo.GetSegmentID(segmentsToAdd[i])
		err = s.repo.AddSegmentToUser(user, segmentsToAdd[i])
		if err != nil {
			return err
		}
	}

	for i := range delete {
		id, is, err := s.repo.IsSegmentUser(delete[i], user)
		if err != nil {
			return err
		}

		if !is {
			continue
		}
		delete[i].Id = id
		err = s.repo.DeleteSegmentFromUser(user, delete[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DynamicSegmentService) ShowUserSegments(user repository.User) ([]repository.Segment, error) {
	is, err := s.repo.IsUserExists(user.Id)
	if err != nil {
		return []repository.Segment{}, err
	}

	if !is {
		return []repository.Segment{}, ErrNoUser
	}

	segments, err := s.repo.GetSegmentsOfUser(user)
	if err != nil {
		return []repository.Segment{}, err
	}

	return segments, nil
}
