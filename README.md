# spoonAPI
 Golang wrapper to spoonacular API 


## 

# Introduction

Welcome to the documentation for wrapper Spoonacular API, a  API built using Golang. This document provides detailed information on how to interact with the API, including endpoints, request methods, request parameters, and response formats.

## Base URL

The base URL for all API endpoints is:
    localhost:50052/api/food


## Error Handling

 I need to check error for example if i dont give my correct api secret

## Endpoints

### 1. Endpoint Name

#### Description

This endpoint allows users to search for a specific menu item and retrieve a list of related food types from the Spoonacular API.

#### Endpoint

GET /menu/menuItem

#### Parameters

| Parameter  | Type   | Description                       |
|------------|--------|-----------------------------------|
| `item`     | string | Item need to be a real food       |


#### Request Example

```http
  GET /endpoint-path?param1=value1&param2=value2 HTTP/1.1
  Host: your-api-base-url.com

  Response
  Copy code
  {
    "response_key": "response_value"
  }


  Response Codes
  200 OK - Request successful
  400 Bad Request - Invalid request parameters
  401 Unauthorized - Authentication failure
  404 Not Found - Resource not found
  500 Internal Server Error - Server error

```

# Docker

1. Created a dockerfile to run create the image that will run golang code inside a docker container running linux base system.

2. Build the image using the Dockerfile created in last step
  $ docker build -t [Name of container] ./

3. Spin up container with docker image. I also use an .env file to add enviornment variables to docker container name the container foodapi
  $ docker run -p 50052:50052 -d  --env-file .env foodapi  

4. Make sure you make a name or new tag of image to have correct name as our repo
 $ docker tag foodapi rudyjcruz831/foodapi:v1.0 

5. To added to repohub we need to run a command
 $ docker push rudyjcruz831/foodapi:v1.0

 # GCP 
 
 1. Created an account using san jose state email 
 2. I used GCP console to create a GCP instance
    1. Name, region, zone, 
    2. Most of the setting are default
    3. Under Firewall i allowed http and https traffic
 3. Then ssh to GCP instance and download docker to VM
    1. Set up Docker's apt repository.
```bash 
    # Add Docker's official GPG key: 
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update 
```
```
    2. Install the Docker packages.
``` bash 
   sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

# Cirlce CI
1. Make account
2. make ssh keys 
  - this part i had trouble figuering it out i am still having trouble with the ssh keys
  - I need to make becasuse I am adding a secretphrase  X - this did not work
  - I created a second project but made sure I follow the instructions