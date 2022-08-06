Feature of yandex-cloud function is that if you receive a binary body (octet-stream),
In unmarshalled Request you will eventually see byte[] slice.
But if you want to send it back as a Response, you will notice that yandex has automatically formatted your body to base64.
Considering this, as body receives thus passed to s3 client.