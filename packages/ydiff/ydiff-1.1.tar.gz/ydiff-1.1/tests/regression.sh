#!/bin/bash

TOP_DIR=$(cd $(dirname $0)/.. && pwd) || exit 1
cd $TOP_DIR || exit 1

YDIFF=./ydiff

# To test with py3k: PYTHON=python3 make test
PYTHON=${PYTHON:-python}

function pass()
{
    if [[ -t 1 ]]; then
        echo -e "\x1b[032mPASS\x1b[0m" "$*"
    else
        echo "PASS" "$*"
    fi
}

function fail()
{
    if [[ -t 1 ]]; then
        echo -e "\x1b[01;31mFAIL\x1b[0m" "$*"
    else
        echo "FAIL" "$*"
    fi
}

function cmp_output()
{
    local input=${1:?}
    local expected_out=${2:?}
    local ydiff_opt=${3:-""}
    local cmd

    cmd=$(printf "%-7s $YDIFF %-24s < %-30s " $PYTHON "$ydiff_opt" "$input")
    printf "$cmd"
    if eval $cmd 2>/dev/null | cmp --silent $expected_out -; then
        pass
        return 0
    else
        fail "!= $expected_out"
        return 1
    fi
}

function main()
{
    local total=0
    local e=0
    local d

    for d in tests/*/; do
        d=${d%/}
        [[ -f $d/in.diff ]] || continue
        cmp_output $d/in.diff $d/out.normal "-c always" || ((e++))
        cmp_output $d/in.diff $d/out.side-by-side "-c always -s" || ((e++))
        cmp_output $d/in.diff $d/out.w70 "-c always -s -w70" || ((e++))
        cmp_output $d/in.diff $d/out.w70.wrap "-c always -s -w70 --wrap" || ((e++))
        cmp_output $d/in.diff $d/in.diff "-c auto" || ((e++))
        cmp_output $d/in.diff $d/in.diff "-c auto -s" || ((e++))
        cmp_output $d/in.diff $d/in.diff "-c auto -s -w70" || ((e++))
        cmp_output $d/in.diff $d/in.diff "-c auto -s -w70 --wrap" || ((e++))
        (( total += 6 ))
    done

    if (( e > 0 )); then
        echo "*** $e out of $total tests failed." >&2
        return 1
    else
        echo "All $total tests passed."
        return 0
    fi
}

main "$@"

# vim:set et sts=4 sw=4:
