package main

type LocationDetails struct {
	Distance    float32  `json:"distance,omitempty"`
	PlaceId     string   `json:"place_id,omitempty"`
	Licence     string   `json:"licence,omitempty"`
	OsmType     string   `json:"osm_type,omitempty"`
	OsmId       string   `json:"osm_id,omitempty"`
	Boundingbox []string `json:"boundingbox,omitempty"`
	Lat         string   `json:"lat,omitempty"`
	Lon         string   `json:"lon,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	Class       string   `json:"class,omitempty"`
	Type        string   `json:"type,omitempty"`
	Importance  float32  `json:"importance,omitempty"`
	Address     Address  `json:"address,omitempty"`
}
