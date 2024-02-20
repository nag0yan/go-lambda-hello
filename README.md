# go hello lambda function
## deploy
*require go and aws cli*
### create exec role
```Shell
aws iam create-role \
--role-name sample-role \
--assume-role-policy-document file://sample-role-truct-policy.json
```

### build and archive
```Shell
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
```  
```Shell
zip helloFunction.zip bootstrap
```

### create lambda function
```Shell
aws lambda create-function --function-name helloFunction --runtime provided.al2023 --handler bootstrap --architectures arm64 --zip-file fileb://helloFunction.zip --role arn:aws:iam::755009017669:role/sample-role
```

### update lambda function code
```Shell
aws lambda update-function-code --function-name helloFunction --zip-file fileb://helloFunction.zip
```