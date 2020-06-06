-- Create database.
CREATE DATABASE IF NOT EXISTS pismo;
USE pismo;

CREATE TABLE `account` (
  `account_id` int(11) NOT NULL AUTO_INCREMENT,
  `document_number` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`account_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `transaction` (
  `transaction_id` int(11) NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL,
  `operation_type` int(11) DEFAULT NULL,
  `amount` decimal(10,5) DEFAULT NULL,
  `event_date` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`transaction_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
