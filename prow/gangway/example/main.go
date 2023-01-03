/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"log"

	pb "k8s.io/test-infra/prow/gangway"
	gangwayGoogleClient "k8s.io/test-infra/prow/gangway/google"
)

var (
	addr      = flag.String("addr", "127.0.0.1:50051", "Address of grpc server.")
	apiKey    = flag.String("apiKey", "", "API key.")
	audience  = flag.String("audience", "", "Audience.")
	keyFile   = flag.String("keyFile", "", "Path to a Google service account key file.")
	clientPem = flag.String("clientPem", "", "Path to a client.pem file.")
)

// Use like this:
//
// go run main.go --keyFile=key.json --audience=SERVICE_NAME.endpoints.PROJECT_NAME.cloud.goog --addr=34.27.163.252:443 --apiKey=API_KEY --clientPem=client.pem foo_proto_message

func main() {
	flag.Parse()

	// Create a Prow API gRPC client that's able to authenticate to it.
	prowClient, err := gangwayGoogleClient.NewFromFile(*keyFile, *audience, *apiKey, *addr, *clientPem)
	if err != nil {
		log.Fatalf("Prow API client creation failed: %v", err)
	}
	defer prowClient.Close()

	ctx, err := prowClient.MkContext()
	if err != nil {
		log.Fatalf("could not create a context with embedded credentials: %v", err)
	}

	jobExecution, err := prowClient.GRPC.CreateJobExecution(ctx, &pb.CreateJobExecutionRequest{})
	if err != nil {
		log.Fatalf("could not trigger job: %v", err)
	}

	log.Printf("triggered job: %v", jobExecution)
}
