# Ansible Setup for AWS

1. Login to AWS

2. Go to instance >> Launch Instance

   - OS(ubuntu 18)
   - key pair
   - SG(port 22 from my ip)
   - type the following code

   ```
    #!/bin/bash`
     sudo apt update
     sudo apt install ansible -y
   ```

3. login to the ansible server

4. search for ansible aws documentation

5. Go to IAM role

6. Go to create role

7. under trusted entity(aws service)

8. Under use case (EC2) >> next

9. search for s3fullaccess >> next

10. Tags >> create

11. Go to the instance >> Actions >> security >> Modify Iam role >> select the role >> save

12 To test type the following command on the instance

```
aws sts get-caller-identity
```

13.
