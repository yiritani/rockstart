-- CREATE TABLE IF NOT EXISTS m_artist (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     name VARCHAR(255) NOT NULL,
--     description TEXT,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );

insert into m_artist (name, description) values 
('HIDE', 'X JAPAN'),
('kuboty', 'TOTALFAT'),
('YOSHIKI', 'X JAPAN'),
('TOSHI', 'X JAPAN'),
('TAIJI', 'X JAPAN'),
('PATA', 'X JAPAN'),
('HEATH', 'X JAPAN');