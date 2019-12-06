CREATE TABLE `test`.`user` (
  `id` int unsigned UNIQUE NOT NULL  AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户id',
  `email` varchar(45) NOT NULL  COMMENT '邮箱',
  `password` varchar(45) NOT NULL   COMMENT '密码',
  `nickname` varchar(45) DEFAULT '' COMMENT '昵称',
  `avatar` varchar(45) DEFAULT '' COMMENT '头像',
  `signature` varchar(100) DEFAULT '' COMMENT '个性签名',
  `sex` int unsigned DEFAULT 0 COMMENT '性别',
  `register_time` integer unsigned DEFAULT 0 COMMENT '注册时间',
  `last_login_time` integer unsigned DEFAULT 0 COMMENT '最近登录时间',
  `token` varchar(200) DEFAULT '' COMMENT 'token 7天有效期',

  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL on update CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` TIMESTAMP NULL COMMENT '删除时间',
    # `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

INSERT INTO `test`.`user` (`email`, `password` ) VALUES ( '123@qq.com', 'test123456');
