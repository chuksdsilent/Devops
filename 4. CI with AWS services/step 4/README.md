# AWS CodeBuild for SonarQube Code Analysis

1. Go to CodeBuild

2. Create codeBuild Project

3. type the name

4. select repository

5. select branch (ci-aws)

6. operating system(ubuntu)

7. image(3.0)

8. runtime env (standard)

9. New Service role (type role name)

10. Under buildspec >> select Insert command

11. Copy the command

12. show view buildspecfile

```
cd vprofile-project

git checkout ci-aws

cd aws-files

cat aws-files/sonar_buildspec.yml

```

13. copy the content

14. Replace the buildspec in aws buildsped box

15. change the values with value in parameter groups

16. select cloud watch

17. Enter group name, stream name

18. Create build project

## Update the role

19. Go to IAM service

20. edit >> environment >> cancel

21. Go to IAM service

22 Create policy

23 choose policy (system manager) - check the video (No. 048 at 14:45)

24. Review policy

25. give it a name

26. Select all

27. create policy

28. go to the role and attach policy

start build
