# Creating RDS

## To Create security group for backend services

1. Add a dummy rule(22 from my ip)
2. save
3. Edit sg
4. Remove the dummy
5. Allow all traffic from the itself(its security group)

## To Create RDS

1. Go to RDS
2. Create a subnet Group
3. Create DB Subnet Group
4. Select all the availability zone
5. Go to parameter group
6. Select database version and type the name
7. Go to databases
8. Create Database
9. Select MySQL(Amazon Aurora is recommended)
10. select the version(5.6)
11. select template(Dev/Test)
12. Type Database Name
13. Db instance(db.t2micro. r series is better for production)
14. Storage autoscaling(1000)
15. Availability(Single instance, multi-az for production)
16. select security group for db
17. Backup (7 days)
18. Enable enhance monitoring
19. create database
    20 Copy out your credentials (username, password and endpoint)
