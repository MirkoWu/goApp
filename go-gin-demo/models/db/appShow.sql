CREATE TABLE `test`.`app_show` (
  `id` int unsigned UNIQUE NOT NULL  AUTO_INCREMENT,
  `app_id` int unsigned UNIQUE NOT NULL COMMENT '应用id',
  `name` varchar(20) NOT NULL  COMMENT '名称',
  `intro` varchar(300) NOT NULL   COMMENT '应用介绍',
  `screenshots` varchar(300)  COMMENT '联系方式',
  `link_url` varchar(50) COMMENT 'app展示链接地址',
  `apk_url` varchar(50) COMMENT 'apk下载地址',
  `app_version_code` integer unsigned  NOT NULL   COMMENT '应用版本号',
  `app_version` varchar(20)  NOT NULL   COMMENT '应用版本',
  `apk_size` varchar(10) COMMENT 'apk大小',
  `app_pn` varchar(50)  NOT NULL COMMENT '应用包名',
  `is_show` int NOT NULL DEFAULT '1'  COMMENT '是否显示 0不显示 1显示',

  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL  COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL on update CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` TIMESTAMP NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用推荐表';



