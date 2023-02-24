# k8 with EKS Setup

1. Clone the vprofile project

2. switch to kubernetes-setup

```
git checkout kubernetes-setup
```

3. open eks folder and vagrant file

4. Create AWS Key(vprofile-eks-key)

5. Run the vagrant file

```
vagrant up
```

6. login to vagrant

```
vagrant ssh
```

7. check aws version

```
aws --version
```

8. check eks version

```
eksctl version
```

9. configure aws

```
aws configure
```

10. Go to IAM USER

11. Type the name(eks-admin) >> select programmatic access

12. Attach existing policy directly

13. Select AdminstorAccess

14. Tags >> create user

15. Setup aws cli in vagrant

16. Go to eks-cluster-setup.sh

17. You can change some variables like region, cluster-name, node-name, key

18. Run the script

```
./eks-cluster-setup.sh
```

19.
