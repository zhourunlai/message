# Base image is in https://registry.hub.docker.com/_/golang/
# Refer to https://blog.golang.org/docker for usage
FROM golang:1.3.3
MAINTAINER xiaorun xiaorun95@gmail.com

# ENV GOPATH /go

# Install beego & bee
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee

EXPOSE 8080

# Run shell script
ADD cmd.sh /cmd.sh
RUN chmod +x /cmd.sh
RUN /cmd.sh

# Usage
# docker build -t message
# docker run --name db -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=message -d mysql
# docker run --rm --link db:mysql -v "$(pwd)"/mysql:/home/mysql -it mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p beeblog < message.sql'
# docker run --rm --link db:mysql -v "$(pwd)":/go/src/ -w /go/src xiaorun/beego bee api message -conn="root:root@tcp(mysql:3306)/message"
# docker run --rm --link db:mysql -v "$(pwd)/message":/go/src/message -w /go/src/message -p 8080:8080 --name beego xiaorun/beego bee run -downdoc=true -gendoc=true
# docker run --name nginx --link beego:beego -v "$(pwd)"/nginx:/etc/nginx/conf.d/ -p 80:80 -d nginx
