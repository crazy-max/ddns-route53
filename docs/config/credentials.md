# Credentials configuration

AWS credentials are required to access and manage your Route 53 hosted zone.

!!! note
    The `credentials` section is optional. If it is not set, ddns-route53 uses
    the AWS SDK default credential chain. This includes AWS environment variables,
    the shared AWS config files (`~/.aws/credentials` and `~/.aws/config`),
    and IAM roles such as EC2 instance roles, ECS task roles, or IRSA in Kubernetes.

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
