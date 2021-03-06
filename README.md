# Web Message Systerm

## 目录
* [Demand](#demand)
* [Function](#function)
* [Reference](#reference)
* [Todos](#todos)
* [Systerm Design](#systerm-design)
* [Database Design](#database-design)
* [API Design](#api-design)
* [Coding&Testing](#codingtesting)
* [Notice](#notice)


---



## Demand
实现一个 Web 上的私信系统



## Function

- [x] * 用户可以注册、登录。需要 id（可以自己决定 email 或者 username）和 password  
- [x] * 用户登录后，进入联系人列表页面  
- [x] - 可以看到自己所有的联系人  
- [x] - 每个联系人需要显示对方 id 以及未读私信数量提醒  
- [x] - 用户可以通过 id 添加新联系人（可以不需要对方同意）  
- [x] - 用户可以删除某个联系人，但保留与对方用户的消息等数据。当再次添加新联系人时，消息等数据都还在  
- [x] * 点击一个联系人会进入聊天界面，同时未读消息置为 0  
- [x] - 可以看到和某个用户的历史消息  
- [x] - 能够在这里收发私信（不需要实时，可以刷一下页面才看到新消息）  
- [ ] - 当用户 A 发私信给用户 B 时，如果 A 还不是 B 的联系人，应该自动把 A 添加为 B 的联系人，并能够在 B 的联系人列表正常显示（不需要实时）  
- [x] - 用户可以删除自己发的消息  

加分项：  

- [ ] * 联系人列表页面未读消息数实时更新  
- [ ] * 聊天界面新消息实时接收  
- [ ] * 自动把 A 添加为 B 联系人时，B 实时更新联系人列表  
- [ ] * 部署，可在线演示  



## Reference

1. 「流利说」消息中心
2. 「微信」的个人消息
3. 「知乎」私信
4. *「Telegram」*

![0](http://ww1.sinaimg.cn/large/9f47c048gy1fdmteeuk13j21kw0zkb29)



## Todos

Deadline: 2 days  
Add to Wunderlist  

![1](http://ww1.sinaimg.cn/large/9f47c048gy1fdlms5yaoqj21kw0zk7wh)



## Systerm Design
1. ***前后端分离***  
Front-end: vue2 + vuex + webpack + babel  
Back-end: Beego + WebSocket  
Database: MySQL  

2. ***WebSocket*** 模式  
发布订阅

3. Nginx 静态服务器，并且 load balancing  

4. Docker 部署  

学习：  
[beego 的 MVC 架构介绍](https://beego.me/docs/mvc/)    
[关于轮询、长轮询、长连接、WebSocket等推送方案](https://github.com/Hipponensis/Notes/blob/master/%E5%85%B3%E4%BA%8E%E8%BD%AE%E8%AF%A2%E3%80%81%E9%95%BF%E8%BD%AE%E8%AF%A2%E3%80%81%E9%95%BF%E8%BF%9E%E6%8E%A5%E3%80%81WebSocket%E7%AD%89%E6%8E%A8%E9%80%81%E6%96%B9%E6%A1%88.md)  


![2](http://ww1.sinaimg.cn/large/9f47c048gy1fdlnu5ebs0j21kw16oakv)



## Database Design
>服务重启后数据不丢失

### 1.MySQL  

*usrs* 表  

|字段|类型|注释|
|----|----|----|
|username|varchar|用户名|
|password|char|密码|
|create_time|int|注册时间|
|last_time|int|上次登录时间

*contacts* 表  

|字段|类型|注释|
|----|----|----|
|id|int|联系人关系编号|
|username|varchar|本用户|
|contact|varchar|联系人|

*chats* 表  

|字段|类型|注释|
|----|----|----|
|id|int|消息编号|
|sender|varchar|发送者|
|receiver|varchar|接收者|
|content|text|消息内容|
|send_time|int|发送时间|
|is_del|int|是否删除|
|is_read|int|是否已读|

![3](http://ww1.sinaimg.cn/large/9f47c048gy1fdmgb5gdepj21kw0zk7pr)

*users* => *contacts* 一对多  
*contacts* => *chats* 一对多  


### 2.Redis

|Key|Value|
|----|----|
|暂无|暂无|




## API Design
>RESTful

### 1.登录
```
GET /v1/user/signin
```  

### 2.注册  
```
POST /v1/user/signup
```

### 3.登出
```
GET /v1/user/signout
```

### 4.搜索用户信息（检查用户名是否已经注册、添加联系人）  
```
GET /v1/user/:username
```

### 5.获取联系人信息（显示联系人列表）
```
GET /v1/user/:username/contacts
```

### 6.新增联系人
```
GET /v1/user/:username/contacts/:contact_username
```

### 7.删除联系人
```
DEL /v1/user/:username/contacts/:contact_username  
```

### 8.获取聊天信息（显示历史消息）
```
GET /v1/user/:username/contacts/:contact_username/chats
```

### 9.删除聊天信息
```
DEL /v1/user/:username/contacts/:contact_username/chats/:id
```

### 10.已读回执（未读置0）
```
GET /v1/user/:username/contacts/:contact_username/chats/:id
```

### 11.未读统计（各个联系人的未读消息）
```
GET /v1/user/:username/contacts/:contact_username/chats/unread
```



## Coding&Testing

├── Dockerfile  
├── README.md  
├── cmd.sh          ------------       Docker 启动时执行的命令  
├── conf        
│   └── app.conf  
├── controllers   
│   ├── main.go      ------------      渲染前台  
│   ├── user.go      ------------      API 接口  
│   └── websocket.go ------------      WS 接口  
├── docs  
│   ├── doc.go  
│   └── docs.go  
├── lastupdate.tmp  
├── main.go  
├── message  
├── message.sql  
├── models  
│   ├── user.go      ------------      封装 API 数据包  
│   └── websocket.go ------------      封装 WS 数据包  
├── routers  
│   ├── commentsRouter_message_controllers.go  
│   └── router.go  
├── swagger  
│   └── swagger-1  
├── tests  
│   └── default_test.go  
├── views  
│   ├── dist          ------------     打包生成的 js  
│   ├── index.html    ------------     视图  
│   ├── node_modules  
│   ├── package.json  
│   ├── src  
│   └── webpack.config.js  

```
cd $GOPATH/src/
//安装beego
go get github.com/astaxie/beego
//安装bee工具
go get github.com/beego/bee
git clone https://github.com/zhourunlai/message.git & cd message/
//启动
bee run -downdoc=true -gendoc=true
```
![3](http://ww1.sinaimg.cn/large/9f47c048gy1fdojxyd41aj21kw0zknb4)
![4](http://ww1.sinaimg.cn/large/9f47c048gy1fdm8bk33yxj21kw0zkamm)

```
cd views/
//安装依赖
cnpm install
//测试
cnpm run dev
//打包成js
cnpm run build
```
![5](http://ww1.sinaimg.cn/large/9f47c048gy1fdmuudctk7j21kw0zkdsm)


## Notice
1. 登陆后发送请求包括：获取本用户 info ，获取联系人列表，获取未读消息
2. ws 传输的数据包括：私信，session 正在输入，是否在线
3. 历史消息懒加载
