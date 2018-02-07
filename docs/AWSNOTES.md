# this document holds some quick notes on AWS services that we can consider


## AWS SQS  (Scalable Queue Service)
https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/Welcome.html

Amazon SQS offers regular queue and FIFO

SQS offers ability to delay messages, but if ordering is important where FIFO is needed, you end up delaying all messages.

This means you would need different queues for different delays

standard queue may not be ideal as we cannot guarantee the order


## AMAZON EC2 servers

* Amazon EBS-Backed AMI - You're charged for instance usage,Amazon EBS volume usage, and storing your AMI as an Amazon EBS snapshot

* Amazon Instance Store-Backed AMI - You're charged for instance usage and storing your AMI in Amazon S3.

* However, we don't recommend that you access AWS using the credentials for your AWS account;
we recommend that you use AWS Identity and Access Management (IAM) 

* AWS uses public-key cryptography to secure the login information for your instance. you'll need to create a key pair in each region. pg 23

