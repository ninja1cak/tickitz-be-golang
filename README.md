<div align='center' style="text-align: center;">

<h1 style="border:0;margin:1rem">Tickitz Backend</h1>

Backend Golang for Tickitz

[Suggestion](mailto:hauzan41200@gmail.com)

<hr>
<br>

</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Postman Collection](#postman-collection)
- [Resources](#resources)
- [Contributors](#contributors)
- [Related Project](#related-projects)
- [License](#license)
- [Suggestion](#suggestion)

## Overview

Tickits is a web-based application for ticket booking movies and cinemas.

## Features

- Authentication
- User Role: Booking ticket movie, Edit Profile (Update)
- Admin Role: Create update movie details and movie schedule
- Error Handling
- Email Activation
- etc.

## Technologies Used

- [Golang](https://go.dev/)
- [Gin Gonic](github.com/gin-gonic/gin)
- [JsonWebToken](github.com/golang-jwt/jwt/v5)
- [Postgresql](https://www.postgresql.org/)
- [Cloudinary](https://cloudinary.com/)
- [Gomail](gopkg.in/gomail.v2)
- etc.

## Getting Started

### Installation

1. Clone this repo

   ```bash
   git clone https://github.com/ninja1cak/tickitz-be
   ```

2. Enter the directory

   ```bash
   cd tickitz-be-golang
   ```

3. Install all dependencies

   ```bash
   go get .
   ```

4. Create .env file

   ```env
    DB_HOST = [YOUR DATABASE HOST]
    DB_USER = [YOUR DATABASE USER]
    DB_PASSWORD = [YOUR DATABASE PASSWORD]
    DB_DATABASE = [YOUR DATABASE NAME]
    DB_PORT = [YOUR DATABASE PORT]
    JWT_SECRET = [YOUR JWT_SECRET]
    PORT = [YOUR PORT]
    CLOUDINARY_URL = [YOUR CLOUDINARY API]
    CONFIG_AUTH_EMAIL = [YOUR EMAIL]
    CONFIG_AUTH_PASSWORD = [YOUR EMAIL PASSWORD]

   ```

5. Start the local server

   ```bash
   go run cmd/main.go
   ```



## Postman Collection

You can download in <a href='https://drive.google.com/file/d/1iGTRVvIm_XJAoyn6ciRrdqwhckKD1cm3/view?usp=sharing'> Here </a>

## Contributors

Currently, there are no contributors to this project. If you would like to contribute, you can submit a pull request.

## Related Projects

- [Tickitz Front End](https://github.com/JulianMindria/Frontend_tickitzv2) - Front End

## License

This project is licensed under the ISC License

## Suggestion

If you find bugs / find better ways / suggestions you can pull request.
