## REFACTORING WITH AWS

In this project I refactored a lift and shift project using aws services in order to boost agility or improve business continuity.

I moved application running virtual machine to aws services.

### PROBLEMS

- Operational Overhead
- Struggling with uptime & scaling
- Manual Process / Difficult to automate

### SOLUTIONS

- Automation
- PayAsUGo
- Flexibility
- Easy to manage

### AWS SERVICES

- BEANSTALK
- S3 BUCKET
- RDS
- ELASTIC CACHE
- ACTIVE MQ
- ROUTE53
- CLOUDFRONT

### OBJECTIVE

- FLEXIBLE INFRA
- NOT UPFRONT COST
- IAAC
- PAAS
- SAAS

### FLOW OF EXECUTION

1. Login to aws account
2. Create Key pair for beanstalk instance login
3. Create Security Group for Elasticcache, RDS & ActiveMQ
4. Create
   - RDS
   - Amazon Elastic Cache
   - Amazon Active MQ
5. Create Elastic Beanstalk Environment
6. Update SG of backend to allow traffic from Bean SG.
7. Update sg of the backend to allow traffick
8. launch ec2 instance for db
9. Login to the instance and initialize RDS DB
10. Change healthcheck on beanstalk to /login
11. Add 443 https listner to ELB
12. Build Artifact with Backend Information
13. Deploy Artifact to Beanstalk
14. Create CDN with ssl cert
15. Update Entry in Whogohost

### STEPS


9. Create Elastic Beanstalk
   - Type Name
   - Select the platform (tomcat)
   - select the version(8.5)
   - select sample application
   - select configure more option
   - instances
     - select the security group
     - save
   - Capacity
     - select load balancer under Env. Type
     - instances( select min and max)
     - instance type(t2.micro)
     - AZ(ANY)
     - Scaling trigger(Network out is most popular)
   - Load balancer
     - select application
     - save
   - Deployment
     - Deployment Policy (Rolling)
     - Batch Size(50%)
     - Save
   - Security
     - Select key pair
     - save
   - Enhance Monitoring
     - Basic for test, Enhanced for production
     - Tick Stream the logs
     - Delete after usage
     - save
   - Create App
10. Update backend security group - Allow traffic from beanstalk security group for each of the security group(mysql, memcache, rabbitmq)
11. Update the health check
    - Go to beanstalk >> Environment >> configuration >> Load Balancer
      - Edit
      - Add Listener
      - port(443)
      - select certificate
      - save
      - under health check
      - select default
      - add /login
      - save
      - apply
12. Clone vprofile-project
    - switch to aws refactor(git checkout aws-Refactor)
    - cd vprofile-project/src/main/resources
    - vim application.properties
    - change the properties of db, memcache and rabbitmq
    - save
13. Go to Beanstalk
    - application version
    - Type in the name
    - upload the war file
    - select the version >> action >> deploy
14. Go to Whogohost
    - create a CNAME record with the beanstalk endpoint
15. Enable stickness
    - Go to beanstalk
    - Configuration
    - edit under load balancer
    - select default
    - Edit
    - Select stickness
    - Save
16. Create CloudFront
    - Go to Cloudfront
    - Create Distribution
    - Region(domain name)
    - Allow all the methods
    - select http & https
    - select ssl certificate
