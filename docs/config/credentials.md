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

AWS access key ID.

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

Use the contents of a secret file as the AWS access key ID if `accessKeyID` is not defined.

!!! example "Config file"
    ```yaml
    credentials:
      accessKeyIDFile: /run/secrets/akid
    ```

!!! abstract "Environment variables"
    * `DDNSR53_CREDENTIALS_ACCESSKEYIDFILE`

## `secretAccessKey`

AWS secret access key.

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

Use the contents of a secret file as the AWS secret access key if `secretAccessKey` is not defined.

!!! example "Config file"
    ```yaml
    credentials:
      secretAccessKeyFile: /run/secrets/sak
    ```

!!! abstract "Environment variables"
    * `DDNSR53_CREDENTIALS_SECRETACCESSKEYFILE`
