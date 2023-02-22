# Launch Staging server

1. launch EC2 for windows server

2. Allow traffic from Jenkins also allow traffic windows in jenkins server

3. launch ec2 instance for app01-staging-vprofile

4. Create SG(port 22 from my ip, 8080 from anywhere, 22 from jenkins SG, 8080 from windows server).

5. create SG(port 22 from my ip, all traffic from tomcat)

6. Launch ec2 instance for vprofile-be-staging-sg

7. Go to vprofile project and switch to cd-ansible-jenkins

8. change userdata directory and cat backend.sh

9. Copy the content and use it to provision vprofile-be-staging-sg
