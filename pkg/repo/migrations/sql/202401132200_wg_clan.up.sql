CREATE TABLE wg_clan (
    id BIGINT NOT NULL,
    region ENUM('asia', 'eu', 'na', 'ru') NOT NULL,
    tag VARCHAR(16) NOT NULL DEFAULT '',
    name VARCHAR(255) NOT NULL DEFAULT '',
    updated_at BIGINT NOT NULL DEFAULT 0,
    members_updated_at BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY (id, region)
);
