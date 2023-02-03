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
