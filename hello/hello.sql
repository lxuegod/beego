/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 80025
Source Host           : localhost:3306
Source Database       : hello

Target Server Type    : MYSQL
Target Server Version : 80025
File Encoding         : 65001

Date: 2022-04-10 20:12:17
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '帖子标题',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '帖子描述',
  `cover` varchar(255) NOT NULL DEFAULT 'static/upload/no_p.jpg' COMMENT '帖子封面图',
  `read_num` int NOT NULL DEFAULT '0' COMMENT '帖子阅读数',
  `star_num` int NOT NULL DEFAULT '0' COMMENT '帖子点赞数',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `author_id` int NOT NULL COMMENT '帖子作者',
  `content` varchar(400) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES ('1', '测试', '测试', 'static/upload/no_p.jpg', '11', '22', '2022-03-13 14:59:14', '1', '22');
INSERT INTO `sys_post` VALUES ('13', '测试2', '测试2', 'static/upload/16484502871.png', '0', '0', '2022-03-28 14:51:27', '1', '测试2');
INSERT INTO `sys_post` VALUES ('14', '测试3', '测试3', 'static/upload/1648450310coin.png', '0', '0', '2022-03-28 14:51:51', '1', '测试3');
INSERT INTO `sys_post` VALUES ('15', '测试4', '测试4', 'static/upload/1648450329coin2.png', '0', '0', '2022-03-28 14:52:09', '1', '测试4');
INSERT INTO `sys_post` VALUES ('16', '测试5', '测试5', 'static/upload/1648450343l.jpg', '0', '0', '2022-03-28 14:52:23', '1', '测试5');
INSERT INTO `sys_post` VALUES ('20', '测试9', '测试9', 'static/upload/1648450395coin2.png', '0', '0', '2022-03-28 14:53:16', '1', '测试9');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `is_admin` int NOT NULL DEFAULT '2' COMMENT '1是管理员，2是普通用户',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `cover` varchar(255) NOT NULL DEFAULT 'static/upload/bq3.png' COMMENT '头像',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '1', '2022-03-11 16:57:33', 'static/upload/bq3.png');
