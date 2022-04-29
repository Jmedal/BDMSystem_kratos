/*
Navicat MySQL Data Transfer

Source Server         : 本地数据库
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : rl

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2022-04-29 17:36:57
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rl_finance_bill
-- ----------------------------
DROP TABLE IF EXISTS `rl_finance_bill`;
CREATE TABLE `rl_finance_bill` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `project_id` bigint(255) unsigned NOT NULL DEFAULT '1' COMMENT '所属项目id',
  `time` datetime(6) NOT NULL DEFAULT '2006-01-02 00:00:00.000000' COMMENT '日期',
  `amount` double(255,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '金额',
  `direction` varchar(80) NOT NULL DEFAULT '' COMMENT '转款去向',
  `type` varchar(2) NOT NULL DEFAULT '1' COMMENT '类型，''1''为进账，''2''为转款',
  `state` varchar(2) DEFAULT '1' COMMENT '账目状态，''1''为激活，''2''为冻结',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=202 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rl_finance_bill
-- ----------------------------
INSERT INTO `rl_finance_bill` VALUES ('14', '14', '2019-08-30 00:00:00.000000', '39393.00', '预付款30%', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('15', '14', '2019-09-03 00:00:00.000000', '32900.00', '上林县恒隆建材有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('16', '14', '2019-09-24 00:00:00.000000', '39393.00', '进度款%30', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('17', '14', '2019-09-12 00:00:00.000000', '3939.30', '黄新丽民工专户工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('18', '14', '2019-11-25 00:00:00.000000', '26263.00', '进度款20%', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('19', '14', '2019-12-11 00:00:00.000000', '13131.20', '进度款10%', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('20', '14', '2019-09-26 00:00:00.000000', '3940.00', '黄新丽民工专户工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('21', '14', '2019-09-26 00:00:00.000000', '36900.00', '上林生隆建材有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('22', '14', '2019-12-10 00:00:00.000000', '2626.30', '姜景明民工专户工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('23', '14', '2019-12-12 00:00:00.000000', '19600.00', '南宁市吉惠建筑工程劳务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('24', '14', '2019-12-23 00:00:00.000000', '11200.00', '上林县生隆建材有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('25', '14', '2019-12-11 00:00:00.000000', '2363.60', '支付管理费2%', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('26', '15', '2020-11-18 00:00:00.000000', '1707765.37', '第一次请款36.5%入账', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('27', '15', '2020-11-19 00:00:00.000000', '512330.00', '广西丰俊建筑工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('28', '15', '2020-11-19 00:00:00.000000', '120400.00', '南宁广源生建材贸易有限责任公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('29', '15', '2020-11-19 00:00:00.000000', '250000.00', '广西御华混凝土有限责任公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('30', '15', '2020-11-26 00:00:00.000000', '236933.64', '广西钢林投资有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('31', '15', '2020-11-26 00:00:00.000000', '314790.00', '柳州市鹏华木业有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('32', '15', '2020-12-11 00:00:00.000000', '157720.00', '杜鹏飞—个人名义开票', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('33', '15', '2020-12-11 00:00:00.000000', '84605.49', '融达公司—退增值税费', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('34', '15', '2020-11-19 00:00:00.000000', '109673.01', '融达公司扣除—增值税7%', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('35', '15', '2020-11-19 00:00:00.000000', '13160.76', '融达公司扣除—附加税12%', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('36', '15', '2020-11-19 00:00:00.000000', '34155.31', '融达公司扣除—企业所得税2.5%', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('37', '15', '2020-11-19 00:00:00.000000', '512.33', '融达公司扣除—合同印花税', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('38', '15', '2020-11-19 00:00:00.000000', '42695.81', '融达公司扣除—管理费2.5%', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('39', '16', '2020-11-18 00:00:00.000000', '179935.00', '第一笔请款12.8%入账', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('40', '16', '2020-11-19 00:00:00.000000', '11555.46', '融达公司扣除—增值税7%', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('41', '16', '2020-11-19 00:00:00.000000', '1386.66', '融达公司扣除—附加税12%', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('42', '16', '2020-11-19 00:00:00.000000', '3598.70', '融达公司扣除—企业所得税2.5%', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('43', '16', '2020-11-19 00:00:00.000000', '53.98', '融达公司扣除—印花税', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('44', '16', '2020-11-19 00:00:00.000000', '4498.39', '融达公司扣除—管理费2.5%', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('45', '16', '2020-11-26 00:00:00.000000', '77490.00', '柳州市鹏华木业有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('46', '16', '2020-11-26 00:00:00.000000', '77600.00', '广西元庆建材有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('47', '16', '2020-12-11 00:00:00.000000', '8914.78', '融达公司退—增值税', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('48', '16', '2020-12-11 00:00:00.000000', '12266.61', '杜鹏飞以个人名义开票', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('49', '21', '2021-05-07 00:00:00.000000', '49039.00', '广西芳耘商贸有限公司', '2', '2');
INSERT INTO `rl_finance_bill` VALUES ('50', '21', '2021-04-26 00:00:00.000000', '54487.80', '上林县澄泰乡人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('51', '21', '2021-05-07 00:00:00.000000', '49039.00', '广西芳耘商贸有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('52', '20', '2021-04-26 00:00:00.000000', '179003.10', '上林县澄泰乡人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('53', '20', '2021-05-07 00:00:00.000000', '161102.70', '广西芳耘商贸有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('54', '18', '2021-04-26 00:00:00.000000', '146701.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('55', '18', '2021-04-28 00:00:00.000000', '95350.00', '上林县佳林建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('56', '18', '2021-04-28 00:00:00.000000', '14670.10', '3月份工资专户', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('57', '18', '2021-04-30 00:00:00.000000', '29360.00', '上林县佳林建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('58', '23', '2021-05-10 00:00:00.000000', '183671.50', '上林县西燕镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('59', '23', '2021-05-11 00:00:00.000000', '165300.00', '上林县佳林建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('60', '23', '2021-05-11 00:00:00.000000', '18367.50', '民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('61', '17', '2021-04-26 00:00:00.000000', '32218.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('62', '17', '2021-04-28 00:00:00.000000', '20920.00', '上林县佳林建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('63', '17', '2021-04-28 00:00:00.000000', '3221.80', '王运喜', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('64', '17', '2021-04-30 00:00:00.000000', '6460.00', '上林县佳林建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('65', '19', '2021-04-26 00:00:00.000000', '120262.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('66', '19', '2021-04-28 00:00:00.000000', '78160.00', '上林县佳林建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('67', '19', '2021-04-28 00:00:00.000000', '12026.20', '3月份民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('68', '19', '2021-04-30 00:00:00.000000', '24070.00', '上林县佳林建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('69', '24', '2021-06-09 00:00:00.000000', '353800.38', '上林县镇圩瑶族乡人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('70', '24', '2021-06-21 00:00:00.000000', '35380.04', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('71', '22', '2021-06-10 00:00:00.000000', '28100.00', '上林县森硕建材有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('72', '22', '2021-06-10 00:00:00.000000', '10200.00', '南宁市北部湾广厦建材有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('73', '22', '2021-06-10 00:00:00.000000', '4370.70', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('74', '22', '2021-05-10 00:00:00.000000', '43706.61', '上林县塘红乡人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('75', '24', '2021-06-17 00:00:00.000000', '51800.00', '南宁广源生建材贸易有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('76', '24', '2021-06-18 00:00:00.000000', '150000.00', '广西千莱商贸有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('77', '26', '2021-05-10 00:00:00.000000', '8976.00', '上林县塘红乡人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('78', '25', '2021-02-04 00:00:00.000000', '3064.35', '上林县税务局 12%的税费', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('79', '25', '2021-02-01 00:00:00.000000', '274.50', '上林税务局 企业所得税', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('80', '26', '2021-02-01 00:00:00.000000', '2451.46', '上林县税务局 12%税费', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('81', '26', '2021-02-01 00:00:00.000000', '219.60', '上林税务局 企业所得税', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('82', '26', '2021-06-09 00:00:00.000000', '336.31', '上林税务局 12%税费', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('83', '26', '2021-06-09 00:00:00.000000', '30.13', '上林税务局 企业所得税', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('84', '27', '2021-04-13 00:00:00.000000', '2332.02', '上林税务局 12%税费', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('85', '27', '2021-04-13 00:00:00.000000', '208.89', '上林税务局 企业所得税', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('86', '24', '2021-06-21 00:00:00.000000', '100000.00', '广西林发建筑工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('87', '17', '2021-07-02 00:00:00.000000', '53697.50', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('88', '17', '2021-07-05 00:00:00.000000', '48327.75', '上林佳林建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('89', '17', '2021-07-05 00:00:00.000000', '5369.75', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('90', '19', '2021-07-02 00:00:00.000000', '200438.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('91', '19', '2021-07-05 00:00:00.000000', '180394.20', '上林县佳林建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('92', '19', '2021-07-05 00:00:00.000000', '20043.80', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('93', '18', '2021-07-19 00:00:00.000000', '244503.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('94', '18', '2021-07-20 00:00:00.000000', '24450.30', '毕佳林等你农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('95', '18', '2021-07-21 00:00:00.000000', '195170.15', '佳林建材材料款', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('96', '21', '2021-08-11 00:00:00.000000', '90813.00', '上林县澄泰乡人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('97', '21', '2021-08-12 00:00:00.000000', '9081.30', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('98', '21', '2021-08-13 00:00:00.000000', '77082.88', '广西妍臻商贸有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('99', '30', '2021-08-13 00:00:00.000000', '456389.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('100', '30', '2021-08-16 00:00:00.000000', '120000.00', '北部湾广厦建材', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('101', '30', '2021-08-20 00:00:00.000000', '203392.95', '广西中成建筑工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('102', '23', '2021-08-19 00:00:00.000000', '99182.60', '上林延昌建材经营部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('103', '23', '2021-08-19 00:00:00.000000', '11020.30', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('104', '31', '2021-08-19 00:00:00.000000', '88422.10', '上林县巷贤人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('105', '32', '2021-08-19 00:00:00.000000', '275965.10', '上林县巷贤人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('106', '23', '2021-08-19 00:00:00.000000', '110202.90', '上林县西燕镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('107', '30', '2021-08-20 00:00:00.000000', '86714.00', '广西联友劳务服务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('108', '30', '2021-08-20 00:00:00.000000', '22819.00', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('109', '32', '2021-08-26 00:00:00.000000', '239214.34', '上林县石山水泥砖厂', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('110', '32', '2021-08-26 00:00:00.000000', '27596.51', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('111', '31', '2021-08-26 00:00:00.000000', '76646.80', '上林县石山水泥砖厂', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('112', '31', '2021-08-26 00:00:00.000000', '8842.21', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('113', '24', '2021-08-30 00:00:00.000000', '353800.38', '上林县镇圩瑶族乡人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('114', '24', '2021-08-30 00:00:00.000000', '35380.04', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('115', '24', '2021-08-30 00:00:00.000000', '279720.00', '上林县延昌建材营销部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('116', '30', '2021-08-30 00:00:00.000000', '230042.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('117', '30', '2021-09-01 00:00:00.000000', '63055.10', '广西中成建筑工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('118', '30', '2021-09-01 00:00:00.000000', '100000.00', '广西御华混凝土有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('119', '30', '2021-09-01 00:00:00.000000', '43707.80', '广西联友劳务服务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('120', '20', '2021-09-17 00:00:00.000000', '179003.10', '上林县澄泰人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('121', '20', '2021-09-17 00:00:00.000000', '32603.94', '广西荣粤商贸有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('122', '20', '2021-09-17 00:00:00.000000', '119335.40', '广西正前建筑劳务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('123', '20', '2021-09-17 00:00:00.000000', '17900.40', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('124', '32', '2021-09-29 00:00:00.000000', '275965.10', '上林县巷贤镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('125', '32', '2021-09-01 00:00:00.000000', '133798.50', '覃永华石灰厂', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('126', '32', '2021-09-09 00:00:00.000000', '124980.00', '广厦建材', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('127', '32', '2021-10-09 00:00:00.000000', '27596.51', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('128', '31', '2021-09-29 00:00:00.000000', '88422.10', '上林县巷贤镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('129', '31', '2021-09-29 00:00:00.000000', '73461.87', '宝旭建筑', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('130', '31', '2021-10-09 00:00:00.000000', '8842.21', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('131', '35', '2021-09-30 00:00:00.000000', '214497.11', '上林县乔贤镇人民政府', '1', '2');
INSERT INTO `rl_finance_bill` VALUES ('132', '35', '2021-09-30 00:00:00.000000', '214497.11', '上林县乔贤人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('133', '35', '2021-10-11 00:00:00.000000', '57000.00', '延昌建材', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('134', '35', '2021-10-11 00:00:00.000000', '21449.71', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('135', '36', '2021-10-18 00:00:00.000000', '346108.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('136', '36', '2021-10-19 00:00:00.000000', '207252.95', '广西中成建筑工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('137', '36', '2021-10-19 00:00:00.000000', '103832.40', '广西联友劳务服务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('138', '36', '2021-10-19 00:00:00.000000', '17305.00', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('139', '20', '2021-11-18 00:00:00.000000', '179002.80', '上林县澄泰人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('140', '20', '2021-11-17 00:00:00.000000', '92183.02', '广西芳耘商贸有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('141', '20', '2021-11-17 00:00:00.000000', '35800.56', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('142', '20', '2021-11-18 00:00:00.000000', '41767.54', '广西正前建筑劳务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('143', '37', '2021-12-24 00:00:00.000000', '388816.00', '上林县西燕镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('144', '37', '2021-12-31 00:00:00.000000', '194408.10', '广西宝旭建筑工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('145', '37', '2021-12-31 00:00:00.000000', '38881.60', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('146', '31', '2021-12-22 00:00:00.000000', '88421.80', '上林县巷贤镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('147', '31', '2021-12-31 00:00:00.000000', '14960.23', '广西宝旭建筑工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('148', '31', '2021-12-31 00:00:00.000000', '59992.97', '上林县佳林建材经营部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('149', '31', '2021-12-31 00:00:00.000000', '8842.18', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('150', '22', '2021-12-24 00:00:00.000000', '87413.39', '上林县塘红乡人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('151', '22', '2021-12-29 00:00:00.000000', '43706.62', '广西宝旭建筑工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('152', '22', '2021-12-28 00:00:00.000000', '30447.37', '上林县森硕建材有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('153', '22', '2021-12-27 00:00:00.000000', '8741.34', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('154', '30', '2021-12-01 00:00:00.000000', '228809.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('155', '30', '2022-01-14 00:00:00.000000', '205542.75', '广西联友劳务服务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('156', '30', '2021-12-14 00:00:00.000000', '11440.00', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('157', '32', '2022-01-07 00:00:00.000000', '275965.10', '上林县巷贤镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('158', '32', '2021-12-10 00:00:00.000000', '239624.32', '广西宝旭建筑材料工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('159', '32', '2021-12-09 00:00:00.000000', '27596.51', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('160', '23', '2021-11-26 00:00:00.000000', '36733.60', '上林县西燕镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('161', '23', '2021-12-01 00:00:00.000000', '33000.00', '上林县延昌建材有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('162', '23', '2021-12-02 00:00:00.000000', '3673.36', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('163', '38', '2021-11-18 00:00:00.000000', '38938.86', '上林县城南小学', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('164', '38', '2021-11-18 00:00:00.000000', '36995.77', '上林县佳林建材经营部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('165', '34', '2021-11-05 00:00:00.000000', '126646.00', '上林县大丰镇中心学校', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('166', '34', '2021-11-10 00:00:00.000000', '120100.16', '上林县延昌建材经营部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('167', '19', '2021-12-22 00:00:00.000000', '40088.00', '上林县白圩镇人民政府', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('168', '19', '2021-12-28 00:00:00.000000', '34079.57', '上林县海纳劳务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('169', '19', '2021-12-27 00:00:00.000000', '4008.00', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('170', '17', '2021-12-22 00:00:00.000000', '10740.50', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('171', '17', '2021-12-28 00:00:00.000000', '9130.54', '上林县海纳劳务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('172', '17', '2021-12-27 00:00:00.000000', '1074.00', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('173', '18', '2021-12-22 00:00:00.000000', '48901.00', '上林县白圩镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('174', '18', '2021-12-28 00:00:00.000000', '41570.79', '上林县海纳劳务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('175', '18', '2021-12-27 00:00:00.000000', '4890.00', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('176', '35', '2021-12-08 00:00:00.000000', '357495.18', '上林县乔贤镇人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('177', '35', '2022-01-10 00:00:00.000000', '198000.00', '上林县延昌建材经营部', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('178', '35', '2022-01-10 00:00:00.000000', '35749.52', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('179', '24', '2021-06-01 00:00:00.000000', '353800.38', '上林县镇圩瑶族乡人民政府', '1', '2');
INSERT INTO `rl_finance_bill` VALUES ('180', '24', '2021-12-20 00:00:00.000000', '235866.93', '广西宝旭建筑服务有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('181', '24', '2021-12-16 00:00:00.000000', '318420.36', '上林县瑶族乡人民政府', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('182', '24', '2021-12-20 00:00:00.000000', '35380.04', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('183', '39', '2022-01-20 00:00:00.000000', '340000.00', '上林县水利局', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('184', '39', '2022-01-27 00:00:00.000000', '102000.00', '广西宝旭建筑服务工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('185', '39', '2022-01-28 00:00:00.000000', '210800.00', '广西千莱商贸有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('186', '36', '2022-03-03 00:00:00.000000', '128299.00', '上林交通局', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('187', '36', '2022-03-04 00:00:00.000000', '69221.60', '广西中成建筑工程有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('188', '36', '2022-03-04 00:00:00.000000', '6455.00', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('189', '36', '2022-03-17 00:00:00.000000', '74556.00', '上林民盛建材有限公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('190', '36', '2022-03-16 00:00:00.000000', '28922.60', '陈总走账', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('191', '15', '2022-02-23 00:00:00.000000', '503294.00', '上林县土地管理中心（第二次））', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('192', '15', '2022-02-25 00:00:00.000000', '218421.00', '兴美', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('193', '15', '2022-02-25 00:00:00.000000', '80000.00', '新辉', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('194', '15', '2022-02-25 00:00:00.000000', '150988.20', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('195', '16', '2022-02-24 00:00:00.000000', '672537.00', '上林县土地管理局', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('196', '16', '2022-02-25 00:00:00.000000', '291879.00', '兴美', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('197', '16', '2022-02-25 00:00:00.000000', '110000.00', '新辉', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('198', '16', '2022-02-25 00:00:00.000000', '201761.10', '农民工工资', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('199', '36', '2022-03-31 00:00:00.000000', '79365.00', '上林县交通运输局', '1', '1');
INSERT INTO `rl_finance_bill` VALUES ('200', '36', '2022-04-08 00:00:00.000000', '65518.00', '广西御华贸易有限责任公司', '2', '1');
INSERT INTO `rl_finance_bill` VALUES ('201', '36', '2022-04-08 00:00:00.000000', '9523.00', '农民工工资', '2', '1');

-- ----------------------------
-- Table structure for rl_finance_project
-- ----------------------------
DROP TABLE IF EXISTS `rl_finance_project`;
CREATE TABLE `rl_finance_project` (
  `id` bigint(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(80) NOT NULL DEFAULT '' COMMENT '项目名称',
  `development` varchar(80) NOT NULL DEFAULT '' COMMENT '建设单位',
  `construction` varchar(80) NOT NULL DEFAULT '' COMMENT '施工单位',
  `schedule` varchar(50) DEFAULT '' COMMENT '工期',
  `contract` double(255,2) DEFAULT '0.00' COMMENT '合同价格',
  `signTime` datetime(6) DEFAULT '2006-01-02 00:00:00.000000' COMMENT '签订日期',
  `prepayment` double(255,2) DEFAULT '0.00' COMMENT '预付款',
  `progress` varchar(100) DEFAULT '' COMMENT '进度款',
  `judge` double(255,2) DEFAULT '0.00' COMMENT '审定价格',
  `remark` varchar(100) DEFAULT '' COMMENT '备注',
  `state` varchar(2) DEFAULT '1' COMMENT '项目状态，''1''为激活，''2''为冻结',
  `create_user_id` bigint(255) DEFAULT '1' COMMENT '创建者id',
  `create_time` datetime(6) DEFAULT '2006-01-02 00:00:00.000000' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rl_finance_project
-- ----------------------------
INSERT INTO `rl_finance_project` VALUES ('14', '上林县木山乡琴水村吊水屯水泥路破旧修缮工程', '上林县木山乡人民政府', '广西恒隆建筑有限公司', '', '131312.00', '2021-02-01 00:00:00.000000', '39393.00', '', '1.00', '', '1', '25', '2021-02-23 15:50:08.000000');
INSERT INTO `rl_finance_project` VALUES ('15', '南宁市上林县白圩镇、西燕镇云桃村耕地提质改造（旱改水）项目工程', '上林县土地整理整理中心', '广西融达建设工程有限公司', '', '4671626.47', '2020-10-19 00:00:00.000000', '1.00', '1707765.37', '1.00', '外管证税跨报【2020】5683号', '1', '25', '2021-02-28 16:06:27.000000');
INSERT INTO `rl_finance_project` VALUES ('16', '南宁市上林县巷贤镇等3个乡镇土地开垦项目', '上林县土地整理中心', '广西融达建设工程有限公司', '', '1403117.05', '2020-10-19 00:00:00.000000', '1.00', '179935', '1.00', '外管证税跨报【2020】5684号', '1', '25', '2021-02-28 16:43:37.000000');
INSERT INTO `rl_finance_project` VALUES ('17', '上林县白圩镇大浪村塘岭庄至古楼庄硬化路路面修缮工程', '上林县白圩镇人民政府', '广西兴盈建设工程有限公司', '', '107395.00', '2021-04-12 00:00:00.000000', '32218.00', '80%', '10.00', '毕佳林施工', '1', '25', '2021-04-28 16:14:50.000000');
INSERT INTO `rl_finance_project` VALUES ('18', '上林县白圩镇覃黄村二级公路至清水凉庄道路修缮施工程', '上林县白圩镇人民政府', '广西兴盈建设工程有限公司', '', '489006.00', '2021-04-12 00:00:00.000000', '146701.00', '', '10.00', '毕佳林施工', '1', '25', '2021-04-28 16:17:51.000000');
INSERT INTO `rl_finance_project` VALUES ('19', '上林县白圩镇覃黄村二级公路庄至长塘庄道路修缮工程', '上林县白圩镇人民政府', '广西兴盈建设工程有限公司', '', '400876.00', '2021-04-12 00:00:00.000000', '120262.00', '', '10.00', '毕佳林施工', '1', '25', '2021-04-28 16:19:04.000000');
INSERT INTO `rl_finance_project` VALUES ('20', '上林县澄泰乡高顶村科足庄至石灰窑产业硬化路工程', '上林县澄泰乡人民政府', '广西兴盈建设工程有限公司', '', '596677.00', '2021-04-07 00:00:00.000000', '179003.10', '', '10.00', '何世荣施工', '1', '25', '2021-04-28 16:46:51.000000');
INSERT INTO `rl_finance_project` VALUES ('21', '上林县澄泰乡大坡村弘康农牧火龙果基地产业砂石路工程', '上林县澄泰乡人民政府', '广西兴盈建设工程有限公司', '', '181626.00', '2021-04-07 00:00:00.000000', '54487.80', '', '10.00', '何世荣施工', '1', '25', '2021-04-28 16:51:47.000000');
INSERT INTO `rl_finance_project` VALUES ('22', '上林县塘红乡万福村北俭庄至弄碗道路扩建工程', '上林县塘红乡人民政府', '广西兴盈建设工程有限公司', '', '145688.72', '2021-04-19 00:00:00.000000', '43706.61', '', '10.00', '', '1', '25', '2021-04-28 17:52:11.000000');
INSERT INTO `rl_finance_project` VALUES ('23', '上林县西燕镇寨鹿村内吉外至内吉内产业硬化路工程', '上林县西燕镇人民政府', '广西兴盈建设工程有限公司', '', '367343.00', '2021-03-30 00:00:00.000000', '1.00', '', '1.00', '公司内部项目', '1', '25', '2021-05-08 17:13:30.000000');
INSERT INTO `rl_finance_project` VALUES ('24', '上林县镇圩瑶族乡望河村古良庄至上朝庄屯级道路修缮工程', '上林县镇圩瑶族乡人民政府', '广西兴盈建设工程有限公司', '', '1179334.63', '2021-05-28 00:00:00.000000', '353800.38', '', '1.00', '公司内部项目', '1', '25', '2021-06-21 15:36:29.000000');
INSERT INTO `rl_finance_project` VALUES ('25', '上林县塘红乡马里村弄桃庄饮水安全巩固提升工程', '上林县塘红乡人民政府', '广西国中建设工程有限公司', '', '187000.00', '2020-11-13 00:00:00.000000', '1.00', '', '1.00', '已开票金额149600.00 未拨款', '1', '25', '2021-06-21 16:30:25.000000');
INSERT INTO `rl_finance_project` VALUES ('26', '塘红乡弄陈村树威饮水安全巩固提升工程', '上林县塘红乡人民政府', '广西国中建设工程有限公司', '', '149600.00', '2020-11-13 00:00:00.000000', '1.00', '', '1.00', '已开票金额136100.00', '1', '25', '2021-06-21 16:35:26.000000');
INSERT INTO `rl_finance_project` VALUES ('27', '大丰镇云城村云城庄往展河至那桐水渠防渗工程', '上林县农业农村局', '广西国中建设工程有限公司', '', '379488.00', '2020-04-09 00:00:00.000000', '1.00', '', '1.00', '已开票金额113846.4', '1', '25', '2021-06-21 16:55:00.000000');
INSERT INTO `rl_finance_project` VALUES ('28', '上林县白圩镇覃黄村二级公路庄至长塘庄道路修缮', '上林县白圩镇人民政府', '广西兴盈建设工程有限公司', '', '400876.00', '2021-04-28 00:00:00.000000', '1.00', '', '1.00', '已开票金额200438.00 未拨款  毕佳林项目 毕佳林项目', '2', '25', '2021-06-22 09:16:31.000000');
INSERT INTO `rl_finance_project` VALUES ('29', '上林县白圩镇大浪村塘岭庄至古楼庄硬化路路面修缮', '上林县白圩镇人民政府', '广西兴盈建设工程有限公司', '', '107395.00', '2021-04-14 00:00:00.000000', '1.00', '', '1.00', '已开票金额 53697.50） 未拨款', '2', '25', '2021-06-22 09:18:02.000000');
INSERT INTO `rl_finance_project` VALUES ('30', '上林县白圩镇陆永村陆鞋庄新江桥工程', '上林县交通运输局', '广西兴盈建设工程有限公司', '', '1144051.00', '2021-07-13 00:00:00.000000', '1.00', '', '1.00', '已开票456389', '1', '25', '2021-08-13 12:02:21.000000');
INSERT INTO `rl_finance_project` VALUES ('31', '巷贤镇万嘉村六乐庄至兰花基地道路扩建工程', '上林县巷贤镇人民政府', '广西兴盈建设工程有限公司', '', '294740.33', '2021-06-30 00:00:00.000000', '1.00', '', '1.00', '', '1', '25', '2021-08-17 17:33:31.000000');
INSERT INTO `rl_finance_project` VALUES ('32', '巷贤镇六联村古竹庄至临时终漂点产业硬化路工程', '上林县巷贤镇人民政府', '广西兴盈建设工程有限公司', '', '919883.69', '2021-07-14 00:00:00.000000', '1.00', '', '1.00', '', '1', '25', '2021-08-17 17:34:17.000000');
INSERT INTO `rl_finance_project` VALUES ('33', '上林县白圩镇龙楼村大虫片区优质“双高”糖料蔗基地水利化项目', '上林县水利局', '广西国中建设工程有限公司', '', '723653.50', '2021-04-29 00:00:00.000000', '1.00', '', '1.00', '', '1', '25', '2021-08-18 11:23:55.000000');
INSERT INTO `rl_finance_project` VALUES ('34', '上林县大丰镇中心学校综合楼及围墙维修改造项目', '上林县大丰镇中心学校', '广西润林建筑装饰工程有限公司', '', '158307.56', '2020-12-29 00:00:00.000000', '1.00', '', '1.00', '', '1', '25', '2021-08-20 09:34:54.000000');
INSERT INTO `rl_finance_project` VALUES ('35', '乔贤镇恭睦村下那沫庄通屯道路修缮工程', '上林县乔贤镇人民政府', '广西兴盈建设工程有限公司', '', '714990.37', '2021-08-18 00:00:00.000000', '1.00', '', '1.00', '', '1', '25', '2021-09-08 16:24:08.000000');
INSERT INTO `rl_finance_project` VALUES ('36', '上林县白圩镇朝韦桥工程', '上林交通运输局', '广西兴盈建设工程有限公司', '', '692216.00', '2021-09-22 00:00:00.000000', '1.00', '', '1.00', '', '1', '25', '2021-10-13 17:46:46.000000');
INSERT INTO `rl_finance_project` VALUES ('37', '上林县西燕镇云灵村云童庄拉单桥梁工程', '上林县西燕镇人民政府', '广西兴盈建设工程有限公司', '', '648027.00', '2021-12-22 00:00:00.000000', '1.00', '', '1.00', '', '1', '25', '2022-01-13 11:30:56.000000');
INSERT INTO `rl_finance_project` VALUES ('38', '城南小学维修改造项目工程', '上林县城南小学', '广西润林', '', '38936.86', '2021-11-01 00:00:00.000000', '1.00', '', '38936.86', '', '1', '25', '2022-01-13 11:54:21.000000');
INSERT INTO `rl_finance_project` VALUES ('39', '上林县水利局太阳能路灯采购及安装工程', '上林县水利局', '广西炳耀建筑工程有限公司', '', '340000.00', '2021-12-01 00:00:00.000000', '1.00', '', '340000.00', '', '1', '25', '2022-01-13 15:05:32.000000');