# Learning Go Rest API Excercise

General Hexagonal Architecture with DI set up to swap between in memory and dynamo repo

Basic rest crud operations on a list of vehicles

## Local Server Mode
//for local offline server with in memory or connection to dynamo, main driver is in the cmd folder under server
- cmd/aws placeholder to use sls / serverless funcs aws
- cmd/server - can switch between in memory or direct dynamo access without serverless - set option in main.go switch
everything else lives in pkg

add a local.env to root of project with the following:
* aws_access_key_id=\<access key>
* aws_secret_access_key=\<secret key>
* region=us-east-1

need to manually create db in dynamo if not running sls setup that follows

## Serverless Lambda / Infra as service

cmd/aws - run `make deploy` - this will deploy to defaults / configuration per serverless.yml, deploying routes
