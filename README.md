# Description
**badbucket** is a quick-n-dirty tool to test the configuration of your S3 bucket. It checks for the following:

* Is the bucket listable?
* Are files readable?
* Is the bucket writeable?

**badbucket** is written in Go and uses the [AWS SDK](https://aws.amazon.com/sdk-for-go/)

# Installation
```
go install github.com/zuBux/badbucket/cmd/...
```
# Usage

```
badbucket -b <bucket_name> -r <AWS region>
```

# Coming up next

* ~~Implement writable check~~
* Parse list of s3 buckets
* Automated region detection
* Concurrent checks
