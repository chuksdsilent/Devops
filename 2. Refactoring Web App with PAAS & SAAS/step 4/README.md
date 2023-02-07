# To test RDS

1. Create ec2 instance

2. login to it

3. update the instance

```sudo apt update

```

4. install mysql client

```
 sudo apt install mysql-client -y
```

5. login to the mysql

```
mysql -h mysql-endpoint -u username -pPassword
```

6. clone vprofile-project

7. switch to branch aws-refractor

```
git checkout aws-Refractor
```

8. cd vprofile-project/src/main/resource

9. import mysql((mysql -h mysql-endpoint -u username -pPassword dbname < db_backup.sql)

Note: you must allow rds to accept traffic from ec2 security group

10. Copy credential of each service
   - db endpoint
   - Amazon MQ AMQP