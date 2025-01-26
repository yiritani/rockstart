insert into m_instrument (name, type, artist_id, description) values 
('ESP SNAPPER-HIDE Custom', 'Guitar', (select id from m_artist where name = 'HIDE'), 'Famous pink guitar used by HIDE'),
('ESP MX-220-HIDE', 'Guitar', (select id from m_artist where name = 'HIDE'), 'Yellow heart guitar used by HIDE'),
('Tama Artstar Custom', 'Drums', (select id from m_artist where name = 'YOSHIKI'), 'Clear acrylic drums used by YOSHIKI'),
('Yamaha CF-IIIS', 'Piano', (select id from m_artist where name = 'YOSHIKI'), 'Concert grand piano used by YOSHIKI'),
('ESP HORIZON-III', 'Guitar', (select id from m_artist where name = 'PATA'), 'Main guitar used by PATA'),
('ESP FOREST-GT Custom', 'Bass', (select id from m_artist where name = 'HEATH'), 'Custom bass guitar used by HEATH');

