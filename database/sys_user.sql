/*
Navicat MySQL Data Transfer

Source Server         : 192.168.2.235
Source Server Version : 100339
Source Host           : localhost:3306
Source Database       : merge_v1

Target Server Type    : MYSQL
Target Server Version : 100339
File Encoding         : 65001

Date: 2024-10-15 13:51:58
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `uuid` varchar(64) DEFAULT '',
  `name` varchar(100) DEFAULT '',
  `gender` varchar(100) DEFAULT '',
  `login_user` varchar(191) DEFAULT '',
  `login_pass` longtext DEFAULT '',
  `email` varchar(64) DEFAULT '',
  `phone_num` varchar(11) DEFAULT '',
  `id_num` longtext DEFAULT '',
  `role_uuid` varchar(191) DEFAULT '',
  `team_uuid` varchar(191) DEFAULT '',
  `enable` tinyint(1) DEFAULT 0,
  `picture` longtext DEFAULT '',
  `is_online` tinyint(1) DEFAULT 0,
  `group_uuid` longtext DEFAULT '',
  `class_uuid` varchar(191) DEFAULT '',
  `is_studying` tinyint(1) DEFAULT 0,
  `latest_login_err_time` bigint(20) DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_sys_user_login_user` (`login_user`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`),
  KEY `idx_sys_user_uuid` (`uuid`),
  KEY `idx_sys_user_name` (`name`),
  KEY `idx_sys_user_role_uuid` (`role_uuid`),
  KEY `idx_sys_user_team_uuid` (`team_uuid`),
  KEY `idx_sys_user_class_uuid` (`class_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
SET FOREIGN_KEY_CHECKS=1;
