# Kops for Cluster on AWS

1. You need a valid domain for k8 DNS records

2. Create a linux VM and setup

   - kops, kubectl, ssh keys, awscli

3. lLogin to AWS account and setup

   - s3 bucket, IAM User for AWSCli, Route53 Hosted Zone

4. Create an EC2 instance with

   - t2.micro
   - SG(port 22 from my ip)
   - name (kops)
   -

5. Create s3 bucket which will be used to store the state of kops

6. Go to IAM user

7. Type the username >> Next

8. Attach existing policies directly

9. Select for AdministratorAccess >> next

10. Add tags >> Create user

11. Go to Route53

12. Create a hosted zone

    - name(vprofile.groovy.in)
    - public hosted zone
    - create

13. Add the values to whogohost

14. Attach the user to the EC2 Instance

15. Login to ec2 instance and generate ssh key

16. Install aws cli

```
sudo apt update && apt install awscli -y

```

17. configure AWS

```
aws configure
```

18. Install kubectl

19. Give the file executable permission

```
chmod +x ./kubectl
```

20. move it to /usr/local/bin/

```
sudo mv kubectl /usr/local/bin/
```

21. check if kubectl is install

```
kubectl
```

22. Install kops

```
chmod +x ./kops-linux-amd64
```

23. check if kops is install

```
kops
```

24. To run kops and create a cluster

```
nslookup -type=ns kubevpro.groophy.in
```

25. To create cluster

```
kops create cluster --name=kubevpro.groophy.in --state=s3://bucket-name --zones=us-east-2a,us-east-2b --node-count=2 --node-size=t2.micro --master-size=t2.micro --dns-zone=kubevprofile.groophy.in

```

26. Then update the cluster

```
kops update cluster --name kubevpro.groophy.in --state=s3://bucket-name --yes
```

27. To validate the cluster

```
kops validate cluster --state=s3://bucket-name
```
