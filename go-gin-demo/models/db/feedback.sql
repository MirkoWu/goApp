CREATE TABLE `feedback` (
  `id` int unsigned UNIQUE NOT NULL  AUTO_INCREMENT,
  `created_at` integer DEFAULT '0' COMMENT '创建时间',
  `updated_at` integer DEFAULT '0' COMMENT '更新时间',

  `user_id` int unsigned NOT NULL COMMENT '用户id',
  `title` varchar(50) NOT NULL  COMMENT '标题',
  `content` varchar(200) NOT NULL   COMMENT '反馈内容',
  `contact` varchar(30) DEFAULT '' COMMENT '联系方式',
  `submit_time` integer DEFAULT '0' COMMENT '提交时间',

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='反馈表';

INSERT INTO `test`.`feedback` (`user_id`, `title`, `content` ) VALUES ( '123', 'title', 'content');
