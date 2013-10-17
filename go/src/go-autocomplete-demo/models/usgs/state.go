package usgs

type (
	StateId int
	State   struct {
		Id   StateId
		Code string
	}

	CountyId string
	County   struct {
		Id      CountyId
		Name    string
		StateId StateId
	}

	PopulatedPlaceId int
	PopulatedPlace   struct {
		Id        PopulatedPlaceId
		Name      string
		Latitude  float64
		Longitude float64
		StateId   StateId
		CountyId  CountyId
	}
)
