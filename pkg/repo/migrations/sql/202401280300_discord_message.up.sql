CREATE TABLE discord_message (
    id BIGINT NOT NULL AUTO_INCREMENT,
    is_processed TINYINT(1) NOT NULL DEFAULT 0,
    event_clan_id BIGINT,
    bot_instance_id BIGINT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (event_clan_id) REFERENCES event_clan (id),
    FOREIGN KEY (bot_instance_id) REFERENCES bot_instance (id),
    UNIQUE KEY (event_clan_id, bot_instance_id)
);
