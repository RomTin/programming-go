#!/usr/bin/env bash

function exec () {
  cwd=`pwd`
  cd $1
  # execute test.bash if exists
  if [ -f 'test.bash' ]; then
    echo 'test.bash'
    bash test.bash
  else
    echo 'command'
    go test -v
  fi
  cd $pwd
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
