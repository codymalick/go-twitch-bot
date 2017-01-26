#/usr/bin/bash
cat channels | xargs -P 22 -I @@ ./go-twitch-bot -c @@
