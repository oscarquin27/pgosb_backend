

select antares_id , antares_name, station_id, station_name, station_abbreviation, count(1) from missions.vw_mission_report_summary
	where antares_id is not null 
		and mission_id in (
			20,21,22,23,24,25
		)
	group by (antares_id, antares_name, station_id, station_name, station_abbreviation)

    
select antares_id , antares_name, count(1) from missions.vw_mission_report_summary
	where antares_id is not null
		and mission_id in (
			20,21,22,23,24,25
		)
	group by (antares_id, antares_name)

select station_id, station_abbreviation, station_name, count(1) from missions.vw_mission_report_summary
	where station_id is not null
		and mission_id in (
			20,21,22,23,24,25
		)
	group by (station_id,station_abbreviation,station_name)

select antares_type, count(1) from missions.vw_mission_report_summary
	where mission_id in (
			20,21,22,23,24,25
		)
 group by (antares_type)



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
					'team', COALESCE(u.code, '')
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
				'urbanization', COALESCE(l.urb, '')
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
				'urbanization', COALESCE(l.urb, '')
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
				'urbanization', COALESCE(c.urb, '')
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
					'phone', COALESCE(p.phone, ''),
					'transfer_vehicle', ''
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
			WHERE i.mission_id = m.id
		), '[]'::jsonb) AS infrastructures
		
	FROM missions.mission m
	WHERE m.station_id = $1
	AND ($2::text IS NULL OR m.manual_mission_date >= NULLIF($2, '')::timestamp)
	AND ($3::text IS NULL OR m.manual_mission_date <= NULLIF($3, '')::timestamp)
	ORDER BY m.manual_mission_date DESC