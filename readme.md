# Bundles Application

The Bundles application is an application that allows companies to store sets of bundles in a database and retrieve them according to the specific version.

## Functionalities

- Bundles creation
- Query bundles
- Query a bundle by ID
- Send to S3 a file with all bundle details

## Prerequisites

- Docker
- Docker-compose

## Installation

Clone this repository:
```shell
git clone https://github.com/laurati/projeto01.git
```

## Running Bundles Application
1. Navigate to the application repository:
```shell
cd projeto01
```

2. Inside the repository in root, execute the command:
```shell
docker-compose up -d
```

3. Inside the repository in root, execute the command:
```shell
go run cmd/api/main.go
```

The database and the application will be running on ports 5432 and 8080, respectively

## Usage

### Routes


Method | Endpoint | Description
------ |--------- | -----------
GET    | /        | Check if application is running
POST   | /bundledetails | Create a new bundle detail
GET    | /bundledetails | Query bundle details
GET    | /bundledetails/{id} | Query a bundle detail by ID
POST   | /bundledetails/file | Send to S3 a file with all bundle details 

## Amazon S3

The Amazon S3 is an object storage service that stores data as objects in buckets. An object is a file and any metadata describing the file. A bucket is a container for objects.

To store your data in Amazon S3, create a bucket and specify a bucket name and the AWS Region. Then, upload your data to that bucket as objects in Amazon S3.
Each object has a key (or key name), which is a unique identifier for the object in the bucket