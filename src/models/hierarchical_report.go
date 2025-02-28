package models

import "time"

// HierarchicalReport represents the top-level structure of the hierarchical report
type HierarchicalReport struct {
	Regions []RegionWithStations `json:"regions"`
}

// RegionWithStations represents a region with its stations
type RegionWithStations struct {
	RegionID   int64                 `json:"region_id"`
	RegionName string                `json:"region_name"`
	Stations   []StationWithMissions `json:"stations"`
}

// StationBasic represents basic station information
type StationBasic struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

// StationWithMissions represents a station with its missions
type StationWithMissions struct {
	StationID        int64           `json:"station_id"`
	StationName      string          `json:"station_name"`
	StationShortName string          `json:"station_short_name"`
	Missions         []MissionDetail `json:"missions"`
}

// MissionDetail represents detailed mission information
type MissionDetail struct {
	ID                             int64            `json:"id"`
	Date                           time.Time        `json:"date"`
	Code                           string           `json:"code"`
	Level                          string           `json:"level"`
	IsImportant                    bool             `json:"is_important"`
	PeaceQuadrant                  string           `json:"peace_quadrant"`
	Description                    string           `json:"description"`
	OperativeAreas                 []string         `json:"operative_areas"`
	Unharmed                       int64            `json:"unharmed"`
	Injured                        int64            `json:"injured"`
	Transported                    int64            `json:"transported"`
	Deceased                       int64            `json:"deceased"`
	Units                          []string         `json:"units"`
	Firefighters                   []Firefighter    `json:"firefighters"`
	FirstServiceID                 int64            `json:"first_service_id"`
	FirstServiceAntaresDescription string           `json:"first_service_antares_description"`
	OriginLocation                 Location         `json:"origin_location"`
	DestinationLocation            Location         `json:"destination_location"`
	CarecenterLocation             Location         `json:"carecenter_location"`
	People                         []Person         `json:"people"`
	Vehicles                       []VehicleReport  `json:"vehicles"`
	Infrastructures                []Infrastructure `json:"infrastructures"`
}

// Firefighter represents firefighter information
type Firefighter struct {
	Name       string `json:"name"`
	DocumentID string `json:"document_id"`
	Role       string `json:"role"`
	Team       string `json:"team"`
	Rank       string `json:"rank"`
}

// Location represents location information
type Location struct {
	State        string `json:"state"`
	Municipality string `json:"municipality"`
	Parish       string `json:"parish"`
	Sector       string `json:"sector"`
	Urbanization string `json:"urbanization"`
	Address      string `json:"address"`
}

// Person represents person information
type Person struct {
	PersonState     string `json:"person_state"`
	Condition       string `json:"condition"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	Age             string `json:"age"`
	DocumentID      string `json:"document_id"`
	Phone           string `json:"phone"`
	TransferVehicle string `json:"transfer_vehicle"`
}

// Vehicle represents vehicle information
type VehicleReport struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Plate string `json:"plate"`
	Year  string `json:"year"`
	Color string `json:"color"`
}

// Infrastructure represents infrastructure information
type Infrastructure struct {
	Type       string `json:"type"`
	Occupation string `json:"occupation"`
	Levels     string `json:"levels"`
}
