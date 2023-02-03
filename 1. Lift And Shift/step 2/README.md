## Create Tomcat EC2 Instance(Ubuntu)

1. login to aws and go to ec2 instance then launch instances

![ec2 instance](https://user-images.githubusercontent.com/18073289/216616355-da79fdcf-2976-4fcb-aea2-5a4e00e12e48.PNG)

2. Select t2 micro

![Capture13](https://user-images.githubusercontent.com/18073289/216616490-9991a045-7a01-4c92-a1eb-caa86b544fb3.PNG)

3. Select the exist vprofile-app-sg security group

![Capture14](https://user-images.githubusercontent.com/18073289/216616628-c8d9d796-81f0-4ef7-ab9a-6925bd11b70b.PNG)


4. Copy Tomcat userdata bash script and past in the userdata box

![Capture15](https://user-images.githubusercontent.com/18073289/216616710-6c66634e-c95f-4c86-8407-743b0d846cb8.PNG)


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

11. Go to users >>  Add user

![Capture16](https://user-images.githubusercontent.com/18073289/216618417-2188a0d5-73f4-44ec-a686-0559fa60d5a8.PNG)

12. Type the name(vprofile-s3-admin) and select programmatic access

![Capture17](https://user-images.githubusercontent.com/18073289/216619249-4212215d-4ad7-4df7-883f-13a75993e3cd.PNG)

13. Attach to existing policies and Select s3 bucket full access

![Capture18](https://user-images.githubusercontent.com/18073289/216619402-a55794e1-318c-4222-8f9a-7b98330750c7.PNG)

14. Preview and create

### To create an acccess key credentials for a user

15. Go to the user and open it

![Capture21](https://user-images.githubusercontent.com/18073289/216619933-ef80f6cd-ad5e-42ec-8afd-99b0506bdfcb.PNG)

16. Go to security credentials

![Capture21](https://user-images.githubusercontent.com/18073289/216620369-ac446570-3c98-4e4f-9ca9-ede970a9dec7.PNG)

17. Go to Access Key and create new key

![Capture22](https://user-images.githubusercontent.com/18073289/216620529-0f09c222-9422-41c9-9dfd-f21beb83da2c.PNG)

18. Select Command Line interface and next

19. Type description and create access Key

![Capture23](https://user-images.githubusercontent.com/18073289/216620778-caf02a72-c4a5-4540-9f27-e77527f05770.PNG)


20. Go to your local system and add the access keys

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

1. Go to IAM >> Go to roles >> Next

![Capture25](https://user-images.githubusercontent.com/18073289/216622347-b91877a3-9150-415f-91d1-65bfcec1bdd2.PNG)

2. Create a role , select aws service and select EC2  >> Next

![Capture26](https://user-images.githubusercontent.com/18073289/216622412-87a57e56-abaf-4411-94ee-18bc5a821dc5.PNG)

3. select s3fullacess

![Capture27](https://user-images.githubusercontent.com/18073289/216622842-76f419a8-303c-4282-a6cc-7b8f0c9d7bc8.PNG)

4. Enter role name and create role

![Capture28](https://user-images.githubusercontent.com/18073289/216622893-53e2e1c7-d65b-4602-9168-19fbf901f833.PNG)

5. Go to instances and select app01 >> Go to actions >> instance settings >> modify IAM Role


![Capture29](https://user-images.githubusercontent.com/18073289/216622999-a1a2f227-b84d-47bc-9446-3a60f7bc784c.PNG)

6. select the role and update


7. Login to app01

8. To check tomcat status

```
systemctl status tomcat8
```

9. To

```
cd /var/lib/tomcat8/webapps && ls
```

10. Stop tomcat service

```
sudo systemctl stop tomcat8
```

11. delete the root directory

```
sudo rm -rf ROOT
```

12. Install AWS cli

```
sudo apt install awscli -y
```

13. Download the File from s3 to tomcat(app01) server

```
aws s3 cp s3://vprofile-sam-artifact-storage/vprofile-v2.war /tmp
```

14. Copy The file to tomcat8

```
sudo cp vprofile-v2.war /var/lib/tomcat8/webapps/ROOT.war
```

15. Start the tomcat service

```
sudo systemctl start tomcat8
```

16 change to webapp directory

```
cd /var/lib/tomcat8/webapps/ROOT/WEB-INF/classes  && cat application.properties
```

17. To test for connectivity. Install telnet

```
sudo apt install telnet
```
