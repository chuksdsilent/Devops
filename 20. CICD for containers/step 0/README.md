# Helm chart & git repo Setup

1. clone vprofile project github repo

2. switch to cicd-kube

3. To create a helm chart

```
helm create chart-name
```

4. Change to template directory and delete all sample file

```
rm -rf *
```

5. Copy everything from kubernetes/vpro-app to helm template

```
cp ../../kubernetes/vpro-app/* helm/vprofilecharts/templates/
```

6. Open vproappdeploy and change image name to a variable.

7. Create a namespace

```
kubectl create namespace test
```

8. To read to helm chart and run it

```
helm install --namespace test vprofile-stack helm/vprofilechart --set appimage=image-url:version
```

9. To list all the helm

```
helm list
```

10. To see all

```
kubectl get all --namespace test
```
