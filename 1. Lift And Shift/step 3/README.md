# Create Load Balancer

## Create a target Group

1. Go to target Group
2. type the name and port(8080)
3. health is /login
4. override the port to 8080
5. Reduce threshold to 3 >> Next
6. select the instance and click on include expending below
7. create target group

## Create a Load balancer

1. Go to Load balancers
2. Select application load balancer
3. Type the name
4. select internetfacing
5. select all the zones
6. select http and https
7. choose a certificate from acm
8. select security group from load balancer
9. select target config
10. create the load balancer
11. Copy the url of the load balancer and create a cname record on your domain name host
