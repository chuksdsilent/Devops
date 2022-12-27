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
