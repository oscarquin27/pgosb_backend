package api_models

import (
	"fdms/src/models"
	"fdms/src/utils"
)

// AntaresStationAggregationJson represents the JSON response for the antares_station query
type AntaresStationAggregationJson struct {
	AntaresId           string `json:"antares_id"`
	AntaresName         string `json:"antares_name"`
	StationId           string `json:"station_id"`
	StationName         string `json:"station_name"`
	StationAbbreviation string `json:"station_abbreviation"`
	Unharmed            string `json:"unharmed"`
	Injured             string `json:"injured"`
	Transported         string `json:"transported"`
	Deceased            string `json:"deceased"`
	Count               string `json:"count"`
}

// ModelToAntaresStationAggregationJson converts from database model to JSON model
func ModelToAntaresStationAggregationJson(s models.AntaresStationAggregation) *AntaresStationAggregationJson {
	result := AntaresStationAggregationJson{
		Count: utils.ParseInt64String(s.Count),
	}

	if s.AntaresId.Valid {
		result.AntaresId = utils.ParseInt64String(s.AntaresId.Int64)
	}

	if s.AntaresName.Valid {
		result.AntaresName = s.AntaresName.String
	}

	if s.StationId.Valid {
		result.StationId = utils.ParseInt64String(s.StationId.Int64)
	}

	if s.StationName.Valid {
		result.StationName = s.StationName.String
	}

	if s.StationAbbreviation.Valid {
		result.StationAbbreviation = s.StationAbbreviation.String
	}

	if s.Unharmed.Valid {
		result.Unharmed = utils.ParseInt64String(s.Unharmed.Int64)
	}

	if s.Injured.Valid {
		result.Injured = utils.ParseInt64String(s.Injured.Int64)
	}

	if s.Transported.Valid {
		result.Transported = utils.ParseInt64String(s.Transported.Int64)
	}

	if s.Deceased.Valid {
		result.Deceased = utils.ParseInt64String(s.Deceased.Int64)
	}

	return &result
}

// AntaresAggregationJson represents the JSON response for the antares query
type AntaresAggregationJson struct {
	AntaresId   string `json:"antares_id"`
	AntaresName string `json:"antares_name"`
	Count       string `json:"count"`
}

// ModelToAntaresAggregationJson converts from database model to JSON model
func ModelToAntaresAggregationJson(s models.AntaresAggregation) *AntaresAggregationJson {
	result := AntaresAggregationJson{
		Count: utils.ParseInt64String(s.Count),
	}

	if s.AntaresId.Valid {
		result.AntaresId = utils.ParseInt64String(s.AntaresId.Int64)
	}

	if s.AntaresName.Valid {
		result.AntaresName = s.AntaresName.String
	}

	return &result
}

// StationAggregationJson represents the JSON response for the station query
type StationAggregationJson struct {
	StationId           string `json:"station_id"`
	StationName         string `json:"station_name"`
	StationAbbreviation string `json:"station_abbreviation"`
	Count               string `json:"count"`
}

// ModelToStationAggregationJson converts from database model to JSON model
func ModelToStationAggregationJson(s models.StationAggregation) *StationAggregationJson {
	result := StationAggregationJson{
		Count: utils.ParseInt64String(s.Count),
	}

	if s.StationId.Valid {
		result.StationId = utils.ParseInt64String(s.StationId.Int64)
	}

	if s.StationName.Valid {
		result.StationName = s.StationName.String
	}

	if s.StationAbbreviation.Valid {
		result.StationAbbreviation = s.StationAbbreviation.String
	}

	return &result
}

// AntaresTypeAggregationJson represents the JSON response for the antares_type query
type AntaresTypeAggregationJson struct {
	AntaresType string `json:"antares_type"`
	Count       string `json:"count"`
}

// ModelToAntaresTypeAggregationJson converts from database model to JSON model
func ModelToAntaresTypeAggregationJson(s models.AntaresTypeAggregation) *AntaresTypeAggregationJson {
	result := AntaresTypeAggregationJson{
		Count: utils.ParseInt64String(s.Count),
	}

	if s.AntaresType.Valid {
		result.AntaresType = s.AntaresType.String
	}

	return &result
}

// Batch conversion functions for handling multiple results

func ModelToAntaresStationAggregationJsonList(models []models.AntaresStationAggregation) []AntaresStationAggregationJson {
	result := make([]AntaresStationAggregationJson, 0, len(models))
	for _, m := range models {
		jsonModel := ModelToAntaresStationAggregationJson(m)
		result = append(result, *jsonModel)
	}
	return result
}

func ModelToAntaresAggregationJsonList(models []models.AntaresAggregation) []AntaresAggregationJson {
	result := make([]AntaresAggregationJson, 0, len(models))
	for _, m := range models {
		jsonModel := ModelToAntaresAggregationJson(m)
		result = append(result, *jsonModel)
	}
	return result
}

func ModelToStationAggregationJsonList(models []models.StationAggregation) []StationAggregationJson {
	result := make([]StationAggregationJson, 0, len(models))
	for _, m := range models {
		jsonModel := ModelToStationAggregationJson(m)
		result = append(result, *jsonModel)
	}
	return result
}

func ModelToAntaresTypeAggregationJsonList(models []models.AntaresTypeAggregation) []AntaresTypeAggregationJson {
	result := make([]AntaresTypeAggregationJson, 0, len(models))
	for _, m := range models {
		jsonModel := ModelToAntaresTypeAggregationJson(m)
		result = append(result, *jsonModel)
	}
	return result
}
