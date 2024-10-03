package main

import "time"

type MovieSchedule interface {
	Clone() MovieSchedule
	SetMovie(movie Movie)
	SetTime(t time.Time)
}

type StandardSchedule struct {
	movie Movie
	time  time.Time
}

func (s *StandardSchedule) Clone() MovieSchedule {
	return &StandardSchedule{
		movie: s.movie,
		time:  s.time,
	}
}

func (s *StandardSchedule) SetMovie(movie Movie) {
	s.movie = movie
}

func (s *StandardSchedule) SetTime(t time.Time) {
	s.time = t
}
