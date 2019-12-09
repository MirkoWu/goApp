CREATE TABLE `test`.`app_tab` (
  `id` int unsigned UNIQUE NOT NULL  AUTO_INCREMENT,
  `tab_id` int unsigned UNIQUE NOT NULL COMMENT 'tab id',
  `title` varchar(50) NOT NULL  COMMENT '标题',
  `type` int NOT NULL DEFAULT '0'  COMMENT '类型 1首页，2抽签 3.游戏 4我的',
  `is_show` int NOT NULL DEFAULT '1'  COMMENT '是否显示 0不显示 1显示',

  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL  COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL on update CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` TIMESTAMP NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='底部Tab表';

