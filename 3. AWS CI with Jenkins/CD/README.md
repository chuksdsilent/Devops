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

1. Login to AWS
2. Go to CodeCommit
   - CodeCommit >> Repository >> Create Repository
   - Type the Repository Name
   - Create/Save
   - You can access the repo with https or ssh but ssh is preferrable
3. Go to IAM user

   - IAM user >> Add User
   - Type the username
   - select programmatic access
   - select attach policies
   - select create policy
     - choose the service(CodeCommit)
     - Select all action
     - select resources
     - Add ARN
     - select the region
     - Type the repo name
     - Add
     - Review Policy
     - Type the name
     - create policy
   - search for the policy name
   - next >> Create User
   - Go to the user created >> Security credentials
     - Delete the keys
     - Upload ssh public key
     - open gitbash
       - create ssh key (ssh-keygen)
       - cat the public key
     - Go to IAM >> user created >> security
       - upload public key
       - paste in the public key copied
       - upload
       - copy ssh key id
   - Go to gitbash
     - Go to /.ssh
     - create ssh_config (vim config)
     - paste the following code
       Host git-codecommit.\*.amazonaws.com
       User ssh key id from aws
       IdentityFile `/.ssh/file-name
     - Change mode to 600(chmod 600)
     - To test (ssh -v git-codecommit.region.amazonaws.com)
   - Clone the aws respository
     - Go To codecommit >> respository >> code
     - copy the url
     - git clone url
   - Clone the the vprofile project
     - switch to vprofile project
     - To check the repository (cat .git/config)
     - switch to master (git checkout master)
     - To view all branches(git branch -a | grep -v READ | cut d '/' f3 | grep -v master > /tmp/branches.txt)
     - To checkout all the branches (for i in cat /tmp/branches.txt'; do git checkout $i; done)
     - to remove current remote origin (git remote rm origin)
     - To add aws remote origin (git remote add aws url)
     - git push origin --all
     - git push --tags
     - To check (Go to codecommit >> repository >> code refresh)

4. Create a repository on codeartifact
   - Create Repository
   - Type the name
   - select maven central store
   - next
   - type the domain name (Any Name)
   - next
   - Review
   - Create respository
   - select the first repository
     - select view connection
     - select mvn
     - copy first link
     - create an iam user with programmatic access and select codeartifact admin access policy
     - install and configure awscli(choco install awscli -y)
     - run the link you coppied on gitbash
     - change settings.xml file
     - switch to ci-aws
     - open settings.xml
     - change the url, domain name. ensure to add / after the url
     - Update pom.xml
     - change the url
     - git add and git commit
     - git push origin ci-aws
   - Setup sonarcloud
   - Go to sonarcloud.io
   - login with git cred
   - Generate a token
     - Profile picture >> my account >> security
     - type token name
     - generate
     - click on the plus symbol(+) >> Analyse new project >> create project manually
     - type name
     - type key
     - select public
     - click setup
   - Go to AWS >> system manager service>> parameter store
     - create parameter
     - create parameter for organization, url, name of project and sonar token from sonar cloud
     - use secure string for token
     - save to each one
     - go to gitbash and type the following command
     - echo $CODEARTIFACT AUTH TOKEN
     - copy the token and create parameter using secure string
   - To create Code Build
   - Go to CodeBuild Project
   - Create codeBuild
   - type the name
   - select repository
   - select branch
   - operating system(ubuntu)
   - image(3.0)
   - runtime env (standard)
   - next
   - Buildspec file (switch to insert build commands)
   - switch to editor
     - Go to gitbash
     - go to branch ci-aws
     - cd aws-file/sonar_buildspec.yml
     - cat buildspec
     - change the values with value in parameter groups
     - group name
     - stream name
     - select cloud watch
     - select build job
   - Update the role
     - Go to IAM service
     - edit >> environment
     - cancel
     - go to IAM service
     - Create policy
     - choose policy (system manager) - check the video (No. 048 at 14:45)
     - Review policy
     - give it a name
     - create policy
     - go to the role and attach policy
   - start build
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
