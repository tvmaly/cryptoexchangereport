# TODO list for development

## Exchange API stats

### Architecture

* Document Code Layout
** Update README.md to reflect new directories in the code
* Dependencies
** setup dep for project
* Query architecture design
* Database and Archive design
* Configuration management and secrets
* Deployment 
* Logging
* Monitoring 
* Client API design


### Parsing

* Quote parsing for all exchanges
** investigate Cloudflare errors
** poloniex need a different structure for quotes
* Currency conversion considerations
* Timezone conversion considerations
* Exchange info 
* Trades
* Depth of Book


## Landing page templates

https://blackrockdigital.github.io/startbootstrap-landing-page/  ( this is free MIT license and I like how clean it is https://github.com/BlackrockDigital/startbootstrap-landing-page/)

### GIN server

Lets Encrypt Support
https://github.com/gin-gonic/autotls

Store assets in Binary
https://github.com/elazarl/go-bindata-assetfs

### AWS Parameter store for secrets

https://aws.amazon.com/blogs/compute/managing-secrets-for-amazon-ecs-applications-using-parameter-store-and-iam-roles-for-tasks/

https://aws.amazon.com/blogs/mt/the-right-way-to-store-secrets-using-parameter-store/

### AWS S3 managing lists

* given eventual consistency, we should store a last updated timestamp in UTC with objects

https://docs.aws.amazon.com/AmazonS3/latest/dev/Introduction.html

https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingObjects.html

https://docs.aws.amazon.com/sdk-for-go/api/service/s3/

https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-s3-with-go-sdk.html




