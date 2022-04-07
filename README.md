# webup

![Build Status](https://github.com/amila-ku/webup/actions/workflows/build.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/amila-ku/webup)](https://goreportcard.com/report/github.com/amila-ku/webup)

S3 is a very cost effective way to host static web content, but it requires multiple configuration changes to properly prepare a s3 bucket for web hosting. This is a simple cli tool to make it easier to prepare s3 bucket for static website hosting.

## What it does

1. Creates a s3 bucket for a given dns name(This should be the name you want to expose your content).
2. Enables s3 bucket webhosting and sets up default files as index.html and default error file as error.html.
3. Create DNS entries in route53 for the given dns name.
## How to create a s3 buckt ready for webhosting.

```
webup create -n www.testwebsite.devops.lk -z Z1TI4H711TUXXX
Starting to set up bucket
2022/04/05 21:52:13 Bucket www.testwebsite.devops.lk created 
2022/04/05 21:52:14 DNS www.testwebsite.devops.lk created 
Done setting up bucket
```

### How to upload content

```
webup upload -w  www.testwebsiteamilaku870306.devops.lk
```

output

```
2022/04/08 00:12:20  create s3 client new s3 upload
2022/04/08 00:12:20 Opened file: webcontent/index.html
2022/04/08 00:12:20 Trying to upload file: index.html to s3
2022/04/08 00:12:21 Upload file: index.html to s3, object version: 0xc000550020
```


## Functionality

- create s3 bucket
- configure for web hosting
- create r53 entries
- initialize directory upload content
- upload files to s3


# Todo

- Upload content

This currently only updates index.html file in webcontent folder and sets public access ACL in S3 bucket, it needs to be able to upload any file available in webcontent folder.

