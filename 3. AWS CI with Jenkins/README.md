# Continuous Integration

## Problems

1. In An Agile SDLC, there will be frequent code changes.

2. Not so frequently code will be tested

3. Which accumulates bugs and error in the code.

4. Developers need to reqork to fix these bugs and error

5. Mual Build & Releases

## Solutions

1. Build & test for every commit.

2. Automated Process

3. Notify for every build status.

4. Fix code if bugs or error found instantly rather than waiting.

## Benefits

1. Short MTTR

2. CI Pipeline

3. Agile

4. No Human intervention

5. Fault Isolation

## Tools

1. Jenkins

2. Git(Github)

3. Maven

4. Checkstyle

5. Slack

6. Nexus

7. Sonarqube

8. AWS EC2

## Objectives

1. Fault isolation

2. Short MTTR

3. Fast turn around on feature changes

4. Less disruptive

## Flow of Execution

1. Login to aws account

2. Create Login key

3. Create SG

   - Jenkins
   - Nexus
   - Sonar

4. Create Ec2 Instances with userdata

   - Jenkins
   - Nexus
   - Sonar

5. Jenkins Post Installation

6. Nexus Repository setup

7. Sonarqube post installation

8. jenkins Steps

   - Build Job
   - Setup Slack Notification
   - Checkstyle code analysis job
   - setup sonar integration
   - sonar code analysis job
   - artifact upload job

9. Connect all jobs together with BuildPipeline

10 Set automatic build trigger

11. Test with intellij

12 Test
