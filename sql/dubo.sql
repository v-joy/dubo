 create database dubo;
use dubo;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '����',
  `title` varchar(100) NOT NULL COMMENT '��Ŀ',
  `image` varchar(100) NOT NULL COMMENT 'ͼƬ',
  `content` varchar(4000) NOT NULL COMMENT '����',
  `categoryid` tinyint(3) NOT NULL COMMENT '���id',
  `source` varchar(50) NOT NULL COMMENT '��Դ',
  `read_count` int(11) NOT NULL COMMENT '�Ķ���',
  `post_time` int(11) NOT NULL COMMENT '����ʱ��',
  `create_time` int(11) NOT NULL COMMENT '����ʱ��',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '����',
  `name` varchar(100) NOT NULL COMMENT '��Ŀ',
  `image` varchar(100) NOT NULL COMMENT 'ͼƬ',
  `subname` varchar(4000) NOT NULL COMMENT '����',
  `create_time` int(11) NOT NULL COMMENT '����ʱ��',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
