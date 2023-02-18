# Continuous Delivery

1. Go to jenkins jobs >> change the branch to cd-jenkins

## Creating security group

1. Login to aws

2. Go to ec2

3. Go to security group

4. create SG (port: 22, 8080 from my ip, port 8080 from jenkins SG)

5. create SG for windows server test (All traffic from jenkins SG)

6. create SG for Jenkins test (All traffic from windows SG)

7. Go to vprofile-project change switch to cd-jenkins branch

8. change to user data

9. copy the tomcat provision.sh code, backend-stack.sh, windows-node

## Change quality gate to sonar default quality gate

1. Go to sonar

2. Go the project >> project settings >> project >> quality gate >> change the quality gave value.
