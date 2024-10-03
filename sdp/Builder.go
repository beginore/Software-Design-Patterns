package main

type TicketBooking struct {
	movieTitle string
	seatNumber string
	snackCombo string
}

type TicketBookingBuilder struct {
	booking TicketBooking
}

func (b TicketBookingBuilder) SetMovieTitle(title string) *TicketBookingBuilder {
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
