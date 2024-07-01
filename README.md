# Odyssey LMS
A simple demo LMS written in Go with a SvelteKit frontend

## How to run this project

First, clone this repository and ```cd web/ui``` into the ui directory.

Then run:
```
npm install
npm run build
```

The app requires a few environment variables:

```
DB_VARIANT
CONNECTION_STRING
LISTEN_ON
```

```DB_VARIANT``` Takes the following values:

```
sqlite
mysql
postgresql
```
It's not mandatory that you specify a variant. if no variant is provided, ```sqlite``` is used by default. However, if you specity a certain database, you must provide the ```CONNECTION_STRING```

Example values are shown below

```
DB_VARIANT=mysql
CONNECTION_STRING=<username>:<password>@<protocol>(<host>:<port>)/<database>?parseTime=true

DB_VARIANT=postgresql
CONNECTION_STRING=postgres://<username>:<password>@<host>/<database>?sslmode=disable
```

```LISTEN_ON``` is the port the app should listen on, it's optional, with a default value of ```:8080```. You can define this variable like this:
```
LISTEN_ON=:8000
```
