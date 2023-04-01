#!/bin/bash


echo "###############################"
date
sudo cat /var/run/httpd/httpd.pid &> /dev/null

if [ $? -eq 0 ]
then
  echo "httpd is running"
else
  echo "httpd process is not running"
  echo "Starting httpd process"
  sudo systemctl start httpd
  if [ $? -eq 0 ]
  then
     echo "Process started successfully.."
  else
     echo "Proccess not started. Contact adming"
  fi
fi
