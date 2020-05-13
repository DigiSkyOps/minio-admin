

# minio-admin

## Description
Minio-admin is a management tool for minio users and policies, which has been changed from the native command line mode of minio to a simple and convenient UI interface. The project is divided into front and rear end projects, front end vue, and back end golang gin.
  
## Function Introduction
- Users add, delete, disable, bind policy, batch delete users, batch disable users.
- Policies are added, deleted, modified, and bulk deleted. When adding policies, it is supported to customize the policy content in json format and generate the policy content by clicking on the UI interface.


## Project Structure
- Fed: minio-admin-fed。
- Platform: minio-admin-platform。


## Deploy
Container deployment is used by default and can be deployed directly on kubernetes. Here is a brief introduction to how docker run can be deployed.
- Build and Deploy: minio-admin-platform
```
$ cd minio-admin-platform

# Replace minio's ACCESSKEY, SECRETKEY, MINIO_ENDPOINT variables in the following command
$ docker build -t minio-admin-platform --build-arg ACCESSKEY="xxx" --build-arg SECRETKEY="xxx"  --build-arg MINIO_ENDPOINT="xxx"  .

# The local port mapped here is 9000
$ docker run -itd --name minio-admin-platform -p  9000:9000 minio-admin-platform 
```
- Build and Deploy:  minio-admin-fed
```
$ cd minio-admin-fed
# Replace the value of the PLATFORM_URL variable in the following command with the address of the deployed minio-admin-platform in the previous step
$ docker build -t minio-admin-fed  --build-arg PLATFORM_URL="xxx" .

$ docker run -itd --name minio-admin-fed -p 8000:8000   minio-admin-fed 
```
- Open Browser
```

Use your browser to open localhost8000

```

## Referenced Document

- minio ：https://min.io/
- github：https://github.com/minio/


## Licensing
minio-admin is under the Apache 2.0 license.