CREATE TABLE patterns (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    icon TEXT,
    instrument_type SMALLINT NOT NULL,
    user_id BIGINT REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO patterns (name, instrument_type) 
VALUES 
('Голова и плечи', 1), 
('Перевернутая голова и плечи', 1), 
('Двойная вершина', 1), 
('Двойное дно', 1), 
('Тройная вершина', 1), 
('Тройное дно', 1), 
('Клин', 1), 
('Бриллиант', 1), 
('Прямоугольник', 1), 
('Флаг', 1), 
('Равносторонний треугольник', 1), 
('Восходящий Треугольник', 1), 
('Нисходящий треугольник', 1);