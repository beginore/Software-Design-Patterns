package main

//// Singleton --- only 1 instance of CinemaConfig()
//CinemaConfig class

type CinemaConfig struct {
	cinemaName      string
	numberOfScreens int
	operatingHours  string
}

var instance *CinemaConfig

func GetInstance() *CinemaConfig {
	if instance == nil {
		instance = &CinemaConfig{
			cinemaName:      "Default Cinema",
			numberOfScreens: 5,
			operatingHours:  "09:00 - 22:00",
		}
	}
	return instance
}

// setters getters

func (c *CinemaConfig) GetCinemaName() string {
	return c.cinemaName
}

func (c *CinemaConfig) SetCinemaName(name string) {
	c.cinemaName = name
}

func (c *CinemaConfig) GetNumberOfScreens() int {
	return c.numberOfScreens
}

func (c *CinemaConfig) SetNumberOfScreens(count int) {
	c.numberOfScreens = count
}

func (c *CinemaConfig) GetOperatingHours() string {
	return c.operatingHours
}

func (c *CinemaConfig) SetOperatingHours(hours string) {
	c.operatingHours = hours
}
