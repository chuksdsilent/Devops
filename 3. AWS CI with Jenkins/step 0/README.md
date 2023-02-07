# Creating Security Group

1. create security group for jenkins(port 8080 from anywhere, All traffic from sonarqube SG)

2. create security group for nexus(port 8081, port 8081 from jenkins SG)

3. create security group for sonarqube (port 9000, port 80 from my ip, port 80 from jenkins SG)
