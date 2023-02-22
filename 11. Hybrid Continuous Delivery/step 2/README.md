# Creating IAM user & Parameter Store

1. Go to IAM user

2. Add user

3. Type the username >> select programmatic access

4. Attach existing policies directly

   - AWSCodeBuildAdminAccess
   - s3ullAccess
   - AWSElasticBeanstalkFullAccess

5. Next >> Add tags >> Create User

6. Copy the credentials

7. Go to system manager >> parameter store

8. Create a new parameter

9. Go to vprofile-project

10. switch to cd-aws-jenkins

11. open aws_files/softTest-BuildSpec.yml >> copey the content

12. use vprofile-stage-url as the parameter store project name

13. copy and paste beanstalk url as the value

14. create parameter

15. Go to build job >> configure >> change the brach cd-aws-jenkins

16. Add Build step >> Execute Shell

17. Type the following code

```
cat << EOT > src/main/resources/application.properties

copy the application properties file from vprofile-project and paste

change the mysql url, username and password
```

18. save

19. Go to integration test , code analysis and test job and change the url to cd-aws-jenkins
