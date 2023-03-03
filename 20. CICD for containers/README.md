# CICD For Containers

## Problems

1. Operation team incharge of managing containers gets continuous deployment requests

2. Manual Deployment creates dependency

## Solutions

1. Automate Build & Release Processes

2. Build Docker Images & Deploy Continuously as fast as the code commits.

3. Continuous Deployment Process

## Tools

1. Kubernetes

2. Docker

3. jenkins

4. Docker hub

5. Helm

6. Git

7. Maven

8. Sonarqube

## Objectives

1. Continuous Integration and continuous delivery for containers

## Flow of Execution

1. Continuous Integration Setup

   - Jenkins, Sonarqube & nexus (continuous Integration Project)

2. Dockerhub Account (Containerization Project)

3. Store Dockerhub credentials in Jenkins

4. Setup Docker Engine in jenkins

5. Install Plugin in Jenkins

   - Docker-pipeline
   - Docker
   - Pipeline utitlity

6. Create Kubernets Cluster with Kops

7. Install Helm in Kops VM

8. Create Helm Chart

9. Add Kops VM as Jenkins Slave

10. Create Pipeline code [Delaravive]

12 Update Git Repository with - Helm Charts - Dockerfile - Jenkinsfile

13. Create Jenkins job for pipeline

14. Run & Test them
