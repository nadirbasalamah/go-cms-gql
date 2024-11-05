#!/bin/sh

# Assign command-line arguments to variables
MONGO_INITDB_ROOT_USERNAME=$1
MONGO_INITDB_ROOT_PASSWORD=$2
ADMIN_NAME=$3
ADMIN_EMAIL=$4
ADMIN_PASSWORD=$5

# Log into the MongoDB service container, check if an admin role already exists, and insert if not
docker exec -i mongodb-service mongosh -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" <<EOF

use cmsapp

if (db.users.findOne({ role: "admin" })) {
    print("Error: An admin user already exists in the database.")
    quit(1)
} else {
    db.users.insertOne({
      username: "$ADMIN_NAME",
      email: "$ADMIN_EMAIL",
      password: "$ADMIN_PASSWORD",
      role: "admin"
    })
    print("New admin user created successfully in the database.")
}

exit
EOF
