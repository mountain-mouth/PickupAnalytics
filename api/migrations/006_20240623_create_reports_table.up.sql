CREATE TABLE reports (
    id BIGSERIAL PRIMARY KEY,
    user_id INT,
    area_id INT,
    sub_area_id INT,
    k INT,
    o INT,
    l INT,
    t INT,
    s INT[],
    targets INT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- INSERT INTO reports (user_id, area_id, sub_area_id, k, o, l, t, s, targets)
-- VALUES
--     (1, 1, 1, 10, 20, 30, 40, '{1, 2, 3}', '{101, 102, 103}', '2024-06-23 14:30:00', '2024-06-23 14:30:00'),
--     (2, 2, 2, 15, 25, 35, 45, '{4, 5}', '{104, 105}', '2024-06-23 15:45:00', '2024-06-23 15:45:00');