# Credentials configuration

AWS credentials are required for getting access and manage your Route53 zone.

!!! note
    When running ddns-route53 on an EC2 instance or in a Kubernetes cluster configured
    with IAM roles for service accounts (IRSA) AWS credentials are not required.
    ddns-route53 will automatically detect the presence of the configured IAM role
    and use it to authenticate with Route 53.

```yaml
credentials:
  accessKeyID: "ABCDEFGHIJKLMNO123456"
  secretAccessKey: "abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH"
```

## `accessKeyID`

AWS Access Key.

!!! example "Config file"
    ```yaml
    credentials:
      accessKeyID: "ABCDEFGHIJKLMNO123456"
    ```

!!! abstract "Environment variables"
    * `DDNSR53_CREDENTIALS_ACCESSKEYID`
    * `AWS_ACCESS_KEY_ID` _(through AWS env provider)_
    * `AWS_ACCESS_KEY` _(through AWS env provider)_

## `accessKeyIDFile`

Use content of secret file as AWS Access Key if `accessKeyID` not defined.

!!! example "Config file"
    ```yaml
    credentials:
      accessKeyIDFile: /run/secrets/akid
    ```

!!! abstract "Environment variables"
    * `DDNSR53_CREDENTIALS_ACCESSKEYIDFILE`

## `secretAccessKey`

AWS Secret Key.

!!! example "Config file"
    ```yaml
    credentials:
      secretAccessKey: "abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH"
    ```

!!! abstract "Environment variables"
    * `DDNSR53_CREDENTIALS_SECRETACCESSKEY`
    * `AWS_SECRET_ACCESS_KEY` _(through AWS env provider)_
    * `AWS_SECRET_KEY` _(through AWS env provider)_

## `secretAccessKeyFile`

Use content of secret file as AWS Secret Key if `secretAccessKey` not defined.

!!! example "Config file"
    ```yaml
    credentials:
      secretAccessKeyFile: /run/secrets/sak
    ```

!!! abstract "Environment variables"
    * `DDNSR53_CREDENTIALS_SECRETACCESSKEYFILE`
