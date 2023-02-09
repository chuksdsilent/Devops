# AWS CodeArtifact Setup & AWS Systems Manager Parameter

## Create a repository on codeartifact

1.  Create Repository

2.  Type the name

3.  select maven central store >> next

4.  type the domain name (Any Name) >> next

5.  Click Review

6.  Create repository

7.  open the repository

8.  select view connection instruction

9.  select mvn

10. Copy the link

11. create an iam user with programmatic access and select codeartifact admin access policy

12 install and configure awscli

```
choco install awscli -y
```

13. configure aws

```
aws configure
```

14. Run the link you copied on gitbash

15. Display code

```
echo $CODEARTIFACT_AUTH_TOKEN
```

16. Go to vprofile project >> change to ci-aws branch

17. Change settings.xml file

18. switch to ci-aws

19. open settings.xml

20. change the url, domain name. ensure to add / after the url

21. Update pom.xml

22. change the url

23. git add and git commit

24. git push origin ci-aws
