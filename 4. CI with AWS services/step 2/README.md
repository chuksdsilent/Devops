## Setup Sonarcloud

1. Go to sonarcloud.io

2. login with git cred

## Generate a token

1. Profile picture >> my account >> security

2. type token name >> generate

3. click on the plus symbol(+) >> Analyse new project >> create project manually

4. type name

5. type key

6. select public >> setup

7. Go to AWS >> system manager service (ssm) >> parameter store

8. Parameter Store >> create parameter

9. Type Organization from sonar cloud in name box and name of organization in value box from sonar cloud >> Create parameter

10. Create another parameter use HOST from sonar cloud as name and use url in the value box from the sonar cloud

11. Create another parameter use project from sonar cloud as name and use project name in the value box from the sonar cloud

12. another parameter use secure_artifact_token as name and use token in the value box from the sonar cloud and select secure string
