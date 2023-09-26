
INSERT INTO users (login, password) VALUES ('user1', 'password1');
INSERT INTO users (login, password) VALUES ('user2', 'password2');
INSERT INTO users (login, password) VALUES ('user3', 'password3');

INSERT INTO spectrums (freq , len, Description,  Image  ,is_delete) VALUES (2439.7, 2439.7, 'opis 1', 'IRB.jpeg',false);
INSERT INTO spectrums (freq , len, Description,  Image  ,is_delete) VALUES (2439.7, 6051.8, 'opis 2', 'relict.jpeg',false);
INSERT INTO spectrums (freq , len, Description,  Image  ,is_delete) VALUES (2439.7, 6371.0,'aboba', 'xrb.jpeg',false);

INSERT INTO analysis_requests (date_start, date_end, status, satellite, user_id) VALUES ('2022-01-01', '2022-01-10', 'создан', 'AMS123', 1);
INSERT INTO analysis_requests (date_start, date_end, status, satellite, user_id) VALUES ('2022-02-01', '2022-02-10', 'в работе', 'AMS456', 2);
INSERT INTO analysis_requests (date_start, date_end, status, satellite, user_id) VALUES ('2022-03-01', '2022-03-10', 'отменён', 'AMS789', 3);

INSERT INTO spectrum_requests (ar_id, spectrum_id) VALUES (1, 1);
INSERT INTO spectrum_requests (ar_id, spectrum_id) VALUES (2, 2);
INSERT INTO spectrum_requests (ar_id, spectrum_id) VALUES (3, 3);
