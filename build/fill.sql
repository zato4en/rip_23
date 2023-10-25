
INSERT INTO users (login, password, is_admin) VALUES ('user', 'user', false);
INSERT INTO users (login, password, is_admin) VALUES ('admin', 'admin', true);
INSERT INTO users (login, password, is_admin) VALUES ('user2', 'user2', false);




INSERT INTO spectrums (len , freq, Description,  Image  ,is_delete) VALUES (1.9, 157.78, 'CMB1 - имеет длину волны 1.9mm и частоту 157.78Ghz', 'CMB1.jpeg',false);
INSERT INTO spectrums (len , freq, Description,  Image  ,is_delete) VALUES (2.72, 110.08, 'CMB2 - имеет длину волны 2.72mm и частоту 110.08Ghz', 'CMB2.jpeg',false);
INSERT INTO spectrums (len , freq, Description,  Image  ,is_delete) VALUES (7.35, 40.86,'CMB3 - имеет длину волны 7.35mm и частоту 40.86Ghz', 'CMB3.jpeg',false);



INSERT INTO satellites (date_created, date_formed, date_accepted, status, satellite, user_id, moder_id) VALUES ('2021-01-01','2022-01-01', '2022-01-10', 'создан', 'Planck', 1,2);
INSERT INTO satellites (date_created, date_formed, date_accepted, status, satellite, user_id, moder_id) VALUES ('2021-05-01','2022-02-01', '2022-02-10', 'создан', 'COBE', 3,2);
INSERT INTO satellites (date_created, date_formed, date_accepted, status, satellite, user_id, moder_id) VALUES ('2021-04-01','2022-03-01', '2022-03-10', 'создан', 'WMAP', 3,2);



INSERT INTO spectrum_requests (satellite_id, spectrum_id, satellite_number) VALUES (1, 1, 1);
INSERT INTO spectrum_requests (satellite_id, spectrum_id, satellite_number) VALUES (2, 2, 2);
INSERT INTO spectrum_requests (satellite_id, spectrum_id, satellite_number) VALUES (3, 3, 3);
