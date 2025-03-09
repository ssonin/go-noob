package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tu TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%v %s", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
func (md MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %v%% Humidity", md.location, md.temperature, md.windDirection, md.windSpeed, md.humidity)
}
