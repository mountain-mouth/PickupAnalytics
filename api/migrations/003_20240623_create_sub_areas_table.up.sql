CREATE TABLE sub_areas (
    id SERIAL PRIMARY KEY,
    area_id INT,
    name TEXT,
    url TEXT,
    FOREIGN KEY (area_id) REFERENCES areas(id)
);

INSERT INTO sub_areas (area_id, name, url) VALUES
    (1, 'Shinjuku', 'https://www.shinjuku.jp'),
    (1, 'Shibuya', 'https://www.shibuya.jp'),
    (2, 'Umeda', 'https://www.umeda.osaka.jp'),
    (2, 'Namba', 'https://www.namba.osaka.jp');