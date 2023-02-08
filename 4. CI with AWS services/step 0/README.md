# AWS CodeCommit Setup

1. Login to AWS
2. Go to CodeCommit
3. CodeCommit >> Repository >> Create Repository
4. Type the Repository Name(vprofile-code-repo) >> Create/Save
   Note: You can access the repo with https or ssh but ssh is preferrable

## Create IAM user

1. Go to IAM user >> Add User

2. Type the username

3. select programmatic access

4. select attach policies

5. select create policy

6. Search for codecommit

7. select all under Access Level

8. Select specific >> ARN under Resources

9. Type the Region and repo name >> Add

10. Review policy

11. Type the policy name

12. Refresh and Select the policy create on iam policies >> next

13. Create user

14. Open the user >> security >> upload ssh public key

15. Generate ssh on your local system

```
ssh-keygen
```

14. cat and copy the public key

15. upload ssh public key >> paste >> click upload ssh public key

## Create SSH_CONFIG FILE

16. Go to gitbash

17. Go to /.ssh (/c/Users/computer-name/.ssh)

18. create ssh_config (vim config)

19. paste the following code

```
 Host git-codecommit.\*.amazonaws.com
 User ssh-key-id(eg 3833383)
 IdentityFile `/.ssh/private-file-name
```

20. Change mode to 600(chmod 600)

21. To test

```
ssh -v git-codecommit.region.amazonaws.com
```

22. Clone the aws respository

23. Go To codecommit >> respository >> code

24. copy the url

25. git clone url

26. Clone the the vprofile project

27. switch to ci-aws

```
git checkout ci-aws
```

28. To check the repository

```
cat .git/config)
```

29 switch to master

```
git checkout master
```

30 To view all branches

```
git branch -a | grep -v READ | cut d '/' f3 | grep -v master > /tmp/branches.txt)
```

31. To checkout all the branches

```
for i in `cat /tmp/branches.txt`; do git checkout $i; done
```

32.To remove current remote origin

```
git remote rm origin code-commit-ssh-url
```

33. To add aws remote origin

```
git remote add aws url
```

34. To push to aws

```
 git push origin --all
```

35. To push tags

```
 git push --tags
```

36. To check

```
Go to codecommit >> repository >> code refresh
```
