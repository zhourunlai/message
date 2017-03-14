#Web Message Systerm

##Demand：
实现一个 Web 上的私信系统



##Function:

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



##Reference：

1. 「流利说」消息中心
2. 「微信」的个人消息
3. 「知乎」私信



##Todos  

Deadline: 2 days  
Add to Wunderlist  

![1](http://ww1.sinaimg.cn/large/9f47c048gy1fdlms5yaoqj21kw0zk7wh)



##Systerm Design  
1. ***前后端分离***  
Front-end: vue2 + vuex + vue router + webpack + babel  
Back-end: Beego  
Database: MySQL  

2. ***WebSocket*** 模式  

3. Nginx 静态服务器，并且 load balancing  

4. Docker 部署  

![2](http://ww1.sinaimg.cn/large/9f47c048gy1fdlnu5ebs0j21kw16oakv)



##Coding & Testing

![3](http://ww1.sinaimg.cn/large/9f47c048gy1fdm8bk33yxj21kw0zkamm)
![4](http://ww1.sinaimg.cn/large/9f47c048gy1fdm8dm5g2pj21kw0zkkfg)



##Database Design  
>服务重启后数据不丢失

###1.MySQL  

*chats*

|字段|类型|注释|
|----|----|----|
|id|int|消息编号|
|sender|varchar|发送者|
|receiver|varchar|接收者|
|content|text|消息内容|
|send_time|int|发送时间|
|is_del|int|是否删除|
|is_read|int|是否已读|

*usrs*

|字段|类型|注释|
|----|----|----|
|username|varchar|用户名|
|password|char|密码|
|create_time|int|注册时间|
|last_time|int|上次登录时间

*contacts*

|字段|类型|注释|
|----|----|----|
|username|varchar|本用户|
|contact|varchar|联系人|

###2.Redis

|Key|Value|
|----|----|
|暂无|暂无|



##API Design  
>RESTful

###1.登录
```
GET /v1/user/signin
```  

###2.注册  
```
POST /v1/user/signup
```

###3.登出
```
GET /v1/user/signout
```

###4.获取用户信息  
```
GET /v1/user/:username
```

###5.获取联系人信息
```
GET /v1/user/:username/contact
```

###6.新增联系人
```
GET /v1/user/:username/contacts/:username
```

###7.删除联系人
```
DEL /v1/user/:username/contacts/:username  
```

###8.获取聊天信息  
```
GET /v1/user/:username/contacts/:username/chats
```

###9.删除聊天信息
```
DEL /v1/user/:username/contacts/:username/chats/:id
```

###10.已读回执
```
GET /v1/user/:username/contacts/:username/chats/:id
```
