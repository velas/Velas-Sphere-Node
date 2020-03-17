# Velas Sphere

Velas Sphere is an opensource initiative and ecosystem to allow usual customers perform resource-demanding tasks using Storage, CPU and GPU sharing concepts.

## Architecture 

Velas Sphere is a P2P network of nodes communicating between each other using gRPC. Each node can act as a requester and provider and has some plugins. In order to allow plugins be written in any general-purpose language, they are physically decoupled from the provider service using again gRPC and each plugin is basically a separate service.

## Economy

Requesters pay providers for the shared resources using Ethereum Network and smart-contracts.

## Building, Installing and Running Velas Sphere locally

Since whole Velas Sphere node is implemented using Golang exclusively, it is very simple to build and run.

```bash
$ go build .
$ ./velas-sphere plugin
$ ./velas-sphere provider
$ ./velas-sphere requester
```

These binaries will be as well dockerized in the nearest future for easier and consistent deployments.