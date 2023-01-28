1. Login to AWS
![capture](https://user-images.githubusercontent.com/18073289/215264000-6219481b-72e5-4e99-8579-2e55c03f6348.PNG)
2. Go to security groups
![Capture2](https://user-images.githubusercontent.com/18073289/215264011-e02a839b-0f8c-4156-acc7-0b3fa78b0732.PNG)
3. Create a security group for load balancer(Allow port 80 and 443)
![Capture4](https://user-images.githubusercontent.com/18073289/215264042-ea328be0-645d-40d6-bded-bbc342115a57.PNG)
4. Create security group for Tomcat instance(Allow traffic from load balancer and open port 8080)
![Capture7](https://user-images.githubusercontent.com/18073289/215264083-f761b890-1847-4be5-b06a-764947ac575e.PNG)
5. Create security group for Backend Services(Allow traffic from tomcat open port 3306 for mysql, 11211 for memcache, 5672 from rabbitmq, allow traffic from its security id so it can interact with itself)
![Capture6](https://user-images.githubusercontent.com/18073289/215264057-081a7d1d-8bc9-4138-864d-7270c1906148.PNG)
6. Go to keypair and click on create keypair
![Capture8](https://user-images.githubusercontent.com/18073289/215264099-1f7e58a9-3c7b-4c7a-901a-661682e74a51.PNG)
7. Select pem for git bash and ppk for putty
8. Type key name and download the key then save
9. Clone the vprofile project from https://github.com/chuksdsilent/vprofile-project.git
10. change to lift and shift directory (git checkout aws-LiftAndShift)
    11 Change directory to userdata
11. Go to ec2 instance and launch instance

- Provision mysql instance(Ubuntu)
  - Select t2 micro
  - Select the exist backend security group
  - Copy mysql userdata bash script and past in the userdata box
  - Launch Instance
  - Login to the instance and check if database is created
- Provision memcache instance(Centos)
  - Select t2 micro
  - Select the exist backend security group
  - Copy memcache userdata bash script and past in the userdata box
  - Launch Instance
- Provision rabbitmq instance(Centos)
  - Select t2 micro
  - Select the exist backend security group
  - Copy rabbitmq userdata bash script and past in the userdata box
  - Launch Instance
