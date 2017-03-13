#实现一个 Web 上的私信系统  

##功能：  

- [ ] * 用户可以注册、登录。需要 id（可以自己决定 email 或者 username）和 password  
- [ ] * 用户登录后，进入联系人列表页面  
- [ ] - 可以看到自己所有的联系人  
- [ ] - 每个联系人需要显示对方 id 以及未读私信数量提醒  
- [ ] - 用户可以通过 id 添加新联系人（可以不需要对方同意）  
- [ ] - 用户可以删除某个联系人，但保留与对方用户的消息等数据。当再次添加新联系人时，消息等数据都还在  
- [ ] * 点击一个联系人会进入聊天界面，同时未读消息置为 0  
- [ ] - 可以看到和某个用户的历史消息  
- [ ] - 能够在这里收发私信（不需要实时，可以刷一下页面才看到新消息）  
- [ ] - 当用户 A 发私信给用户 B 时，如果 A 还不是 B 的联系人，应该自动把 A 添加为 B 的联系人，并能够在 B 的联系人列表正常显示（不需要实时）  
- [ ] - 用户可以删除自己发的消息  

加分项：  

- [ ] * 联系人列表页面未读消息数实时更新  
- [ ] * 聊天界面新消息实时接收  
- [ ] * 自动把 A 添加为 B 联系人时，B 实时更新联系人列表  
- [ ] * 部署，可在线演示  

##progress  

###Plan  
>Deadline: 2 days  
>Add to Wunderlist  

![1](http://ww1.sinaimg.cn/large/9f47c048gy1fdlms5yaoqj21kw0zk7wh)

###Prototype  

![2](http://ww1.sinaimg.cn/large/9f47c048gy1fdlnu5ebs0j21kw16oakv)


###Database design  

```
CREATE TABLE `chats` (
  `id` int(11) NOT NULL COMMENT '消息编号',
  `sender` varchar(64) NOT NULL COMMENT '发送者',
  `receiver` varchar(64) NOT NULL COMMENT '接收者',
  `content` text NOT NULL COMMENT '消息内容',
  `send_time` int(11) NOT NULL COMMENT '发送时间',
  `is_del` int(11) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `is_read` int(11) NOT NULL DEFAULT '0' COMMENT '是否已读'
);

CREATE TABLE `contacts` (
  `username` varchar(32) NOT NULL COMMENT '本用户',
  `contact` varchar(32) NOT NULL COMMENT '联系人'
);

CREATE TABLE `users` (
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password` char(32) NOT NULL COMMENT '密码',
  `create_time` int(11) NOT NULL COMMENT '注册时间',
  `last_time` int(11) DEFAULT NULL COMMENT '上次登录时间'
);
```

###API design  

1. 登录
```
POST /v1/signin
```  

2. 注册  
```
POST /v1/signup
```


3. 获取用户信息  
```
GET /v1/[username]
```

4. 获取联系人信息
```
GET /v1/[username]/contacts
```

5. 新增联系人
```
GET /v1/[username]/contacts/[username]
```

5. 删除联系人
```
DEL /v1/[username]/contacts/[username]  
```


5. 获取与某一个联系人的聊天信息  
```
GET /v1/[username]/contacts/[username]/chats
```
