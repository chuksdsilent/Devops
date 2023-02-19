# Ansible for AWS VPC

## Problem

1. VPC Consist of Many Moving parts

2. Subnets, Nat, Int GW, Route Tables, NACL, Sec Grp

3. Bastion Host

4. Human Error can lead to non functional or exposed VPC

5. Managing manually is time consuming

## Solution

1. Configuration Management of VPC

2. Automatic Setup (NO Human Errors)

3. Centralize Change Managment

4. Version Control[IAAC]

## Tools And Cloud

1. Ansible

2. Aws vpc setup with bastion host

## Flow of Execution

1. Login to AWS

2. Create ec2 instance to run ansible playbook

3. Install Ansible

4. Install boto

5. Setup ec2 Role for ansible

6. Create a project directory

7. Sample cloud task (with key pair)

8. Create variables File for VPC & Bastion host

9. Create VPC setup playbook

10. site.yml playbook to call both playbook at once
