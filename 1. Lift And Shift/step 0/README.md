1. Login to AWS
2. Go to security groups
3. Create a security group for load balancer(Allow port 80 and 443)
4. Create security group for Tomcat instance(Allow traffic from load balancer and open port 8081)
5. Create security group for Backend Services(Allow traffic from tomcat open port 3306 for mysql, 11211 for memcache, 5672 from rabbitmq, allow traffic from its security id so it can interact with itself)
6. Go to keypair and click on create keypair
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
