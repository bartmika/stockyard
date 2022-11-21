-- +goose Up

CREATE TABLE IF NOT EXISTS entities (
    id BIGSERIAL PRIMARY KEY,
    uuid VARCHAR (36) NOT NULL,
    name TEXT NOT NULL DEFAULT '',
    data_type SMALLINT NOT NULL DEFAULT 0,
    meta TEXT NOT NULL DEFAULT ''
);

SELECT create_distributed_table('entities', 'id');

CREATE TABLE IF NOT EXISTS observations (
    entity_id BIGINT NOT NULL,
    meta TEXT NOT NULL DEFAULT '',
    timestamp TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
    value FLOAT NULL DEFAULT 0,
    FOREIGN KEY (entity_id) REFERENCES entities(id) ON DELETE CASCADE,
    PRIMARY KEY (entity_id, timestamp)
);

SELECT create_distributed_table('observations', 'entity_id');

-- DEVELOPERS NOTE:
-- Ideally we want columnar storage for the storage benefits, however currently `citus` has the following limitations:
-- (1) Append-only (no UPDATE/DELETE support)
-- (2) No support for foreign keys, unique constraints, or exclusion constraints
-- As per developer documentation via https://docs.citusdata.com/en/v11.1/admin_guide/table_management.html#limitations.
-- As a result we'll leave this code commented out and uncomment it in the future.
-- SELECT alter_table_set_access_method('observations', 'columnar');

CREATE TABLE IF NOT EXISTS observation_analyzer_requests (
    entity_id BIGINT NOT NULL,
    uuid VARCHAR (36) NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
    type SMALLINT NOT NULL DEFAULT 0,
    observation TEXT NULL DEFAULT '',
    FOREIGN KEY (entity_id) REFERENCES entities(id) ON DELETE CASCADE,
    PRIMARY KEY (entity_id, uuid)
);
SELECT create_distributed_table('observation_analyzer_requests', 'entity_id');

CREATE TABLE IF NOT EXISTS timekeys (
    entity_id BIGINT NOT NULL,
    meta TEXT NOT NULL DEFAULT '',
    timestamp TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
    value TEXT NULL DEFAULT '',
    FOREIGN KEY (entity_id) REFERENCES entities(id) ON DELETE CASCADE,
    PRIMARY KEY (entity_id, timestamp)
);

SELECT create_distributed_table('timekeys', 'entity_id');
-- DEVELOPERS NOTE: Do not make `columnar` because we want the ability to delete and citus does not support it.

CREATE TABLE IF NOT EXISTS observation_counts (
    entity_id BIGINT NOT NULL,
    start TIMESTAMPTZ NOT NULL,
    finish TIMESTAMPTZ NOT NULL,
    day INT NOT NULL,
    week INT NOT NULL,
    month INT NOT NULL,
    year INT NOT NULL,
    frequency SMALLINT NOT NULL,
    result DOUBLE PRECISION NOT NULL,
    FOREIGN KEY (entity_id) REFERENCES entities(id) ON DELETE CASCADE,
    PRIMARY KEY (entity_id, frequency, start, finish)
);

SELECT create_distributed_table('observation_counts', 'entity_id');
-- Do not use "columnar" on `observation_counts` b/c currently `update` not supported by citus, will be supported in future, please read: https://docs.citusdata.com/en/v11.1/admin_guide/table_management.html#limitations

CREATE TABLE IF NOT EXISTS observation_summations (
    entity_id BIGINT NOT NULL,
    start TIMESTAMPTZ NOT NULL,
    finish TIMESTAMPTZ NOT NULL,
    day INT NOT NULL,
    week INT NOT NULL,
    month INT NOT NULL,
    year INT NOT NULL,
    frequency SMALLINT NOT NULL,
    result DOUBLE PRECISION NOT NULL,
    FOREIGN KEY (entity_id) REFERENCES entities(id) ON DELETE CASCADE,
    PRIMARY KEY (entity_id, frequency, start, finish)
);

SELECT create_distributed_table('observation_summations', 'entity_id');
-- Do not use "columnar" on `observation_summations` b/c currently `update` not supported by citus, will be supported in future, please read: https://docs.citusdata.com/en/v11.1/admin_guide/table_management.html#limitations

CREATE TABLE IF NOT EXISTS observation_averages (
    entity_id BIGINT NOT NULL,
    start TIMESTAMPTZ NOT NULL,
    finish TIMESTAMPTZ NOT NULL,
    day INT NOT NULL,
    week INT NOT NULL,
    month INT NOT NULL,
    year INT NOT NULL,
    frequency SMALLINT NOT NULL,
    result DOUBLE PRECISION NOT NULL,
    FOREIGN KEY (entity_id) REFERENCES entities(id) ON DELETE CASCADE,
    PRIMARY KEY (entity_id, frequency, start, finish)
);

SELECT create_distributed_table('observation_averages', 'entity_id');
-- Do not use "columnar" on `observation_averages` b/c currently `update` not supported by citus, will be supported in future, please read: https://docs.citusdata.com/en/v11.1/admin_guide/table_management.html#limitations


-- +goose Down

DROP TABLE observation_averages CASCADE;
DROP TABLE observation_counts CASCADE;
DROP TABLE observation_summations CASCADE;
DROP TABLE timekeys CASCADE;
DROP TABLE observations CASCADE;
DROP TABLE entities CASCADE;
