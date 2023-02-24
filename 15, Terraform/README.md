# Terraform

1. Create a file with extenstion of .tf

2. While in the same folder run the command

```
terraform init
```

3. To validate your terraform files

```
terraform validate
```

4. To format your terraform file

```
terraform fmt
```

5. To preview what will happen if you apply

```
terraform plan
```

6. To apply

```
terraform apply
```

7. To destroy

```
terraform destroy
```

To maintain the state of terraform

1. Create s3 bucket

2. create a folder inside the bucket

3. Write the following code

```
terraform {
    backend "s3" {
        bucket = "terra-state-dove"
        key = "terraform/backend"
        region = "us-east-2
    }
}
```
