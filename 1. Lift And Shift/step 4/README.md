# Create Auto Scaling Group

## Create AMI of the Tomcat Instance

1. Select the instance >> Action >> Image >> Create Image
2. Type the name of the image
3. Create Image
4. Go to launch configurateion
5. Create launch configuration
6. Type the name
7. Select the AMI
8. select the instance type(t2.micro)
9. select iam role
10. select monitoring with cloud watch
11. select security group of apps
12. select key pair
13. create launch configuration

## Autoscaling Group

1. Type the name
2. select the launch configuration
3. select all subnet
4. select the target group
5. select health check
6. Type in the min and max
7. Select track scaling
8. Create AutoScaling Group
