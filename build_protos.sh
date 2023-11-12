#!/bin/bash

protoc -I=pkg --go_out=pkg/ --go_opt=paths=source_relative pkg/config/proto/*.proto
