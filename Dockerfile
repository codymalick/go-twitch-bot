FROM golang
MAINTAINER Cody Malick <github.com/codymalick>

ADD . /go/src/github.com/codymalick/go-twitch-bot
WORKDIR /go/src/github.com/codymalick/go-twitch-bot/
RUN go get -d -v
RUN go install -v
# RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 0C49F3730359A14518585931BC711F9BA15703C6
# RUN echo "deb [ arch=amd64,arm64 ] http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.4 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-3.4.list
# RUN apt-get update
# RUN apt-get install -y mongodb-org
# RUN mkdir -p /data/db
# RUN service mongod start

# Environment Variables
# ARG DB
# ARG USER
# ARG CHAN
# ARG AUTH

# Build the application
#RUN go build

# Mongo port
# EXPOSE 27017

# Start the application attached to a channel (NALCS1)
ENTRYPOINT go-twitch-bot -auth oauth:mkvyz3e7o203nmk0si2cwal8ft6e6e -c twitchpresents -db go-twitch-bot -user cmallocbot
