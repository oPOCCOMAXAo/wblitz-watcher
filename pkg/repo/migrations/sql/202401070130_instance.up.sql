CREATE TABLE bot_instance (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    server_id VARCHAR(64) NOT NULL,
    channel_id VARCHAR(64) NOT NULL,
    type ENUM('clan') NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    UNIQUE KEY (server_id, type)
);
