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


## Functionality

- create s3 bucket
- configure for web hosting
- create r53 entries
- initialize directory upload content
- upload files to s3
