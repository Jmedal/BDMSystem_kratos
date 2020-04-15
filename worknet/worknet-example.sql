/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : worknet-example

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2020-04-15 13:01:52
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_administrator
-- ----------------------------
DROP TABLE IF EXISTS `sys_administrator`;
CREATE TABLE `sys_administrator` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `account` varchar(255) NOT NULL DEFAULT '' COMMENT '用户账号',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',
  `power` int(20) NOT NULL DEFAULT '0' COMMENT '权限',
  PRIMARY KEY (`id`,`account`) USING BTREE,
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='管理员表';

-- ----------------------------
-- Records of sys_administrator
-- ----------------------------
INSERT INTO `sys_administrator` VALUES ('1', 'admin', 'admin', '3');

-- ----------------------------
-- Table structure for sys_admin_check_company
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin_check_company`;
CREATE TABLE `sys_admin_check_company` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '特殊账号注册审核记录id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '申请帐号id',
  `administrator_id` bigint(255) unsigned NOT NULL COMMENT '审核管理员id',
  `apply_time` datetime(6) NOT NULL COMMENT '申请时间',
  `check_time` datetime(6) DEFAULT NULL COMMENT '审核时间',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '当前审核状态',
  `result` tinyint(1) unsigned DEFAULT '0' COMMENT '审核结果',
  PRIMARY KEY (`id`,`user_id`,`administrator_id`) USING BTREE,
  KEY `fk_user_id` (`user_id`) USING BTREE,
  KEY `fk_administrator_id` (`administrator_id`) USING BTREE,
  CONSTRAINT `fk_administrator_id` FOREIGN KEY (`administrator_id`) REFERENCES `sys_administrator` (`id`),
  CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='特殊账号注册审核记录表';

-- ----------------------------
-- Records of sys_admin_check_company
-- ----------------------------

-- ----------------------------
-- Table structure for sys_company
-- ----------------------------
DROP TABLE IF EXISTS `sys_company`;
CREATE TABLE `sys_company` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '公司id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '公司帐号id',
  `field` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '公司从事方向',
  `name` varchar(50) DEFAULT '' COMMENT '公司名称',
  `introduction` varchar(100) DEFAULT '' COMMENT '公司简介',
  `address` varchar(50) DEFAULT '' COMMENT '公司地址',
  `communication` varchar(30) DEFAULT NULL COMMENT '公司联系方式',
  `website` varchar(50) DEFAULT NULL COMMENT '公司官方网站',
  `register_time` datetime(6) DEFAULT NULL COMMENT '注册时间',
  PRIMARY KEY (`id`,`user_id`) USING BTREE,
  KEY `fk_sys_company_sys_user_1` (`user_id`) USING BTREE,
  KEY `id` (`id`) USING BTREE,
  CONSTRAINT `fk_sys_company_sys_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='公司信息表';

-- ----------------------------
-- Records of sys_company
-- ----------------------------
INSERT INTO `sys_company` VALUES ('1', '11', '互联网', '百度', '百度\n牛逼\n\n牛逼', '北京', '021-55555555', 'http://www.baidu.com', '2019-05-07 00:28:05.000000');
INSERT INTO `sys_company` VALUES ('2', '12', '互联网', '腾讯', '腾讯', '深圳', '021-55555555', 'https://www.qq.com/', '2019-05-25 20:52:58.000000');
INSERT INTO `sys_company` VALUES ('3', '13', '互联网', '阿里巴巴', '阿里巴巴', '杭州', '021-55555555', 'https://www.alibabagroup.com/cn/global/home', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('4', '14', '互联网', '蚂蚁金服', '蚂蚁金服', '杭州', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('5', '15', '互联网', '美团', '美团', '北京', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('6', '16', '互联网', '小米', '小米', '北京', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('7', '40', '互联网', '京东', '京东', '北京', '021-55555555', 'https://www.jd.com/', '2019-07-08 14:30:26.000000');
INSERT INTO `sys_company` VALUES ('8', '17', '互联网', '拼多多', '拼多多', '上海', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('9', '18', '互联网', '网易', '网易', '广州', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('10', '19', '互联网', '滴滴', '滴滴', '北京', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('11', '20', '互联网', '字节跳动', '字节跳动', '北京', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('12', '21', '互联网', '虎牙', '虎牙', '广州', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('13', '22', '互联网', '酷狗', '酷狗', '广州', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('14', '23', '互联网', '哈罗出行', '哈罗出行', '上海', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('15', '24', '互联网', 'bilibili', 'bilibili', '上海', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('16', '25', '互联网', '小红书', '小红书', '上海', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('18', '27', '互联网', '爱奇艺', '爱奇艺', '上海', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('19', '30', '互联网', '华为', '华为', '深圳', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('20', '31', '互联网', '贝壳', '贝壳', '西安', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('21', '32', '互联网', '作业帮', '作业帮', '西安', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('22', '33', '互联网', '大禹网络', '大禹网络', '苏州', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('23', '34', '互联网', '星火教育', '星火教育', '苏州', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('24', '35', '互联网', '怪兽充电', '怪兽充电', '武汉', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('25', '36', '互联网', '芒果TV', '芒果TV', '长沙', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('26', '28', '互联网', 'oppo', 'oppo', '深圳', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('27', '38', '互联网', '找你科技', '找你科技', '重庆', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('28', '39', '互联网', '中移在线', '中移在线', '郑州', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('29', '40', '互联网', '嗨学', '嗨学', '成都', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('30', '26', '互联网', '菜鸟网络', '菜鸟网络', '上海', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_company` VALUES ('31', '29', '互联网', '顺丰', '顺丰', '深圳', '021-55555555', 'https://www.xxxxx.com/', '2020-03-05 11:57:18.000000');

-- ----------------------------
-- Table structure for sys_company_contest
-- ----------------------------
DROP TABLE IF EXISTS `sys_company_contest`;
CREATE TABLE `sys_company_contest` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '公司笔试id',
  `company_id` bigint(255) unsigned NOT NULL COMMENT '公司id',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '笔试标题',
  `start_time` datetime(6) NOT NULL COMMENT '笔试开始时间',
  `end_time` datetime(6) NOT NULL COMMENT '笔试结束时间',
  `activity` tinyint(10) NOT NULL COMMENT '是否激活',
  PRIMARY KEY (`id`,`company_id`) USING BTREE,
  KEY `fk_sys_company_contest_sys_company_1` (`company_id`) USING BTREE,
  KEY `id` (`id`) USING BTREE,
  CONSTRAINT `fk_sys_company_contest_sys_company_1` FOREIGN KEY (`company_id`) REFERENCES `sys_company` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='公司笔试表';

-- ----------------------------
-- Records of sys_company_contest
-- ----------------------------
INSERT INTO `sys_company_contest` VALUES ('1', '2', '腾讯校招', '2019-07-05 10:53:57.000000', '2020-12-04 20:54:02.000000', '1');
INSERT INTO `sys_company_contest` VALUES ('2', '2', '腾讯校招', '2019-06-04 10:53:57.000000', '2020-12-04 20:54:02.000000', '1');
INSERT INTO `sys_company_contest` VALUES ('3', '1', '百度校招', '2019-08-10 10:53:57.000000', '2020-12-04 20:54:02.000000', '1');
INSERT INTO `sys_company_contest` VALUES ('4', '3', '阿里校招', '2020-01-07 10:53:57.000000', '2020-12-04 20:54:02.000000', '1');
INSERT INTO `sys_company_contest` VALUES ('5', '3', '阿里校招', '2019-09-01 10:53:57.000000', '2020-12-04 20:54:02.000000', '1');
INSERT INTO `sys_company_contest` VALUES ('6', '3', '阿里校招', '2019-12-04 10:53:57.000000', '2020-12-04 20:54:02.000000', '1');
INSERT INTO `sys_company_contest` VALUES ('7', '3', '阿里校招', '2020-01-01 10:53:57.000000', '2020-01-01 20:54:02.000000', '1');
INSERT INTO `sys_company_contest` VALUES ('8', '2', '腾讯校招', '2019-09-01 10:53:57.000000', '2020-12-04 20:54:02.000000', '1');

-- ----------------------------
-- Table structure for sys_company_contest_apply
-- ----------------------------
DROP TABLE IF EXISTS `sys_company_contest_apply`;
CREATE TABLE `sys_company_contest_apply` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '报名id',
  `contest_id` bigint(255) unsigned NOT NULL COMMENT '笔试id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `permission` tinyint(10) NOT NULL DEFAULT '1' COMMENT '是否有效（考试时间、作弊等）',
  `sign_up_time` datetime(6) NOT NULL COMMENT '报名时间',
  PRIMARY KEY (`id`,`contest_id`,`user_id`) USING BTREE,
  KEY `fk_sys_comany_contest_apply_sys_company_contest_1` (`contest_id`) USING BTREE,
  KEY `fk_sys_comany_contest_apply_sys_user_1` (`user_id`) USING BTREE,
  CONSTRAINT `fk_sys_comany_contest_apply_sys_company_contest_1` FOREIGN KEY (`contest_id`) REFERENCES `sys_company_contest` (`id`),
  CONSTRAINT `fk_sys_comany_contest_apply_sys_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='公司笔试报名表';

-- ----------------------------
-- Records of sys_company_contest_apply
-- ----------------------------
INSERT INTO `sys_company_contest_apply` VALUES ('1', '1', '4', '1', '2019-05-25 20:55:42.000000');
INSERT INTO `sys_company_contest_apply` VALUES ('2', '1', '3', '1', '2019-05-27 15:36:40.000000');
INSERT INTO `sys_company_contest_apply` VALUES ('3', '2', '3', '1', '2019-05-27 15:36:40.000000');
INSERT INTO `sys_company_contest_apply` VALUES ('4', '3', '3', '1', '2019-05-27 15:36:40.000000');
INSERT INTO `sys_company_contest_apply` VALUES ('5', '5', '3', '1', '2019-05-27 15:36:40.000000');
INSERT INTO `sys_company_contest_apply` VALUES ('6', '4', '3', '1', '2019-05-27 15:36:40.000000');
INSERT INTO `sys_company_contest_apply` VALUES ('7', '7', '3', '1', '2019-05-27 15:36:40.000000');
INSERT INTO `sys_company_contest_apply` VALUES ('8', '1', '3', '1', '2019-05-27 15:36:40.000000');
INSERT INTO `sys_company_contest_apply` VALUES ('9', '4', '3', '1', '2019-05-27 15:36:40.000000');
INSERT INTO `sys_company_contest_apply` VALUES ('10', '8', '3', '1', '2019-05-27 15:36:40.000000');

-- ----------------------------
-- Table structure for sys_company_contest_choice_option
-- ----------------------------
DROP TABLE IF EXISTS `sys_company_contest_choice_option`;
CREATE TABLE `sys_company_contest_choice_option` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '笔试选项id',
  `question_id` bigint(255) unsigned NOT NULL COMMENT '问题id',
  `option` varchar(255) NOT NULL COMMENT '选项内容',
  PRIMARY KEY (`id`,`question_id`) USING BTREE,
  KEY `fk_sys_company_contest_choice_option_company_contest_question_1` (`question_id`) USING BTREE,
  CONSTRAINT `fk_sys_company_contest_choice_option_company_contest_question_1` FOREIGN KEY (`question_id`) REFERENCES `sys_company_contest_question` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='公司笔试选项表';

-- ----------------------------
-- Records of sys_company_contest_choice_option
-- ----------------------------
INSERT INTO `sys_company_contest_choice_option` VALUES ('1', '1', '111');
INSERT INTO `sys_company_contest_choice_option` VALUES ('2', '1', '1111');
INSERT INTO `sys_company_contest_choice_option` VALUES ('3', '1', '11111');
INSERT INTO `sys_company_contest_choice_option` VALUES ('4', '1', '111111');

-- ----------------------------
-- Table structure for sys_company_contest_question
-- ----------------------------
DROP TABLE IF EXISTS `sys_company_contest_question`;
CREATE TABLE `sys_company_contest_question` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '问题id',
  `contest_id` bigint(255) unsigned NOT NULL COMMENT '所属笔试id',
  `question` varchar(255) NOT NULL COMMENT '问题内容',
  `correct_option_id` bigint(255) unsigned NOT NULL COMMENT '正确选项id',
  `order_no` int(255) NOT NULL COMMENT '问题序号',
  `question_type` tinyint(10) NOT NULL DEFAULT '0' COMMENT '问题类型',
  PRIMARY KEY (`id`,`contest_id`,`correct_option_id`) USING BTREE,
  KEY `fk_company_contest_question_sys_company_contest_1` (`contest_id`) USING BTREE,
  KEY `fk_sys_company_contest_question_company_contest_choice_option_1` (`correct_option_id`) USING BTREE,
  CONSTRAINT `fk_company_contest_question_sys_company_contest_1` FOREIGN KEY (`contest_id`) REFERENCES `sys_company_contest` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='公司笔试问题表';

-- ----------------------------
-- Records of sys_company_contest_question
-- ----------------------------
INSERT INTO `sys_company_contest_question` VALUES ('1', '1', '123', '1', '1', '5');

-- ----------------------------
-- Table structure for sys_company_contest_result
-- ----------------------------
DROP TABLE IF EXISTS `sys_company_contest_result`;
CREATE TABLE `sys_company_contest_result` (
  `id` bigint(255) unsigned NOT NULL COMMENT '笔试回答id',
  `contest_id` bigint(255) unsigned NOT NULL COMMENT '回答所属笔试id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '回答用户id',
  `question_id` bigint(255) unsigned NOT NULL COMMENT '问题id',
  `start_time` datetime(6) NOT NULL COMMENT '开始回答时间',
  `end_time` datetime(6) NOT NULL COMMENT '结束回答时间',
  `answer_id` bigint(255) unsigned NOT NULL COMMENT '回答选项id',
  PRIMARY KEY (`id`,`user_id`,`contest_id`,`question_id`) USING BTREE,
  KEY `fk_company_contest_result_sys_company_contest_1` (`contest_id`) USING BTREE,
  KEY `fk_company_contest_result_sys_user_1` (`user_id`) USING BTREE,
  KEY `fk_company_contest_result_company_contest_question_1` (`question_id`) USING BTREE,
  CONSTRAINT `fk_company_contest_result_company_contest_question_1` FOREIGN KEY (`question_id`) REFERENCES `sys_company_contest_question` (`id`),
  CONSTRAINT `fk_company_contest_result_sys_company_contest_1` FOREIGN KEY (`contest_id`) REFERENCES `sys_company_contest` (`id`),
  CONSTRAINT `fk_company_contest_result_sys_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='公司笔试回答表';

-- ----------------------------
-- Records of sys_company_contest_result
-- ----------------------------

-- ----------------------------
-- Table structure for sys_company_cv
-- ----------------------------
DROP TABLE IF EXISTS `sys_company_cv`;
CREATE TABLE `sys_company_cv` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '简历id',
  `company_profession_id` bigint(255) unsigned NOT NULL COMMENT '招聘id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `name` varchar(20) DEFAULT NULL COMMENT '姓名',
  `sex` int(1) unsigned DEFAULT '1' COMMENT '性别',
  `birth` datetime(6) DEFAULT NULL COMMENT '出生年月',
  `native_place` varchar(10) DEFAULT NULL COMMENT '籍贯',
  `identity` varchar(50) DEFAULT NULL COMMENT '身份证号',
  `qualification` varchar(10) DEFAULT NULL COMMENT '学历',
  `specialty` varchar(20) DEFAULT NULL COMMENT '专业',
  `university` varchar(50) DEFAULT NULL COMMENT '毕业院校',
  `tel` char(11) DEFAULT NULL COMMENT '电话',
  `experience` varchar(255) DEFAULT NULL COMMENT '实习经历',
  `mailbox` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `introduction` varchar(255) DEFAULT NULL COMMENT '自我介绍',
  `diploma` varchar(255) DEFAULT NULL COMMENT '获奖情况',
  `current_location` varchar(100) DEFAULT NULL COMMENT '当前所在地',
  `in_job_time` varchar(50) DEFAULT NULL COMMENT '最快入职时间',
  `head_path` varchar(255) DEFAULT '\\avatar\\default\\avatar.jpg' COMMENT '证件照',
  `status` tinyint(5) DEFAULT '0' COMMENT '包含（待审核，审核中，拒绝，接受）',
  `last_edit_time` datetime(6) DEFAULT NULL COMMENT '上次提交简历时间',
  PRIMARY KEY (`id`,`company_profession_id`,`user_id`) USING BTREE,
  KEY `fk_company_profession_id` (`company_profession_id`) USING BTREE,
  KEY `fk_comany_cv_user_id` (`user_id`) USING BTREE,
  CONSTRAINT `fk_comany_cv_user_id` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`),
  CONSTRAINT `fk_company_profession_id` FOREIGN KEY (`company_profession_id`) REFERENCES `sys_company_profession` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='简历表';

-- ----------------------------
-- Records of sys_company_cv
-- ----------------------------
INSERT INTO `sys_company_cv` VALUES ('1', '1', '3', '张萨姆', '1', '1997-05-07 00:00:00.000000', '上海市', '124124124', '3', '计算机', '上海大学', '18888888888', '2008-2011 白宫洗碗三年', '88888888@qq.com', '我是一个xxxxxx\\n\\n哈哈哈', '2018-03-21 洗碗全国大奖', '上海', '暑期开始', '\\avatar\\default\\avatar.jpg', '1', '2019-07-10 01:51:21.000000');
INSERT INTO `sys_company_cv` VALUES ('2', '17', '3', '姜云杰', '0', '1998-07-08 00:00:00.000000', '安徽省', '11111111111111111111111', '3', '计算机科学与技术', '上海理工大学', '18321882894', '啦啦啦', '784203082@qq.com', '啦啦啦', '啦啦啦', '北京', '本周', '\\company_cv_avatar\\2019\\7b5982d9-cf9f-47fe-b3bc-019666aa7029.png', '1', '2019-07-11 13:32:21.000000');
INSERT INTO `sys_company_cv` VALUES ('3', '18', '3', '姜云杰', '0', '1998-07-08 00:00:00.000000', '安徽省', '11111111111111111111111', '3', '计算机科学与技术', '上海理工大学', '18321882894', '啦啦啦', '784203082@qq.com', '啦啦啦', '啦啦啦', '北京', '本周', '\\company_cv_avatar\\2019\\7b5982d9-cf9f-47fe-b3bc-019666aa7029.png', '1', '2019-07-11 13:32:21.000000');
INSERT INTO `sys_company_cv` VALUES ('19', '3', '3', '姜云杰', '0', '1998-07-08 00:00:00.000000', '安徽省', '11111111111111111111111', '3', '计算机科学与技术', '上海理工大学', '18321882894', '啦啦啦55555', '784203082@qq.com', '啦啦啦', '啦啦啦', '北京', '三个月内', '\\company_cv_avatar\\2019\\bc8c4b63-a899-4b72-9988-e6ca2c546c98.png', '0', '2019-07-10 10:46:59.000000');
INSERT INTO `sys_company_cv` VALUES ('20', '3', '3', '姜云杰', '0', '1998-07-08 00:00:00.000000', '安徽省', '11111111111111111111111', '3', '计算机科学与技术', '上海理工大学', '18321882894', '啦啦啦', '784203082@qq.com', '啦啦啦', '啦啦啦', '北京', '本周', '\\company_cv_avatar\\2019\\7b5982d9-cf9f-47fe-b3bc-019666aa7029.png', '2', '2019-07-11 13:32:22.000000');
INSERT INTO `sys_company_cv` VALUES ('21', '15', '3', '姜云杰', '0', '1998-07-08 00:00:00.000000', '安徽省', '11111111111111111111111', '3', '计算机科学与技术', '上海理工大学', '18321882894', '啦啦啦', '784203082@qq.com', '啦啦啦', '啦啦啦', '北京', '本周', '\\company_cv_avatar\\2019\\7b5982d9-cf9f-47fe-b3bc-019666aa7029.png', '1', '2019-07-11 13:32:21.000000');
INSERT INTO `sys_company_cv` VALUES ('22', '2', '3', '姜云杰', '0', '1998-07-08 00:00:00.000000', '安徽省', '11111111111111111111111', '3', '计算机科学与技术', '上海理工大学', '18321882894', '啦啦啦', '784203082@qq.com', '啦啦啦', '啦啦啦', '北京', '本周', '\\company_cv_avatar\\2019\\7b5982d9-cf9f-47fe-b3bc-019666aa7029.png', '1', '2019-07-11 13:32:21.000000');

-- ----------------------------
-- Table structure for sys_company_invitation
-- ----------------------------
DROP TABLE IF EXISTS `sys_company_invitation`;
CREATE TABLE `sys_company_invitation` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '邀约id',
  `company_id` bigint(255) unsigned NOT NULL COMMENT '公司id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `company_profession_id` bigint(255) unsigned NOT NULL COMMENT '招聘id',
  `status` tinyint(5) DEFAULT NULL COMMENT '邀约状态（未回复 0，同意 1，拒绝 2）',
  `invite_time` datetime(6) NOT NULL COMMENT '邀请时间',
  PRIMARY KEY (`id`,`company_id`,`user_id`,`company_profession_id`) USING BTREE,
  KEY `invitation_company_id` (`company_id`) USING BTREE,
  KEY `invitation_user_id` (`user_id`) USING BTREE,
  KEY `invitation_profession_id` (`company_profession_id`) USING BTREE,
  CONSTRAINT `fk_sys_company_invitation_sys_company_1` FOREIGN KEY (`company_id`) REFERENCES `sys_company` (`id`),
  CONSTRAINT `invitation_profession_id` FOREIGN KEY (`company_profession_id`) REFERENCES `sys_company_profession` (`id`),
  CONSTRAINT `invitation_user_id` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='面试邀约表';

-- ----------------------------
-- Records of sys_company_invitation
-- ----------------------------
INSERT INTO `sys_company_invitation` VALUES ('1', '2', '3', '1', '2', '2019-07-07 12:48:17.000000');
INSERT INTO `sys_company_invitation` VALUES ('2', '2', '3', '2', '1', '2019-07-07 12:49:15.000000');
INSERT INTO `sys_company_invitation` VALUES ('3', '2', '3', '3', '1', '2019-07-07 12:49:37.000000');
INSERT INTO `sys_company_invitation` VALUES ('4', '2', '3', '1', '1', '2019-07-10 00:55:14.000000');
INSERT INTO `sys_company_invitation` VALUES ('6', '2', '3', '1', '2', '2019-07-10 12:11:25.000000');
INSERT INTO `sys_company_invitation` VALUES ('7', '2', '4', '2', '0', '2019-07-10 12:11:34.000000');

-- ----------------------------
-- Table structure for sys_company_profession
-- ----------------------------
DROP TABLE IF EXISTS `sys_company_profession`;
CREATE TABLE `sys_company_profession` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '招聘id',
  `company_id` bigint(255) unsigned NOT NULL COMMENT '公司id',
  `profession_type_id` bigint(255) unsigned NOT NULL COMMENT '职业大类id',
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '标题',
  `introduction` varchar(255) DEFAULT NULL COMMENT '职位描述',
  `requirement` varchar(255) DEFAULT NULL COMMENT '职位要求',
  `salary` varchar(20) DEFAULT NULL COMMENT '薪水',
  `start_time` datetime(6) DEFAULT NULL COMMENT '招聘发布时间',
  `end_time` datetime(6) DEFAULT NULL COMMENT '招聘结束时间',
  `state` int(1) DEFAULT NULL COMMENT '招聘状态',
  `location` varchar(30) DEFAULT NULL COMMENT '工作地',
  `is_practice` tinyint(1) unsigned DEFAULT NULL COMMENT '是否是实习',
  `duration` varchar(30) DEFAULT NULL COMMENT '实习时间要求',
  `chance_to_formal` varchar(5) DEFAULT NULL COMMENT '转正机会',
  PRIMARY KEY (`id`,`company_id`,`profession_type_id`) USING BTREE,
  KEY `company` (`company_id`) USING BTREE,
  KEY `profession` (`profession_type_id`) USING BTREE,
  KEY `id` (`id`) USING BTREE,
  CONSTRAINT `fk_sys_company_profession_sys_company_1` FOREIGN KEY (`company_id`) REFERENCES `sys_company` (`id`),
  CONSTRAINT `profession` FOREIGN KEY (`profession_type_id`) REFERENCES `sys_profession_type` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='企业发布岗位表';

-- ----------------------------
-- Records of sys_company_profession
-- ----------------------------
INSERT INTO `sys_company_profession` VALUES ('1', '2', '1', '腾讯校招', '腾讯校招', '腾讯校招', '月薪：5000', '2019-07-04 15:22:44.000000', null, '0', '深圳', '1', '暑期', '有');
INSERT INTO `sys_company_profession` VALUES ('2', '3', '2', '腾讯校招', '腾讯校招', '腾讯校招', '月薪：1450', '2019-08-01 15:22:44.000000', null, '1', '深圳', '1', '暑期', '无');
INSERT INTO `sys_company_profession` VALUES ('3', '4', '3', '腾讯夏季校招', '软件程序工程师', '熟悉各种语言即可', '月薪：10000', '2019-09-01 09:40:47.000000', null, '1', '深圳', '1', '暑期', '有');
INSERT INTO `sys_company_profession` VALUES ('4', '2', '7', '腾讯体育招聘', '腾讯体育服务器维护工作', '熟悉数据库，和网络编程', '月薪：3000', '2019-09-01 09:38:18.000000', '2020-03-05 13:50:54.000000', '1', '北京', '1', '五个月', '有');
INSERT INTO `sys_company_profession` VALUES ('5', '4', '7', '腾讯体育招聘', '腾讯体育服务器维护工作', '熟悉数据库，和网络编程', '月薪：3000', '2019-09-01 09:38:18.000000', '2020-03-05 13:51:26.000000', '1', '北京', '1', '五个月', '有');
INSERT INTO `sys_company_profession` VALUES ('6', '3', '2', '腾讯校招', '腾讯校招', '腾讯校招', '月薪：1450', '2019-08-01 15:22:44.000000', '2020-03-05 13:56:10.000000', '1', '深圳', '1', '暑期', '无');
INSERT INTO `sys_company_profession` VALUES ('7', '4', '3', '腾讯夏季校招', '软件程序工程师', '熟悉各种语言即可', '月薪：10000', '2019-09-01 09:40:47.000000', '2020-03-05 13:57:04.000000', '1', '深圳', '1', '暑期', '有');
INSERT INTO `sys_company_profession` VALUES ('15', '2', '2', '腾讯微博测试工程师招聘', '测试腾讯微博系列产品', '熟练使用各类测试工具', '11', '2019-09-01 09:38:10.000000', null, '1', '杭州', '0', '111', '1111');
INSERT INTO `sys_company_profession` VALUES ('16', '7', '7', '腾讯体育招聘', '腾讯体育服务器维护工作', '熟悉数据库，和网络编程', '月薪：3000', '2019-09-01 09:38:18.000000', null, '1', '北京', '1', '五个月', '有');
INSERT INTO `sys_company_profession` VALUES ('17', '2', '1', '腾讯视频招聘', '腾讯视频算法岗位', '熟悉c++，了解网络编程', '月薪5000', '2019-10-01 09:38:23.000000', null, '1', '广州', '0', '23', '23');
INSERT INTO `sys_company_profession` VALUES ('18', '8', '1', 'Tencent外包招聘', '软件架构师', '吃苦耐劳', '月薪：15000', '2019-10-01 09:42:36.000000', null, '1', '北京', '0', null, null);

-- ----------------------------
-- Table structure for sys_course
-- ----------------------------
DROP TABLE IF EXISTS `sys_course`;
CREATE TABLE `sys_course` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '课程id',
  `teacher_id` bigint(255) unsigned NOT NULL COMMENT '发布者id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '课程名称',
  `introduction` varchar(255) NOT NULL DEFAULT '' COMMENT '课程介绍',
  `stars` float(50,0) NOT NULL DEFAULT '0' COMMENT '综合评分',
  `upload_time` datetime(6) NOT NULL COMMENT '课程开设时间',
  `picture_path` varchar(255) NOT NULL DEFAULT '\\course_picture\\default\\default.jpg' COMMENT '封面图片路径',
  `activity` tinyint(10) unsigned NOT NULL DEFAULT '1' COMMENT '是否激活',
  PRIMARY KEY (`id`,`teacher_id`) USING BTREE,
  KEY `fk_sys_course_sys_teacher_info_1` (`teacher_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_sys_teacher_info_1` FOREIGN KEY (`teacher_id`) REFERENCES `sys_teacher_info` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='课程表';

-- ----------------------------
-- Records of sys_course
-- ----------------------------
INSERT INTO `sys_course` VALUES ('1', '1', '独立完成企业级Java电商网站开发', 'Web', '3', '2019-05-09 00:00:00.000000', '\\course_picture\\2019\\58f57d200001461105400300-360-202.jpg', '1');
INSERT INTO `sys_course` VALUES ('2', '1', 'Java生产环境下调优详解', 'JAVA', '3', '2019-05-10 00:00:00.000000', '\\course_picture\\2019\\5b384772000132c405400300-360-202.jpg', '1');
INSERT INTO `sys_course` VALUES ('3', '1', 'go语言实战', 'go', '2', '2020-02-24 15:08:52.000000', '\\course_picture\\default\\default.jpg', '0');
INSERT INTO `sys_course` VALUES ('4', '1', 'python', 'python', '3', '2020-01-01 20:50:07.000000', '\\course_picture\\default\\default.jpg', '1');

-- ----------------------------
-- Table structure for sys_course_click
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_click`;
CREATE TABLE `sys_course_click` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '课程点击记录id',
  `user_id` bigint(255) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `course_id` bigint(255) unsigned NOT NULL DEFAULT '0' COMMENT '课程id',
  `click_time` datetime(6) NOT NULL COMMENT '点击时间',
  PRIMARY KEY (`id`,`user_id`,`course_id`) USING BTREE,
  KEY `fk_sys_course_click_sys_user_1` (`user_id`) USING BTREE,
  KEY `fk_sys_course_click_sys_course_1` (`course_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_click_sys_course_1` FOREIGN KEY (`course_id`) REFERENCES `sys_course` (`id`),
  CONSTRAINT `fk_sys_course_click_sys_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='课程点击记录表';

-- ----------------------------
-- Records of sys_course_click
-- ----------------------------
INSERT INTO `sys_course_click` VALUES ('1', '3', '1', '2019-05-20 11:09:34.000000');
INSERT INTO `sys_course_click` VALUES ('2', '4', '1', '2019-05-20 11:09:44.000000');
INSERT INTO `sys_course_click` VALUES ('3', '6', '2', '2019-05-20 11:09:57.000000');

-- ----------------------------
-- Table structure for sys_course_comment
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_comment`;
CREATE TABLE `sys_course_comment` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '评论者id',
  `course_id` bigint(255) unsigned NOT NULL COMMENT '课程id',
  `time` datetime(6) NOT NULL COMMENT '评论时间',
  `stars` float(50,0) NOT NULL DEFAULT '5' COMMENT '课程打分',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '评价内容',
  PRIMARY KEY (`id`,`course_id`,`user_id`) USING BTREE,
  KEY `fk_course_comment_user_1` (`user_id`) USING BTREE,
  KEY `fk_sys_course_comment_sys_course_1` (`course_id`) USING BTREE,
  CONSTRAINT `fk_course_comment_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`),
  CONSTRAINT `fk_sys_course_comment_sys_course_1` FOREIGN KEY (`course_id`) REFERENCES `sys_course` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='课程评价表';

-- ----------------------------
-- Records of sys_course_comment
-- ----------------------------
INSERT INTO `sys_course_comment` VALUES ('6', '3', '1', '2019-05-20 10:23:15.000000', '3', '还行！');
INSERT INTO `sys_course_comment` VALUES ('7', '3', '1', '2019-05-20 10:23:27.000000', '4', '二刷');
INSERT INTO `sys_course_comment` VALUES ('8', '1', '1', '2019-05-20 10:24:26.000000', '5', '太棒了');
INSERT INTO `sys_course_comment` VALUES ('10', '3', '2', '2019-05-27 22:10:28.000000', '2', 'just so so.');
INSERT INTO `sys_course_comment` VALUES ('11', '3', '1', '2019-05-27 22:11:08.000000', '4', '6666');
INSERT INTO `sys_course_comment` VALUES ('12', '3', '2', '2019-06-04 17:46:17.000000', '4', '还行');
INSERT INTO `sys_course_comment` VALUES ('13', '3', '1', '2019-06-04 18:00:34.000000', '3', '666');
INSERT INTO `sys_course_comment` VALUES ('14', '3', '2', '2019-06-04 18:15:18.000000', '5', 'six');
INSERT INTO `sys_course_comment` VALUES ('15', '3', '2', '2019-06-04 18:21:03.000000', '1', 'zzzz');
INSERT INTO `sys_course_comment` VALUES ('16', '3', '1', '2019-06-04 18:38:19.000000', '1', '测试');

-- ----------------------------
-- Table structure for sys_course_contest
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_contest`;
CREATE TABLE `sys_course_contest` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '课程单元测试id',
  `course_index_id` bigint(255) unsigned NOT NULL COMMENT '所属课程单元id',
  `contest_name` varchar(255) DEFAULT NULL COMMENT '测试标题',
  PRIMARY KEY (`id`,`course_index_id`) USING BTREE,
  KEY `fk_sys_course_contest_sys_course_index_1` (`course_index_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_contest_sys_course_index_1` FOREIGN KEY (`course_index_id`) REFERENCES `sys_course_index` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='课程单元测试表';

-- ----------------------------
-- Records of sys_course_contest
-- ----------------------------

-- ----------------------------
-- Table structure for sys_course_contest_option
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_contest_option`;
CREATE TABLE `sys_course_contest_option` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '选项id',
  `question_id` bigint(255) unsigned NOT NULL COMMENT '问题id',
  `option_content` varchar(255) NOT NULL COMMENT '选项内容',
  PRIMARY KEY (`id`,`question_id`) USING BTREE,
  KEY `fk_sys_course_contest_option_sys_course_contest_question_1` (`question_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_contest_option_sys_course_contest_question_1` FOREIGN KEY (`question_id`) REFERENCES `sys_course_contest_question` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='选项表';

-- ----------------------------
-- Records of sys_course_contest_option
-- ----------------------------

-- ----------------------------
-- Table structure for sys_course_contest_question
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_contest_question`;
CREATE TABLE `sys_course_contest_question` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '题目id',
  `contest_question_id` bigint(255) unsigned NOT NULL COMMENT '所属课程单元测试id',
  `correct_option_id` bigint(255) unsigned NOT NULL COMMENT '正确答案id',
  `question_content` varchar(255) NOT NULL COMMENT '问题内容',
  PRIMARY KEY (`id`,`contest_question_id`,`correct_option_id`) USING BTREE,
  KEY `fk_sys_course_contest_question_sys_course_contest_1` (`contest_question_id`) USING BTREE,
  KEY `fk_sys_course_contest_question_sys_course_contest_option_1` (`correct_option_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_contest_question_sys_course_contest_1` FOREIGN KEY (`contest_question_id`) REFERENCES `sys_course_contest` (`id`),
  CONSTRAINT `fk_sys_course_contest_question_sys_course_contest_option_1` FOREIGN KEY (`correct_option_id`) REFERENCES `sys_course_contest_option` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='问题表';

-- ----------------------------
-- Records of sys_course_contest_question
-- ----------------------------

-- ----------------------------
-- Table structure for sys_course_contest_result
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_contest_result`;
CREATE TABLE `sys_course_contest_result` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '单元测试结果id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `course_contest_id` bigint(255) unsigned NOT NULL COMMENT '所属课程测试id',
  `start_time` datetime(6) DEFAULT NULL COMMENT '开始测试时间',
  `end_time` datetime(6) DEFAULT NULL COMMENT '结束测试时间',
  `score` int(255) unsigned DEFAULT NULL COMMENT '测试得分',
  PRIMARY KEY (`id`,`user_id`,`course_contest_id`) USING BTREE,
  KEY `fk_sys_course_contest_result_sys_course_contest_1` (`course_contest_id`) USING BTREE,
  KEY `fk_sys_course_contest_result_sys_user_1` (`user_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_contest_result_sys_course_contest_1` FOREIGN KEY (`course_contest_id`) REFERENCES `sys_course_contest` (`id`),
  CONSTRAINT `fk_sys_course_contest_result_sys_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='单元测试结果表';

-- ----------------------------
-- Records of sys_course_contest_result
-- ----------------------------

-- ----------------------------
-- Table structure for sys_course_contest_score
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_contest_score`;
CREATE TABLE `sys_course_contest_score` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '单元测试成绩id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `course_index_id` bigint(255) unsigned NOT NULL COMMENT '所属课程单元id',
  `score` float(255,0) NOT NULL DEFAULT '0' COMMENT '测试成绩',
  `take_time` datetime(6) NOT NULL COMMENT '参考时间',
  PRIMARY KEY (`id`,`user_id`,`course_index_id`) USING BTREE,
  KEY `fk_sys_course_contest_score_sys_course_index_1` (`course_index_id`) USING BTREE,
  KEY `fk_sys_course_contest_score_sys_user_1` (`user_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_contest_score_sys_course_index_1` FOREIGN KEY (`course_index_id`) REFERENCES `sys_course_index` (`id`),
  CONSTRAINT `fk_sys_course_contest_score_sys_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='单元测试成绩表';

-- ----------------------------
-- Records of sys_course_contest_score
-- ----------------------------

-- ----------------------------
-- Table structure for sys_course_contest_user_answer
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_contest_user_answer`;
CREATE TABLE `sys_course_contest_user_answer` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '单元测试答题id',
  `contest_result_id` bigint(255) unsigned NOT NULL COMMENT '单元测试结果id',
  `question_id` bigint(255) unsigned NOT NULL COMMENT '问题id',
  `answer_id` bigint(255) unsigned NOT NULL COMMENT '回答选项id',
  PRIMARY KEY (`id`,`contest_result_id`) USING BTREE,
  KEY `fk_sys_course_contest_user_answer_sys_course_contest_question_1` (`question_id`) USING BTREE,
  KEY `fk_sys_course_contest_user_answer_sys_course_contest_option_1` (`answer_id`) USING BTREE,
  KEY `fk_sys_course_contest_user_answer_sys_course_contest_result_1` (`contest_result_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_contest_user_answer_sys_course_contest_option_1` FOREIGN KEY (`answer_id`) REFERENCES `sys_course_contest_option` (`id`),
  CONSTRAINT `fk_sys_course_contest_user_answer_sys_course_contest_question_1` FOREIGN KEY (`question_id`) REFERENCES `sys_course_contest_question` (`id`),
  CONSTRAINT `fk_sys_course_contest_user_answer_sys_course_contest_result_1` FOREIGN KEY (`contest_result_id`) REFERENCES `sys_course_contest_result` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户单元测试答题表';

-- ----------------------------
-- Records of sys_course_contest_user_answer
-- ----------------------------

-- ----------------------------
-- Table structure for sys_course_index
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_index`;
CREATE TABLE `sys_course_index` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '课程目录id',
  `course_id` bigint(255) unsigned NOT NULL COMMENT '所属课程id',
  `scale` int(255) NOT NULL COMMENT '目录等级',
  `index_order` int(255) NOT NULL COMMENT '目录位置',
  `index_title` varchar(255) NOT NULL DEFAULT '' COMMENT '目录标题',
  PRIMARY KEY (`id`,`course_id`) USING BTREE,
  KEY `fk_sys_course_index_sys_course_1` (`course_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_index_sys_course_1` FOREIGN KEY (`course_id`) REFERENCES `sys_course` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='课程目录表';

-- ----------------------------
-- Records of sys_course_index
-- ----------------------------
INSERT INTO `sys_course_index` VALUES ('1', '1', '1', '1', 'JAVA语言');
INSERT INTO `sys_course_index` VALUES ('2', '1', '2', '1', 'JAVA语言');
INSERT INTO `sys_course_index` VALUES ('3', '1', '2', '2', 'JAVA语言');
INSERT INTO `sys_course_index` VALUES ('4', '1', '2', '3', 'JAVA语言');
INSERT INTO `sys_course_index` VALUES ('5', '1', '1', '2', 'JAVA语言');
INSERT INTO `sys_course_index` VALUES ('6', '1', '1', '3', 'JAVA语言');
INSERT INTO `sys_course_index` VALUES ('7', '1', '2', '1', 'JAVA语言');
INSERT INTO `sys_course_index` VALUES ('8', '1', '2', '2', 'JAVA语言');
INSERT INTO `sys_course_index` VALUES ('9', '1', '1', '4', 'JAVA语言');

-- ----------------------------
-- Table structure for sys_course_recommend
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_recommend`;
CREATE TABLE `sys_course_recommend` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '推荐id',
  `profession_id` bigint(255) unsigned NOT NULL COMMENT '职业id',
  `curriculum_id` bigint(255) unsigned NOT NULL COMMENT '科目分类id',
  PRIMARY KEY (`id`,`profession_id`,`curriculum_id`) USING BTREE,
  KEY `fk_sys_course_recommend_sys_profession_type_1` (`profession_id`) USING BTREE,
  KEY `fk_sys_course_recommend_sys_curriculum_1` (`curriculum_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_recommend_sys_curriculum_1` FOREIGN KEY (`curriculum_id`) REFERENCES `sys_curriculum` (`id`),
  CONSTRAINT `fk_sys_course_recommend_sys_profession_type_1` FOREIGN KEY (`profession_id`) REFERENCES `sys_profession` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='职业推荐课程表';

-- ----------------------------
-- Records of sys_course_recommend
-- ----------------------------
INSERT INTO `sys_course_recommend` VALUES ('1', '5', '1');
INSERT INTO `sys_course_recommend` VALUES ('2', '5', '2');
INSERT INTO `sys_course_recommend` VALUES ('3', '5', '3');
INSERT INTO `sys_course_recommend` VALUES ('4', '5', '4');
INSERT INTO `sys_course_recommend` VALUES ('5', '5', '6');

-- ----------------------------
-- Table structure for sys_course_studied
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_studied`;
CREATE TABLE `sys_course_studied` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户学习课程记录id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `course_id` bigint(255) unsigned NOT NULL COMMENT '课程id',
  `score` float(100,0) DEFAULT '0' COMMENT '课程综合成绩',
  `is_finished` tinyint(10) NOT NULL DEFAULT '0' COMMENT '是否完成学习',
  `start_time` datetime(6) NOT NULL COMMENT '开始学习时间',
  PRIMARY KEY (`id`,`user_id`,`course_id`) USING BTREE,
  KEY `fk_sys_course_studied_sys_user_1` (`user_id`) USING BTREE,
  KEY `fk_sys_course_studied_sys_course_1` (`course_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_studied_sys_course_1` FOREIGN KEY (`course_id`) REFERENCES `sys_course` (`id`),
  CONSTRAINT `fk_sys_course_studied_sys_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户学习课程记录表';

-- ----------------------------
-- Records of sys_course_studied
-- ----------------------------
INSERT INTO `sys_course_studied` VALUES ('1', '8', '1', '80', '1', '2019-06-09 09:29:02.000000');
INSERT INTO `sys_course_studied` VALUES ('2', '4', '1', '50', '0', '2019-04-09 09:29:44.000000');
INSERT INTO `sys_course_studied` VALUES ('3', '4', '2', '60', '1', '2019-09-09 09:30:33.000000');
INSERT INTO `sys_course_studied` VALUES ('4', '6', '1', '30', '0', '2019-02-09 09:30:52.000000');
INSERT INTO `sys_course_studied` VALUES ('5', '7', '2', '60', '1', '2019-03-09 09:31:50.000000');
INSERT INTO `sys_course_studied` VALUES ('6', '3', '1', '68', '1', '2019-04-24 00:00:00.000000');
INSERT INTO `sys_course_studied` VALUES ('6', '3', '2', '65', '1', '2019-12-25 20:20:17.000000');
INSERT INTO `sys_course_studied` VALUES ('7', '10', '1', '55', '0', '2019-12-06 00:00:00.000000');

-- ----------------------------
-- Table structure for sys_course_video
-- ----------------------------
DROP TABLE IF EXISTS `sys_course_video`;
CREATE TABLE `sys_course_video` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `course_Index_id` bigint(255) unsigned NOT NULL COMMENT '所属课程id',
  `video_title` varchar(255) NOT NULL DEFAULT '' COMMENT '课程标题',
  `update_time` datetime(6) NOT NULL DEFAULT '2018-12-05 00:00:00.000000' COMMENT '发布时间',
  `video_len` bigint(255) unsigned NOT NULL DEFAULT '30' COMMENT '视频时长',
  `video_path` varchar(255) NOT NULL DEFAULT '\\course_view\\default\\001.mkv' COMMENT '视频文件路径',
  PRIMARY KEY (`id`,`course_Index_id`) USING BTREE,
  KEY `fk_sys_course_video_sys_course_index_1` (`course_Index_id`) USING BTREE,
  CONSTRAINT `fk_sys_course_video_sys_course_index_1` FOREIGN KEY (`course_Index_id`) REFERENCES `sys_course_index` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='课程视频表';

-- ----------------------------
-- Records of sys_course_video
-- ----------------------------
INSERT INTO `sys_course_video` VALUES ('1', '2', 'JAVA语言', '2018-12-05 00:00:00.000000', '30', '\\course_view\\default\\001.mkv');
INSERT INTO `sys_course_video` VALUES ('2', '3', 'JAVA语言', '2018-12-05 00:00:00.000000', '30', '\\course_view\\default\\001.mkv');
INSERT INTO `sys_course_video` VALUES ('3', '4', '', '2018-12-05 00:00:00.000000', '30', '\\course_view\\default\\001.mkv');
INSERT INTO `sys_course_video` VALUES ('4', '7', '', '2018-12-05 00:00:00.000000', '30', '\\course_view\\default\\001.mkv');
INSERT INTO `sys_course_video` VALUES ('5', '8', '', '2018-12-05 00:00:00.000000', '30', '\\course_view\\default\\001.mkv');

-- ----------------------------
-- Table structure for sys_curriculum
-- ----------------------------
DROP TABLE IF EXISTS `sys_curriculum`;
CREATE TABLE `sys_curriculum` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '科目id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '科目名称',
  `introduction` varchar(255) NOT NULL DEFAULT '' COMMENT '科目介绍',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='科目表';

-- ----------------------------
-- Records of sys_curriculum
-- ----------------------------
INSERT INTO `sys_curriculum` VALUES ('1', 'JAVA语言', 'JAVA基础');
INSERT INTO `sys_curriculum` VALUES ('2', 'JavaWeb', 'Web基础');
INSERT INTO `sys_curriculum` VALUES ('3', '数据结构', '数据结构');
INSERT INTO `sys_curriculum` VALUES ('4', '数据库', '数据库基础');
INSERT INTO `sys_curriculum` VALUES ('5', 'PHP语言', 'PHP基础');
INSERT INTO `sys_curriculum` VALUES ('6', 'C语言', 'C语言');
INSERT INTO `sys_curriculum` VALUES ('7', '软件工程', '软件工程');
INSERT INTO `sys_curriculum` VALUES ('8', 'python语言', 'python语言');
INSERT INTO `sys_curriculum` VALUES ('9', '数据分析', '数据分析');
INSERT INTO `sys_curriculum` VALUES ('10', '爬虫技术', '爬虫技术');
INSERT INTO `sys_curriculum` VALUES ('11', '人工智能', '人工智能');
INSERT INTO `sys_curriculum` VALUES ('12', 'C++语言', 'C++语言');
INSERT INTO `sys_curriculum` VALUES ('13', 'Go语言', 'Go语言');

-- ----------------------------
-- Table structure for sys_curriculum_course
-- ----------------------------
DROP TABLE IF EXISTS `sys_curriculum_course`;
CREATE TABLE `sys_curriculum_course` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '关系id',
  `course_id` bigint(255) unsigned NOT NULL COMMENT '课程id',
  `curriculum_id` bigint(255) unsigned NOT NULL COMMENT '科目id',
  PRIMARY KEY (`id`,`curriculum_id`,`course_id`) USING BTREE,
  KEY `fk_curriculum_course_curriculum_1` (`curriculum_id`) USING BTREE,
  KEY `fk_curriculum_course_course_1` (`course_id`) USING BTREE,
  CONSTRAINT `fk_curriculum_course_course_1` FOREIGN KEY (`course_id`) REFERENCES `sys_course` (`id`),
  CONSTRAINT `fk_curriculum_course_curriculum_1` FOREIGN KEY (`curriculum_id`) REFERENCES `sys_curriculum` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='科目_课程关系表';

-- ----------------------------
-- Records of sys_curriculum_course
-- ----------------------------
INSERT INTO `sys_curriculum_course` VALUES ('0', '1', '1');
INSERT INTO `sys_curriculum_course` VALUES ('1', '2', '2');
INSERT INTO `sys_curriculum_course` VALUES ('3', '4', '8');
INSERT INTO `sys_curriculum_course` VALUES ('2', '3', '13');

-- ----------------------------
-- Table structure for sys_curriculum_tree
-- ----------------------------
DROP TABLE IF EXISTS `sys_curriculum_tree`;
CREATE TABLE `sys_curriculum_tree` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '结点id',
  `curriculum_id` bigint(255) unsigned NOT NULL COMMENT '科目id',
  `pre_curriculum_id` bigint(255) unsigned NOT NULL COMMENT '前置科目id',
  PRIMARY KEY (`id`,`curriculum_id`,`pre_curriculum_id`) USING BTREE,
  KEY `fk_curriculum_tree_curriculum_1` (`curriculum_id`) USING BTREE,
  KEY `fk_curriculum_tree_curriculum_2` (`pre_curriculum_id`) USING BTREE,
  CONSTRAINT `fk_curriculum_tree_curriculum_1` FOREIGN KEY (`curriculum_id`) REFERENCES `sys_curriculum` (`id`),
  CONSTRAINT `fk_curriculum_tree_curriculum_2` FOREIGN KEY (`pre_curriculum_id`) REFERENCES `sys_curriculum` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='课程树表';

-- ----------------------------
-- Records of sys_curriculum_tree
-- ----------------------------
INSERT INTO `sys_curriculum_tree` VALUES ('3', '1', '3');
INSERT INTO `sys_curriculum_tree` VALUES ('1', '2', '1');
INSERT INTO `sys_curriculum_tree` VALUES ('2', '2', '4');
INSERT INTO `sys_curriculum_tree` VALUES ('6', '2', '3');
INSERT INTO `sys_curriculum_tree` VALUES ('7', '2', '7');
INSERT INTO `sys_curriculum_tree` VALUES ('4', '3', '6');
INSERT INTO `sys_curriculum_tree` VALUES ('5', '5', '3');
INSERT INTO `sys_curriculum_tree` VALUES ('12', '8', '3');
INSERT INTO `sys_curriculum_tree` VALUES ('8', '9', '8');
INSERT INTO `sys_curriculum_tree` VALUES ('9', '10', '8');
INSERT INTO `sys_curriculum_tree` VALUES ('10', '11', '8');
INSERT INTO `sys_curriculum_tree` VALUES ('11', '12', '6');
INSERT INTO `sys_curriculum_tree` VALUES ('13', '13', '3');

-- ----------------------------
-- Table structure for sys_learner_cv
-- ----------------------------
DROP TABLE IF EXISTS `sys_learner_cv`;
CREATE TABLE `sys_learner_cv` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '简历id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `resume_name` varchar(50) DEFAULT NULL COMMENT '简历名称',
  `name` varchar(20) DEFAULT NULL COMMENT '姓名',
  `sex` int(1) DEFAULT NULL COMMENT '性别',
  `birth` datetime(6) DEFAULT NULL COMMENT '出生年月',
  `native_place` varchar(10) DEFAULT NULL COMMENT '籍贯',
  `identity` varchar(50) DEFAULT NULL COMMENT '身份证号',
  `qualification` varchar(10) DEFAULT NULL COMMENT '学历',
  `specialty` varchar(20) DEFAULT NULL COMMENT '专业',
  `university` varchar(50) DEFAULT NULL COMMENT '毕业院校',
  `tel` char(20) DEFAULT NULL COMMENT '电话',
  `experience` varchar(255) DEFAULT NULL COMMENT '实习经历',
  `mailbox` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `introduction` varchar(255) DEFAULT NULL COMMENT '自我介绍',
  `diploma` varchar(255) DEFAULT NULL COMMENT '获奖情况',
  `current_location` varchar(100) DEFAULT NULL COMMENT '当前所在地',
  `in_job_time` varchar(50) DEFAULT NULL COMMENT '最快入职时间',
  `head_path` varchar(255) NOT NULL DEFAULT '\\avatar\\default\\avatar.jpg' COMMENT '证件照',
  `last_edit_time` datetime(6) DEFAULT NULL COMMENT '上次修改简历时间',
  PRIMARY KEY (`id`,`user_id`) USING BTREE,
  UNIQUE KEY `fk_sys_learner_cv_id` (`id`) USING BTREE,
  KEY `fk_sys_learner_cv_user_id` (`user_id`) USING BTREE,
  CONSTRAINT `fk_sys_learner_cv_user_id` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='简历表';

-- ----------------------------
-- Records of sys_learner_cv
-- ----------------------------
INSERT INTO `sys_learner_cv` VALUES ('22', '3', '简历2', '姜云杰', '0', '1998-07-08 00:00:00.000000', '安徽省', '11111111111111111111111', '3', '计算机科学与技术', '上海理工大学', '18321882894', '啦啦啦', '784203082@qq.com', '啦啦啦', '啦啦啦', null, null, '\\learner_cv_avatar\\2019\\7b5982d9-cf9f-47fe-b3bc-019666aa7029.png', '2019-07-10 11:18:45.000000');

-- ----------------------------
-- Table structure for sys_learner_info
-- ----------------------------
DROP TABLE IF EXISTS `sys_learner_info`;
CREATE TABLE `sys_learner_info` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '学习者id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `nickname` varchar(255) NOT NULL COMMENT '昵称',
  `realname` varchar(255) DEFAULT NULL COMMENT '真实姓名',
  `sex` varchar(10) DEFAULT '男' COMMENT '性别',
  `age` int(20) unsigned DEFAULT '0' COMMENT '年龄',
  `signature` varchar(255) DEFAULT NULL COMMENT '个性签名',
  `vacation` varchar(30) DEFAULT '本科一年级' COMMENT '学历',
  `github` varchar(255) DEFAULT NULL COMMENT 'github帐号',
  `email` varchar(255) DEFAULT NULL COMMENT '电子邮箱',
  `phone` varchar(255) DEFAULT NULL COMMENT '电话号码',
  `hobby` varchar(255) DEFAULT NULL COMMENT '爱好',
  `professional` varchar(255) DEFAULT NULL COMMENT '目标职业',
  `learner_cv_id` bigint(255) unsigned DEFAULT '0' COMMENT '学习者默认简历id',
  PRIMARY KEY (`id`,`user_id`) USING BTREE,
  KEY `fk_sys_learner_info_sys_user_1` (`user_id`) USING BTREE,
  KEY `id` (`id`) USING BTREE,
  CONSTRAINT `fk_sys_learner_info_sys_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='学习者信息表';

-- ----------------------------
-- Records of sys_learner_info
-- ----------------------------
INSERT INTO `sys_learner_info` VALUES ('1', '3', 'T49Jml', '姜云杰', '男', '20', '996', '本科一年级', '123', '123', '123', '吃喝玩乐', null, '22');
INSERT INTO `sys_learner_info` VALUES ('2', '4', '猪大哥', '候懿', '男', '38', '我是猪大哥', '本科一年级', null, null, null, null, null, null);
INSERT INTO `sys_learner_info` VALUES ('12', '5', 'Dds45', '魏智伟', '男', '0', null, '本科一年级', null, null, null, null, null, '0');
INSERT INTO `sys_learner_info` VALUES ('13', '6', 'DAFjj55', '谢志阳', '男', '0', null, '本科一年级', null, null, null, null, null, '0');
INSERT INTO `sys_learner_info` VALUES ('14', '7', 'kjjrER', '陆施涛', '男', '0', null, '本科一年级', null, null, null, null, null, '0');
INSERT INTO `sys_learner_info` VALUES ('15', '8', '44fgd', '洪林建', '男', '0', null, '本科一年级', null, null, null, null, null, '0');
INSERT INTO `sys_learner_info` VALUES ('16', '10', 'dffere', '金轩城', '男', '0', null, '本科一年级', null, null, null, null, null, '0');
INSERT INTO `sys_learner_info` VALUES ('17', '41', '富婆', '李尚昱', '女', '18', null, '本科一年级', 'sds', 'sad', 'asdas', 'ewqe', null, '0');

-- ----------------------------
-- Table structure for sys_message
-- ----------------------------
DROP TABLE IF EXISTS `sys_message`;
CREATE TABLE `sys_message` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '消息id',
  `sender_id` bigint(255) unsigned NOT NULL COMMENT '发送者id',
  `receiver_id` bigint(255) unsigned NOT NULL COMMENT '接受者id',
  `send_time` datetime(6) NOT NULL COMMENT '发送时间',
  `message_content` varchar(255) NOT NULL COMMENT '消息内容',
  `readed` tinyint(10) NOT NULL DEFAULT '0',
  `type` tinyint(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`,`sender_id`,`receiver_id`) USING BTREE,
  KEY `fk_sys_message_sys_user_1` (`sender_id`) USING BTREE,
  KEY `fk_sys_message_sys_user_2` (`receiver_id`) USING BTREE,
  CONSTRAINT `fk_sys_message_sys_user_1` FOREIGN KEY (`sender_id`) REFERENCES `sys_user` (`id`),
  CONSTRAINT `fk_sys_message_sys_user_2` FOREIGN KEY (`receiver_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='消息记录表';

-- ----------------------------
-- Records of sys_message
-- ----------------------------

-- ----------------------------
-- Table structure for sys_profession
-- ----------------------------
DROP TABLE IF EXISTS `sys_profession`;
CREATE TABLE `sys_profession` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '职业id',
  `name` varchar(255) NOT NULL COMMENT '职业名称',
  `salary` float(255,0) NOT NULL COMMENT '职业薪资',
  `type_id` bigint(255) unsigned NOT NULL COMMENT '职业分类id',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_sys_professtion_sys_profession_type_1` (`type_id`) USING BTREE,
  CONSTRAINT `fk_sys_professtion_sys_profession_type_1` FOREIGN KEY (`type_id`) REFERENCES `sys_profession_type` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='职业表';

-- ----------------------------
-- Records of sys_profession
-- ----------------------------
INSERT INTO `sys_profession` VALUES ('1', 'php前端工程师', '15000', '5');
INSERT INTO `sys_profession` VALUES ('2', 'php后端工程师', '20000', '2');
INSERT INTO `sys_profession` VALUES ('3', 'MySql工程师', '25000', '3');
INSERT INTO `sys_profession` VALUES ('4', 'Oracle工程师', '25000', '3');
INSERT INTO `sys_profession` VALUES ('5', 'Java工程师', '15000', '4');
INSERT INTO `sys_profession` VALUES ('6', 'Spring开发工程师', '15000', '6');
INSERT INTO `sys_profession` VALUES ('7', 'goland架构工程师', '25000', '1');

-- ----------------------------
-- Table structure for sys_profession_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_profession_type`;
CREATE TABLE `sys_profession_type` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '职业id',
  `type_name` varchar(255) NOT NULL COMMENT '职业分类名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='职业分类';

-- ----------------------------
-- Records of sys_profession_type
-- ----------------------------
INSERT INTO `sys_profession_type` VALUES ('1', '研发');
INSERT INTO `sys_profession_type` VALUES ('2', '测试');
INSERT INTO `sys_profession_type` VALUES ('3', '数据');
INSERT INTO `sys_profession_type` VALUES ('4', '算法');
INSERT INTO `sys_profession_type` VALUES ('5', '前端');
INSERT INTO `sys_profession_type` VALUES ('6', '产品');
INSERT INTO `sys_profession_type` VALUES ('7', '运营');
INSERT INTO `sys_profession_type` VALUES ('8', '其他');

-- ----------------------------
-- Table structure for sys_teacher_info
-- ----------------------------
DROP TABLE IF EXISTS `sys_teacher_info`;
CREATE TABLE `sys_teacher_info` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '讲师id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '讲师用户账号id',
  `realname` varchar(255) DEFAULT NULL COMMENT '真实姓名',
  `sex` varchar(10) DEFAULT '男' COMMENT '性别',
  `age` int(20) unsigned DEFAULT '0' COMMENT '年龄',
  `identity` varchar(255) DEFAULT NULL COMMENT '身份证号',
  `phone` varchar(255) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `tag` varchar(255) DEFAULT '' COMMENT '讲师头衔',
  `introduction` varchar(255) DEFAULT '' COMMENT '经历简介',
  `github` varchar(255) DEFAULT NULL COMMENT 'github帐号',
  PRIMARY KEY (`id`,`user_id`) USING BTREE,
  KEY `fk_teacher_info_user_1` (`user_id`) USING BTREE,
  CONSTRAINT `fk_teacher_info_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='讲师信息登记表';

-- ----------------------------
-- Records of sys_teacher_info
-- ----------------------------
INSERT INTO `sys_teacher_info` VALUES ('1', '1', '陆施涛', '男', null, '3434************', '152**********', '78********', '天才', '天才', '没有');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `account` varchar(255) NOT NULL DEFAULT '' COMMENT '用户账号',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',
  `head_path` varchar(255) NOT NULL DEFAULT '\\avatar\\default\\avatar.jpg' COMMENT '头像路径',
  `role` tinyint(10) NOT NULL DEFAULT '3' COMMENT '身份标记（讲师/学生/公司）',
  `activity` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '帐号是否激活',
  `createtime` datetime(6) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`,`account`) USING BTREE,
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'admin', 'admin', '\\avatar\\default\\avatar.jpg', '2', '1', '2020-02-22 00:31:48.000000');
INSERT INTO `sys_user` VALUES ('3', '1617440712', '1617440712', '\\avatar\\2019\\18c444c3-c532-4dca-a6b8-ca7aceef48cf.png', '3', '1', '2020-02-19 00:31:51.000000');
INSERT INTO `sys_user` VALUES ('4', '123', 'J123456789', '\\avatar\\default\\avatar.jpg', '3', '1', '2019-12-03 00:31:54.000000');
INSERT INTO `sys_user` VALUES ('5', 'jmedal', '1617440712J', '\\avatar\\default\\avatar.jpg', '3', '1', '2017-01-22 00:32:48.000000');
INSERT INTO `sys_user` VALUES ('6', '789147', '123456789A', '\\avatar\\default\\avatar.jpg', '3', '1', '2018-12-19 00:32:01.000000');
INSERT INTO `sys_user` VALUES ('7', '16175', 'A1617511', '\\avatar\\default\\avatar.jpg', '3', '1', '2020-02-22 00:32:15.000000');
INSERT INTO `sys_user` VALUES ('8', '156786', 'AA156786', '\\avatar\\default\\avatar.jpg', '3', '0', '2019-11-01 00:32:19.000000');
INSERT INTO `sys_user` VALUES ('10', 'lst', 'Lstlstlst', '\\avatar\\2019\\778c8477-4dc7-469c-b6aa-caa14abb787a.png', '3', '0', '2019-01-01 00:00:00.000000');
INSERT INTO `sys_user` VALUES ('11', 'jd', 'jingdong', '\\avatar\\2019\\d51ba82a-ac54-4845-8b0c-4cd78f99f099.png', '1', '1', '2020-02-22 00:32:31.000000');
INSERT INTO `sys_user` VALUES ('12', 'tencent', 'tencent', '\\avatar\\2019\\4df89cbf-f8d3-49b5-881f-549ef8a6d9d8.jpg', '1', '0', '2016-01-22 00:32:33.000000');
INSERT INTO `sys_user` VALUES ('13', 'alibaba', 'alibaba', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:53:28.000000');
INSERT INTO `sys_user` VALUES ('14', 'myjf', 'myjf', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('15', 'mt', 'mt', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('16', 'xiaomi', 'xiaomi', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('17', 'pdd', 'pdd', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('18', 'wy', 'wy', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('19', 'didi', 'didi', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('20', 'zztd', 'zztd', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('21', 'huya', 'huya', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('22', 'kugou', 'kugou', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('23', 'haluo', 'haluo', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('24', 'bilibili', 'bilibili', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('24', 'weimeng', 'weimeng', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('25', 'xhs', 'xhs', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('26', 'cnwl', 'cnwl', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('26', 'qiy', 'qiy', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('27', 'oppo', 'oppo', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('27', 'qiy', 'qiy', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('28', 'oppo', 'oppo', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('29', 'huawei', 'huawei', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('29', 'sf', 'sf', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('30', 'huawei', 'huawei', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('31', 'beike', 'beike', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('32', 'dywl', 'dywl', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('32', 'zyb', 'zyb', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('33', 'dywl', 'dywl', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('33', 'xhjy', 'xhjy', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('34', 'xhjy', 'xhjy', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('35', 'gscd', 'gscd', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('36', 'mgtv', 'mgtv', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('38', 'znkj', 'znkj', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('39', 'zyzx', 'zyzx', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('40', 'baidu', 'baidu', '\\avatar\\default\\avatar.jpg', '1', '1', '2018-01-22 00:32:42.000000');
INSERT INTO `sys_user` VALUES ('40', 'haixue', 'haixue', '\\avatar\\default\\avatar.jpg', '1', '1', '2020-03-05 11:57:18.000000');
INSERT INTO `sys_user` VALUES ('41', 'lsy', 'lsy', '\\avatar\\default\\avatar.jpg', '3', '1', '2020-02-23 12:03:44.000000');

-- ----------------------------
-- Table structure for sys_user_profession
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_profession`;
CREATE TABLE `sys_user_profession` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户选择id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `profession_id` bigint(255) unsigned NOT NULL COMMENT '职业id',
  PRIMARY KEY (`id`,`user_id`,`profession_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `profession_id` (`profession_id`) USING BTREE,
  CONSTRAINT `profession_id` FOREIGN KEY (`profession_id`) REFERENCES `sys_profession` (`id`),
  CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户选择职业表';

-- ----------------------------
-- Records of sys_user_profession
-- ----------------------------
INSERT INTO `sys_user_profession` VALUES ('1', '3', '5');
INSERT INTO `sys_user_profession` VALUES ('4', '3', '5');
INSERT INTO `sys_user_profession` VALUES ('2', '4', '3');
INSERT INTO `sys_user_profession` VALUES ('7', '4', '5');
INSERT INTO `sys_user_profession` VALUES ('3', '5', '4');
INSERT INTO `sys_user_profession` VALUES ('6', '5', '1');
INSERT INTO `sys_user_profession` VALUES ('5', '6', '4');
INSERT INTO `sys_user_profession` VALUES ('8', '7', '2');

-- ----------------------------
-- Table structure for sys_video_discussion
-- ----------------------------
DROP TABLE IF EXISTS `sys_video_discussion`;
CREATE TABLE `sys_video_discussion` (
  `id` bigint(255) NOT NULL AUTO_INCREMENT COMMENT '讨论id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `video_id` bigint(255) unsigned NOT NULL COMMENT '视频id',
  `reply_to_id` bigint(255) NOT NULL COMMENT '回复对象id（0：回复视频；非0：回复评论（值为最上级评论id））',
  `reply_user_id` bigint(255) unsigned NOT NULL COMMENT '回复用户id（0：回复视频或者回复最上级评论；非0：回复评论（值为被回复的用户id））',
  `reply_time` datetime(6) NOT NULL COMMENT '回答时间',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '讨论内容',
  PRIMARY KEY (`id`,`user_id`,`video_id`,`reply_to_id`,`reply_user_id`) USING BTREE,
  KEY `fk_video_discussion_user_1` (`user_id`) USING BTREE,
  KEY `fk_sys_video_discussion_sys_course_video_1` (`video_id`) USING BTREE,
  KEY `fk_sys_video_discussion_sys_user_1` (`reply_user_id`) USING BTREE,
  KEY `id` (`id`) USING BTREE,
  KEY `fk_video_discussion_user_2` (`reply_to_id`) USING BTREE,
  CONSTRAINT `fk_video_discussion_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='讨论表';

-- ----------------------------
-- Records of sys_video_discussion
-- ----------------------------
INSERT INTO `sys_video_discussion` VALUES ('33', '1', '1', '0', '0', '2019-05-20 10:33:42.000000', '牛逼');
INSERT INTO `sys_video_discussion` VALUES ('34', '1', '1', '33', '1', '2019-05-20 10:33:47.000000', '牛逼');
INSERT INTO `sys_video_discussion` VALUES ('35', '3', '1', '33', '1', '2019-05-20 10:34:21.000000', '牛逼');
INSERT INTO `sys_video_discussion` VALUES ('36', '3', '1', '33', '1', '2019-05-20 10:34:28.000000', '牛逼');
INSERT INTO `sys_video_discussion` VALUES ('37', '3', '1', '0', '0', '2019-05-20 10:34:36.000000', '牛逼');

-- ----------------------------
-- Table structure for sys_video_watched
-- ----------------------------
DROP TABLE IF EXISTS `sys_video_watched`;
CREATE TABLE `sys_video_watched` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '播放记录id',
  `user_id` bigint(255) unsigned NOT NULL COMMENT '用户id',
  `video_id` bigint(255) unsigned NOT NULL COMMENT '视频id',
  `videl_watched_length` bigint(255) unsigned NOT NULL COMMENT '观看时长',
  PRIMARY KEY (`id`,`user_id`,`video_id`) USING BTREE,
  KEY `fk_sys_video_watched_sys_user_1` (`user_id`) USING BTREE,
  KEY `fk_sys_video_watched_sys_course_video_1` (`video_id`) USING BTREE,
  CONSTRAINT `fk_sys_video_watched_sys_course_video_1` FOREIGN KEY (`video_id`) REFERENCES `sys_course_video` (`id`),
  CONSTRAINT `fk_sys_video_watched_sys_user_1` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='播放记录表';

-- ----------------------------
-- Records of sys_video_watched
-- ----------------------------
