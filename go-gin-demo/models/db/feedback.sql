CREATE TABLE `test`.`feedback` (
  `id` int unsigned UNIQUE NOT NULL  AUTO_INCREMENT,
  `feedback_id` int unsigned UNIQUE NOT NULL COMMENT '反馈id',
  `user_id` int unsigned NOT NULL COMMENT '用户id',
  `title` varchar(50) NOT NULL  COMMENT '标题',
  `content` varchar(200) NOT NULL   COMMENT '反馈内容',
  `contact` varchar(30) DEFAULT '' COMMENT '联系方式',
  `submit_time` integer unsigned DEFAULT 0 COMMENT '提交时间',

  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL  COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL on update CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` TIMESTAMP NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='反馈表';

INSERT INTO `test`.`feedback` (`user_id`, `title`, `content` ) VALUES ( '123', 'title', 'content');
