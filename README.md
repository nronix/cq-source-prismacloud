# cq-source-prismacloud
 
Sample Config.yaml
```
kind: source
spec:
  name: prisma
  registry: local
  path: /pathto/cq-source-prismacloud
  version: v0.0.1
  destinations: ["file"]
  tables: ["prismacloud_alerts"]
  spec:
    prisma:
      - ENDPOINT: "api2.prismacloud.io"
        ACCOUNT: "Prisma Cloud"
        API_KEY: ""
        API_SECRET: ""
        ALERT_POLICIES: ["AWS S3 bucket policy does not enforce HTTPS request only", "AWS Access logging not enabled on S3 buckets","AWS S3 Object Versioning is disabled"]
        INV_QUERY_STRING: "config from cloud.resource where api.name = 'aws-ec2-describe-instances'"
      - ENDPOINT: "api.gov.prismacloud.io"
        ACCOUNT: "PRISMA Cloud GOV"
        API_KEY: ""
        API_SECRET: ""
        ALERT_POLICIES: [ "AWS S3 bucket policy does not enforce HTTPS request only", "AWS Access logging not enabled on S3 buckets","AWS S3 Object Versioning is disabled" ]
        INV_QUERY_STRING: "config from cloud.resource where api.name = 'aws-ec2-describe-instances'"

```
