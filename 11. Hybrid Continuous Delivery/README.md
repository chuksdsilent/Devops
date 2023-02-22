# Hybrid Continuous Delivery

This project will use AWS services like SAAS and PAAS

## Tools

1. Github

2. Jenkins

3. Code Build

4. Sonar Cloud

## AWS Services

1. Beanstalk

2. RDS

## Flow of Execution

1. Continuous Integration Setup

2. RDS SG

3. Beanstalk & RDS Setup

4. RDS Initialization and SG update

5. ElasticCache & Active MQ (Optional)

6. Sonar Cloud Token

7. Sonar Cloud Integration with Jenkins

   - Sonar Cloud Token
   - Sonar organization in properties

8. Sonar Cloud Scanning Job

9. IAM User to Jenkins authorization

   - Codebuild policy
   - s3 bucket policy

10. Update Bean URL in Parameter Store for CodeBuild job

11. Deploy to Beanstalk Job[Stagging]

12. Software Testing Job with AWS Code Build

    - Plugin
    - Job Setup

13. Changes in CodeBuild

    - cd-aws-jenkins
      = update Buildspec file
    - Store URL variable in parmeter store

14. Connect all jobs in jenkins pipeline

15. Test

16. Prod deploy job(Options.)
