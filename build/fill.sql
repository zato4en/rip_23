--
-- INSERT INTO users (login, password, is_admin) VALUES ('user1', 'user1', false);
-- INSERT INTO users (login, password, is_admin) VALUES ('user4', 'admin', true);
-- INSERT INTO users (login, password, is_admin) VALUES ('user2', 'user2', false);
-- INSERT INTO users (login, password, is_admin) VALUES ('user3', 'user3',false);





INSERT INTO spectrums (name, len , freq, Description,  Image  ,is_delete) VALUES ('CMB1', 1.9, 157.78, 'CMB1 - имеет длину волны 1.9мм и частоту 157.78 Ггц', 'http://127.0.0.1:9000/spectrumbucket/CMB1.jpeg',false);
INSERT INTO spectrums (name, len , freq, Description,  Image  ,is_delete) VALUES ('CMB2', 2.72, 110.08, 'CMB2 - имеет длину волны 2.72мм и частоту 110.08 Ггц', 'http://127.0.0.1:9000/spectrumbucket/CMB2.jpeg',false);
INSERT INTO spectrums (name, len , freq, Description,  Image  ,is_delete) VALUES ('CMB3', 7.35, 40.86,'CMB3 - имеет длину волны 7.35мм и частоту 40.86 Ггц', 'http://127.0.0.1:9000/spectrumbucket/CMB3.jpeg',false);


--
-- INSERT INTO satellites (id, date_created, date_formed, date_accepted, status, satellite, user_id, moder_id, user_login) VALUES (1,'2010-01-01','2011-01-01', '2022-01-10', 'создан', 'Planck', 1,2, 'user1');
-- INSERT INTO satellites (id, date_created, date_formed, date_accepted, status, satellite, user_id, moder_id, user_login) VALUES (2,'2015-05-01','2016-02-01', '2022-02-10', 'создан', 'COBE', 3,2, 'user2');
-- INSERT INTO satellites (id, date_created, date_formed, date_accepted, status, satellite, user_id, moder_id, user_login) VALUES (3,'2017-04-01','2018-03-01', '2022-03-10', 'создан', 'WMAP', 3,2, 'user3');
-- INSERT INTO satellites (id, date_created, date_formed, date_accepted, status, satellite, user_id, moder_id, user_login) VALUES (4,'2021-04-01','2022-03-01', '2022-03-10', 'в работе', 'WMAP', 3,2, 'user4');


--
-- INSERT INTO spectrum_requests (satellite_id, spectrum_id, satellite_number) VALUES (1, 1, 1);
-- INSERT INTO spectrum_requests (satellite_id, spectrum_id, satellite_number) VALUES (2, 2, 2);
-- INSERT INTO spectrum_requests (satellite_id, spectrum_id, satellite_number) VALUES (3, 3, 3);
