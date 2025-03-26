import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class InfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);
    const defaultBucket = new cdk.aws_s3.Bucket(this, "ecommerce-public", {
      bucketName: "ecommerce-public",
      cors: [
        {
          allowedMethods: [
            cdk.aws_s3.HttpMethods.GET,
            cdk.aws_s3.HttpMethods.PUT,
          ],
          allowedOrigins: ["*"],
          allowedHeaders: ["*"],
        },
      ],
      encryption: cdk.aws_s3.BucketEncryption.S3_MANAGED,
      accessControl: cdk.aws_s3.BucketAccessControl.PUBLIC_READ,
      enforceSSL: true,
    });
    new cdk.aws_cloudfront.Distribution(this, "ecommerce-distribution", {
      defaultBehavior: {
        origin:
          cdk.aws_cloudfront_origins.S3BucketOrigin.withOriginAccessControl(
            defaultBucket,
          ),
      },
    });

    // The code that defines your stack goes here

    // example resource
    // const queue = new sqs.Queue(this, 'InfraQueue', {
    //   visibilityTimeout: cdk.Duration.seconds(300)
    // });
  }
}
