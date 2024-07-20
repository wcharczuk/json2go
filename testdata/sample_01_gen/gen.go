package autogen

// AutoGenerated is an autogenerated type from JSON.
type AutoGenerated struct {
	Context    []any                  `json:"@context,omitempty"`
	Geometry   AutoGenerated_Geometry `json:"geometry,omitempty"`
	ID         string                 `json:"id,omitempty"`
	Properties struct {
		ID                  string  `json:"@id,omitempty"`
		Type                string  `json:"@type,omitempty"`
		County              string  `json:"county,omitempty"`
		Cwa                 string  `json:"cwa,omitempty"`
		FireWeatherZone     string  `json:"fireWeatherZone,omitempty"`
		Forecast            string  `json:"forecast,omitempty"`
		ForecastGridData    string  `json:"forecastGridData,omitempty"`
		ForecastHourly      string  `json:"forecastHourly,omitempty"`
		ForecastOffice      string  `json:"forecastOffice,omitempty"`
		ForecastZone        string  `json:"forecastZone,omitempty"`
		GridID              string  `json:"gridId,omitempty"`
		GridX               float64 `json:"gridX,omitempty"`
		GridY               float64 `json:"gridY,omitempty"`
		ObservationStations string  `json:"observationStations,omitempty"`
		RadarStation        string  `json:"radarStation,omitempty"`
		RelativeLocation    struct {
			Geometry   AutoGenerated_Geometry `json:"geometry,omitempty"`
			Properties struct {
				Bearing  AutoGenerated_Properties_RelativeLocation_Properties_Bearing `json:"bearing,omitempty"`
				City     string                                                       `json:"city,omitempty"`
				Distance AutoGenerated_Properties_RelativeLocation_Properties_Bearing `json:"distance,omitempty"`
				State    string                                                       `json:"state,omitempty"`
			} `json:"properties,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"relativeLocation,omitempty"`
		TimeZone string `json:"timeZone,omitempty"`
	} `json:"properties,omitempty"`
	Type string `json:"type,omitempty"`
}

// AutoGenerated_Geometry is an aliased type from JSON.
type AutoGenerated_Geometry struct {
	Coordinates []float64 `json:"coordinates,omitempty"`
	Type        string    `json:"type,omitempty"`
}

// AutoGenerated_Properties_RelativeLocation_Properties_Bearing is an aliased type from JSON.
type AutoGenerated_Properties_RelativeLocation_Properties_Bearing struct {
	UnitCode string  `json:"unitCode,omitempty"`
	Value    float64 `json:"value,omitempty"`
}
