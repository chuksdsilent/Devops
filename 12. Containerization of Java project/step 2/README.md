# Setup Docker Engine

1. Go to gitbash

2. create a vagrant file

3. open the vagrant file

4. Uncomment private_network and change the ip address

5. vagrant up >> vagrant ssh

6. Go to docker documentation and install docker on the vagrant.

7. Add vagrant user into sudoers

```
sudo usermod -a -G docker vagrant

```

7. Exit and Login

8. check status of docker

```
sudo systemctl status docker
```
