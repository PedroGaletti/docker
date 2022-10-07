## Docker

Docker is a set of platform-as-a-service products that use operational-level virtualization to deliver software in packages called containers.

## What is Containers?

- Containers are isolated from each other and bundle their own software, libraries, and configuration files.
- Containers are great for continuous integration and continuous delivery (CI/CD) workflows.

## Why Docker?

- Docker streamlines the development lifecycle by allowing developers to work in standardized environments using local containers which provide your applications and services.

## Stack

- [Docker](https://www.docker.com) - Develop faster. Run anywhere.

## About

I dedicated this repository to my Docker study, but if you want to contribute with some content, feel free.

- [Golang](https://github.com/PedroGaletti/golang/tree/main/arrays/main.go) - Learn how to build images and containers for a Golang application

## How to use

It's necessary to run the following commands to test **Golang** Dockerfile:

```
docker build . -t my-server
docker run -p 8080:8080 my-server
```
