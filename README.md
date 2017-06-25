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
badbucket -b <bucket_name>
```
badbucket will automatically detect the region of your bucket

# Coming up next

* ~~Implement writable check~~
* Parse list of s3 buckets
* ~~Automatic region detection~~
* Concurrent checks
