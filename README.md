# Golang and gRPC rent-vs-buy calculator

### Prerequisites:
1. This project needs a working setup of Docker.
2. This project needs a working setup of Docker-compose.
3. This project needs a working setup of Go tools.
4. You will need a working version of "protoc" compiler for changing ".proto" files, or adding new ".proto" files and
compiling them into ".pb.go" files.

## Setup instructions

### Step 1:
To have the project working correctly, you should have an ".env" file, located in the root of the project.
For this you could copy the ".env.example" file to ".env":
```bash
$ cp .env.example .env
```

### Step 2:
You can (re)build ".proto" files using the following command:

```bash
$ make proto_gen
```

### Step 3:
To spin up the gRPC server using docker, open up a terminal, change directory to the root of the project and run the 
following command:

```bash
$ make up
```

### Step 4:
To shut down the gRPC server and stop the app container, run the following command:

```bash
$ make down
```

### Step 5:
If you wish to run the gRPC server on your own machine, run the following command:

```bash
$ make run
```

### Step 6:
To test the gRPC API, you need a grpc client. Installing "grpcurl" is recommended. To do so, run the following command:

```bash
$ go get github.com/fullstorydev/grpcurl/...
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```
This installs the "grpcurl" utility executable into your $GOPATH/bin directory. Make sure this directory is in your PATH 
environment variable, to be accessed from your terminal.

### Step 6:
Now you can test the gRPC services. To do so, run the following command:

- With empty request(It will calculate the response using default options):
```bash
$ grpcurl -plaintext localhost:9000 v1.breakeven.BreakevenService/CalcBreakeven
```

- With request:
```bash
$ grpcurl -plaintext -d '{"houseValue": 750000, "downPayment": 10, "propertyTaxRate": 1, "propertyTransferTaxRate": 1, "livingYears": 30, "mortgageInterestRate": 1, "mortgageTerm": 30, "monthlyCommonCost": 300}' localhost:9000 v1.breakeven.BreakevenService/CalcBreakeven
```

### Tips
- The "-d" flag is for your input request. You can form your request into a JSON string to test the services, as shown above.
- The "-plaintext" flag is for insecure connections, when the gRPC server does not accept TLS.

## Considerations
- No tests yet! The reason for this was the lack of time and to respect the deadline.

## TODO
- Tests