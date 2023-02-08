## AWS JAVA CONTINUOUS DELIVERY

In this project I will be creating a continuous delivery of java application to AWS using AWS cloud native services

## BENEFITS

- Automation
- Short MTTR
- Faulty Isolation
- Less Disruptive

## TOOLS

- Jenkins
- GIT
- Maven
- Checkstyle
- Slack
- Nexus
- Sonarqube
- Tomcat
- selenium
- Windows server
- AWS EC2

## Architecture

## FLOW OF EXECUTION

- Login to aws account
- Create key pair
- Create SG

  - Jenkins
  - Nexus
  - Sonarqube

- Create EC2 Instances with userdata
  - Jenkins
  - Nexus
  - Sonarqube
- Jenkins Post installation steps
- Nexus Repostory setup
  - 3 Respository
- Sonarqube post installation
- Jenkins
  - Build Job
  - Setup Slack Notification
  - Checkstyle code analysis
  - setup sonar integration
  - sonar code analysis job
  - upload artifact to nexus
- Connect all jos together with BuildPipeline
- Set Automatic build trigger
- Test with git
- Create SG

  - windows server
  - tomcat & backend servers

* Setup tomat & backend server on ec2 with userdata
* create jenkins job to deploy
* create a jenkins jo to deploy to artifact to staging
* Add windows node as slave to jenkins
* create a job to run software Tests(Selenium from windows server)
* Create jo to run software tests(selenium) from windows server
* connect all jos with Build Pipeline
* Test it by committing the code to github

## STEPS

1. Login to aws
2. Go to keypairs
   - Create a key pair
3. Go to security group
   - create 3 security group
     - Jenkins (8080, 22 from my ip, all traffic from sonarqube )
     - Nexus (8081, 22 from my ip, allow 8081 from jenkins sg)
     - Sonarqube (80 from my ip, 80 from jenkins sg, 9000)
4. Launch ec2 instances for

   - Nexus(centos7)

     - Click on signin
     - enter username(admin) and password
       - ssh into nexus server
       - cat the url to get the password
       - paste the new password
       - next
       - set a new password
       - enable anonymous access
       - create
         - Setting >> repository
         - create 3 repositories
           - maven hosted zone
           - create repository
           - Type the name
           - create
           - maven 2 proxy
           - create repository
           - Type the name
           - Enter the maven repo url(https://repo1.maven.org/maven2/)
           - create
           - maven 2 group
           - create repository
           - Type the name
           - maven snapshot
           - create repository
           - Type the name
           - Change releast to snapshot
           - create
           - Under Available move the following
             - vpro-maven-central
             - maven-releases
             - vprofile-snapshot
           - create
       - Go to settings.xml which is located in .m2

   - Jenkins(Ubuntu)

     - Login into jenkins and cat the url to the password
     - complete installation
     - create a jenkins jo
       - type the nme
       - select freestyle
       - ok
       - under source code management
         - select git
         - enter the url
         - enter the branch (ci-jenkins)
         - Add build step >> Invoke top-level Maven targets
         - Goals (install -DskipTests)
         - settings file
           - Settings file in filesystem
           - file path (settings.xml)

   - Sonarqube (ubuntu)
   - for setup go to ci-jenkins on vprofile projects and copy the userdata setup
