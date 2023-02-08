# Upload Artifact to Nexus Server

1. Go to Jenkins >> Manage Jenkins >> Manage plugin >> available >> search for copy artifact, zentimestamp and nexus >> check mark them >> install withouth restart

2. Go to Jenkins >> New Item >> Type the job name(deploy_to_nexus)

3. Under build >> Add build step >> select copy artifact from another project.

4. Type the job name(build)

5. Artifact to copy(\*_/_.war)

6. Add another build step >> nexus artifact uploader
   - version(Nexus 3)
   - protocol(http)
   - nexus url(private ip)
   - credentials(username and password), id >> Add
   - GroupID(QA)
   - version(V$BUILD_ID)
   - Repository(vprofile-release)
   - Artifact(Add)
     - ArtifactId($BUILD_TIMESTAMP)
     - Type(war)
     - target/vprofile-v2.war
7. Scroll up >> check mark change patter for the BUILD_TIMESTAMP(yy-MM-dd_HH:mm)

8. Under Post build action >> Slack Notification >> Type the Organization name and channel name(vprofile-jenkins)

9. Save

10. connect the step 5 to step 6

## To create a beautiful view

1. Go to jenkins >> manage Jenkins >> Manage plugin

2. search for build pipeline >> install without restart

3. Go to Jenkins

4. Click on the plus sign(vprofile continuous integration)

5. select build pipeline view >> ok

6. Under Pipeline flow >> select build project

7. under display option >> select 5 in No of Display Builds

8. save

9. Run
