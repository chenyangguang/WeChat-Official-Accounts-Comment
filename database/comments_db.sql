drop database comments if exists comments;
create database wechat_comments;

--
use wechat_comments;

-- 用户
drop table users if exists  users;
create table users(
`id` int(11) unsigned not null,
 `uid` varchar(255) UNSIGNED NOT NULL ,
 `openid` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户openid',
 `nickname` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户昵称',
 `gender` tinyint(1) unsigned not null  default 0 comment '性别',
 primary key (`id`),
 unique key (`uid`)
)ENGINE=InnoDB default charset utf8 comment '用户表';

-- 留言
drop table comments if exists comments;
create table comments (
`id` INT(11) UNSIGNED NOT NULL,
`content` varchar(255) not null default '' comment '内容',
`article_id` varchar(255) not null default '' comment '文章id',
`comment_uid` varchar(255) not null default '' comment '评论者',
`is_top` tinyint(1) unsigned not null default 0 comment '置顶状态, 0:默认,1:置顶',
`status` tinyint(1) unsigned not null default 0 comment '状态, 0:默认, 1:屏蔽, 2:删除',
`created_at` int(11) unsigned not null default 0 comment '生成时间',
`updated_at` int(11) unsigned not null default 0 comment '更新时间',
PRIMARY  KEY (`id`),
KEY (`article_id`) USING BTREE
)ENGINED=InnoDB default charset utf8 comment '评论表';

-- 点赞


