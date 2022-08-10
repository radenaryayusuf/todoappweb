<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">Todo App(Proyek Individu Golang DTS Kominfo [DTS152308829101-499])</h3>

  <p align="center">
  Aplikasi Todo sederhana dengan menggunakan bahasa pemrograman Go!
  </p>
</p>

## Installation

### Prerequisite

1. Install [PostgreSQL](https://www.postgresql.org/) 
1. [Install Go 1.18 or later](https://go.dev/dl/)

### Manually from source

1. Clone the repo and cd into the directory

```sh 
git clone https://github.com/radenaryayusuf/todoappweb.git && cd todoappweb
```

1. Run sql script to setup db 

```sh
psql postgres -U <USERNAME> -f scripts/db.sql
```

### Setup Env File

1. Rename file extension .env.example to .env
2. Modify .env and edit DB_USER and DB_PASS for Postgres credential

### Running the server

1. Change directory to the root of the project

```sh
cd path/to/todoapp
```

2. Run server with default config

```sh 
go run main.go
```
OR run the server with air package
```sh 
air
```

3. Go to 127.0.0.1:3388 in your browser
