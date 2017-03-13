FROM golang
MAINTAINER Cody Malick <github.com/codymalick>

# Copy the files over
ADD . /go/src/github.com/codymalick/go-twitch-bot
ADD ./config /go/src/github.com/codymalick/go-twitch-bot/config

# Install dependencies
RUN go get github.com/thoj/go-ircevent
RUN go get gopkg.in/mgo.v2

# Mongodb steps
# RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 0C49F3730359A14518585931BC711F9BA15703C6
# RUN echo "deb [ arch=amd64,arm64 ] http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.4 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-3.4.list
# RUN apt-get update
# RUN apt-get install -y mongodb-org
#
# RUN mkdir -p /data/db
#
# # Start mongo
# RUN service mongod start

# Build the application
RUN go build github.com/codymalick/go-twitch-bot

# Mongo port
# EXPOSE 27017

# Start the application attached to a channel (NALCS1)
ENTRYPOINT /go/src/github.com/codymalick/go-twitch-bot -c NALCS1
