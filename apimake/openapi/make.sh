#!/bin/bash

oapi-codegen --package=server --generate chi-server,types -o server_gen.go mauth.yaml
oapi-codegen --package=types --generate types -o types_gen.go mauth.yaml
