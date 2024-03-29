# spoonAPI
 Golang wrapper to spoonacular API 


## Introduction

Welcome to the documentation for wrapper Spoonacular API, a  API built using Golang. This document provides detailed information on how to interact with the API, including endpoints, request methods, request parameters, and response formats.

## Base URL

The base URL for all API endpoints is:
    localhost:50052/api/food


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

``` http
  GET /api/food/menu/menuItem HTTP/1.1
  Host: localhost

{
  "menuItems":
    [
      {
        "id":306187,
        "title":"Burger",
        "restaurantChain":"Garfield's Restaurant &  Pub",
        "image":"https://spoonacular.com/menuItemImages/hamburger.jpg",
        "imageType":"jpg",
        "servings":{
            "totalMenuItems":0,
            "size":0,
            "unit":""
          }
      },
      {
        "id":247615,
        "title":"Tony's Bodacious Burger w/ Fries, Max",
        "restaurantChain":"Max & Erma's",
        "image":"https://spoonacular.com/menuItemImages/hamburger.jpg",
        "imageType":"jpg",
        "servings":{
          "totalMenuItems":0,
          "size":0,
          "unit":""
        }
      },
    ]

}

  Response Codes
  200 OK - Request successful
  400 Bad Request - Invalid request parameters
  500 Internal Server Error - Server error

```

## Docker

1. Created a dockerfile to run create the image that will run golang code inside a docker container running linux base system.

2. Build the image using the Dockerfile created in last step
  $ docker build -t [Name of container] ./

3. Spin up container with docker image. I also use an .env file to add enviornment variables to docker container name the container spoonapi
``` bash
  $ docker run -p 50052:50052 -d  --env-file .env rudyjcruz831/spoonapi  
```
  - Do not forget to create the .env file
4. Make sure you make a name or new tag of image to have correct name as our repo
``` bash
 $ docker tag spoonapi rudyjcruz831/spoonapi:v1.0 
```
5. To added to repohub we need to run a command
 $ docker push rudyjcruz831/spoonapi:lates

6. Dockerhub image: rudyjcruz831/spoonapi


## GCP 
 
 1. Created an account using san jose state email 
 2. I used GCP console to create a GCP instance
    1. Name, region, zone, 
    2. Most of the setting are default
    3. Under Firewall I allowed http and https traffic
 3. Then ssh to GCP instance and download docker to VM
    1. Set up Docker's apt repository.



## Installing and running Docker in GCP instance 

1. Run the following command to uninstall all conflicting packages: 
```bash 
 for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt-get remove $pkg; done
```
2. Set up Docker's apt repository.
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
    2. Install the Docker packages.
``` bash 
   sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```
The images below will show how I ran the docker container in the GCP instance

![image info](./img/GCP_Screenshot1.png)
![image info](./img/GCP_Screenshot2.png)



## Cirlce CI
1. Make account
2. make ssh keys add the public key to github and added the private key to circleci
3. the config file is attached to this project under the directory .circleci/

#### Screenshots of circleci
![image info](./img/Circle_Screenshot1.png)
![image info](./img/Circle_Screenshot2.png)


