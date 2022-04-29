/*
Navicat MySQL Data Transfer

Source Server         : 本地数据库
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : bdms

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2022-04-29 17:36:41
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
) ENGINE=InnoDB AUTO_INCREMENT=94 DEFAULT CHARSET=utf8;

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
INSERT INTO `sys_menu` VALUES ('18', '17', '2', '添加用户', 'el-icon-plus', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('19', '17', '4', '编辑用户', 'el-icon-edit', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('20', '17', '5', '删除用户', 'el-icon-delete', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('21', '17', '6', '分配角色', 'el-icon-setting', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('22', '16', '2', '角色管理', 'el-icon-postcard', '/index/role', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('23', '22', '2', '添加角色', 'el-icon-plus', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('24', '22', '3', '编辑角色', 'el-icon-edit', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('25', '22', '4', '删除角色', 'el-icon-delete', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('26', '22', '5', '分配权限', 'el-icon-setting', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('27', '22', '6', '删除权限', 'el-icon-delete', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('28', '16', '3', '菜单管理', 'el-icon-folder-opened', '/index/menu', '2', null, null, null);
INSERT INTO `sys_menu` VALUES ('29', '28', '2', '添加菜单', 'el-icon-folder-add', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('30', '28', '3', '编辑菜单', 'el-icon-edit', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('31', '28', '4', '删除菜单', 'el-icon-folder-delete', '/null', '3', null, null, null);
INSERT INTO `sys_menu` VALUES ('32', '1', '8', 'Pholcus爬虫', 'el-icon-coffee-cup', '/index/pholcus', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('33', '1', '10', '配置管理', 'el-icon-setting', '/index/config', '1', null, null, null);
INSERT INTO `sys_menu` VALUES ('70', '16', '4', '通知管理', 'el-icon-s-promotion', '/index/message', '2', '', null, null);
INSERT INTO `sys_menu` VALUES ('71', '70', '2', '发布通知', 'el-icon-message', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('72', '70', '3', '修改通知', 'el-icon-edit-outline', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('73', '70', '5', '撤回通知', 'el-icon-refresh-left', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('74', '17', '3', '更新用户状态', 'el-icon-open', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('75', '70', '4', '切换通知置顶状态', 'el-icon-top', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('76', '1', '7', 'Imooc实战', 'el-icon-aim', '/index7', '1', '', null, null);
INSERT INTO `sys_menu` VALUES ('77', '76', '1', '课程分析', 'el-icon-data-board', '/index/imooc_course', '2', '', null, null);
INSERT INTO `sys_menu` VALUES ('78', '76', '2', '课程数据', 'el-icon-coin', '/index/imooc_course_data', '2', '', null, null);
INSERT INTO `sys_menu` VALUES ('79', '1', '11', '润林公司', 'el-icon-office-building', '/rl', '1', '', null, null);
INSERT INTO `sys_menu` VALUES ('80', '79', '1', '项目管理', 'el-icon-s-finance', '/index/rl_finance', '2', '', null, null);
INSERT INTO `sys_menu` VALUES ('81', '80', '1', '新建项目', 'el-icon-folder-add', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('82', '80', '6', '新添账目', 'el-icon-plus', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('83', '80', '7', '修改账目流水', 'el-icon-edit', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('84', '80', '8', '删除账目流水', 'el-icon-delete', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('85', '80', '2', '完成项目', 'el-icon-folder-checked', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('86', '80', '3', '重启项目', 'el-icon-refresh', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('87', '80', '4', '项目管理', 'el-icon-folder-opened', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('88', '80', '5', '修改项目信息', 'el-icon-edit-outline', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('89', '1', '1', '润林首页', 'el-icon-house', '/index/rl_house', '1', '', null, null);
INSERT INTO `sys_menu` VALUES ('90', '17', '1', '查看用户', 'null', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('91', '22', '1', '查看角色', 'null', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('92', '28', '1', '查看菜单', 'null', '/null', '3', '', null, null);
INSERT INTO `sys_menu` VALUES ('93', '70', '1', '查看通知', 'null', '/null', '3', '', null, null);

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
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_message
-- ----------------------------
INSERT INTO `sys_message` VALUES ('18', '1', 'BDMS System 1.0上线', 'el-icon-info', '<h2><span style=\"color: rgb(230, 0, 0);\">系统部分</span></h2><p>包含用户管理，角色管理，权限管理，通知管理，</p><h2><span style=\"color: rgb(194, 133, 255);\">业务部分</span></h2><p>包含用户，课程，学科（未完成），职业，公司</p><h2><span style=\"color: rgb(255, 194, 102);\">爬虫部分（待完善，敬请期待）</span></h2>', '2020-03-26 10:43:24.000000', '2');
INSERT INTO `sys_message` VALUES ('21', '1', 'BDMS System 1.1 更新时间/内容', 'el-icon-info', '<h2><span style=\"color: rgb(230, 0, 0);\">下一次更新估计将在半个月后</span></h2><p><br></p><ol><li><em style=\"color: rgb(0, 102, 204);\">主要涉及爬虫模块的探索 和 与之匹配的可视化页面的实现</em></li><li><em style=\"color: rgb(0, 102, 204);\">其次还有学科模块的实现</em></li></ol>', '2020-03-27 00:03:59.000000', '2');
INSERT INTO `sys_message` VALUES ('29', '1', 'BDMS System 1.1上线', 'el-icon-info', '<h1><span style=\"color: rgb(0, 102, 204);\">BDMS system 1.1</span></h1><h3><span style=\"color: rgb(0, 102, 204);\">完善了科目模块（即学科模块）</span></h3><p><br></p><h3><span style=\"color: rgb(230, 0, 0);\">BDMS系统业务部分</span></h3><h4>包含用户管理，角色管理，权限管理，通知管理（管理员权限）</h4><h3><span style=\"color: rgb(194, 133, 255);\">Worknet数据分析业务</span></h3><h4>包含用户，课程，科目，职业，公司</h4><h3><span style=\"color: rgb(255, 194, 102);\">爬虫部分（待完善，敬请期待）</span></h3>', '2020-04-15 15:29:18.000000', '2');
INSERT INTO `sys_message` VALUES ('30', '1', 'BDMS System 1.2 更新时间/内容', 'el-icon-info', '<h2><strong style=\"color: rgb(230, 0, 0);\">下次更新可能将在五月上旬</strong></h2><ol><li><em style=\"color: rgb(0, 102, 204);\">将会更新慕课网课程数据分析实战</em></li><li><em style=\"color: rgb(0, 102, 204);\">修复一部分bug</em></li></ol>', '2020-04-15 15:47:10.000000', '2');
INSERT INTO `sys_message` VALUES ('34', '1', 'BDMS System 1.2上线了！', 'el-icon-info', '<h1><span style=\"color: rgb(0, 102, 204);\">BDMS system 1.2</span></h1><h3><span style=\"color: rgb(0, 102, 204);\">完善了爬虫部分，</span></h3><h3><span style=\"color: rgb(0, 102, 204);\">上线Imooc实战模块，具有课程分析和课程数据一览功能。</span></h3><p><br></p><h3><span style=\"color: rgb(230, 0, 0);\">BDMS系统业务部分</span></h3><h4>包含用户管理，角色管理，权限管理，通知管理（管理员权限）</h4><h3><span style=\"color: rgb(194, 133, 255);\">Worknet数据分析业务</span></h3><h4>包含用户，课程，科目，职业，公司</h4><h3><span style=\"color: rgb(102, 185, 102);\">Imooc实战</span></h3><h4>包含课程可视化分析、课程数据一览</h4><h3><span style=\"color: rgb(255, 194, 102);\">爬虫部分</span></h3><h4>包含对各类网站的爬虫功能</h4>', '2020-04-26 15:13:28.000000', '2');
INSERT INTO `sys_message` VALUES ('35', '1', '停更通知', 'el-icon-info', '<h2>	<strong style=\"color: rgb(153, 51, 255);\">系统已完成既定功能要求。</strong></h2><p><br></p><h3>			<strong style=\"color: rgb(230, 0, 0);\">无限期停更~</strong></h3>', '2020-04-26 15:14:32.000000', '2');
INSERT INTO `sys_message` VALUES ('36', '1', 'BDMSystem 2.0发布！', 'el-icon-info', '<h1><span style=\"color: rgb(0, 102, 204);\">BDMS system 2.0</span></h1><h3><strong style=\"color: rgb(0, 102, 204);\">System2.0系统为润林公司单独打造了财务流水功能！</strong></h3><h3><strong style=\"color: rgb(0, 102, 204);\">得益于微服务优势，使得系统拓展新业务可以顺利进行！</strong></h3><p><br></p><h3><span style=\"color: rgb(230, 0, 0);\">BDMS系统业务</span></h3><h4>包含用户管理，角色管理，权限管理，通知管理（管理员权限）</h4><h3><span style=\"color: rgb(194, 133, 255);\">Worknet数据分析业务</span></h3><h4>包含用户，课程，科目，职业，公司</h4><h3><span style=\"color: rgb(102, 185, 102);\">Imooc实战</span></h3><h4>包含课程可视化分析、课程数据一览</h4><h3><span style=\"color: rgb(255, 194, 102);\">爬虫业务</span></h3><h4>包含对各类网站的爬虫功能</h4><h3><span style=\"color: rgb(255, 153, 0);\">润林公司业务</span></h3><h4>包含公司项目流水管理</h4><h3><br></h3><p><br></p>', '2021-02-24 00:00:00.000000', '1');
INSERT INTO `sys_message` VALUES ('37', '1', 'BDMS system 2.1.0', 'el-icon-info', '<h3><strong style=\"color: rgb(0, 102, 204);\">BDMS system 2.1.0 更新</strong></h3><p><br></p><p><strong style=\"color: rgb(255, 153, 0);\">更新内容：</strong></p><h4>润林公司业务，项目信息列表的条目修改，添加了累计转款，转款次数，收支总计。</h4><h1><br></h1><p><br></p>', '2021-02-25 15:40:54.000000', '2');
INSERT INTO `sys_message` VALUES ('38', '1', 'BDMS system 2.1.1 更新', 'el-icon-info', '<h3><strong style=\"color: rgb(0, 102, 204);\">BDMS system 2.1.1 更新</strong></h3><p><br></p><p><strong style=\"color: rgb(255, 153, 0);\">更新内容：</strong></p><h4>润林公司业务，优化了项目信息列表中所有的金额条目，金额数量可以格式化显示。</h4><p><br></p>', '2021-02-27 13:38:09.000000', '2');
INSERT INTO `sys_message` VALUES ('39', '1', '润林公司业务Android已上线！', 'el-icon-info', '<h4><strong style=\"color: rgb(0, 102, 204);\">润林公司业务Android端已上线！</strong></h4><p>	<em># 请尽快更新数据</em></p>', '2021-02-27 18:44:50.000000', '2');
INSERT INTO `sys_message` VALUES ('40', '1', 'bug修复', 'el-icon-info', '<p>修复了润林财务管理中的页面跳转问题</p>', '2021-05-07 12:31:28.000000', '2');

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
) ENGINE=InnoDB AUTO_INCREMENT=211 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_message_user
-- ----------------------------
INSERT INTO `sys_message_user` VALUES ('71', '18', '2', '1');
INSERT INTO `sys_message_user` VALUES ('72', '18', '4', '1');
INSERT INTO `sys_message_user` VALUES ('73', '18', '5', '1');
INSERT INTO `sys_message_user` VALUES ('74', '18', '6', '1');
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
INSERT INTO `sys_message_user` VALUES ('146', '34', '1', '2');
INSERT INTO `sys_message_user` VALUES ('147', '34', '2', '1');
INSERT INTO `sys_message_user` VALUES ('148', '34', '4', '1');
INSERT INTO `sys_message_user` VALUES ('149', '34', '5', '1');
INSERT INTO `sys_message_user` VALUES ('150', '34', '6', '1');
INSERT INTO `sys_message_user` VALUES ('151', '34', '13', '1');
INSERT INTO `sys_message_user` VALUES ('152', '34', '20', '2');
INSERT INTO `sys_message_user` VALUES ('153', '35', '1', '2');
INSERT INTO `sys_message_user` VALUES ('154', '35', '2', '1');
INSERT INTO `sys_message_user` VALUES ('155', '35', '4', '1');
INSERT INTO `sys_message_user` VALUES ('156', '35', '5', '1');
INSERT INTO `sys_message_user` VALUES ('157', '35', '6', '1');
INSERT INTO `sys_message_user` VALUES ('158', '35', '13', '1');
INSERT INTO `sys_message_user` VALUES ('159', '35', '20', '1');
INSERT INTO `sys_message_user` VALUES ('160', '36', '1', '2');
INSERT INTO `sys_message_user` VALUES ('161', '36', '2', '1');
INSERT INTO `sys_message_user` VALUES ('162', '36', '5', '1');
INSERT INTO `sys_message_user` VALUES ('163', '36', '13', '1');
INSERT INTO `sys_message_user` VALUES ('164', '36', '21', '1');
INSERT INTO `sys_message_user` VALUES ('165', '36', '22', '1');
INSERT INTO `sys_message_user` VALUES ('166', '36', '23', '1');
INSERT INTO `sys_message_user` VALUES ('167', '36', '24', '1');
INSERT INTO `sys_message_user` VALUES ('168', '36', '25', '2');
INSERT INTO `sys_message_user` VALUES ('169', '37', '1', '2');
INSERT INTO `sys_message_user` VALUES ('170', '37', '2', '1');
INSERT INTO `sys_message_user` VALUES ('171', '37', '5', '1');
INSERT INTO `sys_message_user` VALUES ('172', '37', '13', '1');
INSERT INTO `sys_message_user` VALUES ('173', '37', '21', '1');
INSERT INTO `sys_message_user` VALUES ('174', '37', '22', '1');
INSERT INTO `sys_message_user` VALUES ('175', '37', '23', '1');
INSERT INTO `sys_message_user` VALUES ('176', '37', '24', '1');
INSERT INTO `sys_message_user` VALUES ('177', '37', '25', '2');
INSERT INTO `sys_message_user` VALUES ('178', '37', '26', '1');
INSERT INTO `sys_message_user` VALUES ('179', '38', '1', '2');
INSERT INTO `sys_message_user` VALUES ('180', '38', '2', '1');
INSERT INTO `sys_message_user` VALUES ('181', '38', '5', '1');
INSERT INTO `sys_message_user` VALUES ('182', '38', '13', '1');
INSERT INTO `sys_message_user` VALUES ('183', '38', '21', '1');
INSERT INTO `sys_message_user` VALUES ('184', '38', '22', '1');
INSERT INTO `sys_message_user` VALUES ('185', '38', '23', '1');
INSERT INTO `sys_message_user` VALUES ('186', '38', '24', '1');
INSERT INTO `sys_message_user` VALUES ('187', '38', '25', '2');
INSERT INTO `sys_message_user` VALUES ('188', '38', '26', '1');
INSERT INTO `sys_message_user` VALUES ('189', '39', '1', '2');
INSERT INTO `sys_message_user` VALUES ('190', '39', '2', '1');
INSERT INTO `sys_message_user` VALUES ('191', '39', '5', '1');
INSERT INTO `sys_message_user` VALUES ('192', '39', '13', '1');
INSERT INTO `sys_message_user` VALUES ('193', '39', '21', '1');
INSERT INTO `sys_message_user` VALUES ('194', '39', '22', '1');
INSERT INTO `sys_message_user` VALUES ('195', '39', '23', '1');
INSERT INTO `sys_message_user` VALUES ('196', '39', '24', '1');
INSERT INTO `sys_message_user` VALUES ('197', '39', '25', '2');
INSERT INTO `sys_message_user` VALUES ('198', '39', '26', '1');
INSERT INTO `sys_message_user` VALUES ('199', '39', '27', '1');
INSERT INTO `sys_message_user` VALUES ('200', '40', '1', '2');
INSERT INTO `sys_message_user` VALUES ('201', '40', '2', '1');
INSERT INTO `sys_message_user` VALUES ('202', '40', '5', '1');
INSERT INTO `sys_message_user` VALUES ('203', '40', '13', '1');
INSERT INTO `sys_message_user` VALUES ('204', '40', '21', '1');
INSERT INTO `sys_message_user` VALUES ('205', '40', '22', '1');
INSERT INTO `sys_message_user` VALUES ('206', '40', '23', '1');
INSERT INTO `sys_message_user` VALUES ('207', '40', '24', '1');
INSERT INTO `sys_message_user` VALUES ('208', '40', '25', '1');
INSERT INTO `sys_message_user` VALUES ('209', '40', '26', '1');
INSERT INTO `sys_message_user` VALUES ('210', '40', '27', '1');

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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('1', '1', '超级管理员', '该角色拥有整个系统所有权限', null, null, '0');
INSERT INTO `sys_role` VALUES ('2', '2', '一般管理员', '该角色可以进行一般管理操作', '', null, null);
INSERT INTO `sys_role` VALUES ('3', '4', '运维管理人员', '该角色可以查看浏览监控项目进展情况', '', null, null);
INSERT INTO `sys_role` VALUES ('6', '5', '游客角色', '可以浏览所有功能的用户，但不能进行任何操作', '', null, null);
INSERT INTO `sys_role` VALUES ('7', '3', '项目经理', '该角色可以进行用户管理和检查项目进展情况', '', null, null);
INSERT INTO `sys_role` VALUES ('8', '6', '润林公司项目经理', '该角色可以对润林公司经营项目进行管理', '', null, null);
INSERT INTO `sys_role` VALUES ('9', '5', '润林公司总经理', '该角色可以对润林公司经营项目和公司人员进行管理', '', null, null);
INSERT INTO `sys_role` VALUES ('10', '8', '润林公司会计师', '该角色可以对润林公司经营项目的账目进行编辑处理', '', null, null);
INSERT INTO `sys_role` VALUES ('11', '7', '润林公司出纳', '该角色可以对润林公司经营项目的账目进行核对', '', null, null);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` bigint(100) unsigned DEFAULT NULL COMMENT '角色id',
  `menu_id` bigint(80) unsigned DEFAULT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=571 DEFAULT CHARSET=utf8;

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
INSERT INTO `sys_role_menu` VALUES ('258', '2', '6');
INSERT INTO `sys_role_menu` VALUES ('260', '2', '18');
INSERT INTO `sys_role_menu` VALUES ('261', '2', '5');
INSERT INTO `sys_role_menu` VALUES ('264', '2', '23');
INSERT INTO `sys_role_menu` VALUES ('265', '2', '7');
INSERT INTO `sys_role_menu` VALUES ('266', '2', '16');
INSERT INTO `sys_role_menu` VALUES ('267', '2', '9');
INSERT INTO `sys_role_menu` VALUES ('268', '2', '17');
INSERT INTO `sys_role_menu` VALUES ('272', '2', '8');
INSERT INTO `sys_role_menu` VALUES ('274', '2', '20');
INSERT INTO `sys_role_menu` VALUES ('279', '2', '2');
INSERT INTO `sys_role_menu` VALUES ('280', '2', '4');
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
INSERT INTO `sys_role_menu` VALUES ('409', '1', '2');
INSERT INTO `sys_role_menu` VALUES ('411', '1', '77');
INSERT INTO `sys_role_menu` VALUES ('412', '1', '78');
INSERT INTO `sys_role_menu` VALUES ('414', '6', '15');
INSERT INTO `sys_role_menu` VALUES ('419', '6', '7');
INSERT INTO `sys_role_menu` VALUES ('420', '6', '5');
INSERT INTO `sys_role_menu` VALUES ('421', '6', '32');
INSERT INTO `sys_role_menu` VALUES ('423', '6', '3');
INSERT INTO `sys_role_menu` VALUES ('424', '6', '8');
INSERT INTO `sys_role_menu` VALUES ('427', '6', '77');
INSERT INTO `sys_role_menu` VALUES ('429', '6', '6');
INSERT INTO `sys_role_menu` VALUES ('432', '6', '78');
INSERT INTO `sys_role_menu` VALUES ('436', '6', '33');
INSERT INTO `sys_role_menu` VALUES ('438', '6', '4');
INSERT INTO `sys_role_menu` VALUES ('439', '6', '11');
INSERT INTO `sys_role_menu` VALUES ('440', '6', '13');
INSERT INTO `sys_role_menu` VALUES ('443', '6', '12');
INSERT INTO `sys_role_menu` VALUES ('445', '6', '10');
INSERT INTO `sys_role_menu` VALUES ('446', '6', '9');
INSERT INTO `sys_role_menu` VALUES ('447', '6', '76');
INSERT INTO `sys_role_menu` VALUES ('449', '6', '14');
INSERT INTO `sys_role_menu` VALUES ('454', '2', '77');
INSERT INTO `sys_role_menu` VALUES ('455', '2', '78');
INSERT INTO `sys_role_menu` VALUES ('456', '2', '76');
INSERT INTO `sys_role_menu` VALUES ('457', '2', '30');
INSERT INTO `sys_role_menu` VALUES ('458', '2', '28');
INSERT INTO `sys_role_menu` VALUES ('459', '7', '76');
INSERT INTO `sys_role_menu` VALUES ('460', '7', '78');
INSERT INTO `sys_role_menu` VALUES ('461', '7', '33');
INSERT INTO `sys_role_menu` VALUES ('462', '7', '77');
INSERT INTO `sys_role_menu` VALUES ('464', '7', '17');
INSERT INTO `sys_role_menu` VALUES ('465', '3', '76');
INSERT INTO `sys_role_menu` VALUES ('466', '3', '33');
INSERT INTO `sys_role_menu` VALUES ('467', '3', '78');
INSERT INTO `sys_role_menu` VALUES ('468', '3', '77');
INSERT INTO `sys_role_menu` VALUES ('469', '7', '18');
INSERT INTO `sys_role_menu` VALUES ('470', '7', '19');
INSERT INTO `sys_role_menu` VALUES ('472', '7', '20');
INSERT INTO `sys_role_menu` VALUES ('473', '7', '74');
INSERT INTO `sys_role_menu` VALUES ('477', '1', '76');
INSERT INTO `sys_role_menu` VALUES ('478', '8', '1');
INSERT INTO `sys_role_menu` VALUES ('479', '9', '1');
INSERT INTO `sys_role_menu` VALUES ('481', '9', '79');
INSERT INTO `sys_role_menu` VALUES ('482', '9', '80');
INSERT INTO `sys_role_menu` VALUES ('484', '8', '79');
INSERT INTO `sys_role_menu` VALUES ('485', '8', '80');
INSERT INTO `sys_role_menu` VALUES ('486', '10', '1');
INSERT INTO `sys_role_menu` VALUES ('487', '1', '79');
INSERT INTO `sys_role_menu` VALUES ('488', '1', '80');
INSERT INTO `sys_role_menu` VALUES ('489', '1', '86');
INSERT INTO `sys_role_menu` VALUES ('490', '1', '82');
INSERT INTO `sys_role_menu` VALUES ('491', '1', '81');
INSERT INTO `sys_role_menu` VALUES ('492', '1', '84');
INSERT INTO `sys_role_menu` VALUES ('493', '1', '85');
INSERT INTO `sys_role_menu` VALUES ('494', '1', '83');
INSERT INTO `sys_role_menu` VALUES ('495', '1', '87');
INSERT INTO `sys_role_menu` VALUES ('497', '9', '87');
INSERT INTO `sys_role_menu` VALUES ('501', '9', '85');
INSERT INTO `sys_role_menu` VALUES ('505', '8', '86');
INSERT INTO `sys_role_menu` VALUES ('508', '8', '85');
INSERT INTO `sys_role_menu` VALUES ('509', '8', '87');
INSERT INTO `sys_role_menu` VALUES ('510', '10', '80');
INSERT INTO `sys_role_menu` VALUES ('511', '10', '81');
INSERT INTO `sys_role_menu` VALUES ('512', '10', '87');
INSERT INTO `sys_role_menu` VALUES ('513', '10', '84');
INSERT INTO `sys_role_menu` VALUES ('514', '10', '79');
INSERT INTO `sys_role_menu` VALUES ('515', '10', '86');
INSERT INTO `sys_role_menu` VALUES ('516', '10', '85');
INSERT INTO `sys_role_menu` VALUES ('517', '10', '82');
INSERT INTO `sys_role_menu` VALUES ('518', '10', '83');
INSERT INTO `sys_role_menu` VALUES ('519', '9', '74');
INSERT INTO `sys_role_menu` VALUES ('521', '9', '71');
INSERT INTO `sys_role_menu` VALUES ('522', '9', '18');
INSERT INTO `sys_role_menu` VALUES ('524', '9', '70');
INSERT INTO `sys_role_menu` VALUES ('525', '9', '73');
INSERT INTO `sys_role_menu` VALUES ('526', '9', '16');
INSERT INTO `sys_role_menu` VALUES ('527', '9', '72');
INSERT INTO `sys_role_menu` VALUES ('528', '9', '75');
INSERT INTO `sys_role_menu` VALUES ('529', '9', '17');
INSERT INTO `sys_role_menu` VALUES ('530', '1', '88');
INSERT INTO `sys_role_menu` VALUES ('534', '10', '88');
INSERT INTO `sys_role_menu` VALUES ('535', '9', '89');
INSERT INTO `sys_role_menu` VALUES ('536', '8', '89');
INSERT INTO `sys_role_menu` VALUES ('537', '10', '89');
INSERT INTO `sys_role_menu` VALUES ('538', '1', '89');
INSERT INTO `sys_role_menu` VALUES ('539', '9', '86');
INSERT INTO `sys_role_menu` VALUES ('540', '11', '1');
INSERT INTO `sys_role_menu` VALUES ('541', '11', '86');
INSERT INTO `sys_role_menu` VALUES ('542', '11', '87');
INSERT INTO `sys_role_menu` VALUES ('543', '11', '79');
INSERT INTO `sys_role_menu` VALUES ('544', '11', '80');
INSERT INTO `sys_role_menu` VALUES ('545', '11', '89');
INSERT INTO `sys_role_menu` VALUES ('546', '11', '85');
INSERT INTO `sys_role_menu` VALUES ('547', '8', '88');
INSERT INTO `sys_role_menu` VALUES ('548', '9', '88');
INSERT INTO `sys_role_menu` VALUES ('549', '9', '19');
INSERT INTO `sys_role_menu` VALUES ('550', '9', '21');
INSERT INTO `sys_role_menu` VALUES ('554', '6', '16');
INSERT INTO `sys_role_menu` VALUES ('557', '6', '17');
INSERT INTO `sys_role_menu` VALUES ('560', '6', '90');
INSERT INTO `sys_role_menu` VALUES ('561', '1', '93');
INSERT INTO `sys_role_menu` VALUES ('562', '1', '92');
INSERT INTO `sys_role_menu` VALUES ('563', '1', '91');
INSERT INTO `sys_role_menu` VALUES ('564', '1', '90');
INSERT INTO `sys_role_menu` VALUES ('565', '6', '28');
INSERT INTO `sys_role_menu` VALUES ('566', '6', '70');
INSERT INTO `sys_role_menu` VALUES ('567', '6', '92');
INSERT INTO `sys_role_menu` VALUES ('568', '6', '91');
INSERT INTO `sys_role_menu` VALUES ('569', '6', '22');
INSERT INTO `sys_role_menu` VALUES ('570', '6', '93');

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
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'null', 'root', 'root1', '姜云杰', '1998-07-08 00:00:00.000000', '男', '784203082@qq.com', '18321882894', '1', '0', '1', '2020-03-23 00:00:00.000000', null);
INSERT INTO `sys_user` VALUES ('2', 'null', 'roots', 'root2', '姜云杰', '1998-07-08 00:00:00.000000', '男', '784203082@qq.com', '18321882894', '1', '0', '1', '2020-03-23 00:00:00.000000', null);
INSERT INTO `sys_user` VALUES ('5', 'null', 'roots2', 'roots2', '姜云杰', '1998-07-08 00:00:00.000000', '男', '784203082@qq.ocm', '18321882894', '7', '0', '1', '2020-03-23 20:01:01.046430', null);
INSERT INTO `sys_user` VALUES ('13', 'null', 'liangping', 'liangping', '梁平', '1997-05-16 00:00:00.000000', '女', '6574561657@qq.com', '13888888888', '1', '0', '1', '2020-03-29 01:31:02.431545', null);
INSERT INTO `sys_user` VALUES ('21', 'null', 'lstnb', 'lstnb', '陆施涛', '2020-04-27 00:00:00.000000', '男', 'lstnb@qq.com', '18311111111', '2', '0', '1', '2020-04-27 15:45:47.591716', null);
INSERT INTO `sys_user` VALUES ('22', 'null', 'hhfnb', 'hhfnb', '黄宏峰', '2020-04-27 00:00:00.000000', '男', 'hhfnb@qq.com', '18311111111', '2', '0', '1', '2020-04-27 15:46:23.969626', null);
INSERT INTO `sys_user` VALUES ('23', 'null', '13978185336', '13978185336', '杜文建', '1973-06-01 00:00:00.000000', '男', 'xxxx@xx.com', '13978185336', '9', '0', '1', '2021-02-18 01:57:24.483383', null);
INSERT INTO `sys_user` VALUES ('24', 'null', '13517516737', '13517516737', '毕佳林', '1990-01-01 00:00:00.000000', '男', '798341573@qq.com', '13517516737', '8', '0', '1', '2021-02-18 02:15:17.417501', null);
INSERT INTO `sys_user` VALUES ('25', 'null', 'rlkjzyf', '123456', '张艳芳', '2021-02-01 00:00:00.000000', '女', '827464733@qq.com', '15296383201', '10', '0', '1', '2021-02-22 18:10:49.662105', null);
INSERT INTO `sys_user` VALUES ('26', 'null', 'rlcnhxl', '123456', '黄兴丽', '2021-02-24 00:00:00.000000', '女', 'xxxx@qq.com', '13888888888', '11', '0', '1', '2021-02-24 11:46:02.586289', null);
INSERT INTO `sys_user` VALUES ('27', 'null', '13557582520', '13557582520', '杜鹏飞', '2021-02-01 00:00:00.000000', '男', 'xxxx@xx.com', '13888888888', '8', '0', '1', '2021-02-27 18:29:01.117786', null);
INSERT INTO `sys_user` VALUES ('28', 'null', 'test', 'test', 'test', '2022-04-01 00:00:00.000000', '男', 'xxxx@xx.com)', '13888888888', '6', '0', '1', '2022-04-26 11:13:42.163361', null);
