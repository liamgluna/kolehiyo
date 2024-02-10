CREATE TABLE IF NOT EXISTS universities (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) WITH time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    founded date NOT NULL,
    location text NOT NULL,
    campuses text [],
    website text NOT NULL,
    version integer NOT NULL DEFAULT 1
);