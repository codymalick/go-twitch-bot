#/usr/bin/bash
cat channels | xargs -P 10 -I @@ ./TwitchEmojis -c @@
