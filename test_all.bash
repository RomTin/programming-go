#!/usr/bin/env bash

function exec () {
  cwd=`pwd`
  cd $1
  # execute test.bash if exists
  if [ -f 'test.bash' ]; then
    bash test.bash
  else
    go test -v
  fi
  cd $cwd
}

dirs=`ls | egrep 'ch[0-9]{2}'`
workdir=`pwd`

for dir in $dirs; do
  cd $dir
  exs=`ls | egrep 'ex[0-9]{2}'`
  for ex in $exs; do
    exec $ex
  done
  cd $workdir
done
