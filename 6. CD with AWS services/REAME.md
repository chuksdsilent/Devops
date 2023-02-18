# Continuous Integration with AWS Services

## Problems

- Manual code deployment is time consuming
- Developer and Testers not equipped with OPS Knowledge
- Need ot hire ops professionals
- capital intensive
- operational overhead to maintain server like jenkins, nexus, sonar, git, QA

## Solution

- PASS && SAAS
- Disposable Environment
- Automate CI/CD Process
- Build Test, Deploy & Test for every commit.

## Flow of Execution

1. login to AWS account

2. Code Commit

   - Create a codecommit repo
   - Sync it with local repository

3. Code Artifact

   - Create repository
   - Update Settings.xml file in source code top level directory
   - Update pom.xml file with repo details
   - Generate token and store in SSM parameter store

4. Sonar Setup

   - Create sonar cloud account
   - Generate token and store in ssm parameter store
   - Create build project
   - Update codebuild role to access SSM parameter store

5. Create Notifications for sns or slack

6. Build Project

   - Create variables in SSM => Parameter store
   - create build project

7. Create Pipeline

   - Codecommit
   - Test code
   - Build

8. Create Beanstalk & RDS

9. Update the RDS SG

10. Deploy DB in rds

11. Switch to cd-aws branch

12. Update settings.xml & pom.xml

13. Create another build job to create artifact with buildspec file in cd-aws

14. Create a deploy job to beanstalk

15. Create a build Job for software testing

16. Upload screenshot to s3 bucket.

17. Update Pipeline
    - Codecommit
    - Test code
    - Build & Store
    - Deploy to s3 bucket
    - Build & Release
    - Deploy to Beanstalk
    - Build Job for selenium test scripts
    - Upload result to s3
