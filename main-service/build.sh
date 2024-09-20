# !/bin/bash
# run /bin/chmod +x build.sh

go build -o build cmd/main/*.go && ./build
# ./build -goenv=development -dbname=bookings -dbuser=orololuwa