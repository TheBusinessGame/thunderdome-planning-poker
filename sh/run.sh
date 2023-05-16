#!/bin/bash

export ENV_VAR1="DB_HOST=localhost"
export ENV_VAR2="DB_PORT=5432"
export ENV_VAR3="DB_USER=thor"
export ENV_VAR4="DB_PASS=odinson"
export ENV_VAR5="DB_NAME=thunderdome"
export ENV_VAR6="SMTP_HOST=localhost"
export ENV_VAR7="SMTP_PORT=1025"
export ENV_VAR8="ADMIN_EMAIL=test.test@gmail."
export ENV_VAR9="ADMIN_PASSWORD=test1234"

make dev-go