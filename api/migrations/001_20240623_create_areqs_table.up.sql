DROP TABLE IF EXISTS areas;

CREATE TABLE areas (
    id SERIAL PRIMARY KEY,
    name TEXT,
    url TEXT
);

INSERT INTO areas (name, url) VALUES
    ('Tokyo', 'https://www.tokyo.jp'),
    ('Osaka', 'https://www.city.osaka.lg.jp'),
    ('Kyoto', 'https://www.city.kyoto.lg.jp');