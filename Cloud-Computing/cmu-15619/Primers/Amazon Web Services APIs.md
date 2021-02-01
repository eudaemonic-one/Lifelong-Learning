# Amazon Web Services APIs

## Using Amazon Web Service APIs to create EC2 Instance

### APIs and SDK access for AWS

The Amazon Web Services SDK includes different kinds of API packages which developers can use to create and manage resources and applications running on AWS. This allows developers to programmatically automate management tasks on AWS. Almost all of the functionality in AWS is accessible via the AWS Command Line Interface (CLI), APIs or SDK.

## AWS SDK for Java

The [AWS SDK for Java](http://aws.amazon.com/sdk-for-java/) has been developed by Amazon and is a fully featured SDK.

[Video 1](https://youtu.be/6Ru_f9WVIno): Introduction to AWS SDKs and Java SDK Example

### API Access in AWS

Programmatically provision, manage and destroy AWS resources.

Entire AWS Suite is included in API access:

* EC2
* S3
* Elastic Map Reduce
* DynamoDB
* etc.

### Under the Hood

AWS is designed with a RESTful Web API.

Make requests to a URL Endpoint.

### Getting Started with AWS SDKs

* Find the appropriate SDKs to use
* Find the library within the SDK
  * S3 in Java: com.amazonaws.services.s3.AmazonS3
  * EC2 in Python: boto.ec2
* Write the Program
  * Initiate a connection to an AWS with your credentials
  * Use the API calls to create/manage/destroy resources
  * Close the AWS connection when done.

## Amazon EC2 Java API

```java
public class LaunchEC2Instance {
    private static final String AMI_ID         = "ami-cd0f5cb6";
    private static final String INSTANCE_TYPE  = "t2.micro";
    private static final String KEY_NAME       = "project1_test";
    private static final String SECURITY_GROUP = "MySecurityGroup";

    public static void main(String[] args) {
        /*
         * http://docs.aws.amazon.com/AWSJavaSDK/latest/javadoc/com/amazonaws/auth/DefaultAWSCredentialsProviderChain.html
         *
         * AWS credentials provider chain that looks for credentials in this order:
         *   1. Environment Variables - AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
         *   2. Java System Properties - aws.accessKeyId and aws.secretKey
         *   3. Credential profiles file at the default location (~/.aws/credentials) shared by all AWS SDKs and the AWS CLI
         *   4. Credentials delivered through the Amazon EC2 container service if AWS_CONTAINER_CREDENTIALS_RELATIVE_URI"
         *       environment variable is set and security manager has permission to access the variable
         *   5. Instance profile credentials delivered through the Amazon EC2 metadata service
         */
        AWSCredentialsProvider credentialsProvider = new DefaultAWSCredentialsProviderChain();

        // Create an Amazon EC2 Client
        AmazonEC2 ec2 = AmazonEC2ClientBuilder
                .standard()
                .withCredentials(credentialsProvider)
                .withRegion(Regions.US_EAST_1)
                .build();

        // Create a Run Instance Request
        RunInstancesRequest runInstancesRequest = new RunInstancesRequest()
                .withImageId(AMI_ID)
                .withInstanceType(INSTANCE_TYPE)
                .withMinCount(1)
                .withMaxCount(1)
                .withKeyName(KEY_NAME)
                .withSecurityGroups(SECURITY_GROUP);

        // Execute the Run Instance Request
        RunInstancesResult runInstancesResult = ec2.runInstances(runInstancesRequest);

        // Return the Object Reference of the Instance just Launched
        Instance instance = runInstancesResult.getReservation()
                .getInstances()
                .get(0);

        System.out.printf("Launched instance with Instance Id: [%s]!\n", instance.getInstanceId());

        ec2.shutdown();
    }
}
```

To run this sample, execute the following commands:

```text
$ mkdir aws_intro && cd aws_intro

$ wget -O AWS-Java-LaunchInstance.tar.gz https://s3.amazonaws.com/cmucc-public/aws-api/AWS-Java-LaunchInstance.tar.gz
...
2017-08-20 20:55:14 (155 MB/s) - ‘AWS-Java-LaunchInstance.tar.gz’ saved [1767/1767]

$ tar xfz AWS-Java-LaunchInstance.tar.gz

$ mvn compile && mvn exec:java -Dexec.mainClass="edu.cmu.cs.cloud.samples.aws.LaunchEC2Instance"
 ...
[INFO] --- exec-maven-plugin:1.6.0:java (default-cli) @ AWS-API-Samples ---
Launched instance with Instance Id: [i-0b8c7ed1909c63a90]!
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
...

$ ssh -i "project1_test.pem" ubuntu@ec2-51-221-179-191.compute-1.amazonaws.com
Welcome to Ubuntu 16.04.2 LTS (GNU/Linux 4.4.0-1022-aws x86_64)

ubuntu@ip-172-31-9-237:~$
```

## AWS SDK for Python (Boto3)

For Python developers, AWS supports a third-party SDK called [boto3](https://boto3.readthedocs.io/en/latest/), which can be used to make API requests to AWS from within Python.

[Video 2](https://youtu.be/7IOsOHJKxvY): AWS Boto3 SDK for python

## Amazon EC2 Python API

```python

import boto3

# Refer to the Boto3 documentation:
#    http://boto3.readthedocs.io/en/latest/guide/quickstart.html
#
# Your AWS credentials must be configured in accordance with:
#    http://boto3.readthedocs.io/en/latest/guide/configuration.html

IMAGE_ID = 'ami-cd0f5cb6'
INSTANCE_TYPE = 't2.micro'
KEY_NAME = 'project1_test'
SECURITY_GROUP = 'MySecurityGroup'

# Create an EC2 Client
ec2_client = boto3.client("ec2",
                          region_name="us-east-1")

# Launching instance
#
# http://boto3.readthedocs.io/en/latest/reference/services/ec2.html#EC2.Client.run_instances
response = ec2_client.run_instances(
    ImageId=IMAGE_ID,
    InstanceType=INSTANCE_TYPE,
    KeyName=KEY_NAME,
    MaxCount=1,
    MinCount=1,
    SecurityGroups=[
        SECURITY_GROUP,
    ]
)

instance = response.get('Instances')[0]

print("Launched instance with Instance Id: [{}]!".format(instance.get('InstanceId')))
```

To run this sample, execute the following steps:

```text
$ wget -O AWS-Python-LaunchInstance.tar.gz https://s3.amazonaws.com/cmucc-public/aws-api/AWS-Python-LaunchInstance.tar.gz
...
2017-08-20 21:56:58 (36.4 MB/s) - ‘AWS-Python-LaunchInstance.tar.gz’ saved [679/679]

$ tar xfv AWS-Python-LaunchInstance.tar.gz
launch_ec2_instance.py

$ virtualenv env
New python executable in /home/ubuntu/env/bin/python
Installing setuptools, pip, wheel...done.

$ source env/bin/activate

(env)$ pip install boto3

(env)$ python launch_ec2_instance.py
Launched instance with Instance Id: [i-0f227cdd6a25bfa17]!

$ ssh -i "project1_test.pem" ubuntu@ec2-51-221-179-191.compute-1.amazonaws.com
Welcome to Ubuntu 16.04.2 LTS (GNU/Linux 4.4.0-1022-aws x86_64)

ubuntu@ip-172-31-9-237:~$
```

## Amazon Command Line Interface Tools

[AWS CLI Tool](http://aws.amazon.com/cli/)

[Video 3](https://youtu.be/OSGjoMeHc2A): The AWS CLI Tool

## Using the AWS CLI to launch instances

To launch instance, you should first configure the AWS CLI:

```text
$ aws configure
AWS Access Key ID [None]: YOUR AWS ACCESS KEY
AWS Secret Access Key [None]: YOUR AWS SECRET ACCESS KEY
Default region name [None]: us-east-1
Default output format [None]: json
```

Next, use following command to create a Security Group, Key Pair, and Role for the EC2 Instance

```text
$ aws ec2 create-security-group --group-name devenv-sg --description "security group for development environment in EC2"

$ aws ec2 authorize-security-group-ingress --group-name devenv-sg --protocol tcp --port 22 --cidr 0.0.0.0/0

$ aws ec2 create-key-pair --key-name devenv-key --query 'KeyMaterial' --output text > devenv-key.pem

$ chmod 400 devenv-key.pem
```

Finally, you are ready to launch an instance and connect to it. Note that you need to provide AMI_ID (ami-xxxxx) and id of your defined security group (e.g. sg-xxxxx).

```text
$ aws ec2 run-instances \
    --image-id YOUR_AMI_ID \
    --security-group-ids SECURITY_GROUP_ID \
    --count 1 \
    --instance-type t2.micro \
    --key-name devenv-key \
    --query 'Instances[0].InstanceId'
```

## CloudWatch

Amazon CloudWatch enables developers to monitor various facets of their AWS resources. Developers can use it to collect and track metrics from various AWS resources that are running on the AWS Cloud. Using APIs, CloudWatch also allows you to programmatically retrieve monitoring data, which can be used to track resources, spot trends and take automated action based on the state of your cloud resources on AWS. CloudWatch also allows you to set alarms, which constitute a set of instructions to be followed in case of an event that is tracked by CloudWatch is triggered.

CloudWatch can be used to monitor various types of AWS resources including:

* EC2 instances
* EBS volumes
* EMR job flows etc
* ELB Loads

For EC2 instances, CloudWatch allows you to monitor CPU, memory and disk utilization.

For more information on CloudWatch please refer to the [Amazon CloudWatch documentation](http://aws.amazon.com/documentation/cloudwatch/).
