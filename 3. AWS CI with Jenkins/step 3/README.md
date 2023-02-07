# Creating Jenkins Job

1. Open jenkins

2. Create a new Job

3. Enter the name and select freestyle >> ok

4. Under source code Management select git >> type the git url and branch(ci-jenkins)

5. Add your credentials

6. Under build select top level invoke target

7. Under goal type

```
install -DskipTest
```

8. click on advance

9. under Settings
