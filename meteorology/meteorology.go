package meteorology

import "fmt"

type TemperatureUnit int

func (tu TemperatureUnit) String() string {
	if tu == 0 {
		return "°C"
	}

	return "°F"
}

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t *Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit.String())
}

// Add a String method to the Temperature type

type SpeedUnit int

func (s SpeedUnit) String() string {
	if s == 0 {
		return "km/h"
	}

	return "mph"
}

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit.String())
}

// Add a String method to Speed

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m *MeteorologyData) String() string {
	return fmt.Sprintf(
		"%s: %s, Wind %s at %s, %d%% Humidity",
		m.location,
		m.temperature.String(),
		m.windDirection,
		m.windSpeed.String(),
		m.humidity,
	)
}

// Add a String method to MeteorologyData
