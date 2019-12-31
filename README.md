[![Build Status](https://travis-ci.org/mradile/rssfeeder.svg?branch=master)](https://travis-ci.org/mradile/rssfeeder)
[![Go Report Card](https://goreportcard.com/badge/github.com/mradile/rssfeeder?style=flat-square)](https://goreportcard.com/report/github.com/mradile/rssfeeder)



# rssfeeder
A client / server app for providing rss feeds written in go.

## Overview
This project consists of two separate applications, a server and a client.
  
The client can add feed entries in different feeds on the server. Each feed can be retrieved as atom or rss feed or as a simple json list.
The data is stored in a bolt db file. 
 
## Usage
To use these applications a server is needed and a client that can access the server.


## Server
To start the server you can use the docker-compose file. Make sure to replace at lease the user credentials and the secret. The secret is used for signing the JWTs.  

The data is stored in a single file. If loosing the data is not option make sure to use a mounted volume.

The following ENV variables are available:


| ENV variable  | Description                                           |Example Values             |
|---------------|-------------------------------------------------------|---------------------------|
|LOG_DEBUG      |print debug level log messages                         |1                          |
|LOG_VERBOSE    |print info level log messages                          |1                          |
|DB             |Path where the db file will be stored.                 |/data                      |
|PORT           |Port to listen for requests                            |3000                       |
|HOST           |Hostname of the service                                |http://rssfeeder.com:3000  |
|SECRET         |Secret string used to sign JWT                         |someSecret                 |
|SESSION_TTL    |Time after a user has to relogin                       |24h                        |
|CREATE_USER    |Create a user on startup. Needs LOGIN and PASSWORD     |1                          |
|LOGIN          |Login of the user to create on startup                 |martin                     |
|PASSWORD       |Password of the user to create on startup              |martin                     |

## Client

