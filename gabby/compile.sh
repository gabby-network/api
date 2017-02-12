#!/bin/bash

protoc -I . gabby.proto --go_out=plugins=grpc:.
