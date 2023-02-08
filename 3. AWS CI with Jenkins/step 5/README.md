# Test, Code Analysis and Sonarqube Integration

## Test

1. Go to Jenkins >> New Item

2. Enter the name(test)

3. Copy from older job(build)

4. scroll down Under build inside goals box type

```
test
```

5. click on advanced >> close the archaive artifact

6. click on save.

7. Build Now

8 open the build job >> scroll down and under post-build action >> select build other jobs >> inside projects to build box type Test.

9. Save

## Integration

1. Go to Jenkins >> New Item

2. Enter the name(integration)

3. Copy from older job(test)

4. scroll down Under build inside goals box type

```
verify -DskipUnitTests
```

5. open the build job >> scroll down and under post-build action >> select build other jobs >> inside projects to build box type Integration Test.

6. click on save.

7. Build Now

## Code Analysis

1. Go to Jenkins >> manage jenkins >> manage plugins >> avaialble >> search for a plugin(checkstyle)

2. check mark the checkstyle and violation >>install without restart

3. Go to jenkins >> New item >> type the name(Code Analysis)

4. copy from older job (build) >> ok

5. scroll down Under build inside goals box type

```
checkstyle/checkstyle
```

6. save

7. Build Now.

### To setup Violation

1. Go to the project >> workspace >> target >> copy the file name

2. Go to configure >> scroll down >> under post build action >> select Report violation

3. Paste the filename(target/checkstyle-result.xml) on the checkstylebox

4. Go to integration Test >> configure

5. Under post build action >> select build another job

6. Save

## Sonarqube Server

1. Copy sonarqube server ip address

2. run on browser >> click on login

3. default sername: admin, password: admin

4. click on the profile picture >> My account

5. Go to security >> Enter token name >> generate

6. Go to jenkins >> manage jenkins >> manage plugin >> search for sonar >> check mark sonarqube server and sonar quality >> install without restart

7. Go to jenkins >> manage jenkins >> global tool configuration >> scroll down and add sonarqube scanner >>select the version, type the name(SONAR-4.4.2170) >> save

8. Check mark Enable injection of sonarqube

9. Add sonarqube >> Type the name, copy sonarqube private ip and paste with port, add credentials.

10. Under kind select secret text.

11. Copy the token from sonarqube and paste on the box then type id and description

12. Scroll down and under quality gate click on add sonar instance >> type the name, enter private of sonarqube, paste token of sonarqube or type the username and password

13. Save

## Sonar Scanner Analysis Job

1. Go to Jenkins >> New item

2. Type the name(sonarScanner-CodeAnalysis)

3. copy from build job >> ok

4. Under build >> Remove -DskipTests

5. Remove archive artifact

6. under post build actions >> select execute sonarqube scanner

7. undder execute sonarqube scanner >> in the analysis properties >> go to vprofile project >> inside the ci-jenkins branch >> open the userdata open sonar-analysis-properties >> copy the properties >> paste inside the analysis properties box

8. under post build actions >> invoke top-level Maven targets >> type checkstyle/checkstyle in the goal box >> them move it above execute sonarqube

9. Save

10. Build Now

## Quality Gate

1. Go to sonarqube >> quality Gate >> Create

2. Type the name >> save

3. Add condition

4. Do your configuration >> Save.

5. Go to Jenkins >> open the project >> configure

6. Under post build action >> select quality gate sonarqube

7. Enter the project name >> select if the gate did not happen what happens.

8. save >> Build Now
