# Prerequisites

I assume you have already created a [Route 53 Hosted Zones](https://console.aws.amazon.com/route53/home#hosted-zones:)
as a **Public Hosted Zone** type and setted Amazon name servers in your domain
name registrar.

!!! note
    If you are using [EC2 IAM instance roles](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html)
    or [AWS IAM roles for service accounts (IRSA)](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html),
    creating a user is not required. Configuration of IAM roles is beyond the scope of this
    document, and is discussed in the official AWS documentation linked above.
    For all other use cases, follow the steps below to create a policy and user.

Go to the [IAM Policies page](https://console.aws.amazon.com/iam/home#/policies)
and click on **Create Policy**.

Then click on **JSON** tab and paste the following content:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "route53:ChangeResourceRecordSets",
                "route53:ListResourceRecordSets"
            ],
            "Effect": "Allow",
            "Resource": "arn:aws:route53:::hostedzone/<HOSTED_ZONE_ID>"
        }
    ]
}
```

!!! tip
    Replace `<HOSTED_ZONE_ID>` with your current hosted zone id you can find on
    [Route 53 Hosted Zones page](https://console.aws.amazon.com/route53/home#hosted-zones:).

Enter a **Policy Name** and click **Create Policy**.

Now go to the [IAM Users page](https://console.aws.amazon.com/iam/home#/users)
and click the **Add user** button.

Enter a **User name**, check **Programmatic access** for _Access type_ and
click **Next: Permissions**.

Choose the last option **Attach existing policies directly** and fill in the
**Search** field with the name of the policy you created before and click
**Next: Review** then **Create user**.

An _Access Key ID_ and a _Secret Access key_ will be displayed. This is the [credentials](../config/credentials.md)
needed for ddns-route53. Save them somewhere since you will need them in the
[configuration](../config/index.md) step.
