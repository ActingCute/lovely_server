-- phpMyAdmin SQL Dump
-- version 4.7.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: 2019-04-30 04:52:55
-- 服务器版本： 5.5.56
-- PHP Version: 5.6.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `lovely`
--

-- --------------------------------------------------------

--
-- 表的结构 `comment`
--

CREATE TABLE `comment` (
  `id` int(64) NOT NULL COMMENT '自增ID',
  `user_id` int(64) NOT NULL COMMENT '用户ID',
  `url` varchar(255) NOT NULL COMMENT '页面url',
  `father_id` int(64) DEFAULT '0' COMMENT '父级ID',
  `son_id` int(64) DEFAULT '0' COMMENT '子级ID',
  `content` text NOT NULL COMMENT '评论内容',
  `update_time` datetime NOT NULL COMMENT '评论时间',
  `delete_time` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `comment`
--

INSERT INTO `comment` (`id`, `user_id`, `url`, `father_id`, `son_id`, `content`, `update_time`, `delete_time`) VALUES
(1, 1, '/Blog/20190418/23133944/', 0, 0, '爱你哟', '2019-04-28 11:40:35', NULL),
(2, 1, '/Blog/20190418/23133944/', 0, 0, '爱你哟', '2019-04-28 13:37:08', NULL),
(3, 0, '/Message/', 0, 0, '<h2><a id=\"_0\"></a>很认真的想了想</h2>\n<h3><a id=\"_1\"></a>可是不知道说啥</h3>\n<h1><a id=\"_2\"></a>走开走开：很太</h1>\n', '2019-04-30 10:08:52', NULL),
(4, 1, '/Message/', 0, 0, '<h2><a id=\"_0\"></a>很认真的想了想</h2>\n<h3><a id=\"_1\"></a>可是不知道说啥</h3>\n<h1><a id=\"_2\"></a>走开走开：很太</h1>\n', '2019-04-30 10:17:28', NULL),
(5, 1, '/Message/', 0, 0, '<h2><a id=\"_0\"></a>很认真的想了想</h2>\n<h3><a id=\"_1\"></a>可是不知道说啥</h3>\n<h1><a id=\"_2\"></a>走开走开：很太</h1>\n', '2019-04-30 10:18:55', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `count`
--

CREATE TABLE `count` (
  `id` int(64) NOT NULL COMMENT '自增ID',
  `url` varchar(255) NOT NULL COMMENT '页面url',
  `count` int(64) DEFAULT '1' COMMENT '阅读数',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` int(64) NOT NULL COMMENT '自增ID',
  `website` varchar(255) DEFAULT NULL COMMENT '网站',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `name` varchar(255) NOT NULL COMMENT '名字',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `website`, `email`, `name`, `delete_time`) VALUES
(1, 'www.haibarai.com', 'string', 'ActingCute酱', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `comment`
--
ALTER TABLE `comment`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `count`
--
ALTER TABLE `count`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `comment`
--
ALTER TABLE `comment`
  MODIFY `id` int(64) NOT NULL AUTO_INCREMENT COMMENT '自增ID', AUTO_INCREMENT=6;
--
-- 使用表AUTO_INCREMENT `count`
--
ALTER TABLE `count`
  MODIFY `id` int(64) NOT NULL AUTO_INCREMENT COMMENT '自增ID';
--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` int(64) NOT NULL AUTO_INCREMENT COMMENT '自增ID', AUTO_INCREMENT=2;COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
