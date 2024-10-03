package main

import (
	"fmt"
	"time"
)

type Movie interface {
	GetTitle() string
	GetType() string
}

type RegularMovie struct {
	title string
}

//Interfaces

func (m *RegularMovie) GetTitle() string {
	return m.title
}

func (m *RegularMovie) GetType() string {
	return "Regular"
}

type IMAXMovie struct {
	title string
}

func (m *IMAXMovie) GetTitle() string {
	return m.title
}

func (m *IMAXMovie) GetType() string {
	return "IMAX"
}

type _3DMovie struct {
	title string
}

func (m *_3DMovie) GetTitle() string {
	return m.title
}

func (m *_3DMovie) GetType() string {
	return "3D"
}

// Fabric
type MovieFactory interface {
	CreateMovie(title string) Movie
}

type RegularMovieFactory struct{}

func (f *RegularMovieFactory) CreateMovie(title string) Movie {
	return &RegularMovie{title: title}
}

type IMAXMovieFactory struct{}

func (f *IMAXMovieFactory) CreateMovie(title string) Movie {
	return &IMAXMovie{title: title}
}

type _3DMovieFactory struct{}

func (f *_3DMovieFactory) CreateMovie(title string) Movie {
	return &_3DMovie{title: title}
}

// Abstract fabric
type Button interface {
	Render()
}

type DarkThemeButton struct{}

func (b *DarkThemeButton) Render() {
	println("Rendering Dark Theme Button")
}

type LightThemeButton struct{}

func (b *LightThemeButton) Render() {
	println("Rendering Light Theme Button")
}

type UIFactory interface {
	CreateButton() Button
}

type DarkThemeFactory struct{}

func (f *DarkThemeFactory) CreateButton() Button {
	return &DarkThemeButton{}
}

type LightThemeFactory struct{}

func (f *LightThemeFactory) CreateButton() Button {
	return &LightThemeButton{}
}

// Builder
type TicketBooking struct {
	movieTitle string
	seatNumber string
	snackCombo string
}

type TicketBookingBuilder struct {
	booking TicketBooking
}

func NewTicketBookingBuilder() *TicketBookingBuilder {
	return &TicketBookingBuilder{}
}

func (b *TicketBookingBuilder) SetMovieTitle(title string) *TicketBookingBuilder {
	b.booking.movieTitle = title
	return b
}

func (b *TicketBookingBuilder) SetSeatNumber(seat string) *TicketBookingBuilder {
	b.booking.seatNumber = seat
	return b
}

func (b *TicketBookingBuilder) SetSnackCombo(snack string) *TicketBookingBuilder {
	b.booking.snackCombo = snack
	return b
}

func (b *TicketBookingBuilder) Build() TicketBooking {
	return b.booking
}

// Prototype
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

// Singleton
type CinemaConfig struct {
	cinemaName string
}

var instance *CinemaConfig

func GetInstance() *CinemaConfig {
	if instance == nil {
		instance = &CinemaConfig{}
	}
	return instance
}

func (c *CinemaConfig) SetCinemaName(name string) {
	c.cinemaName = name
}

func (c *CinemaConfig) GetCinemaName() string {
	return c.cinemaName
}

// Main func
func main() {

	config := GetInstance()
	config.SetCinemaName("Starlight Cinemas")
	fmt.Println("Cinema Name:", config.GetCinemaName())

	regularFactory := &RegularMovieFactory{}
	movie := regularFactory.CreateMovie("Inception")
	fmt.Printf("Movie: %s, Type: %s\n", movie.GetTitle(), movie.GetType())

	uiFactory := &DarkThemeFactory{}
	button := uiFactory.CreateButton()
	button.Render()

	bookingBuilder := NewTicketBookingBuilder().
		SetMovieTitle("Inception").
		SetSeatNumber("A1").
		SetSnackCombo("Popcorn and Soda")

	booking := bookingBuilder.Build()

	fmt.Printf("Booking for movie: %s, Seat: %s, Snack: %s\n", booking.movieTitle, booking.seatNumber, booking.snackCombo)

	template := &StandardSchedule{}
	template.SetTime(time.Now())
	eveningSchedule := template.Clone()
	eveningSchedule.SetMovie(movie)

	fmt.Println("Cinema Management System initialized.")
}
