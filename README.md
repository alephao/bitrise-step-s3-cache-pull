# S3 Cache Pull

A bitrise step to download your cache from a s3 bucket using custom keys with fallback.

Should be used with [S3 Cache Push](https://github.com/alephao/bitrise-step-s3-cache-push)

### Inputs

- **aws_access_key_id**: Your aws access key id. You can set an environment var `AWS_ACCESS_KEY_ID` instead of using this input.
- **aws_secret_access_key**: Your aws secret access key. You can set an environment var `AWS_SECRET_ACCESS_KEY` instead of using this input. 
- **aws_region**: The region of your S3 bucket. E.g.: `us-east-1 `. You can set an environment var `AWS_S3_REGION` instead of using this input.
- **bucket_name**: The name of your S3 bucket. E.g.: `mybucket`. You can set an environment var `S3_BUCKET_NAME` instead of using this input.
- **restore_keys**: The list of key with fallbacks to restore the cache. You can use `{{ checksum "path/to/file" }}` to use the checksum of a file as part of the cache key. E.g.:
```
carthage-{{ checksum "Cartfile.resolved" }}
carthage-
```
- **path**: Path to extract the file or directory cached. For instance, if you used [S3 Cache Push](https://github.com/alephao/bitrise-step-s3-cache-push) with the path `./Carthage` then this value should be `./`.