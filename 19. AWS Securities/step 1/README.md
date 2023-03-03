# Encrypt your data on cloud | Protecting Data at Rest

1. Encryption Keys

   - AWS managed Keys
   - KMS
   - KMS with custom keys
   - AWS CloudHSM

2. Data to encrypt
   - s3
   - EBS
   - RDS
   - Glacier
   - Dynamo db
   - EMR

## Creating encrypted data with KMS

### KMS with s3 bucket

1. Login to AWS

2. Create users in IAM user, give them AWS Management console access and do not give them any priviledge

3. Create 2 groups (one for admin and the other for testers)

   - Give admin adminstrator address
   - testers s3fullaccess, ReadOnlyAccess

4. Add users to group

5. Go to KMS

6. Create a key

7. Select symmetric encryption >> Advance option

   - select KMS

8. Type the name >> next

9. Select the admin user you created >> next

10. Select the account that can encrypt and decrypt >> next

11. Finish

12. Create a bucket

    - select the same region you have the key >>
    - server-side encryption(Enable)
    - Encryption key type(AWS KMS key)
    - Select the key

13. Create bucket.

14. Upload a text file

15. click on upload

16. check mark on acknowledge you dont have versioning >> upload

### Make the file public

17. open the file >> permissions >> Edit >> uncheck Block all public access >> make public

### KMS with EBS

1. Go to ec2 instance

2. Under Encryption select the Encryption

3.
