# AWS CodeBuild for Build Artifact

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

cat aws-files/build_buildspec.yml

```

13. copy the content

14. Replace the buildspec in aws buildsped box

15. change the values with value in parameter groups and use CODEARTIFACT_TOKEN created in the parameter store

16. select cloud watch

17. Enter group name(use the same one used in step 4), stream name

18. Create build project

## Update the role

19. Go to IAM service

20. edit >> environment >> cancel

21. Go to IAM service

22 go to the role and attach policy (use the policy created in step 4)

23. start build
