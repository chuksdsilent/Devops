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

![Capture62](https://user-images.githubusercontent.com/18073289/216981054-9b5b190e-0594-4a2e-872e-bbe0fb6a1c86.PNG)

3. Type name of subnet group and create

![Capture63](https://user-images.githubusercontent.com/18073289/216981279-b514309b-4ac8-4097-ab94-d7c47d12435b.PNG)

4. Select all the availability zone

![Capture64](https://user-images.githubusercontent.com/18073289/216981312-3472c516-5fe5-44fd-acee-d74f43fa94f8.PNG)

5. Go to parameter group

![Capture65](https://user-images.githubusercontent.com/18073289/216981374-778a9a1c-a215-4a47-a5b9-b8a77a158b48.PNG)

6. Select database version and type the name

![Capture67](https://user-images.githubusercontent.com/18073289/216981507-53cf2f57-aa3e-4409-b1c7-b793654a681e.PNG)

7. Go to databases >> Create Database

![Capture68](https://user-images.githubusercontent.com/18073289/216981558-c5bbc3cf-8ee3-41df-8558-4afc6ae8ad5b.PNG)

8. Select MySQL(Amazon Aurora is recommended)

![Capture69](https://user-images.githubusercontent.com/18073289/216981678-0aca0e6c-78f4-460e-be76-1e71403c331c.PNG)

9. select the version(5.6) then vselect template(Dev/Test)

![Capture70](https://user-images.githubusercontent.com/18073289/216981807-621ddd58-7a78-468f-b348-8434e0df45df.PNG)

10. Type Database Name

![Capture71](https://user-images.githubusercontent.com/18073289/216982147-e23f4a7f-8295-4b6c-b684-36fad7eacd38.PNG)

11. Db instance(db.t2micro. r series is better for production)

![Capture72](https://user-images.githubusercontent.com/18073289/216982201-12cc28c6-9677-4dc3-9295-d6f57d2632c1.PNG)

12. Seect Public Access

![Capture73](https://user-images.githubusercontent.com/18073289/216982754-c7cf687f-2f22-4249-a50c-778082a65fc3.PNG)

13. select security group for db, Backup (7 days), Enable enhance monitoring then create database
14.  Copy out your credentials (username, password and endpoint)

![Capture74](https://user-images.githubusercontent.com/18073289/216983090-da14487c-f6e6-4554-b4f8-3ccd24f90ab7.PNG)


