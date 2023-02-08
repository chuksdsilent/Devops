# Post Installation

## Jenkins

1. Open Jenkins ip address (port 8080)

2. ssh into jenkins server

3. Change to root user

```
 sudo -i
```

4. Copy the jenkins password path

```
cat jenkins password path
```

5. Copy the password and paste on the jenkins box >> Next

6. Select Install suggested plugins

7. Create admin user >> save and continue

8. Start using jenkins

## Nexus

1. Copy the ip address (port 8081)

2. Click on sign in

3. ssh into nexus

4. change to root user

```
sudo -i
```

5. copy the password path

```
cat nexus password path
```

6. Paste the password on the nexus box

7. Change the password

8. Select Enable anonymous access

9. Click on settings >> Repository >> Create Repository

10. Select maven hosted Repository

11. Type the name >> create repository

12. Create repository >> maven 2 proxy repository

13. Type the name

14. Under Remote storage type the repository url it will download dependencies from

```
https://repo1.maven.org/maven2
```

15. Create

16. Click on Create repository >> maven2(group)

17. Type the name

18. Click on Create repository >> maven2(group)

19. Type the name >> Create

20. Under Member repositories move your repo (vpro-maven-central), maven releases and snapshot from available to members

21. Go to vprofile project >> open .m2 >> open settings.xml
