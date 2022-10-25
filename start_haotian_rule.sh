#!/bin/bash
BUILD_ID=dontKillMe
chmod +x hao_tian_gin_rule

sh stop_haotian_rule.sh
sleep 1
./hao_tian_gin_rule  > rule-api.log 2>&1 &
sleep 1

pid_count=`ps -aux | grep hao_tian_gin_rule | grep api | grep -v grep   | wc -l `
if [ 0 == $pid_count ];then
  echo "hao_tian_gin_rule api pid no exit"
  exit 1
fi
