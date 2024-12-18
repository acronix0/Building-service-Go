CREATE TABLE buildings (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    year INT NOT NULL,
    floors INT NOT NULL
);

CREATE INDEX idx_buildings_city ON buildings (city);
CREATE INDEX idx_buildings_year ON buildings (year);
CREATE INDEX idx_buildings_floors ON buildings (floors);
