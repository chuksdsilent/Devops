# Launching EC2 Instances

1. Go to vprofile project

2. change to ci-jenkins. You will find the userdata directory.

3. open it and find the script to create jenkins, nexus and sonarqube

4. use the script to create this application on ec2 instances

5. Login to aws

6. Go to launch instances >> create instances
   - Jenkins(Ubuntu 18, t2 small, select already created sg and key pair)
   - Nexus (centos7, t2 small, select already created sg and key pair)
   - Sonarqube(Ubuntu 18, t2 small, select already created sg and key pair)
