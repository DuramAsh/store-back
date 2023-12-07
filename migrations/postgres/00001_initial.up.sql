CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS products (
	created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	id          VARCHAR PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
	title   VARCHAR NOT NULL,
	image   VARCHAR NOT NULL,
	price   NUMERIC NOT NULL,
	category VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
	created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	id          VARCHAR PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
	user_id    VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
	created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	id          VARCHAR PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
	email   	VARCHAR NOT NULL,
	password   VARCHAR NOT NULL
);

-- INSERT INTO products (title, image, price, category) VALUES
-- ('Стеганая куртка', 'https://i.ibb.co/jbnZK7W/photo-2023-01-20-00-02-21.jpg', 9500, 'Куртки'),
-- ('Флисовый костюм', 'https://i.ibb.co/fdnC2rt/photo-2023-01-20-00-02-24.jpg', 7500, 'Для мальчиков'),
-- ('Стильная двойка', 'https://i.ibb.co/Mks0Wsq/photo-2023-01-20-00-02-28.jpg', 8000, 'Для мальчиков'),
-- ('Свитшот', 'https://i.ibb.co/bgThWQb/photo-2023-01-20-00-02-31.jpg', 5000, 'Для мальчиков'),
-- ('Свободный спортивный костюм', 'https://i.ibb.co/88Xh9jz/photo-2023-01-20-00-02-31-2.jpg', 9000, 'Для мальчиков'),
-- ('Однотонные брюки', 'https://i.ibb.co/ypcgnTK/photo-2023-01-20-00-02-33-2.jpg', 5500, 'Брюки'),
-- ('Спортивные штаны', 'https://i.ibb.co/Hg5cQdc/photo-2023-01-20-00-02-35-2.jpg', 5000, 'Брюки'),
-- ('Рубашка с длинными рукавами', 'https://i.ibb.co/8ckRyLP/photo-2023-01-20-00-02-37-2.jpg', 5000, 'Для мальчиков'),
-- ('Рубашка с воротником', 'https://i.ibb.co/vz4qpn7/photo-2023-01-20-00-02-39-2.jpg', 5000, 'Для мальчиков'),
-- ('Модные кофточки ', 'https://i.ibb.co/D5x4vFT/photo-2023-01-20-00-02-42-2.jpg', 3500, 'Для мальчиков'),
-- ('Куртка', 'https://i.ibb.co/tYWsCV1/photo-2023-01-20-00-02-45-2.jpg', 7000, 'Куртки'),
-- ('Безрукавка', 'https://i.ibb.co/XSmykxS/photo-2023-01-20-00-02-47-2.jpg', 9000, 'Куртки'),
-- ('Спортивный костюм', 'https://i.ibb.co/wCpbTTB/photo-2023-01-20-00-02-49.jpg', 10000, 'Для мальчиков'),
-- ('Бархатная куртка', 'https://i.ibb.co/pzqgCTZ/photo-2023-01-20-00-02-50-2.jpg', 9000, 'Куртки'),
-- ('Шапка', 'https://i.ibb.co/n7SH370/photo-2023-01-20-00-02-52-2.jpg', 2000, 'Головные уборы'),
-- ('Сумка', 'https://i.ibb.co/MS0CpJp/photo-2023-01-20-00-02-55-2.jpg', 2500, 'Сумки'),
-- ('Стеганая куртка', 'https://i.ibb.co/d72tLvD/photo-2023-01-20-00-02-58-2.jpg', 12000, 'Куртки'),
-- ('Куртка', 'https://i.ibb.co/Kr5tbBt/photo-2023-01-20-00-03-00-2.jpg', 7000, 'Куртки'),
-- ('Спортивная двойка', 'https://i.ibb.co/fpGLjMC/photo-2023-01-20-00-03-07.jpg', 7500, 'Для девочек'),
-- ('Стильная двойка', 'https://i.ibb.co/CbPftZ3/photo-2023-01-20-00-03-10-2.jpg', 10000, 'Для девочек'),
-- ('Вельветовая юбка', 'https://i.ibb.co/wgG7kDS/photo-2023-01-20-00-03-13.jpg', 4000, 'Для девочек'),
-- ('Джинсовая куртка', 'https://i.ibb.co/WD7YBFM/photo-2023-01-20-00-03-15.jpg', 5000, 'Куртки'),
-- ('Кожаная куртка', 'https://i.ibb.co/QKMtDyP/photo-2023-01-20-00-03-16-2.jpg', 10000, 'Куртки'),
-- ('Двойка', 'https://i.ibb.co/BnMvvJL/photo-2023-01-20-00-03-18.jpg', 12000, 'Для девочек'),
-- ('Модная двойка', 'https://i.ibb.co/dDTqLNN/photo-2023-01-20-00-03-21-2.jpg', 7000, 'Для девочек'),
-- ('Широкие джинсы', 'https://i.ibb.co/0jkpkH6/photo-2023-01-20-00-03-22.jpg', 6000, 'Для девочек'),
-- ('Стильная двойка юбка + кофта', 'https://i.ibb.co/71s5Jng/photo-2023-01-20-00-03-25.jpg', 10000, 'Для девочек'),
-- ('Двойка юбка + кофта', 'https://i.ibb.co/Yc54M2P/photo-2023-01-20-00-03-26-2.jpg', 10000, 'Для девочек'),
-- ('Джинсовая юбка', 'https://i.ibb.co/pQnJv4z/photo-2023-01-20-00-03-29.jpg', 4000, 'Для девочек'),
-- ('Сумка', 'https://i.ibb.co/PQXQpf6/photo-2023-01-20-00-03-29-2.jpg', 3500, 'Сумки'),
-- ('Сумка', 'https://i.ibb.co/YW8rLJz/photo-2023-01-20-00-03-32.jpg', 3500, 'Сумки');