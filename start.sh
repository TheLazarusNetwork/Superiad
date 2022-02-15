#!/bin/bash

docker run --name="mtwallet" --rm -d -p 5432:5432 \
-e POSTGRES_PASSWORD=revotic \
-e POSTGRES_USER=revotic \
-e POSTGRES_DB=mtwallet \
postgres -c log_statement=all