# Update Private IPs of the instances to Route53 private DNS ZONES

## To create Hosted zone

1. Go to Route53 and select hosted zone
2. Type in domain name(sam.org), Select the Region, Select default vpc and Create hosted zone

## To create records

1. Create Simple Records
   - Select simple routing
   - Type the sub domain (db01, mc01,bmq01) - Enter ip address of the corresponding ec2 instance
