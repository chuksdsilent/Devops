## Deploy artifact to Tomcat Server

1. Go to manage jenkins >> manage plugins >> available >> type deploy to container >> install without restart

2. Go to new item >> Type the name >> select freestyle project >> ok

3. Under build select copy artifact from another project >> type the project name >> under artifact to copy type \*_/_.war

4. post build action >> deploy to container >>

5. under WAR/EAR files (target/project-name.war)

6. under container select version

7. elect username with password under kind >>type the username and password from the provisioning script >> add

8. Type the private ip of tomcat server >>under context path type any name(vprofileapp) >> save

9. Build Now.

10. Configure

11. Post build action

12. execute shell >> move it on to of the build action >> Tye the following command

```

cat << EOT > src/main/resources/application.properties

(copy application.properties file from profile project and paste)

EOT

```

Then change to ip address of the backend(mysql, memcache, rabbitmq)

16. Post build Action >> build other project >> vprofile-staging-deployment.

## Checking the url

1. Manage Jenkins >> Manage plugins >> available >> http request

2. New item >> project Name(Test-url)

3. Add build step >> execute shell >> Type the following command

```
sleep 60
```

4. Add build step >> http request >> type jenkins ipaddress and port(8080), checkmark if the build stable.

5. In the deployment project >> post build action >> build other project >> project-name(test-url)
