#/usr/bin/bash
cat channels | xargs -P 20 -I @@ ./TwitchEmojis -c @@
