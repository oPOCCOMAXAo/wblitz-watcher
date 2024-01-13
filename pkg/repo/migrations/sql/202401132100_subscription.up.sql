CREATE TABLE subscription_clan (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    instance_id BIGINT NOT NULL,
    clan_id BIGINT NOT NULL,
    region ENUM('asia', 'eu', 'na', 'ru') NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    UNIQUE KEY (instance_id, clan_id, region),
    FOREIGN KEY (instance_id) REFERENCES bot_instance (id) ON DELETE CASCADE
);
