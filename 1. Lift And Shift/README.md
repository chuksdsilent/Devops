# LIFT AND SHIFT PROJECT

This project moves application that runs on premise to cloud

## AWS SERVICES

- EC2 INSTANCES
- ELB [LOAD BALANCER]
- AUTOSCALING
- S3/EFS STORAGE
- ROUTE 53
- IAM

## OBJECTIVES

- Flexible
- No Upfront Cost
- Modernize Effectively
- IAAC

## ARCHITECTURE OF OUR PROJECT

## FLOW OF EXECUTION

- Login to AWS Account
- Create Key Pairs
- Launch Ec2 Instances
- Update IP to name mapping in route 53
- Build Application from source code
- Upload Artifact to s3 Bucket
- Setup ELB WITH https
- Map ELB Endpoint to website name in whogohost
- Autoscaling group for ec2 instances

6. Change Directory to vprofile project/userdata which contains a lot of bash scripts

7. Go to ec2 instance and launch instance
   - Provision mysql instance(Ubuntu)
     - Select t2 micro
     - Select the exist backend security group
     - Copy mysql userdata bash script and past in the userdata box
     - Launch Instance
     - Login to the instance and check if database is created
   - Provision memcache instance(Centos)
     - Select t2 micro
     - Select the exist backend security group
     - Copy memcache userdata bash script and past in the userdata box
     - Launch Instance
   - Provision rabbitmq instance(Centos)
     - Select t2 micro
     - Select the exist backend security group
     - Copy rabbitmq userdata bash script and past in the userdata box
     - Launch Instance
8. Note the backend service private IP address
9. Go to Route53
   - Create a hosted zone
     - Type in domain name(sam.org)
     - Select the Region
     - Select default vpc
     - Create hosted zone
   - Create Simple Records
     - Select simple routing
     - Type the sub domain (db01, mc01,bmq01) - Enter ip address of the corresponding ec2 instance
10. Copy the sub domains (db01.sam.org) and paste in the application.properties file
11. Create Tomcat EC2 Instance(Ubuntu)
    - Select t2 micro
    - Select the exist backend security group
    - Copy Tomcat userdata bash script and past in the userdata box
    - Launch Instance
12. Install JDK and Maven using Powershell
    - choco install jdk8
    - choco install maven
13. navicate to vprofile project/src/main/resources/
14. Open Application.properties file and update
    - db host as db01.sam.org
    - rabbitmq host as rmq01.sam.org
    - memcache host as mc01.sam.org
15. Run mvn install
16. Create an IAM user
    - Go to users
    - Add user
    - select programmatic access
    - Attach to existing policies
    - Select s3 bucket full access
    - Download the credentials
17. Upload the artifact to s3 bucket
    - install AWS Cli
      - choco install awscli
    - Configure AWS Cli
    - Create aws bucket
    - copy the artifact
18. Create a role to download the artifact
    - Go to IAM
    - create role
      - Select s3 bucket full access
      - Type in the role name
      - create role
      - Go to ec2 instance >> Action >> Setting >> Modify Access role >> Select role >> Modify
19. Login to Tomcat instance
    - sudo -i
    - cd /var/lib/tomcat8/webapp
    - stop tomcat (systemctl stop tomcat8)
    - rm -rf ROOT
    - install aws cli (apt install awscli -y)
20. Download the artifact from aws s3 bucket
    - aws s3 ls s3//<bucketname>
    - aws s3 cp s3://<bucket-name>/file-name.war /tmp/file-name.war
    - cp /tmp/file-name.war /var/lib/tomcat8/webapps/ROOT.war
    - systemctl start tomcat8
    - apt install telnet -y
    - telnet db01.sam.org 3306
21. Create a Target Group

    - Go to Target Group
    - Create a target group
    - type the name and port
    - health is /login
    - override the port to 8080
    - select the instance
    - include expending below
    - create target group

22. Setup Load Balancer

    - Go to load balancer
    - select application load balancer
      - Type the name
      - select internetfacing
      - select all the zones
      - select http and https
      - choose a certificate from acm
      - select security group from load balancer
      - select target config
      - create the load balancer
      - Copy the url of the load balancer and create a cname record on your domain name host

23. Setup Auto Scaling Group
    - Create AMI of the Tomcat Instance
      - Select the instance >> Action >> Image >> Create Image
      - Type the name of the image
      - Create Image
    - Go to launch configurateion
      - Create launch configuration
      - Type the name
      - Select the AMI
      - select the instance type(t2.micro)
      - select iam role
      - select monitoring with cloud watch
      - select security group of apps
      - select key pair
      - create launch configuration
    - Go to Autoscaling Group
      - Type the name
      - select the launch configuration
      - select all subnet
      - select the target group
      - select health check
      - Type in the min and max
      - Select track scaling
      - Create AutoScaling Group
