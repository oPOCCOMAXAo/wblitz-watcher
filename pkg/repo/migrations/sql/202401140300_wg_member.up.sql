CREATE TABLE wg_clan_member (
    region ENUM('asia', 'eu', 'na', 'ru') NOT NULL,
    clan_id BIGINT NOT NULL,
    account_id BIGINT NOT NULL,
    PRIMARY KEY (region, clan_id, account_id)
);
CREATE TABLE event_clan (
    id BIGINT NOT NULL AUTO_INCREMENT,
    time BIGINT NOT NULL,
    type ENUM('enter', 'leave') NOT NULL,
    region ENUM('asia', 'eu', 'na', 'ru') NOT NULL,
    clan_id BIGINT NOT NULL,
    account_id BIGINT NOT NULL,
    is_processed TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
