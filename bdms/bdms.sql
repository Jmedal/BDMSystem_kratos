/*
Navicat MySQL Data Transfer

Source Server         : jmedal
Source Server Version : 50720
Source Host           : 47.107.116.160:3306
Source Database       : bdms

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2020-04-22 23:04:02
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint(80) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `pid` bigint(80) DEFAULT NULL COMMENT '父级id',
  `num` bigint(80) unsigned DEFAULT NULL COMMENT '菜单序号',
  `menu_name` varchar(100) DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(50) DEFAULT '' COMMENT '菜单图标',
  `path` varchar(150) DEFAULT '' COMMENT '菜单路径',
  `levels` bigint(15) unsigned DEFAULT NULL COMMENT '菜单层级：0为顶级菜单，占位层级',
  `tips` varchar(255) DEFAULT '' COMMENT '菜单备注（保留字段）',
  `status` bigint(15) unsigned DEFAULT NULL COMMENT '菜单状态：  1:启用   0:不启用（保留字段）',
  `is_menu` bigint(11) unsigned DEFAULT NULL COMMENT '是否是菜单（保留字段）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES ('1', '0', '0', '顶级菜单', '', '', '0', '顶级菜单仅作为占位，不使用', '0', '1');
INSERT INTO `sys_menu` VALUES ('2', '1', '1', '首页', 'el-icon-house', '/index/house', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('3', '1', '2', '用户', 'el-icon-s-custom', '/index2', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('4', '3', '1', '类型分析', 'el-icon-pie-chart', '/index/user_type', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('5', '1', '3', '课程', 'el-icon-data-board', '/index3', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('6', '5', '1', '活跃度分析', 'el-icon-data-line', '/index/course_activity', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('7', '5', '2', '学习情况分析', 'el-icon-pie-chart', '/index/course_study', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('8', '5', '3', '排行榜', 'el-icon-trophy', '/index/course_list', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('9', '1', '4', '科目', 'el-icon-collection', '/index4', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('10', '9', '1', '结构分析', 'el-icon-connection', '/index/curriculum_heat', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('11', '1', '5', '职业', 'el-icon-s-cooperation', '/index5', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('12', '11', '1', '热门分析', 'el-icon-data-analysis', '/index/profession_heat', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('13', '1', '6', '公司', 'el-icon-office-building', '/index6', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('14', '13', '1', '分布/招聘分析', 'el-icon-map-location', '/index/company_profession', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('15', '13', '2', '笔试分析', 'el-icon-edit-outline', '/index/company_contest', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('16', '1', '9', '管理员权限', 'el-icon-smoking', '/index9', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('17', '16', '1', '用户管理', 'iconfont icon-users', '/index/user', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('18', '17', '1', '添加用户', 'el-icon-plus', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('19', '17', '3', '编辑用户', 'el-icon-edit', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('20', '17', '4', '删除用户', 'el-icon-delete', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('21', '17', '5', '分配角色', 'el-icon-setting', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('22', '16', '2', '角色管理', 'el-icon-postcard', '/index/role', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('23', '22', '1', '添加角色', 'el-icon-plus', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('24', '22', '2', '编辑角色', 'el-icon-edit', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('25', '22', '3', '删除角色', 'el-icon-delete', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('26', '22', '4', '分配权限', 'el-icon-setting', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('27', '22', '5', '删除权限', '', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('28', '16', '3', '菜单管理', 'el-icon-folder-opened', '/index/menu', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('29', '28', '1', '添加菜单', 'el-icon-folder-add', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('30', '28', '2', '编辑菜单', 'el-icon-edit', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('31', '28', '3', '删除菜单', 'el-icon-folder-delete', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('32', '1', '8', 'Pholcus爬虫', 'el-icon-coffee-cup', '/index/pholcus', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('33', '1', '10', '配置管理', 'el-icon-setting', '/index/config', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('55', '1', '11', '测试菜单', 'el-icon-platform-eleme', '/index/application', '1', '', null, null);
INSERT INTO `sys_menu` VALUES ('70', '16', '4', '通知管理', 'el-icon-s-promotion', '/index/message', '2', '', null, null);
INSERT INTO `sys_menu` VALUES ('71', '70', '1', '发布通知', 'el-icon-message', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('72', '70', '3', '修改通知', 'el-icon-edit-outline', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('73', '70', '4', '撤回通知', 'el-icon-refresh-left', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('74', '17', '2', '更新用户状态', 'el-icon-open', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('75', '70', '2', '切换通知置顶状态', 'el-icon-top', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('76', '1', '7', 'Imooc实战', 'el-icon-aim', '/index7', '1', '', null, null);
INSERT INTO `sys_menu` VALUES ('77', '76', '1', '课程分析', 'el-icon-data-board', '/index/imooc_course', '2', '', null, null);
INSERT INTO `sys_menu` VALUES ('78', '76', '2', '课程数据', 'el-icon-coin', '/index/imooc_course_data', '2', '', null, null);

-- ----------------------------
-- Table structure for sys_message
-- ----------------------------
DROP TABLE IF EXISTS `sys_message`;
CREATE TABLE `sys_message` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` bigint(255) unsigned DEFAULT NULL COMMENT '发布者id',
  `title` varchar(50) DEFAULT '' COMMENT '通知标题',
  `icon` varchar(50) DEFAULT '' COMMENT '通知图标',
  `content` varchar(15000) DEFAULT '' COMMENT '通知内容',
  `release_time` datetime(6) DEFAULT NULL COMMENT '发布时间',
  `is_top` int(2) unsigned DEFAULT '2' COMMENT '是否置顶（置顶：1，不置顶：2）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_message
-- ----------------------------
INSERT INTO `sys_message` VALUES ('18', '1', 'BDMS System 1.0上线', 'el-icon-info', '<h2><span style=\"color: rgb(230, 0, 0);\">系统部分</span></h2><p>包含用户管理，角色管理，权限管理，通知管理，</p><h2><span style=\"color: rgb(194, 133, 255);\">业务部分</span></h2><p>包含用户，课程，学科（未完成），职业，公司</p><h2><span style=\"color: rgb(255, 194, 102);\">爬虫部分（待完善，敬请期待）</span></h2>', '2020-03-26 10:43:24.000000', '2');
INSERT INTO `sys_message` VALUES ('19', '1', '测试', 'el-icon-warning', '<h2><span style=\"color: rgb(230, 0, 0);\">系统部分</span></h2><p><br></p><p>包含用户管理，角色管理，权限管理，通知管理，</p><p><br></p><h2><span style=\"color: rgb(194, 133, 255);\">业务部分</span></h2><p><br></p><p>包含用户，课程，学科（未完成），职业，公司，</p><p><br></p><h2><span style=\"color: rgb(255, 194, 102);\">爬虫部分（待完善，敬请期待）</span></h2><p><br></p>', '2020-03-18 00:00:00.000000', '2');
INSERT INTO `sys_message` VALUES ('21', '1', 'BDMS System 1.1 更新时间/内容', 'el-icon-info', '<h2><span style=\"color: rgb(230, 0, 0);\">下一次更新估计将在半个月后</span></h2><p><br></p><ol><li><em style=\"color: rgb(0, 102, 204);\">主要涉及爬虫模块的探索 和 与之匹配的可视化页面的实现</em></li><li><em style=\"color: rgb(0, 102, 204);\">其次还有学科模块的实现</em></li></ol>', '2020-03-27 00:03:59.000000', '2');
INSERT INTO `sys_message` VALUES ('29', '1', 'BDMS System 1.1上线', 'el-icon-info', '<h1><span style=\"color: rgb(0, 102, 204);\">BDMS system 1.1</span></h1><h3><span style=\"color: rgb(0, 102, 204);\">完善了科目模块（即学科模块）</span></h3><p><br></p><h3><span style=\"color: rgb(230, 0, 0);\">BDMS系统业务部分</span></h3><h4>包含用户管理，角色管理，权限管理，通知管理（管理员权限）</h4><h3><span style=\"color: rgb(194, 133, 255);\">Worknet数据分析业务</span></h3><h4>包含用户，课程，科目，职业，公司</h4><h3><span style=\"color: rgb(255, 194, 102);\">爬虫部分（待完善，敬请期待）</span></h3>', '2020-04-15 15:29:18.000000', '1');
INSERT INTO `sys_message` VALUES ('30', '1', 'BDMS System 1.2 更新时间/内容', 'el-icon-info', '<h2><strong style=\"color: rgb(230, 0, 0);\">下次更新可能将在五月上旬</strong></h2><ol><li><em style=\"color: rgb(0, 102, 204);\">将会更新慕课网课程数据分析实战</em></li><li><em style=\"color: rgb(0, 102, 204);\">修复一部分bug</em></li></ol>', '2020-04-15 15:47:10.000000', '2');

-- ----------------------------
-- Table structure for sys_message_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_message_user`;
CREATE TABLE `sys_message_user` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Iid',
  `message_id` bigint(255) unsigned DEFAULT NULL COMMENT '通知id',
  `user_id` bigint(255) unsigned DEFAULT NULL COMMENT '接收者id',
  `is_read` int(5) unsigned DEFAULT '1' COMMENT '阅读（1：未读，2：已读）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=146 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_message_user
-- ----------------------------
INSERT INTO `sys_message_user` VALUES ('71', '18', '2', '1');
INSERT INTO `sys_message_user` VALUES ('72', '18', '4', '1');
INSERT INTO `sys_message_user` VALUES ('73', '18', '5', '1');
INSERT INTO `sys_message_user` VALUES ('74', '18', '6', '1');
INSERT INTO `sys_message_user` VALUES ('76', '19', '2', '1');
INSERT INTO `sys_message_user` VALUES ('77', '19', '4', '1');
INSERT INTO `sys_message_user` VALUES ('78', '19', '5', '1');
INSERT INTO `sys_message_user` VALUES ('79', '19', '6', '1');
INSERT INTO `sys_message_user` VALUES ('85', '21', '1', '2');
INSERT INTO `sys_message_user` VALUES ('86', '21', '2', '1');
INSERT INTO `sys_message_user` VALUES ('87', '21', '4', '1');
INSERT INTO `sys_message_user` VALUES ('88', '21', '5', '1');
INSERT INTO `sys_message_user` VALUES ('89', '21', '6', '1');
INSERT INTO `sys_message_user` VALUES ('131', '29', '1', '2');
INSERT INTO `sys_message_user` VALUES ('132', '29', '2', '1');
INSERT INTO `sys_message_user` VALUES ('133', '29', '4', '1');
INSERT INTO `sys_message_user` VALUES ('134', '29', '5', '1');
INSERT INTO `sys_message_user` VALUES ('135', '29', '6', '1');
INSERT INTO `sys_message_user` VALUES ('136', '29', '13', '2');
INSERT INTO `sys_message_user` VALUES ('137', '30', '1', '2');
INSERT INTO `sys_message_user` VALUES ('138', '30', '2', '1');
INSERT INTO `sys_message_user` VALUES ('139', '30', '4', '1');
INSERT INTO `sys_message_user` VALUES ('140', '30', '5', '1');
INSERT INTO `sys_message_user` VALUES ('141', '30', '6', '1');
INSERT INTO `sys_message_user` VALUES ('142', '30', '13', '2');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint(100) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `num` bigint(100) unsigned DEFAULT NULL COMMENT '序号',
  `role_name` varchar(255) DEFAULT '' COMMENT '角色名称',
  `role_desc` varchar(255) DEFAULT '' COMMENT '角色描述',
  `tips` varchar(255) DEFAULT '' COMMENT '提示（保留字段）',
  `pid` bigint(100) DEFAULT NULL COMMENT '父角色id（保留字段）',
  `version` bigint(50) DEFAULT NULL COMMENT '等級（保留字段）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('1', '1', '超级管理员', '能够拥有整个系统所有权限', null, null, '0');
INSERT INTO `sys_role` VALUES ('2', '2', '一般管理员', '一般一般吧', '', null, null);
INSERT INTO `sys_role` VALUES ('3', '4', '运维管理人员', '本系统业务功能的主要使用者', '', null, null);
INSERT INTO `sys_role` VALUES ('6', '5', 'abcdefg', 'asdasdasdas', '', null, null);
INSERT INTO `sys_role` VALUES ('7', '3', '项目经理', '检查项目进展', '', null, null);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` bigint(100) unsigned DEFAULT NULL COMMENT '角色id',
  `menu_id` bigint(80) unsigned DEFAULT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=413 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES ('2', '1', '3');
INSERT INTO `sys_role_menu` VALUES ('3', '1', '4');
INSERT INTO `sys_role_menu` VALUES ('4', '1', '5');
INSERT INTO `sys_role_menu` VALUES ('5', '1', '6');
INSERT INTO `sys_role_menu` VALUES ('6', '1', '7');
INSERT INTO `sys_role_menu` VALUES ('7', '1', '8');
INSERT INTO `sys_role_menu` VALUES ('8', '1', '9');
INSERT INTO `sys_role_menu` VALUES ('9', '1', '10');
INSERT INTO `sys_role_menu` VALUES ('10', '1', '11');
INSERT INTO `sys_role_menu` VALUES ('11', '1', '12');
INSERT INTO `sys_role_menu` VALUES ('12', '1', '13');
INSERT INTO `sys_role_menu` VALUES ('13', '1', '14');
INSERT INTO `sys_role_menu` VALUES ('14', '1', '15');
INSERT INTO `sys_role_menu` VALUES ('15', '1', '16');
INSERT INTO `sys_role_menu` VALUES ('16', '1', '17');
INSERT INTO `sys_role_menu` VALUES ('17', '1', '18');
INSERT INTO `sys_role_menu` VALUES ('18', '1', '19');
INSERT INTO `sys_role_menu` VALUES ('19', '1', '20');
INSERT INTO `sys_role_menu` VALUES ('20', '1', '21');
INSERT INTO `sys_role_menu` VALUES ('21', '1', '22');
INSERT INTO `sys_role_menu` VALUES ('22', '1', '23');
INSERT INTO `sys_role_menu` VALUES ('23', '1', '24');
INSERT INTO `sys_role_menu` VALUES ('24', '1', '25');
INSERT INTO `sys_role_menu` VALUES ('25', '1', '26');
INSERT INTO `sys_role_menu` VALUES ('26', '1', '27');
INSERT INTO `sys_role_menu` VALUES ('27', '1', '28');
INSERT INTO `sys_role_menu` VALUES ('28', '1', '29');
INSERT INTO `sys_role_menu` VALUES ('29', '1', '30');
INSERT INTO `sys_role_menu` VALUES ('30', '1', '31');
INSERT INTO `sys_role_menu` VALUES ('31', '1', '32');
INSERT INTO `sys_role_menu` VALUES ('32', '1', '33');
INSERT INTO `sys_role_menu` VALUES ('33', '1', '1');
INSERT INTO `sys_role_menu` VALUES ('173', '2', '1');
INSERT INTO `sys_role_menu` VALUES ('174', '3', '1');
INSERT INTO `sys_role_menu` VALUES ('175', '6', '1');
INSERT INTO `sys_role_menu` VALUES ('176', '7', '1');
INSERT INTO `sys_role_menu` VALUES ('216', '3', '55');
INSERT INTO `sys_role_menu` VALUES ('258', '2', '6');
INSERT INTO `sys_role_menu` VALUES ('259', '2', '29');
INSERT INTO `sys_role_menu` VALUES ('260', '2', '18');
INSERT INTO `sys_role_menu` VALUES ('261', '2', '5');
INSERT INTO `sys_role_menu` VALUES ('262', '2', '30');
INSERT INTO `sys_role_menu` VALUES ('263', '2', '26');
INSERT INTO `sys_role_menu` VALUES ('264', '2', '23');
INSERT INTO `sys_role_menu` VALUES ('265', '2', '7');
INSERT INTO `sys_role_menu` VALUES ('266', '2', '16');
INSERT INTO `sys_role_menu` VALUES ('267', '2', '9');
INSERT INTO `sys_role_menu` VALUES ('268', '2', '17');
INSERT INTO `sys_role_menu` VALUES ('271', '2', '27');
INSERT INTO `sys_role_menu` VALUES ('272', '2', '8');
INSERT INTO `sys_role_menu` VALUES ('274', '2', '20');
INSERT INTO `sys_role_menu` VALUES ('277', '2', '28');
INSERT INTO `sys_role_menu` VALUES ('278', '2', '21');
INSERT INTO `sys_role_menu` VALUES ('279', '2', '2');
INSERT INTO `sys_role_menu` VALUES ('280', '2', '4');
INSERT INTO `sys_role_menu` VALUES ('281', '2', '31');
INSERT INTO `sys_role_menu` VALUES ('282', '2', '22');
INSERT INTO `sys_role_menu` VALUES ('283', '2', '3');
INSERT INTO `sys_role_menu` VALUES ('285', '2', '10');
INSERT INTO `sys_role_menu` VALUES ('286', '2', '25');
INSERT INTO `sys_role_menu` VALUES ('287', '2', '24');
INSERT INTO `sys_role_menu` VALUES ('288', '2', '19');
INSERT INTO `sys_role_menu` VALUES ('290', '7', '15');
INSERT INTO `sys_role_menu` VALUES ('291', '7', '6');
INSERT INTO `sys_role_menu` VALUES ('292', '7', '13');
INSERT INTO `sys_role_menu` VALUES ('293', '7', '7');
INSERT INTO `sys_role_menu` VALUES ('294', '7', '11');
INSERT INTO `sys_role_menu` VALUES ('295', '7', '3');
INSERT INTO `sys_role_menu` VALUES ('296', '7', '5');
INSERT INTO `sys_role_menu` VALUES ('297', '7', '12');
INSERT INTO `sys_role_menu` VALUES ('298', '7', '2');
INSERT INTO `sys_role_menu` VALUES ('299', '7', '9');
INSERT INTO `sys_role_menu` VALUES ('300', '7', '10');
INSERT INTO `sys_role_menu` VALUES ('301', '7', '4');
INSERT INTO `sys_role_menu` VALUES ('302', '7', '14');
INSERT INTO `sys_role_menu` VALUES ('303', '7', '8');
INSERT INTO `sys_role_menu` VALUES ('304', '3', '6');
INSERT INTO `sys_role_menu` VALUES ('305', '3', '7');
INSERT INTO `sys_role_menu` VALUES ('306', '3', '14');
INSERT INTO `sys_role_menu` VALUES ('307', '3', '8');
INSERT INTO `sys_role_menu` VALUES ('308', '3', '5');
INSERT INTO `sys_role_menu` VALUES ('309', '3', '4');
INSERT INTO `sys_role_menu` VALUES ('310', '3', '2');
INSERT INTO `sys_role_menu` VALUES ('311', '3', '13');
INSERT INTO `sys_role_menu` VALUES ('312', '3', '12');
INSERT INTO `sys_role_menu` VALUES ('313', '3', '3');
INSERT INTO `sys_role_menu` VALUES ('314', '3', '11');
INSERT INTO `sys_role_menu` VALUES ('315', '3', '15');
INSERT INTO `sys_role_menu` VALUES ('316', '3', '10');
INSERT INTO `sys_role_menu` VALUES ('317', '3', '9');
INSERT INTO `sys_role_menu` VALUES ('323', '2', '33');
INSERT INTO `sys_role_menu` VALUES ('324', '2', '14');
INSERT INTO `sys_role_menu` VALUES ('325', '2', '13');
INSERT INTO `sys_role_menu` VALUES ('326', '2', '11');
INSERT INTO `sys_role_menu` VALUES ('327', '2', '12');
INSERT INTO `sys_role_menu` VALUES ('328', '2', '15');
INSERT INTO `sys_role_menu` VALUES ('329', '2', '32');
INSERT INTO `sys_role_menu` VALUES ('334', '1', '74');
INSERT INTO `sys_role_menu` VALUES ('335', '2', '73');
INSERT INTO `sys_role_menu` VALUES ('336', '2', '71');
INSERT INTO `sys_role_menu` VALUES ('337', '2', '74');
INSERT INTO `sys_role_menu` VALUES ('338', '2', '70');
INSERT INTO `sys_role_menu` VALUES ('339', '2', '72');
INSERT INTO `sys_role_menu` VALUES ('340', '7', '70');
INSERT INTO `sys_role_menu` VALUES ('341', '7', '18');
INSERT INTO `sys_role_menu` VALUES ('345', '7', '17');
INSERT INTO `sys_role_menu` VALUES ('348', '7', '74');
INSERT INTO `sys_role_menu` VALUES ('350', '7', '16');
INSERT INTO `sys_role_menu` VALUES ('351', '1', '73');
INSERT INTO `sys_role_menu` VALUES ('352', '1', '72');
INSERT INTO `sys_role_menu` VALUES ('353', '1', '70');
INSERT INTO `sys_role_menu` VALUES ('354', '1', '71');
INSERT INTO `sys_role_menu` VALUES ('355', '1', '75');
INSERT INTO `sys_role_menu` VALUES ('356', '2', '75');
INSERT INTO `sys_role_menu` VALUES ('395', '6', '2');
INSERT INTO `sys_role_menu` VALUES ('402', '7', '75');
INSERT INTO `sys_role_menu` VALUES ('403', '7', '73');
INSERT INTO `sys_role_menu` VALUES ('405', '7', '71');
INSERT INTO `sys_role_menu` VALUES ('406', '7', '72');
INSERT INTO `sys_role_menu` VALUES ('408', '7', '22');
INSERT INTO `sys_role_menu` VALUES ('409', '1', '2');
INSERT INTO `sys_role_menu` VALUES ('410', '1', '76');
INSERT INTO `sys_role_menu` VALUES ('411', '1', '77');
INSERT INTO `sys_role_menu` VALUES ('412', '1', '78');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `avatar` varchar(255) DEFAULT 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png' COMMENT '头像',
  `account` varchar(45) NOT NULL DEFAULT '' COMMENT '帐号',
  `password` varchar(45) NOT NULL DEFAULT '' COMMENT '密码',
  `name` varchar(45) DEFAULT '' COMMENT '姓名',
  `birthday` datetime(6) DEFAULT '2006-01-02 00:00:00.000000' COMMENT '生日',
  `sex` varchar(5) DEFAULT '' COMMENT '性别',
  `email` varchar(45) DEFAULT '' COMMENT '电子邮件',
  `phone` varchar(45) DEFAULT '' COMMENT '电话',
  `role_id` bigint(255) unsigned DEFAULT '0' COMMENT '角色id',
  `dept_id` bigint(11) unsigned DEFAULT '0' COMMENT '部门id',
  `status` int(11) DEFAULT '1' COMMENT '状态(1：启用  2：冻结 ）',
  `create_time` datetime(6) DEFAULT '2006-01-02 00:00:00.000000' COMMENT '创建时间',
  `version` int(11) DEFAULT NULL COMMENT '保留字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'null', 'root', 'root', '姜云杰', '1998-07-08 00:00:00.000000', '男', '784203082@qq.com', '18321882894', '1', '0', '1', '2020-02-18 23:52:21.000000', null);
INSERT INTO `sys_user` VALUES ('2', 'null', 'roots', 'roots', '姜云杰', '1998-07-08 00:00:00.000000', '男', '784203082@qq.com', '18321882894', '1', '0', '1', '2020-03-23 00:00:00.000000', null);
INSERT INTO `sys_user` VALUES ('4', 'null', 'sdadad', 'asdsadsad', 'asdasdasd', '2020-03-10 00:00:00.000000', '女', 'asdda@sd.com', '13821882894', '7', '0', '2', '2020-03-23 19:47:25.099788', null);
INSERT INTO `sys_user` VALUES ('5', 'null', 'roots2', 'roots2', '姜云杰', '1998-07-08 00:00:00.000000', '男', '784203082@qq.ocm', '18321882894', '7', '0', '1', '2020-03-23 20:01:01.046430', null);
INSERT INTO `sys_user` VALUES ('6', 'null', 'asdasdasd', 'sadadas', 'sadsa', '2020-03-02 00:00:00.000000', '女', 'asdasda@asasd.com', '18321882894', '6', '0', '2', '2020-03-23 20:17:27.970052', null);
INSERT INTO `sys_user` VALUES ('13', 'null', 'liangping', 'liangping', '梁平', '1997-05-16 00:00:00.000000', '女', '6574561657@qq.com', '13888888888', '1', '0', '1', '2020-03-29 01:31:02.431545', null);
