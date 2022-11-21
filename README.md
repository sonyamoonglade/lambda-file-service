Serverless lambda-function based on Yandex Cloud Functions.

Made for file uploading, deleting, soft-upload (upload and delete previous version of file).

Uses self-written s3-yandex-go library.

Requires several environment variables to work.
- OWNER (yandex s3 bucket owner)
- BUCKET (yandex s3 bucket name)
- AWS_SECRET_KEY_ID (aws credential)
- AWS_SECRET_ACCESS_KEY (aws credential)

