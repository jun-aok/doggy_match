-- 権限追加
-- GRANT ALL PRIVILEGES ON *.* TO 'test'@'%' WITH GRANT OPTION;
-- SequelProからアクセスできるようにする
-- https://qiita.com/ysk1o/items/7f0ca12ced72363f9448
-- ALTER USER 'test'@"%" IDENTIFIED WITH mysql_native_password BY 'test';

CREATE TABLE `user` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `user_code` char(62) NOT NULL,
  `email` varchar(256) NOT NULL,
  `name` varchar(64) NOT NULL,
  `birth_date` date NOT NULL,
  `gender` tinyint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE user_info (
  user_id INT NOT NULL PRIMARY KEY,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME NULL
);

CREATE TABLE `dog` (
  `dog_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `name` varchar(32) NOT NULL,
  `birth_date` date NOT NULL,
  `gender` tinyint NOT NULL,
  `personality` tinyint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`dog_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE check_in (
  user_id INT NOT NULL UNIQUE,
  latitude DECIMAL(9, 6) NOT NULL,
  longitude DECIMAL(9, 6) NOT NULL,
  check_in_time DATETIME NOT NULL
)

CREATE TABLE check_in_dog (
  user_id INT NOT NULL,
  dog_id INT NOT NULL
)

CREATE TABLE check_in_log (
  check_in_log_id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  dog_ids VARCHAR(255) NOT NULL,
  latitude DECIMAL(9, 6) NOT NULL,
  longitude DECIMAL(9, 6) NOT NULL,
  check_in_begin_time DATETIME NOT NULL
  check_in_fimnish_time DATETIME NULL
)