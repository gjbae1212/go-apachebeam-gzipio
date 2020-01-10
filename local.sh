#!/bin/bash

set -e -o pipefail
trap '[ "$?" -eq 0 ] || echo "Error Line:<$LINENO> Error Function:<${FUNCNAME}>"' EXIT
cd `dirname $0`
CURRENT=`pwd`

function test
{
   rm $CURRENT/test/output.read.* || true
   rm $CURRENT/test/output.write.* || true
   go test -v $(go list ./... | grep -v vendor) --count 1 -race -covermode=atomic
}

CMD=$1
shift
$CMD $*
