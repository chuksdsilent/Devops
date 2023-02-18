# BeanStalk Setup

1. login to AWS

2. Go to beanstalk >> Create application

3. Type the app name >> select platform (Tomcat) >> select sample application

4. Click edit capacity >> Env. Type(load balanced) >> min(2), max(4) >> more option

5. Click edit keypair >> select the keypair >> save

6. Edit tags

7. create app

## Creating RDS

1. Go to RDS

2. Create database >> mysql >> version(5.\*) >> free tier >> DB instance identifier(vprofile-cicd-project) >> autogenerate password >> security group (create new) >> initial database(accounts) >> create database >> copy out the password

3. Go to ec2

4. select elastic beanstalk server

5. Go to security group and copy the SGID

6. Go RDS instance >> Security Group >> Allow traffic from elastic beanstalK ID.

7. Copy elastic beanstalk ip and ssh to it

8. change to root

```
sudo -i
```

9. install mysql client and git

```
yum install mysql git -y
```

10. login to mysql

```
mysql -h database-url -u admin -pPassword
```

11. To databases

```
show databases;
```

12. Clone vprofile project

```
git clone vprofile-project-url
```

13. checkout the cd-aws

```
git checkout cd-aws
cd src/main/resources
```

14. Import Database

```
mysql -h database-url -u admin -pPassword < db_backup.sql
use accounts;
show tables;
```

15. After beanstalk provisioning >> go to load balancer >>under processes select default >> under path (/login) >> Apply

16. Go to CodeCommit Env. >> Select ci-aws branch >> open pom.xml >> copy the repository at the bottom

17. Go to CodeCommit Env. >> Select cd-aws branch >> open pom.xml >> edit >> replace the repository

18. Enter Name and Email Address >> Save

19. Go to CodeCommit Env. >> Select ci-aws branch >> open settings.xml >> copy the whole file

20. Go to CodeCommit Env. >> Select cd-aws branch >> open settings.xml >> edit >> replace the whole file.

21. Go to build >> build project >> open vprofile-build >> Edit >> source >> change the branch from ci-aws to cd-aws >> update

22. Go to build >> build project >> open vprofile-build-artifact >> Edit >> source >> change the branch from ci-aws to cd-aws >> update

23. Create build project
    - project-name(vprofile-12-build-release)
    - Repository(vprofile-code-repo)
    - branch(cd-aws)
    - operating system(ubuntu)
    - Runtime(standard)
    - image(3.0)
    - service role(existing >> Go to build projects >> open vprofile build >> edit >> copy the role and paste)
    - Build Specification(Insert build commands(Go to cd-aws branch >> aws files >> open buildAndRelease_buildspec.yml the copy and replace))
    - CloudWatch group name(vprofile-cicd-logs)
    - stream name (Build&Release Job)

### Generate latest token

24. Go to repositories

25. Maven Central store

26. View connection instructions

27. Copy the url and paste on gitbash then enter

28. Type the following command

```
echo $CODEARTIFACT_AUTH_TOKEN
```

It will generate a token then copy the token.

29. Go to system manager

30. click on code artifact token >>edit >> paste inside the value box >> update

31 create parameter for RDS_ENDPOINT, RDS_USER, RDSPASS

32. Go to s3 bucket >> Create Bucket

    - Bucket Name(vprofile-cicd-testoutput56)
    - create

33. Go to Build Project >> Create Build Project
    - project name(software testing)
    - source provider(aws code commit)
    - Repository (vprofile-code-repo)
    - Branch (selenium auto scripts)
    - Operating system (window server 2019)
    - Runtime (Base)
    - image (aws/codebuild/windows-base: 2019)
    - Existing role
    - insert build command(Go to cd-aws branch >> aws files >> win_buildspec.yml >> copy and paste)
    - Artifact type(s3 bucket) >> select the bucket
    - enable sematic versioning
    - select zip
    - use the same log group name
    - Enter stream name
    - create project

34 Go to pipeline >> create pipeline

35. Type Name >> new role >> next

36. source

    - source provider(aws codecommit)
    - branch (cd-aws)
    - next

37. build
    - project name(vprofile14-build&release)
