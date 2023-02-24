# Spin k8 Cluster

1. Login to your cluster

```
kops create cluster --name=kubevpro.groophy.in --state=s3://bucket-name --zones=us-east-2a,us-east-2b --node-count=2 --node-size=t2.micro --master-size=t2.micro --dns-zone=kubevprofile.groophy.in

```

2. update the cluster

```
kops update cluster --name kubevpro.groophy.in --state=s3://bucket-name --yes
```

3. Validate the cluster

```
kops validate cluster --state=s3://bucket-name
```

4. Create github repo >> Clone the repo

5. create EBS Volume aws ec2 create-volume --availability-zone=use-east-2a --size=3 --volume-type=gp2

6. Copy out the volumeId

Note: Make sure the DB pod runs in the same AZ as the EBS volume.

7. To display labels

```
kubectl get nodes --show-labels
```

8. To create labels

```
kubectl get nodes

kubectl describe node node-name to check the region of the nodes on the cluster

kubectl label nodes node-name zone=us-east-2a
```

9. To see the definition files go to vprofile-project > kube-app branch >> kubernetes >> vpro-app

10. Encode database and rabbitMQ password

```
echo -n "password" | base64
```

11. To deploy a deployment manifest on a pod

```
kubectl create -f app-secret.yaml
```

12. To see the secret content deployed

```
kubectl get secret
```

13. To see more information on a pod

```
kubectl describe secret secret-name
```
