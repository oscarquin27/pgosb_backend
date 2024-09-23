CREATE TABLE IF NOT EXISTS missions.authorities_person (
    id bigint NOT NULL DEFAULT nextval('missions.authorities_person_id_seq' :: regclass),
    mission_id bigint NOT NULL,
    authority_id bigint NOT NULL,
    name character varying,
    last_name character varying,
    legal_id character varying,
    identification_number character varying,
    phone character varying,
    gender character varying,
    observations text,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT authorities_person_pkey PRIMARY KEY (id),
    CONSTRAINT fk_mission FOREIGN KEY (mission_id) REFERENCES missions.missions (id) ON DELETE CASCADE,
    CONSTRAINT fk_authority FOREIGN KEY (authority_id) REFERENCES missions.authorities (id) ON DELETE CASCADE
);

-- Create an index on mission_id for better query performance
CREATE INDEX idx_authorities_person_mission_id ON missions.authorities_person(mission_id);

-- Create an index on authority_id for better query performance
CREATE INDEX idx_authorities_person_authority_id ON missions.authorities_person(authority_id);

CREATE TABLE IF NOT EXISTS missions.authorities_vehicle (
    id bigint NOT NULL DEFAULT nextval(
        'missions.authorities_vehicle_id_seq' :: regclass
    ),
    mission_id bigint NOT NULL,
    authority_id bigint NOT NULL,
    type character varying,
    make character varying,
    model character varying,
    plate character varying,
    year character varying,
    color character varying,
    description text,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT authorities_vehicle_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS missions.authorities (
    id bigint NOT NULL DEFAULT nextval(
        'missions.mission_authorities_id_seq' :: regclass
    ),
    mission_id bigint NOT NULL,
    type character varying COLLATE pg_catalog."default",
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    alias character varying COLLATE pg_catalog."default",
    CONSTRAINT mission_authorities_pkey PRIMARY KEY (id)
) CREATE
OR REPLACE VIEW missions.mission_authority_summary AS
SELECT
    a.id,
    a.mission_id,
    a.created_at,
    a.alias,
    COALESCE(v.vehicle_count, 0) AS vehicles,
    COALESCE(p.person_count, 0) AS people,
    a.type
FROM
    missions.authorities a
    LEFT JOIN (
        SELECT
            authority_id,
            COUNT(*) AS vehicle_count
        FROM
            missions.authorities_vehicle
        GROUP BY
            authority_id
    ) v ON a.id = v.authority_id
    LEFT JOIN (
        SELECT
            authority_id,
            COUNT(*) AS person_count
        FROM
            missions.authorities_person
        GROUP BY
            authority_id
    ) p ON a.id = p.authority_id;

CREATE TABLE IF NOT EXISTS missions.mission_authority_service (
    id bigint NOT NULL DEFAULT nextval(
        'missions.mission_authority_service_id_seq' :: regclass
    ),
    mission_id bigint NOT NULL,
    service_id bigint NOT NULL,
    authority_id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT mission_authority_service_pkey PRIMARY KEY (id),
    CONSTRAINT fk_mission FOREIGN KEY (mission_id) REFERENCES missions.missions (id) ON DELETE CASCADE,
    CONSTRAINT fk_service FOREIGN KEY (service_id) REFERENCES missions.services (id) ON DELETE CASCADE,
    CONSTRAINT fk_authority FOREIGN KEY (authority_id) REFERENCES missions.authorities (id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX idx_mission_authority_service_mission_id ON missions.mission_authority_service(mission_id);

CREATE INDEX idx_mission_authority_service_service_id ON missions.mission_authority_service(service_id);

CREATE INDEX idx_mission_authority_service_authority_id ON missions.mission_authority_service(authority_id);