/*
 Navicat Premium Data Transfer

 Source Server         : 华为云
 Source Server Type    : MySQL
 Source Server Version : 80300
 Source Host           : 121.37.143.160:3306
 Source Schema         : ginchat

 Target Server Type    : MySQL
 Target Server Version : 80300
 File Encoding         : 65001

 Date: 05/04/2024 21:02:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for communities
-- ----------------------------
DROP TABLE IF EXISTS `communities`;
CREATE TABLE `communities`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `owner_id` bigint UNSIGNED NULL DEFAULT NULL,
  `img` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `desc` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_communities_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of communities
-- ----------------------------
INSERT INTO `communities` VALUES (18, '2024-04-02 18:52:50.923', '2024-04-02 18:52:50.923', NULL, 'love', 27, './asset/upload/17120551671490748086.jpeg', '');

-- ----------------------------
-- Table structure for contact
-- ----------------------------
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `owner_id` bigint UNSIGNED NULL DEFAULT NULL,
  `target_id` bigint UNSIGNED NULL DEFAULT NULL,
  `type` bigint NULL DEFAULT NULL,
  `desc` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_contact_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 189 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of contact
-- ----------------------------
INSERT INTO `contact` VALUES (185, '2024-04-01 15:29:06.044', '2024-04-01 15:29:06.044', NULL, 26, 27, 1, '');
INSERT INTO `contact` VALUES (186, '2024-04-01 15:29:06.213', '2024-04-01 15:29:06.213', NULL, 27, 26, 1, '');
INSERT INTO `contact` VALUES (187, '2024-04-02 18:52:51.113', '2024-04-02 18:52:51.113', NULL, 27, 18, 2, '');
INSERT INTO `contact` VALUES (188, '2024-04-02 18:54:31.188', '2024-04-02 18:54:31.188', NULL, 26, 18, 2, '');

-- ----------------------------
-- Table structure for group_basic
-- ----------------------------
DROP TABLE IF EXISTS `group_basic`;
CREATE TABLE `group_basic`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `owner_id` bigint UNSIGNED NULL DEFAULT NULL,
  `icon` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `type` bigint NULL DEFAULT NULL,
  `desc` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group_basic_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_basic
-- ----------------------------

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `target_id` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `type` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `media` bigint NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `pic` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `url` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `desc` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `amount` bigint NULL DEFAULT NULL,
  `create_time` bigint NULL DEFAULT NULL,
  `read_time` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_message_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (33, '2024-04-02 20:26:29.447', '2024-04-02 20:26:29.447', NULL, '27', '18', '2', 1, '1', '', '', '', 0, 1712060769862, 0);
INSERT INTO `message` VALUES (34, '2024-04-02 20:29:02.742', '2024-04-02 20:29:02.742', NULL, '26', '27', '1', 1, '1', '', '', '', 0, 1712060938235, 0);

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `pass_word` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `phone` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `email` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `identity` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `client_ip` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `client_port` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `login_time` datetime(3) NULL DEFAULT NULL,
  `heartbeat_time` datetime(3) NULL DEFAULT NULL,
  `login_out_time` datetime(3) NULL DEFAULT NULL,
  `is_logout` tinyint(1) NULL DEFAULT NULL,
  `device_info` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `salt` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_basic_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_basic
-- ----------------------------
INSERT INTO `user_basic` VALUES (26, '2024-04-01 15:28:14.565', '2024-04-02 18:55:10.713', NULL, 'love1', '570f1599c9e66c3f9a19046796a5187a', '', '', '123456', '', '', '2024-04-01 15:28:14.528', '2024-04-01 15:28:14.528', '2024-04-01 15:28:14.528', 0, '', '375768629', './asset/upload/171205530629525436.jpeg');
INSERT INTO `user_basic` VALUES (27, '2024-04-01 15:28:28.557', '2024-04-02 18:53:44.434', NULL, 'love', '8623124b7e19f6cc62df073db367bac3', '', '', '123456', '', '', '2024-04-01 15:28:28.508', '2024-04-01 15:28:28.508', '2024-04-01 15:28:28.508', 0, '', '963259754', './asset/upload/1712055221571377529.jpeg');

SET FOREIGN_KEY_CHECKS = 1;
