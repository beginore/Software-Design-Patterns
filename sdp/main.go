package main

import (
	"fmt"
	"time"
)

func main() {
	// Singleton
	config := GetInstance()
	config.SetCinemaName("Starlight Cinemas")
	fmt.Println("Cinema Name:", config.GetCinemaName())

	// Fabric
	regularFactory := &RegularMovieFactory{}
	movie := regularFactory.CreateMovie("Inception")
	fmt.Printf("Movie: %s, Type: %s\n", movie.GetTitle(), movie.GetType())

	// Abs. Fabric
	uiFactory := &DarkThemeFactory{}
	button := uiFactory.CreateButton()
	button.Render()

	booking := TicketBookingBuilder{}.
		SetMovieTitle("Inception").
		SetSeatNumber("A1").
		SetSnackCombo("Popcorn and Soda").
		Build()

	fmt.Printf("Booking for movie: %s, Seat: %s, Snack: %s\n", booking.movieTitle, booking.seatNumber, booking.snackCombo)

	// Паттерн Прототип
	template := &StandardSchedule{}
	template.SetTime(time.Now())
	eveningSchedule := template.Clone()
	eveningSchedule.SetMovie(movie)
	fmt.Println("Cinema Management System initialized.")
}
