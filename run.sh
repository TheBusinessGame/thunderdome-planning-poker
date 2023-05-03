#!/bin/bash

DIR="/home/$USER/$DIR_NAME"

FILE="$DIR/poker.env"

if [[ -d "$DIR" ]]
then
    echo "$DIR already exists"
else
    mkdir "$DIR"
    echo "$DIR created"
fi

if [[ -f "$FILE" ]]
then
    echo "$FILE already exists"
else
    touch "$FILE"
    echo "$FILE created"
fi

# Set the environment variables in the file
echo "DB_HOST=localhost" > "$FILE"
echo "DB_PORT=5432" >> "$FILE"
echo "DB_USER=thor" >> "$FILE"
echo "DB_PASS=odinson" >> "$FILE"
echo "DB_NAME=thunderdome" >> "$FILE"
echo "SMTP_HOST=localhost" >> "$FILE"
echo "SMTP_PORT=1025" >> "$FILE"
echo "ADMIN_EMAIL=test.test@gmail.com" >> "$FILE"

# Export the environment variables
export $(cat "$FILE" | xargs)

echo "Environment variables set in $FILE and exported"