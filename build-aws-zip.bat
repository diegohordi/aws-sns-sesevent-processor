set GOOS=linux
go build
mkdir target
move .\aws-sns-sesevent-processor .\target\aws-sns-sesevent-processor
build-lambda-zip --output .\target\aws-sns-sesevent-processor.zip .\target\aws-sns-sesevent-processor
