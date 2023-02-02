## Create Tomcat EC2 Instance(Ubuntu)

1. login to aws and go to ec2 instance then launch instances
2. Select t2 micro
3. Select the exist vprofile-app-sg security group
4. Copy Tomcat userdata bash script and past in the userdata box
5. Launch Instance
6. install jdk in your local system using powershell and choco commands

```
choco install jdk8 -y
```

6. install maven in your local system using powershell and choco commands

```
choco install maven -y
```

7. Navigate to your vprofile project folder

8. Change directory to resources

```
 cd vprofile-project/src/main/resources
```

9. update the application.properties file

```
sudo vim application.properties
```

Note - change db01 -> db01.vprofile.in, rmq01 -> rmq01.vprofile.in, mc01 -> mc01.vprofile.in

10. Build the app

```
mvn install
```

10. Install AWS CLI

```
choco install awscli -y
```

### Create IAM user

11. Go to users
12. Add user
13. Type the name(vprofile-s3-admin) and select programmatic access
14. Attach to existing policies and Select s3 bucket full access
15. Preview and create
16. Go to the user and open it
17. Go to security credentials
18. Go to Access Key and create new key
19. Select Command Line interface and next
20. Type description and create new Key
21. Go to your local system and add the access keys

```
aws configure
```

22. Create s3 bucket

```
aws s3 mb s3://bucket-name
```

Note. Bucket name should be unique 23. Navigate to vprofile-project/target

```
cd vprofile-project/target
```

24. copy vprofile-v2.war to s3 bucket

```
aws s3 cp vprofile-v2.war s3://bucket-name
```

25. To see list of file in a bucket

```
aws s3 ls s3://bucket-name
```

### In order to download the target file into ec2 tomcat instance we need to create a role

### To create a role

1. Go to IAM
2. Go to roles
3. Create a role , select aws service and select EC2
4. select s3fullacess
5. Enter role name and create role
6. Go to instances and select app01
7. Go to actions >> instance settings >> modify IAM Role
8. select the role and update
9. Login to app01
   10 To check tomcat status

```
systemctl status tomcat8
```

11. To

```
cd /var/lib/tomcat8/webapps && ls
```

12. Stop tomcat service

```
sudo systemctl stop tomcat8
```

13. delete the root directory

```
sudo rm -rf ROOT
```

14. Install AWS cli

```
sudo apt install awscli -y
```

17. Download the File from s3 to tomcat(app01) server

```
aws s3 cp s3://vprofile-sam-artifact-storage/vprofile-v2.war /tmp
```

18. Copy The file to tomcat8

```
sudo cp vprofile-v2.war /var/lib/tomcat8/webapps/ROOT.war
```

19. Start the tomcat service

```
sudo systemctl start tomcat8
```

20 change to webapp directory

```
cd /var/lib/tomcat8/webapps/ROOT/WEB-INF/classes  && cat application.properties
```

21. To test for connectivity. Install telnet

```
sudo apt install telnet
```
