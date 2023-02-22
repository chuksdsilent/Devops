# Deploy to Beanstalk

1. Go to manage jenkins >> manage plugins >> available

2. search for beanstalk >> install without restart

3. new item >> item name (Deploy-To-Staging-BeanstalkEnv)

4. Check martk change date pattern for the BUILD_TIMESTAMP (build timestamp) variable

5. Under Date and Time Pattern

   - yy-MM-ddH_Hmm

6. Add Build step >> copy artifacts from another project

   - project name(Build job project)
   - Artifacts to copy (\*_/_.war)

7. Add build step >> AWS Elastic Beanstalk

   - Credentials
     - under kind select aws credentials
     - id (Any name)
     - access key and secret key
   - enter region
   - Number of attempts (30)
   - validate credentials

8. Under Application and Environment

   - Application Name (Copy application found name)
   - Environment Name (Beanstalk project name)

9. Under packaging

   - Root Object (File/Directory) - target/vprofile-v2.war

10. Under Versioning

    - Version Label Format (vprofile-V$BUILD_TIMESTAMP)

11. Select the notifications

12. Advanced

    - channel Name
    - credentials

13. Save
