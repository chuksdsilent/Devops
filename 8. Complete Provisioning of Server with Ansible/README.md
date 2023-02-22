# Complete Provisioning of Server with Ansible

## Problems

1. Complete Infra setup is complex

2. Is not Repeatable

3. Difficult to Track

4. High chance of making mistake

5. Chances of Human Error

6. Managing infrastructure manually is time consuming

## Solution

1. Configuration Management of Infrastructure

2. Automatic Setup (No Human Error)

3. Centralize change Management

4. Version Control

5. Repeatable

6. Reusable

## Tools And Cloud

1. Ansible Automation

2. AWS Cloud Platform

## Steps

1. Setup VPC[Secure & HA]

2. Provision EC2 Instnaces, ELB, Sec Grp

3. Provision VPROFILE Stack on EC2 Instances
   - Build Artifact
   - MYSQL
   - Memcache
   - RabbitMQ
   - Tomcat
   - Nginx

## Flow of Execution

1. Login to AWS

2. Create ec2 instance ot run ansible playbook

3. Install Ansible

4. Install boto

5. Setup ec2 Role for ansible

6. Fetch source code from project ansible for AWS

7. Execute the VPC Playbook

8. Playbook to launch ec2, elb, sec grp for vprofile

9. get into vprofile vpc

10.
