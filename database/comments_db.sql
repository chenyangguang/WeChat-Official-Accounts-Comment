drop database comments if exists comments;
create database wechat_comments;

--
use wechat_comments;

-- 用户
drop table  if exists wechat_comments.users;
create table users(
 `uid` bigint(20) UNSIGNED NOT NULL auto_increment,
 `openid` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户openid',
 `nickname` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户昵称',
 `gender` tinyint(1) unsigned not null  default 0 comment '性别',
 `avatar` varchar(255) NOT NULL default '' comment '头像url',
 `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP  comment '生成时间',
 `updated_at` Timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  comment '更新时间',
 `deleted_at` Timestamp  comment '删除时间',
 primary key (`uid`)
)ENGINE=InnoDB default charset utf8 comment '评论用户表';

-- 留言
drop table if exists wechat_comments.comments;
create table comments(
`id` bigint(20) UNSIGNED NOT NULL auto_increment,
`content` varchar(255) not null default '' comment '内容',
`article_id` varchar(255) not null default '' comment '文章id',
`comment_uid` varchar(255) not null default '' comment '评论者',
`parent_id` int(11) unsigned not null default 0 comment '父评论id',
`is_top` tinyint(1) unsigned not null default 0 comment '置顶状态, 0:默认,1:置顶',
`status` tinyint(1) unsigned not null default 0 comment '状态, 0:默认, 1:屏蔽, 2:删除',
`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP  comment '生成时间',
`updated_at` Timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  comment '更新时间',
`deleted_at` Timestamp  comment '删除时间',
PRIMARY  KEY (`id`),
KEY (`article_id`) USING BTREE,
KEY (`comment_uid`) USING BTREE
)ENGINE=InnoDB default charset utf8 comment '评论表';



