
INSERT INTO users (login, password) VALUES ('user1', 'password1');
INSERT INTO users (login, password) VALUES ('user2', 'password2');
INSERT INTO users (login, password) VALUES ('user3', 'password3');

INSERT INTO planets (name, radius, distance, gravity, image, is_delete) VALUES ('Сатурн', 2439.7, 57.9, 3.7, 'saturn.jpg',false);
INSERT INTO planets (name, radius, distance, gravity, image,is_delete) VALUES ('Марс', 6051.8, 108.2, 8.9, 'mars.jpg',false);
INSERT INTO planets (name, radius, distance, gravity, image,is_delete) VALUES ('Луна', 6371.0, 149.6, 9.8, 'moon.jpg',false);

INSERT INTO flight_requests (date_start, date_end, status, AMS, user_id) VALUES ('2022-01-01', '2022-01-10', 'создан', 'AMS123', 1);
INSERT INTO flight_requests (date_start, date_end, status, AMS, user_id) VALUES ('2022-02-01', '2022-02-10', 'в работе', 'AMS456', 2);
INSERT INTO flight_requests (date_start, date_end, status, AMS, user_id) VALUES ('2022-03-01', '2022-03-10', 'отменён', 'AMS789', 3);

INSERT INTO planets_requests (fr_id, planet_id) VALUES (1, 1);
INSERT INTO planets_requests (fr_id, planet_id) VALUES (2, 2);
INSERT INTO planets_requests (fr_id, planet_id) VALUES (3, 3);