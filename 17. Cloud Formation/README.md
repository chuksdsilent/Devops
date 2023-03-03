# Cloud Formation

Cloud Formation is infrastructure as a code [IAAC] just like terraform but it is specific to AWS but terraform can be used in other platforms.

## Terminologies in Cloud Formation

1. Template is a script in cloud formation but it is written in YAML or JSON. Cloud Formation reads the file as input and the output will be the resource that will be created.

2. Stack: The resources created from template is caled the stack. Stack is a single unit of set of resources

3. Change Set allows you to see how changes will impact your existing resources.

## Template Anatomy

```
{
    "AWSTemplateFormatVersion": "version date",
    "Description" : "String",
    "Metadata" : "template metadata",
    "Parameters" : { set of parameters},
    "Mappings" : {set of mappings},
    "conditions" : {set of conditions},
    "Transform" : {set of transforms},
    "Resources": {set of resources},
    "Outputs" : {set of outputs}
}
```

Resource (JSON)

```
{
    "Resources": {
        "Logical ID" : {
            "Type" : "Resource type",
            "Properties", : {
                ... set of properties ...
            }
        }
    }
}
```
