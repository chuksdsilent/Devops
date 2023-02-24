# Kops for Cluster on AWS

1. You need a valid domain for k8 DNS records

2. Create a linux VM and setup

   - kops, kubectl, ssh keys, awscli

3. lLogin to AWS account and setup

   - s3 bucket, IAM User for AWSCli, Route53 Hosted Zone

4. Create an EC2 instance with

   - t2.micro
   - SG(port 22 from my ip)
   - name (kops)
   -

5. Create s3 bucket which will be used to store the state of kops

6. Go to IAM user

7. Type the username >> Next

8. Attach existing policies directly

9. Select for AdministratorAccess >> next

10. Add tags >> Create user

11. Go to Route53

12. Create a hosted zone

    - name(vprofile.groovy.in)
    - public hosted zone
    - create

13. Add the values to whogohost
