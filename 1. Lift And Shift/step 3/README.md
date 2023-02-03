# Create Load Balancer

## Create a target Group

1. Go to target Group

![Capture30](https://user-images.githubusercontent.com/18073289/216625998-9ce7c419-f52f-4577-9795-5b4ae131a435.PNG)

2. type the name and port(8080)

![Capture31](https://user-images.githubusercontent.com/18073289/216626042-7d8d8b6e-acf6-42cd-81b2-7c6717a7606f.PNG)

3. health is /login,  override the port to 8080,  Reduce threshold to 3 >> Next

![Capture32](https://user-images.githubusercontent.com/18073289/216626275-0e66516f-1be9-40ce-88af-df8647b7059c.PNG)

4. select the instance and click on include expending below

![Capture33](https://user-images.githubusercontent.com/18073289/216626383-d8bde916-45b6-41ca-bd61-241436762412.PNG)

5. create target group

## Create a Load balancer

1. Go to Load balancers

![Capture34](https://user-images.githubusercontent.com/18073289/216626457-7d12799f-7b82-45b2-a30b-91f5e602d674.PNG)

2. Select application load balancer

![Capture35](https://user-images.githubusercontent.com/18073289/216626514-409f1ac3-8220-4097-93ec-c94ee5a43f80.PNG)

3. Type the name and select internetfacing

![Capture36](https://user-images.githubusercontent.com/18073289/216626643-3c50ab67-fd32-46e7-99ac-7271cf93aae5.PNG)

4. select all the zones

![Capture37](https://user-images.githubusercontent.com/18073289/216626713-6cd4609e-a4dd-4789-b54f-8b8e38a8a3f9.PNG)

5. select listeners(http and https)

![Capture38](https://user-images.githubusercontent.com/18073289/216626782-9b1b0901-8c52-4312-a44d-b98c2b05533b.PNG)

6. choose a certificate from acm

7. select security group from load balancer
8. select target config
9. create the load balancer
10. Copy the url of the load balancer and create a cname record on your domain name host
