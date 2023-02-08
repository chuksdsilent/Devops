# Data Backup of CI/CD Stack

1. Login to aws

2. Create an s3 bucket

## Create IAM role

3. Go IAM role

4. Go to role >> create role

5. select aws service under trusted entity and ec2 under choose a use case >> next

6. select s3fullaccess policy

7. Enter role name >> Create role

## Attach role to instance

NB: Attach the role to all the instances

8. Go to instances >> select the instance >> actions >> security >> modify IAM role

9. select the IAM role >> save

10 ssh into jenkins instance

11 switch to root user

```
sudo -i
```

12. Stop jenkins service

```
systemctl stop jenkins
```

13. change dir to /var/lib

```
ls ltr
```

14. To see the size

```
du -sh jenkins
```

15. cd into jenkins repository

```
cd jenkins/repository
```

16. delete all

```
rm -rf *
```

17. cd into jenkins workspace

```
cd jenkins/workspace
```

16. delete all

```
rm -rf *
```

## To see how many time you run a job

19. cd into builds

```
cd /var/lib/jenkins/jobs/BuildProject/builds
```

19. Delete the numbers

```
rm -rf 1 10 9
```

20. To archive

```
cd /var/lib
tar -czvf jenkins_cicdjobs.tar.gz jenkins
```

21. install aws cli

```
sudo apt update
sudo apt install awscli -y
```

22. Copy the tar file to s3 bucket

```
aws s3 cp jenkins_cicdjobs.tar.gz s3://bucket-name
```

23. ssh into nexus

24. stop nexus service

```
systemctl stop nexus
```

25. archive nexus

```
tar -czvf nexus-cicd-vpro-pro.tgz nexus
```

26. Copy the tar file to s3 bucket

```
aws s3 cp nexus-cicd-vpro-pro.tgz s3://bucket-name
```

27. ssh to sonarqube server

28. stop sonarqube sonar

```
sudo systemctl stop sonarqube
```

29. cd into opt

```
cd /opt
```

30. Archive /opt

```
tar -czvf sonarqube-vpro-pro-data.tgz sonarqube
```

31. Copy the tar file to s3 bucket

```
aws s3 cp sonarqube-vpro-pro-data.tgz s3://bucket-name
```

### To restore archive programs

32. Launch a new jenkins instance using ubuntu

33. Install Jenkins

34. ssh into the instance

35. change to root user

```
sudo -i
```

36. cd into /var/lib

37. stop jenkins server

```
systemctl stop jenkins
```

38. Attach the s3 bucket IAM role to the jenkins instance

39. install aws cli

40. Download Jenkins artifact from s3 bucket

```
aws s3 cp s3://bucket-name/file-name .
```

41. To extract the file

```
tar xzvf file-name
```

42. check jenkins user

```
ls -ld jenkins
```

43. Change Ownership

```
chown jenkins.jenkins jenkins -R
```

44. start Jenkins

```
systemctl start jenkins
```
