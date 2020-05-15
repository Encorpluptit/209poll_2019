#!/usr/bin/env bash

PROF_FILE="profiling.prof"
PROF_DIR="test/profiling"
#FILES_DIR="test/ex"

function profile() {
    NAME="$1-$2-$4"
    echo "Running --> ${NAME}"
    stack run --profile "$1" "$2" "$3" -- +RTS -p &> /dev/null
    
}

#for k in 1 2 4 8 16; do
#    for c in 0.1 0.5 0.8 1.5; do
#        for f in "$FILES_DIR/medium" "$FILES_DIR/big"; do
#            echo "$k $c $f"
#        done
#    done
#done
