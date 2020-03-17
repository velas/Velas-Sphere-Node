# Velas Sphere

Velas Sphere is an opensource initiative and ecosystem allowing usual customers perform resource-demanding tasks using storage, CPU and GPU sharing concepts.

## Architecture 

Velas Sphere is a P2P network of nodes communicating between each other using gRPC. Each node can act as a requester and provider and has some plugins. In order to allow plugins be written in any general-purpose language, they are physically decoupled from the provider service using again gRPC and each plugin is basically a separate service.

## Economy

Requesters pay providers for the shared resources using Ethereum Network and smart-contracts.

## Building, Installing and Running a Velas Sphere node locally

Since the whole Velas Sphere node is implemented using Golang exclusively, it is very simple to build and run.

```bash
$ go build .
$ ./velas-sphere plugin
$ ./velas-sphere provider
$ ./velas-sphere requester
```

These binaries will be dockerized as well in the nearest future for easier and consistent deployments.

## Contribution Guideline

Basic contribution rules are:

1. Each package serves a single signature. It can contain both instances or their constructors accepting a config
2. All packages are located in the `internal` folder
3. `"internal/resources"` is a special package since it contains the Protobuf defenition and autogenerated files