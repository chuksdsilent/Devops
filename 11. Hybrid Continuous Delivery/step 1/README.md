# Sonar Cloud Integration

1. Login to Sonar Cloud

2. Go to profile picture >> my account

3. Click on Security tab

4. Type any name >> Generate >> Copy out the token

5. Go to Jenkins >> Manage Jenkins

6. Go to configure system

7. scroll down to sonarqube server and Add SonarQube

8. Enter name(SonarHybrid) and url(https://sonarcloud.io)

9. save it and reopen it

10. Add credentials

11. Under kind select Select text

12. Enter the token on the secret box, type the username

13. Save

14. Go to Dashboard >> New item

15. Enter item name(sonarCloud-CodeAnalysis) >> select freestyle

16. copy from SonarScanner-CodeAnalysis >> Ok

17. branch (\*/cd-aws-jenkins)

18. Under Execute SonarQube Scanner

    - SonarQube Installation (sonarHybrid)
    - Analysis Properties (sonar.projectName=vprofile-repoHybrid)
    - Add sonar.organization=organization-name from sonarcloud

19. Remove Quality Gate

20. Add build Step >> Execute shell

21. open bash >> move to vprofile-project >> switch to branch cd-aws-jenkins

```
git checkout cd-aws-jenkins
```

22. open Jenkins_files/sonarQualityGatesCommand.txt

23. Copy the content and paste on Execute shell then change the project key

24. Under Build

    - Settings file (Use default settings)

25. Under Post-build Actions

    - remove build other projects

26. Save >> Build Now.
