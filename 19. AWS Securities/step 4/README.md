# Inspector & Best Practices

## Secure Your Operating Systems and Applications

1. Disable root API access keys and secret key

2. Restrict access to instances from limited IP ranges using Security Groups

3. Password protect the .pem file on user machines

4. Delete keys from authorized_keys file on your instances when someone leaves your organization or no longer requires access

5. Rotate credentials (DB, Access Keys)

6. Regularly run least priviledge checks using IAM user Access Advisor and IAM user Last Used Access Keys

7. Use bastion host to enforce control and availability

8. use Secure AMIs from

   - Center for Internet Security(CIS)
   - Internation Organization for Standardization (ISO)
   - SysAdmin Audit Network Security (SANS) Institute
   - National Institute of Standards Technology (NIST)

9. Patching
   - Bootstrapping
     - UserData
     - Ansible, Puppet, Chef
     - Cloudformation

## Vulnerability Assessment

1. AWS Inspector

   This is used to scan vulnerability of an operating system

### Steps

1. Login to AWS

2. Install agent on your operating system

```
wget https://inspector-agent.amazonaws.com/linux/latest/install

bash install

```

To Check if it is installed

```
/etc/init.d/awsagen status
```

3.  Add some tag to your instance

4.  Go to amazon inspect

5.  Get Started

6.  Check mark network assessment and host assessment

7.  Advance setup

8.  Enter the name

9.  Uncheck all instances >> select the tags attached to your instance

10. Uncheck install agent >> next

11. Enter name >> select

    - security best practices
    - Common Vulnerabilities and exposures
    - Network Reachablility
    - CIS Operating System Security Configuration Benchmarks

12. Next >> Create

13.
