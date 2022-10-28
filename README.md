# nscat

nscat displays the content of a file replacing keys and passwords with Xs

**As a** user **I want** a command-line program **that** displays the content of a file, but removes any 'secret' or sensitive information, like passwords or API keys.

## Use Case

While troubleshooting problems on a server, often the sysadmin in sharing their screen. There often comes a time when someone says, "What's in that .env file?". If the SA cats the file, there is a chance that any secrets in that file will be revealed to anyone viewing the shared screen. It would be nice if there was a program that would 'cat' the file but remove any sensitive information. 

```bash
nscat .env

ABC_DBHOST=db_server_rev.example.com:3337
ABC_DBUSER=appuser
ABC_DBPASS=XXXXXXXXXXXXXXXX
ABC_DBNAME=table_name
DBHOST=db_server.example.com:3337
DBUSER=appuser_db
DBPASS=XXXXXXXXXXXXXXXX
DBNAME=dev_reports
ADAP_HOST=adis-example.ent.romp.com
ADAP_URL=https://ldap.example.com:8090
ADAP_CREDENTIAL=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
HOST_CERTIFICATE=/etc/ssl/cert.pem
BUST_DBUSER=buster
BUST_DBPASS=XXXXXXXXXX
BUST_DBHOST=database.example.com:3337
BUST_DBNAME=dev_buster
```
