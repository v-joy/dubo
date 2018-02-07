 create database dubo;
use dubo;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(100) NOT NULL COMMENT '题目',
  `image` varchar(100) NOT NULL COMMENT '图片',
  `content` varchar(4000) NOT NULL COMMENT '内容',
  `categoryid` tinyint(3) NOT NULL COMMENT '类别id',
  `source` varchar(50) NOT NULL COMMENT '来源',
  `read_count` int(11) NOT NULL COMMENT '阅读数',
  `post_time` int(11) NOT NULL COMMENT '发表时间',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL COMMENT '题目',
  `image` varchar(100) NOT NULL COMMENT '图片',
  `subname` varchar(4000) NOT NULL COMMENT '内容',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
