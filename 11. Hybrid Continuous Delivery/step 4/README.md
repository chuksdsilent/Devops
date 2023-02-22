# Software Testing with selenium

1. Login to jenkins >> manage jenkins >> manage plugin

2. search for codebuildplugin

3. Go to aws >> codebuild >> software testing repo

4. Go to edit >> source >> change the branch to cd-aws-jenkins

5. save

6. Edit >> Buildspecfile

7. Go to vprofile-project >> switch to cd-aws-jenkins >> open aws_files/softTest_buildspec.yml

8. Copy and paste in buildspec.yaml box >> save

9. Go to Edit >> Artifacts >> change the bucket name

10. Under Artifact packaging >> Select Zip Format

11. Update Artifact

12. Edit Environment >> save

13. Go to jenkins >> new item >> project Name (codeBuild-SoftwareTesting) >> select freestyle project >> Ok

14. Add Build step >> AWS CodeBuild

    - check mark credentials from jenkins >> Add Credentials
      - under kind select code build credentials
      - type id
      - aws access and secret key
    - Use project source code

15.
