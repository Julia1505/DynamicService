package handler

import "UserSegmentationService/internal/repository"

type SegmentName string

type RequestSegment SegmentName

type RequestChange struct {
	Add    []SegmentName `json:"add"`
	Delete []SegmentName `json:"delete"`
}

func (r *RequestChange) ToAddSegment() []repository.Segment {
	segments := make([]repository.Segment, 0, len(r.Add))
	for _, segment := range r.Add {
		segments = append(segments, repository.Segment{Name: string(segment)})
	}
	return segments
}

func (r *RequestChange) ToDeleteSegment() []repository.Segment {
	segments := make([]repository.Segment, 0, len(r.Delete))
	for _, segment := range r.Delete {
		segments = append(segments, repository.Segment{Name: string(segment)})
	}
	return segments
}
