# Manage Security Monitoring, Alerting, audit Trail, and Incident Response.

1. Cloudwatch

   - Logs, Metrics & Events

2. CloudTrails

   - Record all API calls and save it to s3 encrypted(KMS)
   - S3 & Lambda Events separately

3. AWS Config
   - Builds config info of your account to track any changes
   - Rules for config change and action based on evaluation

## CloudTrail

1. Login to AWS

2. Go to Cloud trails

3. Create trail

4. Type the name

5. Select

   - All region(yes)
   - apply to organization(No)
   - read/write events(all)
   - log aws kms events(yes)
   - create a new bucket(Yes)
   - encrypt with kms(yes)
   - Enable log file validate(yes)
   - send sns notification (yes)

6. Create

## AWS Config

1. Go to AWS

2. Go to config >> Get started >>next

3. search ssh and select it >> Next >> Create
