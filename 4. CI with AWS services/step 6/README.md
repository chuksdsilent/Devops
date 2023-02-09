# AWSCodePipeline & Notifications

1. Go to SNS >> Topics

2. Create topic

3. type the name >> create topic

4. Create subscription

5. select email and type Email

6. create subscription

7. Go to email >> confirm subscription

8. Go to pipeline >> create pipeline

9. type the name >> next

10. Create a new role

11. select source provider(codecommit)

12. select our repository and branch(ci-aws)

13. select amazon cloudwatch events >> next

14. Build provider(codebuild) 15. select the project >>

15. deploy(aws s3) >> Next >> Review >> create pipeline
16. stop execution 18. you can click on edit to add other

### For Test

21. Add Stage

22. Type Name(test)

23. Add action group

24. Type name

25. action provider(aws codebuild)

26. Project name

27. input artifact(source artifact)

28. Single build >> done

### For Deploy

29. Add Stage

30. Type the Name(deploy)

31. Add action group

32. Type name

33. action provider(s3 bucket)

34. create s3 bucket(in the same region)

35. create a folder inside the s3 bucket

36. source artifact (build Artifact)

37. select the bucket name

38. enter the folder name

39. select extract file before deploy >> done

40. Set the notifications

41 Codepipeline >> pipeline settings >> Notification >>create Notification rule 42. name

43. select the action

44. select target type (SNS topic)

45. choose topic >> submit
