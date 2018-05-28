# GoFun

### migtate db

>自动生成数据库数据

- 1.1安装migrate tool
  - https://db-migrate.readthedocs.io/en/latest/Getting%20Started/installation/
- 1.2配置database.json
  ```json
  {
  "dev": {
    "host": "localhost",
    "user": "root",
    "password" : "123456",
    "database": "func",
    "driver": "mysql",
    "multipleStatements": true
  }
}

  ```
- 1.3 db-migrate up
  - 查看数据库是否写入

- 2.1 make run

### 要配置的文件
/pkg/db/congig.json
database.json
