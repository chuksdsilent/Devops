#!/bin/bash

### Script that displays system info

echo "System Information"

echo 

# Checking system uptime
echo "###########################"
echo "System Uptime"
uptime

echo 
# Memory Utilization
echo "###########################"
echo "Memory Utilization"
free -m
echo 

echo "##########################"
echo "Disk Utiliztion"
df -h


