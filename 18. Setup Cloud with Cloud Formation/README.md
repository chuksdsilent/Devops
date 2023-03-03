# Setup AWS Cloud with Cloud Formation

## Problems

1. High cost of maintaining server

2. Backup through AMI? Storage cost

3. Manually setup is time consuming

4. Chances of Human Error

## Solution

1. Use of CloudFormation to automate AWS stacks

2. Automatic error(No Human Errors)

3. Maintain state of Infrastructure

4. Repeatable

5. Reusable

## Technologies

1. Cloud Formation

2. AWS Cloud Platform

## Flow of Execution

1. Check cicd dat aon s3 bucket

2. Note down bucket name and filenames

3. Create s3 bucket to upload templates, create folder and named

   - stack template

4. Note down bucket name and folder name

5. create key pair

6. Write root template named cicdtemplate.yml

7. Write all child template

   - cicds3role.yaml
   - jenk.yaml
   - nexus.yaml
   - sonar.yaml
   - wintest.yaml
   - app01qa.yaml
   - db0qa.yaml

8. Update Root template with child template paths

9. Upload all child template to s3 bucket in folder

   - stack-template

10. Create Nested stack by using cicdtemp.yaml
