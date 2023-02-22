# Ansible On Jenkins

1. ssh into ansible server

2. change to root user

```
sudo -i
```

3. Install Ansible

```
apt update

sudo apt install software-properties-common

sudo apt-add-repository --yes --update ppa:ansible/ansible

sudo apt install ansible
```

4. Login to Ansible

5. Go to manage Jenkins

6. Go to Manage plugins

7. Go to available

8. Search for ansible >> checkmark ansible >> install without restart

9. Go to Dashboard >> New item

10. Enter name(deploy to staging)

11. choose freestyle

12. copy from build job >> ok

13. change the branch to /cd-ansible/jenkins

14. Uncheck poll SCM

15. Remove execute shell, invoke top level maven and Build other projects. >> save >> Build Now.

16. Go to Workspace >> ansible >> ansible.cfg

17. Go to configure

18. Go to invoke ansible playbook

19. Under invoke ansible playbook

    - Name(ansible/site.yml)
    - Inventory >> select inline content (app01-staging ansible_host=app_server_ip)
    - credentials
      - under kind (username with private key)
      - username(ubuntu)
      - upload private key

20. Under advanced

    - checkmark Double the host SSH Key check

21. click on add extra variables

    - USER( nexus username)
    - PASS(nexus password)- check mark hidden variable in in the logs
    - nexusip(Nexus server private ip)
    - reponame(nexusrepo)
    - groupid()
    - time
    - BUILD
    - $ID
    - vprofile_version

22. Under project check mark this project is parameterized >> Add parameter
    - Name(TIME)
23. Add another parameter

    - Name(ID)

24. Save

25. Allow app server from nexus security group

26. Build Now

27. You can change all the branches to cd-ansible-jenkins

28. Go to Build job >> workspace src/main/resource/application.properties

29. Go to configure >> add build step >> execute shell

```
cat << EOT > src/main/resources/application.properties

paste the application.properties

```

30. copy the be-01-staging-vprofile private ip and replace on application.properties

31. save

32. Go to code analysis >> configure

33. Go to report violation >> inclure the second and third box to 1000. >> Save.

34. Go to sonarqube server and copy the ip

35. Login >> open the project >> project settings >> Quality Gates >> select vprofile-quality-gates >> click on quality gates ate navbar >> click on quality gates >> edit >> change the value to 1000.

36 Save

37. Go to Nexus >> Add post build action >> Trigger parameterized build and other projects

    - type the name of the job
    - Add Parameters
      - TIME($buildTTIMESTAMP)
      - ID $BUILD_ID

38. Create a view for it.
