cd $GOPATH/src/
//安装beego
go get github.com/astaxie/beego
//安装bee工具
go get github.com/beego/bee
//下载项目
git clone https://github.com/zhourunlai/message.git & cd message/
//启动
nohup bee run -downdoc=true -gendoc=true &
