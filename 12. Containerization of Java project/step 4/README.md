Dockerfile App Image [Tomcat]

1. Go to vprofile-project

2. Switch branch to vprofile

3. Open docker-files/

   - App/Dockerfile
   - db/Dockerfile
   - web/Dockerfile
   -

4. Login to vagrant

5. change to directory vagrant and you will see vprofile-project there

```
cd vagrant
```

5. change to vprofile-project directory

```
cd vprofile-project
```

6. install jdk and maven

```
sudo apt install openjdk-8-jdk -y
sudo apt install maven
```

7. Open src/main/resources/application.properties and change the properties to connect to the backend

8. Build the artifact

```
mvn install
```

9. change to target directory

```
cd target
```

10. copy the artifact to app folder

```
cp -r target Docker-files/app/
```

11. Go to app directory

```
cd app
```

12. Build the docker images

```
docker build -t vprofile/vprofileapp:V1 .
```

13. To check

```
docker images
```

14. Go to db directory

```
cd db
```

15. Build the docker images

```
docker build -t vprofile/vprofiledb:V1 .
```

14. Go to web directory

```
cd web
```

15. Build the docker images

```
docker build -t vprofile/vprofileweb:V1 .
```

16. Pull nginx

```
docker pull nginx
```

17. Pull nginx

```
docker pull memcached
```

18. Pull rabbitmq

```
docker pull rabbitmq
```
