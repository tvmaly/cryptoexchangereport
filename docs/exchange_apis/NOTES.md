# Notes on scraping mechanism

Most exchanges have limits on their API endpoints to prevent abuse

Imagine you are pulling data from biance and they changed the limit to a lower value in the middle of the day

How should your scraper behave?

Consider it keeps doing the same thing, your IP gets banned, you cannot get your data.

Consider you change your sampling rate, your data becomes inconsistent.

What options do you have?   

Round Robin on machines with different IP addresses seems to be one possibility, could possibly use AWS Lambda like:

https://stackoverflow.com/questions/30647844/how-to-automate-amazon-aws-ec2-for-scraping

One issue is that AWS Lambda does not allow you to predict which external IP address that will be used.  

https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html

AND

https://ip-ranges.amazonaws.com/ip-ranges.json

You could be using the same IP address, worse, others could be using AWS Lambda for the same purpose and you could trigger limit without knowing it.

# AWS SQS scaling queue messaging service

we could coordinate with the queue service potentially

https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-using-sqs-queue.html

Question:  Can we pull an item out of the queue if there is a rate limit issue or if exchange is down?  

Or should we schedule on demand?

## AWS nano instances

Instead of using AWS Lambda with unknown IP address, we could provision smaller T2 nano instances

https://aws.amazon.com/ec2/instance-types/

more work needed to see true cost




## AWS Lambda

# Pricing Calculations

https://serverless.com/blog/understanding-and-controlling-aws-lambda-costs/

# Regions Availability for AWS Lamda

https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/

