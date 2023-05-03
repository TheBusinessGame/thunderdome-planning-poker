#!/bin/bash

if netstat -tuln | grep ':8080 ' >/dev/null; then
    echo "The port 8080 is listening"

    pid=$(lsof -ti :8080)

    service_name=$(ps -p $pid -o comm=)

    if [ "$service_name" = "thunderdom" ]; then
        echo "The service thunderdom is running with the PID: $pid"
        echo "Killing the service ..."
        kill $pid
    else
        echo "The service is not thunderdom"
    fi
else
    echo "The port 8080 is not listening"
fi