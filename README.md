# Monkey

This repository begun with an exact replica of the final code for
Writing An Interpreter In Go (https://interpreterbook.com/).

It is gradually being modified to suit the needs of a filter language
for my project [trpc](https://github.com/shric/trpc). trpc is a frontend to the
[transmission](https://transmissionbt.com/) bittorrent client and will support 
arbitrary filter expressions such as `trpc list -f 'size < 10 MiB'` to list all
torrents smaller than 10 MiB.
