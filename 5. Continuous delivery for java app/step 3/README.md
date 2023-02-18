# Software Testing with Selenium

1. Login to AWS

2. Go to ec2 instance

3. Select the microsoft windows server 209 base

4. select t2.small

5. copy the user data and past on the user data box

6. select key pair

7. SG(All traffic from jenkins SG, rdp from RDp)

8. Launch

9. Select the instance >> connect

10. RDP client >> password

11. upload the private key

12. decrypt the password >> copy and put on your stick note

13. Go to your local system >> start search for remote desktop connection

14. Type the ip/dns name of the windows server >> connect

15. Type the username(Administrator) and password

16. Select yes

17. copy jenkins private ip >> open the chrome browser >> type the jenkins private ip and port number

18. Go to manage jenkins >> configure system >> change jenkins location to private ip

19. Go to manage Jenkins >> manage Node and clouds

20. Add new Node >> Type the name of the node >> under remote root directory type c:\jenkins >> launch method(launch agent by connecting it to the master) >> custom workDir path is the same as remote root directory>> checkmart use WebSocket >> save

21. Go to manage jenkins >> Configure gloable security >> under agent >> select fixed(8090)

22. Go to manage Nodes

23. open the node >> click on agent.jar and download it.

24. Go to you local system. cd to downloads

25. Execute the command on the node page.

Note do not close the terminal

22. Go to test url project >> configure >> under http request type (tomcat-url:port-number/vproapp/login) >> save

23. Build Now

## To perform software testing

1. Go to New item

2. Enter name(selenium-test-suite)

3. Copy from build job.

4. change branch to selenuiumSuite >> save

5. Go to manage jenkins

6. manage node and cloud

7. Open the project >> configure >> under usage select only build jobs with label expressions matching the node >> save

8. Go to selenium job >> configure >> under label expression type the node-name >> save

9. Build Now.

10. Once the work space is created >> Go to selenium test project >> jenkins_data

11. open the data file >> Copy the command

12. Go to configure >> under build select execute batch command >> paste the command

13. Chang the public ip

14. Invoke top-level Maven targets >> under goals type clean test-DsureFire.SuiteXMLFiles=testng.xml

15. Go to test url project >> configure

16 Post Build Action >> Build Other projects >> type the project-name
