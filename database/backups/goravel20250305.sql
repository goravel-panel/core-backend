/*
 Navicat Premium Data Transfer

 Source Server         : gocmf-mysql
 Source Server Type    : MySQL
 Source Server Version : 50740 (5.7.40-log)
 Source Host           : 123.56.122.150:3306
 Source Schema         : gocmf

 Target Server Type    : MySQL
 Target Server Version : 50740 (5.7.40-log)
 File Encoding         : 65001

 Date: 05/03/2025 15:27:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `category_id` int(10) NOT NULL DEFAULT '0',
  `from` varchar(255) NOT NULL COMMENT '来源',
  `title` varchar(255) NOT NULL COMMENT '新闻标题',
  `content` text NOT NULL COMMENT '新闻内容',
  `cover` text COMMENT '封面地址',
  `is_enable` tinyint(3) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `article/title_index` (`title`),
  KEY `article/is_enable_index` (`is_enable`),
  KEY `article/from_index` (`from`),
  KEY `article/category_id_index` (`category_id`),
  KEY `article/created_at_index` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` (`id`, `category_id`, `from`, `title`, `content`, `cover`, `is_enable`, `created_at`, `updated_at`) VALUES (1, 1, '博泽名车汇', '被315晚会点名后，宝马中国的这份声明，宝马车主能接受吗？', '<p style=\"text-align: start;\">3月15日315晚会如期举行，在本届315晚会中，宝马530Li因为传动轴异响被点名，成为了今年315汽车维权事件关注的焦点。</p><p style=\"text-align: start;\">第一个案例是2023年6月车主李先生花了45万购买的宝马530Li开了没多久底盘就传出了金属摩擦的声音。车辆在从R档切换到D挡的过程中，传动系统会发出尖锐的摩擦声和金属撞击声。</p><p style=\"text-align: start;\">后到宝马汽车4S店进行检测，售后工作人员告知李先生是传动轴异响。是小概率事件，可以按照三包规定进行维修。如果需要更换，会帮忙申请更换传动轴。</p><p style=\"text-align: start;\"><img src=\"https://n.sinaimg.cn/sinakd20111/192/w1024h768/20240316/51ce-b1aa9ec8fd109956150621e1fcc93a5a.jpg\" alt=\"\" data-href=\"\" style=\"\"></p><p style=\"text-align: start;\">对于宝马汽车4S店的这个答复，李先生不能接受。他认为刚买的新车就要更换传动轴，等同于做了“特大手术”。随后李先生致电宝马厂家客服，宝马厂家400客服给他的答复是厂家把售后事宜授权给当地经销商与客户沟通，处理车辆问题要通过当地经销商。</p><p style=\"text-align: start;\">截止到2023年底，关于自己爱车异响以及售后解决方案李先生跟宝马4S店售后人员始终沟通未果。</p><p style=\"text-align: start;\"><img src=\"https://n.sinaimg.cn/sinakd20111/289/w1024h865/20240316/8dfc-0fbd2fcb48413bdca70ab8df49e05ca1.jpg\" alt=\"\" data-href=\"\" style=\"\"></p><p style=\"text-align: start;\">既然上了315晚会，那么宝马530Li传动轴异响的问题肯定不是个例。晚会还提到了陈先生、冯先生等多位宝马530Li的车主，他们所遇到的问题跟李先生一样，也都是因为车辆传动轴异响。</p><p style=\"text-align: start;\">值得注意的是，上述两位车主的宝马530Li也都是去年下半年购买的新车。其中陈先生的宝马530Li购于2023年7月，冯先生的宝马530Li购于2023年9月。</p><p style=\"text-align: start;\"><img src=\"https://n.sinaimg.cn/sinakd20111/576/w1024h1152/20240316/c0b2-25a1e90b88c4554c0a31601178bc0215.jpg\" alt=\"\" data-href=\"\" style=\"\"></p><p style=\"text-align: start;\">跟李先生不同的是，陈先生和冯先生在发现车辆传动轴异响跟宝马4S店交涉后经销商的答复是可以免费更换，同时还可以补偿2000元代金券。</p><p style=\"text-align: start;\">让冯先生不理解的是，宝马经销商在大阳可以免费换轴和补偿代金券的同时，却并不承认问题。</p><p style=\"text-align: start;\">与此同时，在处理宝马530Li传动轴异响的过程中，宝马4S店售后负责人的答复也是五花八样：既有可以开到首保1万公里看看声音有没有去掉，如果声音没有了就皆大欢喜；也有即便是更换传动轴，以后不想用车也不会影响车辆转手的价格。</p><p style=\"text-align: start;\"><img src=\"https://n.sinaimg.cn/sinakd20111/192/w1024h768/20240316/de9a-2ce94a4790307c7eda9653722b1e7df7.jpg\" alt=\"\" data-href=\"\" style=\"\"></p><p style=\"text-align: start;\">另外两种说辞是厂家说属于是功能噪音，如果客户适应不了声音或者强烈要求换是可以换的，现在换了有可能以后还会响；全国统一2000元的补偿代金券，但前提是要换掉传动轴，如果不换就没有别的解决办法，投诉宝马宝马也不会直接有人处理等等等等……<br></p>', 'https://img.zcool.cn/community/016a2e5f110b9fa801215aa097202c.png', 1, '2024-03-06 17:39:29', '2024-04-26 17:28:26');
INSERT INTO `article` (`id`, `category_id`, `from`, `title`, `content`, `cover`, `is_enable`, `created_at`, `updated_at`) VALUES (2, 1, '博泽名车汇', '德国电动汽车销量需增长6倍，才能实现1,500万辆的目标', '<p style=\"text-align: start;\"> &nbsp; &nbsp; 盖世汽车讯，一项研究显示，德国的电动汽车销量需要大幅增长，才有可能实现该国的排放目标。</p><p style=\"text-align: start;\">　　可再生能源行业游说团体BEE称，德国新电动汽车的销量必须在未来三年翻两番，到2030年增长六倍，才能达到德国拥有1，500万辆电动汽车的目标。该组织在一项研究中表示，更有可能出现的结果是，德国将只有1，000万辆电动汽车，而且也不太可能实现其温室气体排放目标（差三分之一左右）。在德国减少污染的努力中，交通领域已经成为一个主要的落后者。</p><p style=\"text-align: start;\"><img src=\"https://n.sinaimg.cn/spider20240315/781/w500h281/20240315/f618-df4ed462581a8ff6b6ae0afb10bc7ed0.png\" alt=\"图片来源：大众\" data-href=\"\" style=\"\"></p><p style=\"text-align: start;\"><br></p><p style=\"text-align: start;\">　　BEE倡导通过公共交通激励措施、增加使用合成燃料和农业生物燃料，以及对德国的高速公路实行限速来遏制排放。该游说团体表示，即使德国成功实现其电动汽车的数量目标，交通运输领域也可能达不到排放目标，需要其他类似措施的支持。</p><p style=\"text-align: start;\">　　上个月，德国经济部长Robert Habeck在参观柏林的一家梅赛德斯-<a href=\"https://db.auto.sina.com.cn/b3.html?c=spr_auto_trackid_1e17744b130ea936\" target=\"_blank\">奔驰</a>工厂时也承认，“到2030年，我们将无法达到1，500万辆的目标”，“但技术发展和社会接受度，并不是线性发展的”。</p><p style=\"text-align: start;\">　　2023年，德国销售了超过52.4万辆纯电动汽车，高出其它欧洲市场的销量；插电式混动汽车销量则减少一半以上，仅为17.6万辆。1月30日，德国汽车协会VDA表示，由于电动汽车的补贴退坡拖累了需求，预计德国今年纯电动汽车的出货量将下降14%，至45.1万辆，将是近8年来首次出现下滑。</p>', 'https://img.zcool.cn/community/016a2e5f110b9fa801215aa097202c.png', 1, '2024-03-17 03:10:33', '2024-04-26 17:28:29');
INSERT INTO `article` (`id`, `category_id`, `from`, `title`, `content`, `cover`, `is_enable`, `created_at`, `updated_at`) VALUES (3, 1, '博泽名车汇', '威马汽车申请破产重整后，10万车主售后问题悬置', '<p style=\"text-align: start;\">去年10月申请破产重整的<a href=\"https://db.auto.sina.com.cn/b180.html?c=spr_auto_trackid_1e17744b130ea936\" target=\"_blank\">威马汽车</a>即将在3月底召开首次债权人会议，从目前进展来看想要恢复正常运营还需相当长的一段时间。而在此之前，接近10万名威马车主的售后问题依然难解。</p><p style=\"text-align: start;\">　　界面新闻查询第三方投诉平台注意到，多位威马车主投诉称，遇到故障后想要维修，却被告知没有配件，也不清楚什么时候配件才能够到位。</p><p style=\"text-align: start;\">　　一名威马网约车司机表示，他在2021年以11.6万元的落地价购买了一辆威马EX5，现在经常为售后问题头疼。此前他遭遇了一场轻微的追尾事故，左后大灯需要维修，只能去找个体维修商解决，单个大灯维修价格高达近3000元。</p><p style=\"text-align: start;\">　　另外由于没有合适的配件，直到现在他的尾灯罩子仍有裂痕。“我去淘宝上搜了配件，结果发过来是威马E5车型的，型号也对不上。”</p><p style=\"text-align: start;\">　　这位威马网约车车主透露，现在不论是自费或在维保期，都没有威马厂家提供的售后服务，只能车主自己寻找小的维修厂商自行解决。但这种小型维修商只能诊断基础问题，如果遇到动力电池故障，将很难得到妥善解决。</p><p style=\"text-align: start;\"><img src=\"https://n.sinaimg.cn/sinakd20240314s/215/w2048h1367/20240314/a1c5-cc5d35424d78ea168b201352dd611704.jpg\" alt=\"威马汽车申请破产重整后，10万车主售后问题悬置\" data-href=\"\" style=\"\"></p><p style=\"text-align: start;\">　　界面新闻分别致电上海两家威马官方授权的维修中心，均表示目前仍正常营业。一家维修中心相关工作人员告诉界面新闻，更换配件可以从威马原有的供应商处调货，只是不再有威马汽车的标识。</p><p style=\"text-align: start;\">　　另据社交平台上车主透露，部分地区授权门店由于配件不全，只能做基础保养服务，或者通过外采零部件的方式解决。还有一些官方门店已经改头换面，成为了其他汽车品牌的4S店。</p><p style=\"text-align: start;\">　　界面新闻在网购平台检索发现，部分汽配厂家会提供威马车型的零部件，多为雨刮器、保险杠装饰条、后护板等简单通用部件，售价在百元以内。并且，一些网购平台的店主会在商品选择栏里细分为原厂和副厂，后者则是非官方授权的配件。</p><p style=\"text-align: start;\">　　一位提供威马配件的店主向界面新闻介绍，原厂配件确保是全新的且由威马官方授权，而副厂配件采购自一家常州的工厂，价格相对更为便宜。</p><p style=\"text-align: start;\">　　从平台显示的店家月销量来看，相当一部分威马车主正在通过网购渠道解决基础售后难题。但是，一些威马车主在评论区里反馈，配件的质量堪忧，或者普遍存在配件型号不准的问题。</p><p style=\"text-align: start;\">　　根据我国颁布的《汽车销售管理实施办法》，生产企业应当及时向社会公布停产或停止销售的车型，并保证至少10年的配件供应以及相应售后服务。但威马汽车目前已经难以保证为10万名车主提供正常的维保服务。</p><p style=\"text-align: start;\">　　此外，不同于传统燃油车主要价值在于汽车本身，智能电动汽车的智能化配置和软件升级服务是车型溢价的重要来源之一。但在威马申请破产重组后，车主普遍反馈手机APP已经停止服务，车机系统运行也出现卡顿等问题。</p><p style=\"text-align: start;\">　　这意味着花更多钱购买更高智能化配置的车主，只能享受和入门款车型相差无几的体验。同时，威马车机系统也不再有专门的团队负责升级维护。</p><p style=\"text-align: start;\">　　相较之下，背靠大集团的汽车品牌在宣告破产后，其车主的售后问题仍能得到妥善处理。2022年<a href=\"https://db.auto.sina.com.cn/b72.html?c=spr_auto_trackid_1e17744b130ea936\" target=\"_blank\">讴歌</a>退出中国市场后，<a href=\"https://db.auto.sina.com.cn/b2.html?c=spr_auto_trackid_1e17744b130ea936\" target=\"_blank\">广汽本田</a>提供其后续服务；同年广汽菲克退市后，母公司Stellantis集团宣布为广汽菲克车主进行售后维保工作。</p>', 'https://img.zcool.cn/community/016a2e5f110b9fa801215aa097202c.png', 1, '2024-03-17 03:18:47', '2024-04-26 17:28:27');
INSERT INTO `article` (`id`, `category_id`, `from`, `title`, `content`, `cover`, `is_enable`, `created_at`, `updated_at`) VALUES (4, 1, '博泽名车汇', '现代起亚在韩国召回近17万辆电动汽车', '<p>3333</p>', 'https://img.zcool.cn/community/016a2e5f110b9fa801215aa097202c.png', 1, '2024-03-17 03:20:09', '2025-02-05 22:36:56');
COMMIT;

-- ----------------------------
-- Table structure for article_category
-- ----------------------------
DROP TABLE IF EXISTS `article_category`;
CREATE TABLE `article_category` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `sort` int(10) NOT NULL DEFAULT '0',
  `is_enable` tinyint(3) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `article_category/name_index` (`name`),
  KEY `article_category/is_enable_index` (`is_enable`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of article_category
-- ----------------------------
BEGIN;
INSERT INTO `article_category` (`id`, `name`, `remark`, `sort`, `is_enable`, `created_at`, `updated_at`) VALUES (1, '社会热点', '', 0, 1, '2024-04-20 22:28:56', '2024-04-26 17:28:16');
INSERT INTO `article_category` (`id`, `name`, `remark`, `sort`, `is_enable`, `created_at`, `updated_at`) VALUES (2, '新闻资讯', '', 0, 1, '2024-04-20 22:29:00', '2024-04-26 17:28:17');
COMMIT;

-- ----------------------------
-- Table structure for sys_admin
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin`;
CREATE TABLE `sys_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `account` varchar(20) NOT NULL DEFAULT '' COMMENT '登录账号',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '管理员昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '管理员头像',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '登录密码',
  `is_enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用，[-1 => 否, 1 => 是]',
  `is_root` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是root用户，[-1 => 否, 1 => 是]',
  `last_login_ip` varchar(20) NOT NULL DEFAULT '' COMMENT '最近登录ip',
  `last_login_time` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin/account_uindex` (`account`) USING BTREE,
  KEY `admin/name_index` (`name`),
  KEY `admin/is_enable_index` (`is_enable`),
  KEY `admin/is_root_index` (`is_root`),
  KEY `admin/created_at_index` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Records of sys_admin
-- ----------------------------
BEGIN;
INSERT INTO `sys_admin` (`id`, `account`, `name`, `avatar`, `password`, `is_enable`, `is_root`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`) VALUES (1, 'admin', '超级管理员', 'http://127.0.0.1:3000/static/images/20250122/f0EK09HT7fV9quaG6bNuKuCnk6FbrYKbQO1XBHE3.gif', '$2a$12$12CMb3vySqWErDuuxrvCt.nznMsgil7tVcsHSKV.0ypLcnh9kz6xG', 1, 1, '127.0.0.1', '2025-03-05 15:12:11', '2024-03-20 15:52:47', '2025-03-05 15:12:11');
INSERT INTO `sys_admin` (`id`, `account`, `name`, `avatar`, `password`, `is_enable`, `is_root`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`) VALUES (6, 'demo', 'demo', 'https://gouguoyin-1300075642.cos.ap-nanjing.myqcloud.com/images/20240512/YiCW0UtLwLk6jEpNLjuM7tKMF8vDDm4SDCjRIxlV.png', '$2a$12$NBYGxlp/GU2ACvHQnGzOzONyiolHeEFzsODY3oWHmDkV4Xyb1fqZ6', 1, 0, '127.0.0.1', '2025-01-25 21:25:28', '2024-06-01 11:22:12', '2025-01-25 21:25:50');
COMMIT;

-- ----------------------------
-- Table structure for sys_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin_role`;
CREATE TABLE `sys_admin_role` (
  `admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  UNIQUE KEY `admin_role/admin_role_id_uindex` (`admin_id`,`role_id`),
  KEY `admin_role/admin_id_index` (`admin_id`),
  KEY `admin_role/role_id_index` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员、角色关联表';

-- ----------------------------
-- Records of sys_admin_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_admin_role` (`admin_id`, `role_id`) VALUES (3, 1);
INSERT INTO `sys_admin_role` (`admin_id`, `role_id`) VALUES (3, 4);
INSERT INTO `sys_admin_role` (`admin_id`, `role_id`) VALUES (6, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_attach
-- ----------------------------
DROP TABLE IF EXISTS `sys_attach`;
CREATE TABLE `sys_attach` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `category_id` int(10) NOT NULL DEFAULT '0' COMMENT '分类ID',
  `type` varchar(10) NOT NULL DEFAULT '' COMMENT '附件类型\nenum[image => 图片, file = 文件， video = 视频,  audio = 音频]',
  `disk` varchar(10) NOT NULL DEFAULT '' COMMENT '上传驱动\nenum[local => 本地, oss = 阿里云存储， cos = 腾讯云储存]',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '附件名称',
  `mime` varchar(10) NOT NULL DEFAULT '' COMMENT '类型标准',
  `ext` varchar(10) NOT NULL COMMENT '附件扩展名',
  `size` int(10) NOT NULL DEFAULT '0' COMMENT '附件大小',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT '附件路径',
  `url` varchar(255) NOT NULL COMMENT '附件地址',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `attach/category_id_index` (`category_id`),
  KEY `attach/type_index` (`type`),
  KEY `attach/disk_index` (`disk`),
  KEY `attach/name_index` (`name`),
  KEY `attach/mime_index` (`mime`),
  KEY `attach/ext_index` (`ext`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4 COMMENT='附件表';

-- ----------------------------
-- Records of sys_attach
-- ----------------------------
BEGIN;
INSERT INTO `sys_attach` (`id`, `category_id`, `type`, `disk`, `name`, `mime`, `ext`, `size`, `path`, `url`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 3, 'image', 'cos', 'logo.png', 'image/png', 'png', 13115, 'images/20240512/YiCW0UtLwLk6jEpNLjuM7tKMF8vDDm4SDCjRIxlV.png', 'https://gouguoyin-1300075642.cos.ap-nanjing.myqcloud.com/images/20240512/YiCW0UtLwLk6jEpNLjuM7tKMF8vDDm4SDCjRIxlV.png', '2024-05-12 17:58:47', '2025-01-02 22:04:38', NULL);
INSERT INTO `sys_attach` (`id`, `category_id`, `type`, `disk`, `name`, `mime`, `ext`, `size`, `path`, `url`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 0, 'image', 'oss', 'login-banner.png', 'image/png', 'png', 22065, 'images/20250104/hXxrjH5EcYFU20FDWgPVHzh983A3U1qLJto5VaCq.png', 'http://cdn.gouguoyin.cn/images/20250104/hXxrjH5EcYFU20FDWgPVHzh983A3U1qLJto5VaCq.png', '2025-01-04 23:03:15', '2025-01-04 23:03:15', NULL);
INSERT INTO `sys_attach` (`id`, `category_id`, `type`, `disk`, `name`, `mime`, `ext`, `size`, `path`, `url`, `created_at`, `updated_at`, `deleted_at`) VALUES (57, 0, 'image', 'local', 'app_icon.0ff07bfb.png', 'image/png', 'png', 11411, 'images/20250122/MrI1HsPgdzpSaYffQll4HT96sHBS2IEjoWOgHvTN.png', 'https://cdn.gouguoyin.cn/images/20250224/WUaolcGLrsNAxhWhu8GG9ByZNylcOanIRTZzgmln.jpg', '2025-01-22 12:09:46', '2025-03-05 15:17:04', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_attach_category
-- ----------------------------
DROP TABLE IF EXISTS `sys_attach_category`;
CREATE TABLE `sys_attach_category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` int(10) NOT NULL DEFAULT '0' COMMENT '父级ID',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '分类名',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序数字',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  `is_enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用，[-1 => 否, 1 => 是]',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `attach_category/parent_id_index` (`parent_id`),
  KEY `attach_category/name_index` (`name`),
  KEY `attach_category/is_enable_index` (`is_enable`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='附件分类表';

-- ----------------------------
-- Records of sys_attach_category
-- ----------------------------
BEGIN;
INSERT INTO `sys_attach_category` (`id`, `parent_id`, `name`, `sort`, `remark`, `is_enable`, `created_at`, `updated_at`) VALUES (3, 0, '头像', 0, '', 1, '2024-12-30 22:51:39', '2025-01-18 14:21:45');
COMMIT;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(20) NOT NULL DEFAULT '' COMMENT '配置类型',
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '配置名称',
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '配置标题',
  `value` json NOT NULL COMMENT '配置值',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序字段',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是默认值，[-1=> 否, 1 => 是]',
  `is_enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用，[-1 => 否, 1 => 是]',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `config/type_name_uindex` (`type`,`name`),
  KEY `config/name_index` (`name`),
  KEY `config/is_enable_index` (`is_enable`),
  KEY `config/type_index` (`type`) USING BTREE,
  KEY `config/is_default_index` (`is_default`),
  KEY `config/type_is_enable_index` (`type`,`is_enable`),
  KEY `config/name_is_enable_index` (`name`,`is_enable`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='配置表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_config` (`id`, `type`, `name`, `title`, `value`, `remark`, `sort`, `is_default`, `is_enable`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'base', 'backend', '基础配置', '{\"logo\": \"https://cdn.gouguoyin.cn/images/20250224/WUaolcGLrsNAxhWhu8GG9ByZNylcOanIRTZzgmln.jpg\", \"name\": \"管理平台\", \"email\": \"245629560@qq.com\", \"copyright\": \"Copyright 2025-2030 © goravel-admin All Rights Reserved.\", \"loginBackground\": \"http://cdn.gouguoyin.cn/images/20250104/hXxrjH5EcYFU20FDWgPVHzh983A3U1qLJto5VaCq.png\", \"useLoginCaptcha\": true}', '', 0, 0, 1, '2024-03-20 16:52:24', '2025-03-05 15:17:50', NULL);
INSERT INTO `sys_config` (`id`, `type`, `name`, `title`, `value`, `remark`, `sort`, `is_default`, `is_enable`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'captcha', 'captcha', '验证码配置', '{\"type\": \"digit\", \"length\": 4, \"noiseCount\": 0, \"digitSource\": \"0123456789\", \"letterSource\": \"ABCDEFGHJKMNOQRSTUVXYZabcdefghjkmnoqrstuvxyz\", \"stringSource\": \"0123456789ABCDEFGHJKMNOQRSTUVXYZabcdefghjkmnoqrstuvxyz\"}', '', 0, 0, 1, '2024-03-31 22:49:26', '2025-01-18 13:57:35', NULL);
INSERT INTO `sys_config` (`id`, `type`, `name`, `title`, `value`, `remark`, `sort`, `is_default`, `is_enable`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'storage', 'local', '本地存储', '{\"name\": \"local\"}', '存储到本服务器上，无需申请', 50, 1, 1, '2024-03-29 11:58:22', '2025-03-05 15:11:16', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `admin_id` int(10) NOT NULL DEFAULT '0',
  `account` varchar(255) NOT NULL DEFAULT '',
  `ip` varchar(20) NOT NULL DEFAULT '',
  `location` varchar(255) NOT NULL DEFAULT '',
  `os` varchar(255) NOT NULL DEFAULT '',
  `browser` varchar(255) NOT NULL DEFAULT '',
  `status` varchar(10) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `login_log/account_index` (`account`),
  KEY `login_log/admin_id_index` (`admin_id`),
  KEY `login_log/ip_index` (`ip`),
  KEY `login_log/status_index` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` int(10) NOT NULL DEFAULT '0' COMMENT '父级ID',
  `type` varchar(10) NOT NULL DEFAULT '' COMMENT '菜单类型：[catalogue => 目录, menu => 菜单, button => 按钮, link => 外链, iframe => 内嵌]',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单图标',
  `route_path` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单路径',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '链接地址',
  `perm` varchar(255) NOT NULL DEFAULT '' COMMENT '权限标识',
  `param` varchar(255) NOT NULL DEFAULT '' COMMENT '路由参数',
  `component` varchar(255) NOT NULL DEFAULT '' COMMENT '组件路径',
  `selected_path` varchar(255) NOT NULL DEFAULT '' COMMENT '高亮菜单路径',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序数字',
  `is_cache` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否缓存：[-1 => 否, 1 => 是]',
  `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示：[-1 => 否, 1 => 是]',
  `is_enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用，[-1 => 否, 1 => 是]',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `menu/parent_id_index` (`parent_id`),
  KEY `menu/type_index` (`type`),
  KEY `menu/name_index` (`name`),
  KEY `menu/perm_index` (`perm`),
  KEY `menu/is_show_index` (`is_show`),
  KEY `menu/is_enable_index` (`is_enable`),
  KEY `menu/is_cache_index` (`is_cache`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (1, 0, 'menu', '工作台', 'el-icon-Monitor', 'workbench', '', 'workbench', '', 'workbench/index', '', 9999, 1, 1, 1, NULL, '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (2, 0, 'catalogue', '权限管理', 'el-icon-Lock', 'permission', '', '', '', '', '', -11000, 1, 1, 1, '2024-04-03 21:07:34', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (3, 2, 'menu', '账号管理', '', 'account', '', '', '', 'permission/account/index', '', 0, 1, 1, 1, '2024-04-03 21:07:38', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (4, 2, 'menu', '角色管理', '', 'role', '', '', '', 'permission/role/index', '', 0, 1, 1, 1, '2024-04-03 21:07:43', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (5, 2, 'menu', '菜单管理', '', 'menu', '', '', '', 'permission/menu/index', '', 0, 1, 1, 1, '2024-04-03 21:07:50', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (7, 3, 'button', '新增账号', '', '', '', 'account.add', '', '', '', 0, 1, 0, 1, '2024-04-15 20:35:37', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (8, 3, 'button', '编辑账号', '', '', '', 'account.edit', '', '', '', 0, 1, 0, 1, '2024-04-15 20:36:06', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (9, 3, 'button', '删除账号', '', '', '', 'account.delete', '', '', '', 0, 1, 0, 1, '2024-04-15 20:37:01', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (10, 4, 'button', '添加角色', '', '', '', 'role.add', '', '', '', 0, 1, 0, 1, '2024-04-15 20:38:53', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (11, 4, 'button', '编辑角色', '', '', '', 'role.edit', '', '', '', 0, 1, 0, 1, '2024-04-15 20:39:13', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (12, 4, 'button', '删除角色', '', '', '', 'role.delete', '', '', '', 0, 1, 0, 1, '2024-04-15 20:39:30', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (13, 5, 'button', '添加菜单', '', '', '', 'menu.add', '', '', '', 0, 1, 0, 1, '2024-04-15 20:45:35', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (14, 5, 'button', '编辑菜单', '', '', '', 'menu.edit', '', '', '', 0, 1, 0, 1, '2024-04-15 20:46:09', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (15, 0, 'catalogue', '系统设置', 'el-icon-Setting', 'setting', '', '', '', '', '', -12000, 1, 1, 1, '2024-04-15 21:22:56', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (16, 15, 'menu', '基础设置', '', 'base', '', '', '', 'setting/base/edit', '', 100, 1, 1, 1, '2024-04-15 21:24:07', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (18, 15, 'menu', '验证码设置', '', 'captcha', '', '', '', 'setting/captcha/edit', '', 70, 1, 1, 1, '2024-04-15 21:41:48', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (19, 0, 'catalogue', '内容管理', 'local-icon-copy', 'article', '', '', '', '', '', 200, 1, 1, 1, NULL, '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (20, 19, 'menu', '文章列表', '', 'list', '', '', '', 'article/index', '', 0, 1, 1, 1, '2024-04-15 23:55:34', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (21, 19, 'menu', '新增文章', '', 'add', '', 'article.add', '', 'article/edit', 'article/list', 0, 1, -1, 1, '2024-04-15 23:56:49', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (22, 19, 'menu', '编辑文章', '', 'edit', '', 'article.edit', '', 'article/edit', 'article/list', 0, 1, -1, 1, '2024-04-16 07:53:45', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (23, 5, 'button', '删除菜单', '', '', '', 'memu.delete', '', '', '', 0, 1, 0, 1, '2024-04-16 23:19:02', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (25, 19, 'menu', '文章分类', '', 'category', '', '', '', 'article/category/index', '', 90, 1, 1, 1, '2024-04-19 22:23:04', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (26, 0, 'iframe', '内嵌页面', 'el-icon-Film', 'iframe', 'https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm', '', '', '', '', -15000, 1, 1, 1, '2024-04-19 23:47:33', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (27, 0, 'link', '外部链接', 'el-icon-Link', '', 'https://element-plus.org/zh-CN', '', '', '', '', -19998, 1, 1, 1, '2024-04-21 11:40:17', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (30, 0, 'catalogue', '附件管理', 'el-icon-FolderRemove', 'attach', '', '', '', '', '', -10000, 1, 1, 1, '2024-04-22 17:56:09', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (31, 30, 'menu', '图库管理', '', 'list', '', '', '', 'attach/index', '', 0, 1, 1, 1, '2024-04-22 17:56:37', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (45, 36, 'menu', '文章1', '', 'list', '', '', '', 'article/index', '', 0, 1, 1, 1, '2025-01-13 17:09:03', '2025-02-04 21:19:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `type`, `name`, `icon`, `route_path`, `url`, `perm`, `param`, `component`, `selected_path`, `sort`, `is_cache`, `is_show`, `is_enable`, `created_at`, `updated_at`) VALUES (46, 36, 'menu', '文章8', '', 'list8', '', '', '{\"type\": 1, \"name\": \"admin\"}', 'logger/login', '', 0, 1, 1, 1, NULL, '2025-02-04 21:19:10');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '角色名',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序数字',
  `is_enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用，[-1 => 否, 1 => 是]',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `role/is_enable_index` (`is_enable`),
  KEY `role/name_index` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`id`, `name`, `remark`, `sort`, `is_enable`, `created_at`, `updated_at`) VALUES (1, '配置管理员', '专门管理配置信息', 4, 1, '2024-04-03 16:01:40', '2025-01-20 11:51:52');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` int(11) NOT NULL DEFAULT '0',
  `menu_id` int(11) NOT NULL DEFAULT '0',
  UNIQUE KEY `role_menu/role_menu_id_uindex` (`role_id`,`menu_id`),
  KEY `role_menu/role_id_index` (`role_id`),
  KEY `role_menu/menu_id_index` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色、菜单关联表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 15);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 16);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 17);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 18);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 24);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 26);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 32);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 33);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 34);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 35);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(50) NOT NULL COMMENT '用户账号',
  `name` varchar(255) NOT NULL COMMENT '用户昵称',
  `memo` varchar(50) NOT NULL DEFAULT '' COMMENT '备注名',
  `avatar` varchar(250) NOT NULL DEFAULT '' COMMENT '用户头像',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码',
  `points` int(10) NOT NULL DEFAULT '0' COMMENT '会员积分',
  `level` int(11) NOT NULL DEFAULT '10' COMMENT '会员等级，备用',
  `is_enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否禁用: [0=否, 1=是]',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user/account_uindex` (`account`) USING BTREE,
  KEY `user/is_enable_index` (`is_enable`),
  KEY `user/level_index` (`level`),
  KEY `user/email_index` (`email`),
  KEY `user/mobile_index` (`mobile`),
  KEY `user/name_index` (`name`),
  KEY `user/created_at_index` (`created_at`),
  KEY `user/mobile_uindex` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
