#!/bin/bash

USR='devops'
for host in `cat remHosts`
do
  echo "#############################"
  echo "connecting to $host"
  scp websetup_with_variable.sh $USR@$host:/tmp
  ssh $USR@$host /tmp/websetup_with_variable.sh
  ssh $USR@$host  rm /tmp/websetup_with_variable.sh
done
