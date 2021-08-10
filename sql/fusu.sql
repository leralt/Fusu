create database Fusu;
use Fusu;
create table `user` (
    `account` char(255) NOT NULL,
    `password` char(255) NOT NULL,
    `nick` char(255) NOT NULL UNIQUE,
    PRIMARY KEY (`account`)
)charset=utf8;
create table `massage` (
    `msg_id` int NOT NULL AUTO_INCREMENT,
    `nick` char(255) NOT NULL,
    `msg` char(255) NOT NULL,
    `time` datetime NOT NULL,
    PRIMARY KEY (`msg_id`)
)charset=utf8;