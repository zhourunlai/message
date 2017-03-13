SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `message`
--

-- --------------------------------------------------------

--
-- 表的结构 `chats`
--

CREATE TABLE `chats` (
  `id` int(11) NOT NULL COMMENT '消息编号',
  `sender` varchar(64) NOT NULL COMMENT '发送者',
  `receiver` varchar(64) NOT NULL COMMENT '接收者',
  `content` text NOT NULL COMMENT '消息内容',
  `send_time` int(11) NOT NULL COMMENT '发送时间',
  `is_del` int(11) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `is_read` int(11) NOT NULL DEFAULT '0' COMMENT '是否已读'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `contacts`
--

CREATE TABLE `contacts` (
  `username` varchar(32) NOT NULL COMMENT '本用户',
  `contact` varchar(32) NOT NULL COMMENT '联系人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password` char(32) NOT NULL COMMENT '密码',
  `create_time` int(11) NOT NULL COMMENT '注册时间',
  `last_time` int(11) DEFAULT NULL COMMENT '上次登录时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `chats`
--
ALTER TABLE `chats`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `contacts`
--
ALTER TABLE `contacts`
  ADD PRIMARY KEY (`username`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`username`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `chats`
--
ALTER TABLE `chats`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '消息编号';
