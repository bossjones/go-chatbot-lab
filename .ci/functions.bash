#!/usr/bin/env bash

function canonicalPath
{
    local path="$1" ; shift
    if [ -d "$path" ]
    then
        echo "$(cd "$path" ; pwd)"
    else
        local b=$(basename "$path")
        local p=$(dirname "$path")
        echo "$(cd "$p" ; pwd)/$b"
    fi
}

# SOURCE: https://stackoverflow.com/questions/2564634/convert-absolute-path-into-relative-path-given-a-current-directory-using-bash
# INFO: python -c "import os.path; print os.path.relpath('/foo/bar', '/foo/baz/foo')"
function relpath() {
    python -c "import os.path; print os.path.relpath('$1','${2:-$PWD}')" ;
}
