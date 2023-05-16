#!/bin/bash

export "DB_HOST=localhost"
export "DB_PORT=5432"
export "DB_USER=thor"
export "DB_PASS=odinson"
export "DB_NAME=thunderdome"
export "SMTP_HOST=localhost"
export "MTP_PORT=1025"
export "ADMIN_EMAIL=test.test@gmail."
export "ADMIN_PASSWORD=test1234"

make dev-go
