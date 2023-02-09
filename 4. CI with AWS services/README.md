# AWS CI/CD

This project creates a CI/CD pipeline using AWS Cloud Native services

## BENEFITS

- Build & Test for every commit
- Automated Process
- Notify For Every Build Status

## AWS SERVICES

- Code Commit
- Code Artifact
- Code Build
- Code Deploy
- SonarCloud
- Checkstyle
- Code Pipeline

## FLOW OF EXECUTION

- Login to AWS account
- Code Commit
  - Create codecommit repo
  - create iam user with codecommit policy
  - generate ssh keys locally
  - Exhange keys with IAM user
  - Put source code from github repo to cc repository and push
- Code Artifact
  - Create an IAM user with code artifact access
  - install AWS cli and configure
  - Export auth token
  - Update settins.xml file in source code top level directory with below details
- Sonar Cloud
  - Create Sonar Cloud Account
  - Generate Token
  - Create SSM parameters with sonar details
  - create build project
  - update code build role to access SSM paramater store
- Create notifications for sns or slack
- Build Project
  - Update pom.xml with artifact version with timestamp
  - create variables in SSM => parametersore
  - Create build project
  - update codebuild role to access SSM parameterstore
- Create a Pipeline
  - codecommit
    -testcode build
    deploy to s3bucket
- Test Pipeline

### STEPS

5. check 049 video
6. Setup Notification
   - Go to SNS >> Topics
   - Create topic
   - type the name
   - create topic
     - Create subscription
     - type Email
     - create subscription
     - Go to email >> confirm subscription
7. Go to pipeline
   - create pipeline
   - type the name
   - next
   - source provider(codecommit)
   - select our repository and branch(ci-aws)
   - select amazon cloudwatch events
   - next
   - Build provider(codebuild)
   - select the project
   - Next
   - deploy(aws s3)
   - Next
   - Review
   - create pipeline
   - stop execution
   - you can click on edit to add other builds
     - For Test
       - Add Stage
         - Name
       - Add action group
         - name
         - action provider(aws codebuild)
         - project name
         - input artifact(source artifact)
         - single build
         - done
     - For Deploy
       - Add Stage
         - Name
       - Add action group
         - name
         - action provider(s3 bucket)
         - create s3 bucket
         - create a folder inside the s3 bucket
         - source artifact (build Artifact)
         - select the bucket name
         - enter the folder name
         - select extract file before deploy
         - done
8. Set the notifications
   - Codepipeline >> pipeline settings >> Notification >>create Notification rule
   - name
   - select the action
   - select target type
   - choose topic
   - submit
