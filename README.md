# Quickounce

Quickounce is a photosharing app written in [Go](https://go.dev) and [React](https://react.dev/)

## Status

The project is currently a work in progress.
The goals are to have the basic features of a photosharing web app,
including creating posts, creating comments, browsing feeds, and following users.

## Inspiration

This project is inspired by [Instagram](https://instagram.com)
and how data is structured in the app

I also used this article [System Design Of Instagram](https://medium.com/@lazygeek78/system-design-series-cfa60db16c27#id_token=eyJhbGciOiJSUzI1NiIsImtpZCI6ImI1MDljNTEzODc2OGY3Y2YyZTgyN2UwNGIyN2U3ZTRjYmM3YmI5MTkiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIyMTYyOTYwMzU4MzQtazFrNnFlMDYwczJ0cDJhMmphbTRsamRjbXMwMHN0dGcuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIyMTYyOTYwMzU4MzQtazFrNnFlMDYwczJ0cDJhMmphbTRsamRjbXMwMHN0dGcuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDM5NzIwNjMzMDE2NzEwNDA3NzUiLCJlbWFpbCI6InRzdG9uZTQ0OUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmJmIjoxNzUzMjAwNjA0LCJuYW1lIjoiVG9tIFN0b25lIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0ltVF9UZS1UMDU1THBvYTlnNmJkT2FTUERiTXFxNE5JdDJuZ0hDQnliTno3cF9XZz1zOTYtYyIsImdpdmVuX25hbWUiOiJUb20iLCJmYW1pbHlfbmFtZSI6IlN0b25lIiwiaWF0IjoxNzUzMjAwOTA0LCJleHAiOjE3NTMyMDQ1MDQsImp0aSI6IjNhMGZmNzIzYWUzOWI0ZTFjYzUxYzlmZWYzM2ZkMDBiODhjZjk1NTMifQ.lPwLfdY9cYCQ4o-KfJr4zy27RhrxFQiixyjQXMHBg8XqJbykJa05ObRCM4xEXFC8OXomrlzqoh00SPmBITZK-zf-uP4lADVmrznIJx_iLoDQxYNUDUwBIToJdNIK8CAOvIJr39omwErO2uKUWXqWZZ0LDjk1uQCRVk4O6iiA1utdy65Njy2hz0Px8MtV7KMjdpoxcvXdhAfp5g4MBTsFh3CHtkCUh-MjtBaVmb6Iws7CAAffcvnr9EwSkxIQCuunnuYHX4xnz460tXN3YtI31M962TJ3gQJcpcuJMz6mPZ93YSp0m4zU4Im0MEBIZ0eSTlWIrVsS8JZkDXRtadBCjA) for the design of the database

## Dependencies

- ```go``` [installation](https://go.dev/dl/)
- ```node``` [installation](https://nodejs.org/en/download)
- ```postgresql``` [installation](https://www.postgresql.org/download/)

## Setup

Fork this repo, and then clone it to your machine

### Go backend server

Download dependecies with ```go mod tidy```

Enter the server directory ```cd server/```

Create a database for quickounce:

- Enter the ```psql``` shell:

  - Mac: ```psql postgres```
  - Linux: ```sudo -u postgres psql```

  You should see a propmpt
  ```postgres=#```

- Create a new database:

  ```CREATE DATABASE quickounce;```

- Connect to the new database:

  ```\c quickounce```

  You should see a new prompt
  ```quickounce=#```

  Create your ```.env``` file:

  ```touch .env```

  Edit the env file to have
  ```DB_URL={YOUR URL}```
  ```PLATFORM=dev```
  ```SECRET={YOUR SECRET}```

  Where ```DB_URL``` is your postgres url

  - Mac: ```postgres://{username}@localhost:5432/quickounce```
  - Linux: ```postgres://postgres:postgres@localhost5432/quickounce```

  And ```SECRET``` is your secret

  You can generate one with ```openssl rand -base64 64```

  Now you can run the backend api server with ```go run .```

### Webapp

Enter the webapp folder ```cd webapp/```

Install dependencies ```npm install```

Run the frontend with ```npm run dev```
