CREATE TABLE targets (
    id BIGSERIAL PRIMARY KEY,
    age INT,
    name TEXT,
    occupation_id INT,
    area_id INT,
    remarks TEXT,
    FOREIGN KEY (occupation_id) REFERENCES occupations(id),
    FOREIGN KEY (area_id) REFERENCES areas(id)
);

INSERT INTO targets (age, name, occupation_id, area_id, remarks) VALUES
    (30, '山田太郎', 1, 1, '営業部門担当'),
    (25, '佐藤花子', 2, 2, '研究開発部門担当'),
    (40, '田中一郎', 3, 1, 'マーケティング部門担当');