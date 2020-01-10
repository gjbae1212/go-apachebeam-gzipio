#!/bin/bash

set -e -o pipefail
trap '[ "$?" -eq 0 ] || echo "Error Line:<$LINENO> Error Function:<${FUNCNAME}>"' EXIT
cd `dirname $0`
CURRENT=`pwd`

function test
{
   rm $CURRENT/test/output.read.* || true
   rm $CURRENT/test/output.write.* || true
   go test -v github.com/gjbae1212/go-apachebeam-gzipio
}

CMD=$1
shift
$CMD $*
