package main

type Movie interface {
	GetTitle() string
	GetType() string
}

/////

type RegularMovie struct {
	title string
}

func (m *RegularMovie) GetTitle() string {
	return m.title
}

func (m *RegularMovie) GetType() string {
	return "Regular"
}

// ///
type IMAXMovie struct {
	title string
}

func (m *IMAXMovie) GetTitle() string {
	return m.title
}

func (m *IMAXMovie) GetType() string {
	return "IMAX"
}

// //
type _3DMovie struct {
	title string
}

func (m *_3DMovie) GetTitle() string {
	return m.title
}

func (m *_3DMovie) GetType() string {
	return "IMAX"
}

type MovieFactory interface {
	CreateMovie(title string) Movie
}

type RegularMovieFactory struct{}

func (f *RegularMovieFactory) CreateMovie(title string) Movie {
	return &RegularMovie{title: title}
}

// //
type IMAXMovieFactory struct{}

func (f *IMAXMovieFactory) CreateMovie(title string) Movie {
	return &IMAXMovie{title: title}
}

////

type _3DMovieFactory struct{}

func (f *_3DMovieFactory) CreateMovie(title string) Movie {
	return &_3DMovie{title: title}
}
