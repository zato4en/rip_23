
INSERT INTO users (login, password, is_admin) VALUES ('user', 'user', false);
INSERT INTO users (login, password, is_admin) VALUES ('admin', 'admin', true);
INSERT INTO users (login, password, is_admin) VALUES ('user2', 'user2', false);




INSERT INTO spectrums (len , freq, Description,  Image  ,is_delete) VALUES (1.9, 157.78, 'CMB1 - has its wavelength 1.9mm and frequency 157.78Ghz', 'CMB1.jpeg',false);
INSERT INTO spectrums (len , freq, Description,  Image  ,is_delete) VALUES (2.72, 110.08, 'CMB2 - has its wavelength 2.72mm and frequency 110.08Ghz', 'CMB2.jpeg',false);
INSERT INTO spectrums (len , freq, Description,  Image  ,is_delete) VALUES (7.35, 40.86,'CMB3 - has its wavelength 7.35mm and frequency 40.86Ghz', 'CMB3.jpeg',false);



INSERT INTO analysis_requests (date_send, date_start, date_end, status, satellite, user_id, moder_id) VALUES ('2021-01-01','2022-01-01', '2022-01-10', 'создан', 'Planck', 1,2);
INSERT INTO analysis_requests (date_send, date_start, date_end, status, satellite, user_id, moder_id) VALUES ('2021-05-01','2022-02-01', '2022-02-10', 'в работе', 'COBE', 3,2);
INSERT INTO analysis_requests (date_send, date_start, date_end, status, satellite, user_id, moder_id) VALUES ('2021-04-01','2022-03-01', '2022-03-10', 'отменён', 'WMAP', 3,2);



INSERT INTO spectrum_requests (ar_id, spectrum_id) VALUES (1, 1);
INSERT INTO spectrum_requests (ar_id, spectrum_id) VALUES (2, 2);
INSERT INTO spectrum_requests (ar_id, spectrum_id) VALUES (3, 3);
