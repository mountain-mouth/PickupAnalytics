DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    mail TEXT,
    pass TEXT,
    age INT,
    experience_years INT,
    is_published BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

INSERT INTO users (name, mail, pass, age, experience_years, is_published, created_at, updated_at)
VALUES ('杉浦裕晃', 'sugiura.doe@example.com', 'password123', 30, 5, TRUE, '2024-06-14 12:00:00+00', '2024-06-14 12:00:00+00');

INSERT INTO users (name, mail, pass, age, experience_years, is_published, created_at, updated_at)
VALUES ('沈徹', 'chin@example.com', 'password456', 25, 3, FALSE, '2024-06-14 12:15:00+00', '2024-06-14 12:15:00+00');

INSERT INTO users (name, mail, pass, age, experience_years, is_published, created_at, updated_at)
VALUES ('山口_猫', 'yamagiuchin@example.com', 'securepass789', 28, 4, TRUE, '2024-06-14 12:30:00+00', '2024-06-14 12:30:00+00');

INSERT INTO users (name, mail, pass, age, experience_years, is_published, created_at, updated_at)
VALUES ('小林弥生', 'kobayashi@example.com', 'password123', 30, 5, TRUE, '2024-06-14 12:00:00+00', '2024-06-14 12:00:00+00');
