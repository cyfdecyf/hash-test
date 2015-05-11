#!/usr/local/bin/awk -f

# Must use GNU awk for asort

{ array[$3]++ }

END {
    n = asorti(array, b)
    for (i = 1; i <= n; i++) {
        print b[i], array[b[i]]
    }
}

