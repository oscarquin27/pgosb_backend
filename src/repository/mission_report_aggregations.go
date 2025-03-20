package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionReportAggregationsRepository struct {
	db *pgxpool.Pool
}

func NewMissionReportAggregationsService(db *pgxpool.Pool) services.MissionReportAggregationsService {
	return &MissionReportAggregationsRepository{
		db: db,
	}
}

// createInClauseWithParams creates an IN clause with numbered parameters
func createInClauseWithParams(startIdx int, count int) string {
	params := make([]string, count)
	for i := 0; i < count; i++ {
		params[i] = fmt.Sprintf("$%d", startIdx+i)
	}
	return fmt.Sprintf("(%s)", strings.Join(params, ","))
}

func stringSliceToInterfaceSlice(ids []string) []interface{} {
	result := make([]interface{}, len(ids))
	for i, id := range ids {
		result[i] = id
	}
	return result
}

// int64SliceToInterfaceSlice converts a slice of int64 to a slice of interface{}
// func int64SliceToInterfaceSlice(ids []int64) []interface{} {
// 	result := make([]interface{}, len(ids))
// 	for i, id := range ids {
// 		result[i] = id
// 	}
// 	return result
// }

// GetAntaresStationByMissionIds implements the antares_station query
// select antares_id, antares_name, station_id, station_name, station_abbreviation, count(1) from missions.vw_mission_report_summary
func (r *MissionReportAggregationsRepository) GetAntaresStationByMissionIds(missionIds []string) ([]models.AntaresStationAggregation, error) {
	ctx := context.Background()

	conn, err := r.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	inClause := createInClauseWithParams(1, len(missionIds))
	query := fmt.Sprintf(`
	SELECT antares_id, antares_name, station_id, station_name, station_abbreviation, unharmed, injured, transported, deceased, count(1), municipality_origin, parish_origin, municipality_destiny, parish_destiny
	FROM missions.vw_mission_report_summary
	WHERE antares_id IS NOT NULL 
		AND mission_id IN %s
	GROUP BY (antares_id, antares_name, station_id, station_name, station_abbreviation, unharmed, injured, transported, deceased, municipality_origin, parish_origin, municipality_destiny, parish_destiny)
	`, inClause)

	args := stringSliceToInterfaceSlice(missionIds)
	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.AntaresStationAggregation])
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetAntaresByMissionIds implements the antares query
// select antares_id, antares_name, count(1) from missions.vw_mission_report_summary
func (r *MissionReportAggregationsRepository) GetAntaresByMissionIds(missionIds []string) ([]models.AntaresAggregation, error) {
	ctx := context.Background()

	conn, err := r.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	inClause := createInClauseWithParams(1, len(missionIds))
	query := fmt.Sprintf(`
	SELECT antares_id, antares_name, count(1)
	FROM missions.vw_mission_report_summary
	WHERE antares_id IS NOT NULL
		AND mission_id IN %s
	GROUP BY (antares_id, antares_name)
	`, inClause)

	args := stringSliceToInterfaceSlice(missionIds)
	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.AntaresAggregation])
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *MissionReportAggregationsRepository) GetMunicipalityOriginByMissionIds(missionIds []string) ([]models.MunicipalityOriginAggregation, error) {
	ctx := context.Background()

	conn, err := r.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	inClause := createInClauseWithParams(1, len(missionIds))
	query := fmt.Sprintf(`
	SELECT municipality_origin, count(1)
	FROM missions.vw_mission_report_summary
	WHERE antares_id IS NOT NULL
		AND mission_id IN %s
	GROUP BY (municipality_origin)
	`, inClause)

	args := stringSliceToInterfaceSlice(missionIds)
	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MunicipalityOriginAggregation])
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *MissionReportAggregationsRepository) GetParishOriginByMissionIds(missionIds []string) ([]models.ParishOriginAggregation, error) {
	ctx := context.Background()

	conn, err := r.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	inClause := createInClauseWithParams(1, len(missionIds))
	query := fmt.Sprintf(`
	SELECT parish_origin, count(1)
	FROM missions.vw_mission_report_summary
	WHERE antares_id IS NOT NULL
		AND mission_id IN %s
	GROUP BY (parish_origin)
	`, inClause)

	args := stringSliceToInterfaceSlice(missionIds)
	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ParishOriginAggregation])
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetStationByMissionIds implements the station query
// select station_id, station_abbreviation, station_name, count(1) from missions.vw_mission_report_summary
func (r *MissionReportAggregationsRepository) GetStationByMissionIds(missionIds []string) ([]models.StationAggregation, error) {
	ctx := context.Background()

	conn, err := r.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	inClause := createInClauseWithParams(1, len(missionIds))
	query := fmt.Sprintf(`
	SELECT station_id, station_abbreviation, station_name, count(1)
	FROM missions.vw_mission_report_summary
	WHERE station_id IS NOT NULL
		AND mission_id IN %s
	GROUP BY (station_id, station_abbreviation, station_name)
	`, inClause)

	args := stringSliceToInterfaceSlice(missionIds)
	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.StationAggregation])
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetAntaresTypeByMissionIds implements the antares_type query
// select antares_type, count(1) from missions.vw_mission_report_summary
func (r *MissionReportAggregationsRepository) GetAntaresTypeByMissionIds(missionIds []string) ([]models.AntaresTypeAggregation, error) {
	ctx := context.Background()

	conn, err := r.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	inClause := createInClauseWithParams(1, len(missionIds))
	query := fmt.Sprintf(`
	SELECT antares_type, count(1)
	FROM missions.vw_mission_report_summary
	WHERE mission_id IN %s
	GROUP BY (antares_type)
	`, inClause)

	args := stringSliceToInterfaceSlice(missionIds)
	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.AntaresTypeAggregation])
	if err != nil {
		return nil, err
	}

	return results, nil
}
