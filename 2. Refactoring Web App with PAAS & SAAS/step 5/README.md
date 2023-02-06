# Creating Elastic Beanstalk

## To Create Elastic Beanstalk

1. Go to Elastic Beanstalk
2. Type Name
3. Select the platform (tomcat)
4. select the version(8.5)
5. select sample application
6. select configure more option
7. instances
8. select the security group
9. save
10. Capacity
11. select load balancer under Env. Type
12. instances( select min and max)
13. instance type(t2.micro)
14. AZ(ANY)
15. Scaling trigger(Network out is most popular)
16. Load balancer
17. select application >> save
18. Deployment
19. Deployment Policy (Rolling)
20. Batch Size(50%) >> Save
21. Security
22. Select key pair >> save
23. Enhance Monitoring
24. Basic for test, Enhanced for production
25. Tick Stream the logs
26. Delete after usage >> save
27. Create App
