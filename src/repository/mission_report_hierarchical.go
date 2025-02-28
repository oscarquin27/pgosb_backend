package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fdms/src/models"
	"fdms/src/services"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionReportHierarchicalRepository struct {
	db *pgxpool.Pool
}

func NewMissionReportHierarchicalService(db *pgxpool.Pool) services.MissionReportHierarchicalService {
	return &MissionReportHierarchicalRepository{
		db: db,
	}
}

// GetHierarchicalReport generates a hierarchical report structure with regions, stations, and missions
func (r *MissionReportHierarchicalRepository) GetHierarchicalReport(missionIds []string) (*models.HierarchicalReport, error) {
	ctx := context.Background()

	// Create the report structure
	report := &models.HierarchicalReport{
		Regions: []models.RegionWithStations{},
	}

	// Get all regions with their stations
	regions, err := r.getRegionsWithStations(ctx)
	if err != nil {
		return nil, err
	}

	// Process each region
	for _, region := range regions {
		regionWithMissions, err := r.getRegionWithMissions(ctx, region, missionIds)
		if err != nil {
			return nil, err
		}

		// Only add regions that have stations with missions
		if len(regionWithMissions.Stations) > 0 {
			report.Regions = append(report.Regions, regionWithMissions)
		}
	}

	return report, nil
}

// getRegionsWithStations fetches all regions with their stations
func (r *MissionReportHierarchicalRepository) getRegionsWithStations(ctx context.Context) ([]struct {
	RegionID   int64
	RegionName string
	Stations   []models.StationBasic
}, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	query := `
	WITH region_stations AS (
		SELECT 
			r.id AS region_id,
			r.description AS region_name,
			jsonb_agg(
				jsonb_build_object(
					'id', s.id,
					'name', COALESCE(s.description, ''),
					'abbreviation', COALESCE(s.abbreviation, '')
				)
			) FILTER (WHERE s.id IS NOT NULL) AS stations
		FROM hq.operative_regions r
		LEFT JOIN hq.stations s ON s.region_id = r.id
		GROUP BY r.id, r.description
	)
	SELECT 
		region_id,
		region_name,
		COALESCE(stations, '[]'::jsonb) as stations
	FROM region_stations
	ORDER BY region_name
	`

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regions []struct {
		RegionID   int64
		RegionName string
		Stations   []models.StationBasic
	}

	for rows.Next() {
		var regionID int64
		var regionName string
		var stationsJSON []byte

		err := rows.Scan(&regionID, &regionName, &stationsJSON)
		if err != nil {
			return nil, err
		}

		var stations []models.StationBasic
		err = json.Unmarshal(stationsJSON, &stations)
		if err != nil {
			return nil, err
		}

		regions = append(regions, struct {
			RegionID   int64
			RegionName string
			Stations   []models.StationBasic
		}{
			RegionID:   regionID,
			RegionName: regionName,
			Stations:   stations,
		})
	}

	return regions, nil
}

// getRegionWithMissions fetches missions for all stations in a region
func (r *MissionReportHierarchicalRepository) getRegionWithMissions(
	ctx context.Context,
	region struct {
		RegionID   int64
		RegionName string
		Stations   []models.StationBasic
	},
	missionIds []string,
) (models.RegionWithStations, error) {
	regionWithMissions := models.RegionWithStations{
		RegionID:   region.RegionID,
		RegionName: region.RegionName,
		Stations:   []models.StationWithMissions{},
	}

	// Process each station
	for _, station := range region.Stations {
		stationWithMissions, err := r.getStationWithMissions(ctx, station, missionIds)
		if err != nil {
			return models.RegionWithStations{}, err
		}

		// Only add stations that have missions
		if len(stationWithMissions.Missions) > 0 {
			regionWithMissions.Stations = append(regionWithMissions.Stations, stationWithMissions)
		}
	}

	return regionWithMissions, nil
}

func missionIdsToQuery(missionIds []string) string {
	return strings.Join(missionIds, ",")
}

// getStationWithMissions fetches missions for a specific station
func (r *MissionReportHierarchicalRepository) getStationWithMissions(
	ctx context.Context,
	station models.StationBasic,
	missionIds []string,
) (models.StationWithMissions, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return models.StationWithMissions{}, err
	}
	defer conn.Release()

	stationWithMissions := models.StationWithMissions{
		StationID:        station.ID,
		StationName:      station.Name,
		StationShortName: station.Abbreviation,
		Missions:         []models.MissionDetail{},
	}

	// Get missions for this station
	missionsQuery := `
	SELECT 
		m.id,
		m.manual_mission_date,
		COALESCE(m.code, '') as code,
		COALESCE(m.level, '') as level,
		m.is_important,
		COALESCE(m.peace_quadrant, '') as peace_quadrant,
		COALESCE(m.description, '') as description,
		COALESCE(m.operative_areas, ARRAY[]::varchar[]) as operative_areas,
		m.unharmed,
		m.injured,
		m.transported,
		m.deceased,
		
		-- Units
		COALESCE((
			SELECT array_agg(COALESCE(u.plate, '')) 
			FROM missions.units mu
			JOIN vehicles.unit u ON u.id = mu.unit_id
			WHERE mu.mission_id = m.id
		), ARRAY[]::varchar[]) AS units,
		
		-- Firefighters
		COALESCE((
			SELECT jsonb_agg(
				jsonb_build_object(
					'name', COALESCE(concat(u.first_name, ' ', u.last_name), ''),
					'document_id', COALESCE(u.legal_id, ''),
					'role', COALESCE(f.service_role, ''),
					'team', COALESCE(u.code, ''),
					'rank', COALESCE(u."rank", '')
					
				)
			)
			FROM missions.firefighters f
			JOIN users."user" u ON u.id = f.user_id
			WHERE f.mission_id = m.id
		), '[]'::jsonb) AS firefighters,
		
		-- First service info
		COALESCE((
			SELECT jsonb_build_object(
				'id', s.id,
				'antares_description', COALESCE(a.description, '')
			)
			FROM services.service s
			LEFT JOIN missions.antares a ON a.id = s.antares_id
			WHERE s.mission_id = m.id
			ORDER BY s.id
			LIMIT 1
		), '{}'::jsonb) AS first_service,
		
		-- Origin location
		COALESCE((
			SELECT jsonb_build_object(
				'state', COALESCE(l.state, ''),
				'municipality', COALESCE(l.municipality, ''),
				'parish', COALESCE(l.parish, ''),
				'sector', COALESCE(l.sector, ''),
				'urbanization', COALESCE(l.urb, ''),
				'address', COALESCE(l.address, '')
			)
			FROM missions.locations l
			WHERE l.id = m.location_id
		), '{}'::jsonb) AS origin_location,
		
		-- Destination location
		COALESCE((
			SELECT jsonb_build_object(
				'state', COALESCE(l.state, ''),
				'municipality', COALESCE(l.municipality, ''),
				'parish', COALESCE(l.parish, ''),
				'sector', COALESCE(l.sector, ''),
				'urbanization', COALESCE(l.urb, ''),
				'address', COALESCE(l.address, '')
			)
			FROM missions.locations l
			WHERE l.id = m.location_destiny_id
		), '{}'::jsonb) AS destination_location,
		
		-- Care center location
		COALESCE((
			SELECT jsonb_build_object(
				'state', COALESCE(c.state, ''),
				'municipality', COALESCE(c.municipality, ''),
				'parish', COALESCE(c.parish, ''),
				'sector', COALESCE(c.sector, ''),
				'urbanization', COALESCE(c.urb, ''),
				'address', COALESCE(c.address, '')
			)
			FROM hq.centers c
			WHERE c.id = m.center_id
		), '{}'::jsonb) AS carecenter_location,
		
		-- People
		COALESCE((
			SELECT jsonb_agg(
				jsonb_build_object(
					'person_state', COALESCE(p.condition, ''),
					'condition', COALESCE(p.person_condition, ''),
					'name', COALESCE(concat(p.first_name, ' ', p.last_name), ''),
					'gender', COALESCE(p.gender, ''),
					'age', CASE 
						WHEN p.age IS NULL THEN ''
						WHEN p.age::text = '' THEN ''
						ELSE p.age::text
					END,
					'document_id', COALESCE(p.legal_id, ''),
					'phone', COALESCE(p.phone, '')
					
				)
			)
			FROM missions.person p
			WHERE p.mission_id = m.id
		), '[]'::jsonb) AS people,
		
		-- Vehicles
		COALESCE((
			SELECT jsonb_agg(
				jsonb_build_object(
					'brand', COALESCE(v.make, ''),
					'model', COALESCE(v.model, ''),
					'plate', COALESCE(v.plate, ''),
					'year', CASE 
						WHEN v.year IS NULL THEN ''
						WHEN v.year::text = '' THEN ''
						ELSE v.year::text
					END,
					'color', COALESCE(v.color, '')
				)
			)
			FROM missions.vehicles v
			WHERE v.mission_id = m.id
		), '[]'::jsonb) AS vehicles,
		
		-- Infrastructures
		COALESCE((
			SELECT jsonb_agg(
				jsonb_build_object(
					'type', COALESCE(i.build_type, ''),
					'occupation', COALESCE(i.build_occupation, ''),
					'levels', CASE 
						WHEN i.levels IS NULL THEN ''
						WHEN i.levels::text = '' THEN ''
						ELSE i.levels::text
					END
				)
			)
			FROM missions.infrastructure i
			WHERE i.mission_id IN (%s)
		), '[]'::jsonb) AS infrastructures
		
	FROM missions.mission m
	WHERE m.station_id = %s
	AND m.id IN (%s)
	ORDER BY m.manual_mission_date DESC
	`

	ids := missionIdsToQuery(missionIds)
	stationId := strconv.FormatInt(station.ID, 10)
	missionsQuery = fmt.Sprintf(missionsQuery, ids, stationId, ids)

	//fmt.Println(missionsQuery)

	rows, err := conn.Query(ctx, missionsQuery)
	if err != nil {
		return models.StationWithMissions{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var mission models.MissionDetail
		var unharmed, injured, transported, deceased sql.NullInt64
		var firstServiceJSON, firefightersJSON, peopleJSON, vehiclesJSON, infrastructuresJSON []byte
		var originLocationJSON, destinationLocationJSON, carecenterLocationJSON []byte
		var units []string
		var code, level, peaceQuadrant, description sql.NullString
		var operativeAreas []sql.NullString

		err := rows.Scan(
			&mission.ID,
			&mission.Date,
			&code,
			&level,
			&mission.IsImportant,
			&peaceQuadrant,
			&description,
			&operativeAreas,
			&unharmed,
			&injured,
			&transported,
			&deceased,
			&units,
			&firefightersJSON,
			&firstServiceJSON,
			&originLocationJSON,
			&destinationLocationJSON,
			&carecenterLocationJSON,
			&peopleJSON,
			&vehiclesJSON,
			&infrastructuresJSON,
		)
		if err != nil {
			return models.StationWithMissions{}, err
		}

		// Set string fields
		mission.Code = code.String
		mission.Level = level.String
		mission.PeaceQuadrant = peaceQuadrant.String
		mission.Description = description.String

		// Convert operative areas
		mission.OperativeAreas = make([]string, 0, len(operativeAreas))
		for _, area := range operativeAreas {
			if area.Valid {
				mission.OperativeAreas = append(mission.OperativeAreas, area.String)
			}
		}

		// Set nullable fields
		if unharmed.Valid {
			mission.Unharmed = unharmed.Int64
		}
		if injured.Valid {
			mission.Injured = injured.Int64
		}
		if transported.Valid {
			mission.Transported = transported.Int64
		}
		if deceased.Valid {
			mission.Deceased = deceased.Int64
		}

		mission.Units = units

		// Unmarshal JSON fields
		if len(firstServiceJSON) > 0 {
			var firstService struct {
				ID                 int64  `json:"id"`
				AntaresDescription string `json:"antares_description"`
			}
			err = json.Unmarshal(firstServiceJSON, &firstService)
			if err == nil {
				mission.FirstServiceID = firstService.ID
				mission.FirstServiceAntaresDescription = firstService.AntaresDescription
			}
		}

		if len(firefightersJSON) > 0 {
			err = json.Unmarshal(firefightersJSON, &mission.Firefighters)
			if err != nil {
				return models.StationWithMissions{}, err
			}
		}

		if len(originLocationJSON) > 0 {
			err = json.Unmarshal(originLocationJSON, &mission.OriginLocation)
			if err != nil {
				return models.StationWithMissions{}, err
			}
		}

		if len(destinationLocationJSON) > 0 {
			err = json.Unmarshal(destinationLocationJSON, &mission.DestinationLocation)
			if err != nil {
				return models.StationWithMissions{}, err
			}
		}

		if len(carecenterLocationJSON) > 0 {
			err = json.Unmarshal(carecenterLocationJSON, &mission.CarecenterLocation)
			if err != nil {
				return models.StationWithMissions{}, err
			}
		}

		if len(peopleJSON) > 0 {
			err = json.Unmarshal(peopleJSON, &mission.People)
			if err != nil {
				return models.StationWithMissions{}, err
			}
		}

		if len(vehiclesJSON) > 0 {
			err = json.Unmarshal(vehiclesJSON, &mission.Vehicles)
			if err != nil {
				return models.StationWithMissions{}, err
			}
		}

		if len(infrastructuresJSON) > 0 {
			err = json.Unmarshal(infrastructuresJSON, &mission.Infrastructures)
			if err != nil {
				return models.StationWithMissions{}, err
			}
		}

		stationWithMissions.Missions = append(stationWithMissions.Missions, mission)
	}

	return stationWithMissions, nil
}
