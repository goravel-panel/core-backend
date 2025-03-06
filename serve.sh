#!/bin/bash

# Get the port number configured in the .env file
source .env
port=$APP_PORT

# Check if the port is being used
if netstat -an | grep "$port" | grep -i listen > /dev/null ; then
    lsof -i:"$port" | grep main
    echo -e "\033[1;31mPort $port is in use. Please enter the PID of the process occupying the port: \033[0m"
    read -r pid
    kill -15 "$pid"
    if [ $? -eq 0 ]; then
        echo -e "\033[1;32mProcess ($pid) terminated successfully. Restarting the service... \033[0m"
        sleep 1
        air
    else
        echo -e "\033[1;31mFailed to terminate process ($pid). Please check if the process ID entered is correct. \033[0m"
    fi
else
   air
fi