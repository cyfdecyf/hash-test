#!/bin/bash

if [ $# != 1 ]; then
    echo "Usage: `basename $0` <input>"
    exit 1
fi

input=$1

mkdir -p result

test() {
    local hash=$1
    local modulo=$2
    ./hash-test $modulo $hash < $input | ./stat.awk > result/$hash-$modulo
}

for i in 3 5 10 20; do
    test crc32 $i
    test djb2 $i
done

