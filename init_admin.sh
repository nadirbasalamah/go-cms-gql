#!/bin/sh

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Log into the MongoDB service container
docker exec -i mongodb-service mongosh -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" <<EOF

use cmsapp

db.users.insertOne({
  username: "$ADMIN_NAME",
  email: "$ADMIN_EMAIL",
  password: "$ADMIN_ENCRYPTED_PASSWORD",
  role: "admin"
})

exit
EOF

echo "New admin user created successfully in the cmsapp database."
