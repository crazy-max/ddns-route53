# Prerequisites

I assume you have already created a [Route 53 Hosted Zones](https://console.aws.amazon.com/route53/home#hosted-zones:)
as a **Public Hosted Zone** type and setted Amazon name servers in your domain name registrar.

Go to the [IAM Policies page](https://console.aws.amazon.com/iam/home#/policies) and click on **Create Policy**.

Then click on **JSON** tab and paste the following content :

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "route53:ChangeResourceRecordSets"
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

Now go to the [IAM Users page](https://console.aws.amazon.com/iam/home#/users) and click the **Add user** button.

Enter a **User name**, check **Programmatic access** for _Access type_ and click **Next: Permissions**.

Choose the last option **Attach existing policies directly** and fill in the **Search** field with the name of the
policy you created before and click **Next: Review** then **Create user**.

An _Access Key ID_ and a _Secret Access key_ will be displayed. This is the [credentials](../config/credentials.md)
needed for ddns-route53. Save them somewhere since you will need them in the
[configuration](../config/index.md) step.
