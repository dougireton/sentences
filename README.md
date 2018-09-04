# sentences


## Requirements

* AWS CLI already configured with at least PowerUser permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

### Installing dependencies

In this example we use the built-in `go get` and the only dependency we need is AWS Lambda Go SDK:

```shell
make install
```

**NOTE:** As you change your application code as well as dependencies during development, you might want to research how to handle dependencies in Golang at scale.

### Building

Golang is a staticly compiled language, meaning that in order to run it you have to build the executeable target.

You can issue the following command in a shell to build it:

```shell
make build
```

**NOTE**: If you're not building the function on a Linux machine, you will need to specify the `GOOS` and `GOARCH` environment variables, this allows Golang to build your function for another system architecture and ensure compatability.

### Local development

**Invoking function locally through local API Gateway**

```bash
make local-api
```

If the previous command ran successfully you should now be able to hit the
following local endpoint to invoke your function:

```shell
 curl --data-binary "@sentences/testdata/basic.txt" http://localhost:3000/sentences | jq .
 ```

### Example Test Output (using httpie)

 ```shell
$ cat sentences/testdata/basic.txt|http POST :3000/sentences
HTTP/1.0 200 OK
Content-Length: 188
Content-Type: application/json
Date: Tue, 04 Sep 2018 04:19:00 GMT
Server: Werkzeug/0.14.1 Python/3.7.0

[
    "This is a sentence with a period.",
    "Is it followed by a sentence with a question mark?",
    "Yes it is!",
    "Now we have a sentence with two spaces before it.",
    "New paragraph with \"quotes\"."
]
```

**SAM CLI** is used to emulate both Lambda and API Gateway locally and uses our `template.yaml` to understand how to bootstrap this environment (runtime, where the source code is, etc.) - The following excerpt is what the CLI will read in order to initialize an API and its routes:

```yaml
...
Events:
    Sentences:
        Type: Api # More info about API Event Source: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
        Properties:
            Path: /sentences
            Method: post
```

## Packaging and deployment


Run the following command to package our Lambda function to S3:

```bash
sam package \
    --template-file template.yaml \
    --output-template-file packaged.yaml \
    --s3-bucket REPLACE_THIS_WITH_YOUR_S3_BUCKET_NAME
    --s3-prefix sentences
```

Next, the following command will create a Cloudformation Stack and deploy your SAM resources.

```bash
sam deploy \
    --template-file packaged.yaml \
    --stack-name sentences \
    --capabilities CAPABILITY_IAM
```

> **See [Serverless Application Model (SAM) HOWTO Guide](https://github.com/awslabs/serverless-application-model/blob/master/HOWTO.md) for more details in how to get started.**

After deployment is complete you can run the following command to retrieve the API Gateway Endpoint URL:

```bash
aws cloudformation describe-stacks \
    --stack-name sentences \
    --query 'Stacks[].Outputs'
```

### Testing

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
make test
```

## Appendix

### Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: https://golang.org/doc/install

A quickstart way would be to use Homebrew, chocolatey or your Linux package manager.

#### Homebrew (Mac)

Issue the following command from the terminal:

```shell
brew install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
brew update
brew upgrade golang
```

#### Chocolatey (Windows)

Issue the following command from the powershell:

```shell
choco install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
choco upgrade golang
```

## AWS CLI commands

AWS CLI commands to package, deploy and describe outputs defined within the cloudformation stack:

```bash
sam package \
    --template-file template.yaml \
    --output-template-file packaged.yaml \
    --s3-bucket REPLACE_THIS_WITH_YOUR_S3_BUCKET_NAME

sam deploy \
    --template-file packaged.yaml \
    --stack-name sentences \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides MyParameterSample=MySampleValue

aws cloudformation describe-stacks \
    --stack-name sentences --query 'Stacks[].Outputs'
```

## Todos

Todos.md contains items for future improvements.