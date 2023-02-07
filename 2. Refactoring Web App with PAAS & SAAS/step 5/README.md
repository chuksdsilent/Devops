# Creating Elastic Beanstalk

## To Create Elastic Beanstalk

1. Go to Elastic Beanstalk

![Capture88](https://user-images.githubusercontent.com/18073289/216986125-f675d0bf-16bd-47ed-8784-79a54ec5aab1.PNG)

2. Type Name

![Capture89](https://user-images.githubusercontent.com/18073289/216986260-13f8d15a-4a73-448a-828c-6b6065934a92.PNG)

3. Select the platform (tomcat), select the version(8.5) and select sample application

![Capture90](https://user-images.githubusercontent.com/18073289/216986283-7b86cb69-4abe-404b-8c32-e346cd24de8a.PNG)

4. select configure more option

### Click on edit Instances

5.  select Genderal purpose, type size and select security group

![Capture92](https://user-images.githubusercontent.com/18073289/216986713-73b77cd6-aa4f-43d5-9795-d637ef61b573.PNG)

6. save

### Click on edit Capacity

7. select load balanced under Env. Type the min and max, instance type(t2.micro), AZ(ANY) and Scaling trigger(Network out is most popular)

![Capture93](https://user-images.githubusercontent.com/18073289/216987571-2ab6f26c-f96f-495e-99f3-cfb521bb3d42.PNG)

### Click on edit Rolling update and deployment

8. Select deployment policy(Rolling and batch size =50%)

![Capture95](https://user-images.githubusercontent.com/18073289/216988700-540bfd85-5da9-42d7-b506-7a47f2e3e512.PNG)

### Click on edit Security

9. Select key pair >> save

### Click on Enhance Monitoring

10. Basic for test, Enhanced for production

11. Tick Stream the logs

12. Delete after usage >> save

13. Create App

14 Go to the environment and open it

![Capture97](https://user-images.githubusercontent.com/18073289/216990516-e87fe76f-f5b0-4605-b936-323f4e8b6107.PNG)

15. Click on the environment name

![Capture98](https://user-images.githubusercontent.com/18073289/216990872-bd62342c-394c-403e-aa66-5c4368df673f.PNG)

16. Click on edit load balancer -> Add listeners

![Capture99](https://user-images.githubusercontent.com/18073289/216990893-388c0b13-4119-464b-af75-0b795263ccce.PNG)

17. change port -> 443, protocol -> https and ssl -> acm ssl

![Capture101](https://user-images.githubusercontent.com/18073289/216991193-b56f0e5b-b8f7-4fd2-a8a4-5e1c9c270f98.PNG)

18. Click on add process, Change path to /login

![Capture102](https://user-images.githubusercontent.com/18073289/216991469-b24780af-2477-4d4b-ab2e-c6ead02e69dc.PNG)

19. save

20. Click on application versions

![Capture104](https://user-images.githubusercontent.com/18073289/216991792-f3f6b0e9-4503-4ef8-9a44-99d627fac430.PNG)

21. Click on upload

![Capture106](https://user-images.githubusercontent.com/18073289/216991952-3b3f1ba5-e550-4e9d-af08-a3da7e6d255a.PNG)

22. Clcik on action >> Deploy

![Capture108](https://user-images.githubusercontent.com/18073289/216992156-1c7e7e03-4e67-49a9-be12-0122682f5b7c.PNG)

> > > > > > > 11ac3228c797b1b294ec132e0f8636d53a3b10b9
