CREATE TABLE strategies (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    user_id BIGINT REFERENCES users (id) ON DELETE CASCADE,
    instrument_type SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO strategies 
(name, instrument_type)
VALUES 
('высокочастотный трейдинг', 1), 
('Скальпинг', 1), 
('Дейтрейдинг', 1), 
('Свинг трейдинг', 1), 
('Позиционная торговля', 1), 
('Торговля на новостях', 1), 
('Стратегии трейдинга', 1), 
('Торговля по тренду', 1), 
('Торговля в консолидации', 1), 
('Торговля на пробой', 1), 
('Торговля на разворот', 1), 
('Макро-экономический трейдинг', 1),
('Керри трейд', 1);