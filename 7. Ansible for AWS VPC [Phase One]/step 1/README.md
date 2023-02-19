# Introduction to AWS Cloud Playbooks

1. Login into the instance

2. create a file. test-aws.yml

```
- hosts: localhost
  connection: local
  gather_facts: False
  tasks:
    - name: sample ec2 key
      ec2_key:
        name: sample
        region: us-east-1
```

3. To Test the playbook

```
ansible-playbook test-aws.yml
```

4. The module depends on boto because we are python 3

```
sudo apt search boto
sudo apt install python3-boto -y
sudo apt install python3-boto3 -y
sudo apt install python3-botocore -y

```
