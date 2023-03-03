# Protecting Data in Transit

1. SSL/TLS Encryption

   - https on ELB
   - Bucket policy for HTTPS
   - ACM can generate and manage Certs

2. IPSec ESP/AH

3. Services
   - s3
   - RDS
   - DynamoDB
   - EMR

# Create a Load balancer

1. Go to ec2

2. Go to load balancer

3. select classic

4. Select

   - https on security group
   - Load Balancer Port (443)
   - Instance Port (80)

5. Choose certificate from ACM

6. Configure health check

   - port (8080)
   - path(/logina)

7. Add Instance >> create

8. Allow traffic from load balancer SG in your jenkins instance on port 8080

9. use the load balancer url to create a CNAME record.

## Encryption on Transit for S3

1. Open the file

2. Go to permissions >> edit

3. Copy the policy and paste

4.
