CREATE TABLE occupations (
    id BIGSERIAL PRIMARY KEY,
    name TEXT
);

INSERT INTO occupations (name) VALUES
    ('ソフトウェアエンジニア'),
    ('データサイエンティスト'),
    ('マーケティングマネージャー');