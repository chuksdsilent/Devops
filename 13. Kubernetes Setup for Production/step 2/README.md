# Kubeadm Setup for Cluster

Steps

1. Login to master & node with root user

2. Prerequisites on all nodes

3. Install Docker engine on all nodes

4. Install kubeadm, kubelet & kubectl

5. Execute kubeadm init --- on Master Node

6. Save node join command(kubeadm join...) returned by master node

7. Execute "kubeadm join" command returned by master setup, on Node

8. Execute user-setup script on Master node

9. Login to ubuntu user and run kubectl get node

Implementation

1. Open vprofile-project

2. switch to kubernetes-setup

```
git checkout kubernetes-setup
```

3. Open kubeadm folder

4. open Vagrantfile

5. Run the vagrant file

6. To destroy/delete vagrant

```
vagrant destroy
```

7.
