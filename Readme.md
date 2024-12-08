# CRUD Generator for Gin & GORM

This project is a CRUD generator for creating basic application scaffolds using the **Gin** framework and **GORM** for database interactions. It supports automatic model, controller, and route generation based on input from an Excel file. Additionally, it supports migrations, seeding, and running a Gin server.

## Features

- **Generate Models, Controllers, and Routes** from an Excel file.
- Supports **PostgreSQL** and **MySQL**.
- **Migrations** and **Seeding** functionalities.
- Can automatically run the **Gin server**.
- Uses **Cobra** for CLI commands.

## Prerequisites

Make sure you have the following installed:
- **Go** (version 1.18 or higher)
- **MySQL** or **PostgreSQL** database
- **Gin** framework
- **GORM** ORM
- **Excelize** library for parsing Excel files

You can install **Go** from the official website: [https://golang.org/dl/](https://golang.org/dl/)

---

## Setup and Installation

### 1. Clone the repository

```bash
git clone https://github.com/AMETORY/ametory-crud.git
cd ametory-crud
```

### 2. Install Dependencies
Run the following command to install required Go dependencies:

```bash
go mod tidy
```

### 3. Configure .env.yaml
Create a ```.env.yaml``` file in the root of the project and define your database credentials. Here is an example:

```yaml
server:
  app_name: "Good Job"
  host: "localhost"
  port: 8080
  app_desc: "Your app description here"
  version: "1.0.0"
  api_url: "http://localhost:8080/api"
  front_end_url: "http://localhost:3000"
  secret_key: "your_secret_key"
  expired_jwt: 30

database:
  type: "postgres"
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "balakutak"
  name: "crud"
  auth_table: "users"

scheduler:
  enabled: true
  interval: 10

mailer:
  smtp_host: "smtp.mailtrap.io"
  smtp_port: 587
  username: "user@mailtrap"
  password: "password"
  sender: "Ametory Support"
  from_to: "support@ametory.id"

s3:
  access_key: "your_access_key"
  secret_key: "your_secret_key"
  bucket: "your_bucket_name"
  region: "us-west-1"

es:
  host: "localhost"
  port: 9200

redis:
  host: "localhost"
  port: 6379
  password: "redis_password"
  db: 0

sms:
  account_sid: "your_account_sid"
  auth_token: "your_auth_token"
  phone_number: "+123456789"

```
For MySQL:

```ini
DB_TYPE=mysql
```
For PostgreSQL:

```ini
DB_TYPE=postgres
```
## Usage
### 1. Run Server
To start the Gin server, use the following command:

```bash
go run main.go run
```
This will start the Gin server locally. You can access it at ```http://localhost:8080```.

### 2. Generate CRUD from Excel
The generate-from-excel command allows you to create models, controllers, and routes from an Excel file. The Excel file should define features, columns, and their types.

#### Excel File Format
Your Excel should have the following columns:

| Model Name | Field Name | Field Type | DB Type |
|------------|------------|------------|---------|
| User       | Name       | string     | varchar(255) |
| User       | Age        | int        | int     |
| Product    | Name       | string     | varchar(255) |
| Product    | Price      | float      | decimal(10,2) |

Each row represents a column in a table for a specific feature.

#### Generate Files from Excel
Run the following command to generate files:

```bash
go run main.go generate-from-excel --path=path/to/your/file.xlsx
```
This will generate the following files for each feature:

- models/Feature.go
- controllers/Feature_controller.go
- routes/Feature_route.go

#### Rename Module
```bash
go run main.go rename new_module  
```

### 3. Migrate Database
To apply migrations to the database, use the following command:

```bash
go run main.go migrate
```
This will apply all migrations to your database, ensuring the tables are created based on the models.

### 4. Seed Database
To seed the database with some initial data, you can use the seed command:

```bash
go run main.go seed
```
This will insert initial data into your database (like default users, etc.).

### 5. Running Migrations and Seeding Automatically
You can also run migrations and seed the database automatically before starting the server:

```bash
go run main.go migrate && go run main.go seed && go run main.go run
```
This will apply migrations, seed the database, and then start the server.

Folder Structure
```plaintext
ametory-crud/
├── cmd/
│   ├── generate.go
│   ├── generate_from_excel.go
│   ├── migrate.go
│   ├── root.go
│   ├── run.go
│   ├── seed.go
├── database/
│   ├── connection.go
│   └── migrate.go
├── models/
│   └── templates/
│       ├── controller.tpl
│       ├── model.tpl
│       └── route.tpl
├── seeders/
│   ├── user_seeder.go
├── main.go
├── go.mod
├── go.sum
└── .env.yaml
```

#### Explanation of Folders and Files
- **cmd/**: Contains the Cobra CLI commands for generating files, migrating the database, seeding the database, and running the server.
- **database/**: Handles database connection and migrations.
- **models/**: Contains the templates for generating model, controller, and route files.
- **seeders/**: Contains the seed logic for populating the database with initial data.
- **main.go**: The entry point to run the application.

#### Example Excel File
Input Example (```data.xlsx```)
| Model Name | Field Name | Field Type | DB Type |
|------------|------------|------------|---------|
| User       | ID         | int        | int     |
| User       | Name       | string     | varchar(255) |
| User       | Age        | int        | int     |
| Product    | ID         | int        | int     |
| Product    | Name       | string     | varchar(255) |
| Product    | Price      | float      | decimal(10,2) |

#### Customization
You can customize the templates for models, controllers, and routes by modifying the files in ```models/templates/```:

- **model.tpl**: Template for generating model structures.
- **controller.tpl**: Template for generating controller functions.
- **route.tpl**: Template for generating route handlers.

Feel free to modify these templates according to your application needs.

Contributing
Feel free to open issues or submit pull requests for improvements or fixes.

#### License
This project is licensed under the MIT License - see the LICENSE file for details.

```yaml

LICENSE
-------

Copyright (c) 2022 xuri

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```