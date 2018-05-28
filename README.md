# GoFun

### migtate db

>自动生成数据库数据

- 安装migrate tool
  - https://db-migrate.readthedocs.io/en/latest/Getting%20Started/installation/
- 配置database.json
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
- db-migrate up
  - 查看数据库是否写入


