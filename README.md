# Hash function test

Run [`test.sh`](test.sh) to see the distributions for hash and then modulo results.

[`extract-sites-from-cow-stat.rb`](extract-sites-from-cow-stat.rb) is for
getting site list from cow stat file.

Currently, only compare djb2 and crc32. crc32 distribute hash results much more
evenly.

