# Create Auto Scaling Group

## Create AMI of the Tomcat Instance

1. Select the instance >> Action >> Image >> Create Image and Type the name of the image

![Capture41](https://user-images.githubusercontent.com/18073289/216647731-f3f32db7-a64e-4808-b0cd-dfe7eada35eb.PNG)

2. Create Image

## To Create Launch Configuration

1. Go to launch configurateio >> Create launch configuration  >> Type the name and Select the AMI
![Capture47](https://user-images.githubusercontent.com/18073289/216647769-49699b02-c82f-428c-86cf-c6145503cf3a.PNG)

2. select the AMI

![Capture43](https://user-images.githubusercontent.com/18073289/216653361-53c04c46-d800-405a-8db6-d78c501826ca.PNG)


3. select the instance type(t2.micro)

![Capture44](https://user-images.githubusercontent.com/18073289/216653461-ab2bd180-b4ff-459e-b001-ec81493ee96e.PNG)

4. select security group of apps

![Capture45](https://user-images.githubusercontent.com/18073289/216653629-1b3db60f-7725-4f53-a2c3-9129488dc1fb.PNG)

5. select key pair

![Capture46](https://user-images.githubusercontent.com/18073289/216653665-94cc5e45-89a7-47a8-9896-736ae611db28.PNG)

9. select iam role

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
