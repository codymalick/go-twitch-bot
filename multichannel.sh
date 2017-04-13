#/usr/bin/bash

AUTH=$1
USER=$2

cat channels | xargs -P 22 -I @@ ./go-twitch-bot -c @@ -auth $AUTH -user $USER
