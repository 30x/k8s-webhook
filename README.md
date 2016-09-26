# k8s-webhook

Sample https golang server that accepts Kubernetes webhook authorization request bodies. Serves on port 8081

Must have a `cert.pem` and `key.pem` in root directory to run. Use the following openSSL command to generate these files:

```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```

make sure you set the Common Name to `127.0.0.1:8081`