#!/bin/bash

# Variable Declaration
PACKAGE="httpd wget unzip"
SVC="httpd"
URL="https://www.tooplate.com/zip-templates/2131_wedding_lite.zip"
ART_NAME="2131_wedding_lite"
TEMPDIR="/tmp/webfiles"

echo "Installing package. please wait..."
sudo yum install $PACKAGE  -y > /dev/null

echo "Starting and enabling httpd. Please wait..."
sudo systemctl  start $SVC
sudo systemctl enable $SVC

mkdir -p $TEMPDIR

cd $TEMPDIR

wget $URL
unzip $ART_NAME.zip
sudo cp -r $ART_NAME/* /var/www/html

sudo systemctl restart $SVC

sudo rm -rf $TEMPDIR > /dev/null

sudo systemctl status $SVC
ls 2131_wedding_lite


