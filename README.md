# database_migrations
Проект для миграции базы данных


```bash
NAME:
   database_migrations - application create (create) template file for sql migration or execute migrations (exec)

USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
   create, c  create new template file for migration with sql format
   exec, exc  execute migrations from source into database
   show, s    show version of database
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

Настройки могут как передаваться в команде, так и быть записаны в database_migration.json в месте запуска приложения

```json
{
    "user" : "{пользователь базы}",
    "password": "{пароль пользователя базы}",
	"host": "{хост базы}",
	"port": {порт},
	"database": "{название базы}",
	"sslmode": "disable",
    "sources": "{путь к файлам миграции}"
}
```