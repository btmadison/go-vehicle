# Learning Go Rest API Excercise

General Hexagonal Architecture with DI set up to swap between in memory and dynamo repo

Basic rest crud operations on a list of vehicles

main driver is in the cmd folder
- cmd/aws placeholder to use sls / serverless funcs aws
- cmd/server - can switch between in memory or direct dynamo access without serverless - set option in main.go switch
everything else lives in pkg

add a local.env to root of project with the following:
* aws_access_key_id=\<access key>
* aws_secret_access_key=\<secret key>
* region=us-east-1