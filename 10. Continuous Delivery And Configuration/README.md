# Continuous Delivery And Configuration Management

# Problems

1. Involves task assignment/ticketing/approvals.

2. Different owners for Different Environment

3. Dependency on Ops & Build & Release Team.

# Solution

1. Build, test, deploy & Test for every commit.

2. Automated Deployment Process

3. Notification at every step in pipeline

4. Integrate Automation Tool

5. Remove/Minimize Human Intervention

6. Fix Code if bugs or error found instantly rather than waiting

## Tools

1. Jenkins

2. Git

3. Maven

4. Checkstyle

5. Slack

6. Nexus

7. Quality Gates

8. Ansible

9. Selenium

10. Windows Server

11. AWS EC2 Instances

## Objectives

1. Deployment Automation

2. Short MTTR

3. Fast turn around on features changes

4. Less Disruptive.

## Flow of Execution

1. Login to aws account

2. create Login key

3. Create SG

   - Jenkins
   - Nexus
   - Sonar

4. Create EC2 instances with user data

   - Jenkins
   - Sonarqube
   - Nexus

5. Jenkins Post Installation

6. Nexus Repositories

7. Sonarqube post installation

8. Jenkins Steps

   - Build Jobs
   - Setup Slack Notification
   - Checkstyle code analysis job
   - Setup Sonar integration
   - Sonar code analysis job
   - Artifact upload job

9. Connect All Jobs together with BuildPipeline

10. Set automatic build trigger

11. Test with git

12. Create SG

    - windows server
    - tomcat & backend servers

13. Create tomcat & backend server on ec2

14. Create Jenkins Job to Run Ansible Playbook

15. Deploy artifact to staging by Ansible

16. Add windows node as slave to jenkins

17. Create job to run software Test(selenium) from windows server

18. Deploy artifact to production tomcat server using Ansible

19. Connect all the jobs with build Pipeline

20. Test it by commiting code to github.
