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

16. Go to Whogohost
    - create a CNAME record with the beanstalk endpoint
17. Enable stickness
    - Go to beanstalk
    - Configuration
    - edit under load balancer
    - select default
    - Edit
    - Select stickness
    - Save
18. Create CloudFront
    - Go to Cloudfront
    - Create Distribution
    - Region(domain name)
    - Allow all the methods
    - select http & https
    - select ssl certificate
