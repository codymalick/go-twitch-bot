FROM golang:1.6
MAINTAINER Cody Malick <github.com/codymalick>

ADD . /go/src/github.com/codymalick/TwitchEmoji
ADD . /go/pkg
RUN go get github.com/thoj/go-ircevent
RUN go get github.com/gopkg.in/mgo.v2
RUN go get github.com/gopkg.in/mgo.v2/bson
RUN go install github.com/codymalick/TwitchEmoji
ENTRYPOINT /go/bin/TwitchEmoji -c amazhs
