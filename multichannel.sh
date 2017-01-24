#/usr/bin/bash
cat channels | xargs -P 21 -I @@ ./go-twitch-bot -c @@
