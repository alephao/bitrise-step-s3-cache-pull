# S3 Cache Pull

A bitrise step to download your cache from a s3 bucket using custom keys with fallback.

Should be used with [S3 Cache Push](https://github.com/alephao/bitrise-step-s3-cache-push)

### Inputs

<table>
    <thead>
        <tr>
            <th>Input</th>
            <th>Environment Var</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>
                <b>aws_access_key_id</b>
            </td>
            <td>
                <tt>AWS_ACCESS_KEY_ID</tt>
            </td>
            <td>
                Your aws access key id
            </td>
        </tr>
        <tr>
            <td>
                <b>aws_secret_access_key</b>
            </td>
            <td>
                <tt>AWS_SECRET_ACCESS_KEY</tt>
            </td>
            <td>
                Your aws secret access key
            </td>
        </tr>
        <tr>
            <td>
                <b>aws_region</b>
            </td>
            <td>
                <tt>AWS_S3_REGION</tt>
            </td>
            <td>
                The region of your S3 bucket. E.g.: <tt>us-east-1</tt>
            </td>
        </tr>
        <tr>
            <td>
                <b>bucket_name</b>
            </td>
            <td>
                <tt>S3_BUCKET_NAME</tt>
            </td>
            <td>
                The name of your S3 bucket. E.g.: <tt>mybucket</tt>
            </td>
        </tr>
        <tr>
            <td>
                <b>restore_keys</b>
            </td>
            <td>-</td>
            <td>
                <span>The list of keys with fallbacks to restore the cache. E.g.:</span>
                <pre>
carthage-{{ branch }}-{{ checksum "Cartfile.resolved" }}
carthage-{{ branch }}
carthage-
                </pre>
            </td>
        </tr>
        <tr>
            <td>
                <b>path</b>
            </td>
            <td>-</td>
            <td>
                Path to extract the file or directory cached. For instance, if you used <a href="https://github.com/alephao/bitrise-step-s3-cache-push">S3 Cache Push</a> with the path <tt>./Carthage</tt> then this value should be <tt>./</tt>
            </td>
        </tr>
    </tbody>
</table>

#### Cache Key

The cache key can contain special values for convenience.

Value|Description
-|-
`{{ branch }}`|The current branch being built. It will use the `$BITRISE_GIT_BRANCH` environment var.
`{{ checksum "path/to/file" }}`|A SHA256 hash of the given file's contents. Good candidates are dependency manifests, such as `Gemfile.lock`, `Carthage.resolved`, and `Mintfile`.