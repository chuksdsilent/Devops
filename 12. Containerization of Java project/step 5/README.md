# Docker Compose

1. Go to docker documentation and install docker compose

2. Go to vprofile-project/docker-compose.yml

3. To inspect an image

```
docker inspect vprofile/vprofiledb:V1
```

4. Run docker composer

```
docker-compose up
```

5. To push all images to docker hub

```
docker login
docker push vprofile/vprofileapp:V1
docker push vprofile/vprofiledb:V1
docker push vprofile/vprofileweb:V1
```
