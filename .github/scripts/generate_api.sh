#!/bin/bash
cd ./blog-backend
buf mod update
buf generate
rm -fr gen/api/go/errors
rm -fr gen/api/go/google
rm -fr gen/api/go/openapiv3
rm -fr gen/api/go/validate