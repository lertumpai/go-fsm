#!/bin/bash

# Replace these variables with your actual MySQL connection details
MYSQL_USER="root"
MYSQL_PASSWORD="root"
DATABASE="my_database"

n_bash=1000

# Number of records to insert
num_records=1000

# Loop to generate values for the insert statement
# shellcheck disable=SC2167
for ((x = 0; x < n_bash; x++)); do
    # Prepare the insert statement
    insert_query="INSERT INTO users (username, email, birthdate, desctiprion) VALUES "

    # shellcheck disable=SC2165
    for i in $(seq $num_records); do
        username="user${i}"
        email="user${i}@example.com"
        birthdate="1990-01-01"
        desctiprion="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

        # Append values to the insert statement
        insert_query+="('$username', '$email', '$birthdate', '$desctiprion'),"
    done

    # Remove the trailing comma
    insert_query="${insert_query%,}"

    # Use the mysql command to execute the insert statement
    mysql -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -D "$DATABASE" -e "$insert_query;"
    echo "Insert bash $x $num_records records"
done