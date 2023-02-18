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
