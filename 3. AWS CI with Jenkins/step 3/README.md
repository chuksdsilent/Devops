# Creating Jenkins Job

1. Open jenkins

2. Create a new Job

3. Enter the name and select freestyle >> ok

4. Under source code Management select git >> type the git url and branch(ci-jenkins)

5. Add your credentials

6. Under build select top level invoke target

7. Under goal type

```
install -DskipTests
```

8. click on advance

9. under Settings file choose settings file in filesystem then type settings.xml

10. Go to settings.xml in vprofile project copy out the varIables(SNAP-REPO, NEXUS-USER, NEXUS-PASS, RELEASE-REPO, CENTRAL-REPO, NEXUS-GRP-REPO, NEXUSIP, NEXUSPORT)

11. Go to nexus >> settings >> vprofile-maven-central you will see the snapshot there.

    - SNAP-REPO = vprofile-snapshot
    - NEXUS-USER = admin
    - NEXUS-PASS = admin123
    - RELEASE-REPO = vprofile-release
    - CENTRAL-repo = vprofile-maven-central
    - NEXUS-GRP-REPO = vprofile-maven-group
    - NEXUSIP=nexus private ip
    - NEXUSPORT= 8081

12. Copy the variable and paste in the maven properties box >> save

13. Build the job

14. if the build is complete Go to nexus >> Browse section >> Browse >> vpro-maven-central. You should see the dependencies getting stored.
