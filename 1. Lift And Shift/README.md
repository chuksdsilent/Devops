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
