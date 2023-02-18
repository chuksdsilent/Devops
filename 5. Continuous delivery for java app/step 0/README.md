# Continuous Integration Project

1. Login to aws

2. Go to keypairs
   - Create a key pair
3. Go to security group

4. create 3 security group

   - Jenkins (8080, 22 from my ip, all traffic from sonarqube )
   - Nexus (8081, 22 from my ip, allow 8081 from jenkins sg)
   - Sonarqube (80 from my ip, 80 from jenkins sg, 9000)

5. Launch ec2 instances for

   - Nexus(centos7)

6. Click on signin

7. Enter username(admin) and password(admin)

8. ssh into nexus server

9. cat the url to get the password

10. paste the new password >> next

11. set a new password

12. enable anonymous access

13. create

14. Setting >> repository

15. create 3 repositories

16. maven hosted zone >> create repository

17. Type the name >> create

18. Click on maven 2 proxy

19. create repository

20. Type the name

21. Enter the maven repo url(https://repo1.maven.org/maven2/) >> create

22. maven 2 group

23. create repository

24. Type the name

25. maven snapshot

26. create repository

27 Type the name

28. Change release to snapshot >> create

29. Under Available move the following

    - vpro-maven-central
    - maven-releases
    - vprofile-snapshot

30. click on create

### Go to settings.xml which is located in .m2

1. Jenkins(Ubuntu)

2. Login into jenkins and cat the url to the password

3. complete installation

4. create a jenkins job

5. type the nme

6. select freestyle

7. ok

8. under source code management

9. select git

10. Enter the url

11. enter the branch (ci-jenkins)

12. Add build step >> Invoke top-level Maven targets

13. Goals (install -DskipTests)

14. settings file >> Settings file in filesystem

15. file path (settings.xml)

16. For properties(Go to nexus and the following properties SNAP-REPO, NEXUS-USER=admin,NEXUS-PASS=nexus-password, RELEASE-REPO,CENTRAL-REPO, NEXUS-GRP-REPO,NEXUSIP=nexus-private-ip, NEXUSPORT=8081)

17. Copy them and put on the property box

18. Build Now.

19. Go nexus >> browse >> vpro-maven-central

## Integration Notification with Slack

1. sign in to slack

2. create a channel

3. create bot

   - Go to api.slack.com
   - start building
   - create app
   - type the app name and select the workspace >> create app

4. OAuth and Permissions

5. Under Bot Token Scopes >> chat write

6. Install app to workspace >> Allow

7. Copy the token and store somewhere safe

8. Go to the channel you created on slack type @Jenkins >> Enter

9. Invite to channel

10. Go to manage Jenkins >> plugins >> available >> search for slack >> slack notification >>install without restart

11. Go to manage Jenkins >> manage credentials >> jenkins >> under kind select secret text >> save

12. Save

14 Go to manage jenkins >> configure system >> under slack checkmark custom slack app bot user

15. Go to your job >> configure >> post build action >> slack notifications >> checkmark the ones you want >> advance >> select slack token and type the channel-name >> test the connection >> ok

16. archive the artifact

- post build action >> archive the artifact
- \*_/_.war
- save

17. Build Now.

## Complete with Project 2
