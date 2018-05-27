# GoFun

### migtate db

>自动生成数据库数据

- 安装migrate tool
  - https://db-migrate.readthedocs.io/en/latest/Getting%20Started/installation/
- db-migrate create add-users
  - 生成migrate文件
- 修改migrate文件
  ```js
    exports.up = function(db) {
    return db.createTable('users',{
        id:{type:'int',primaryKey:true,autoIncrement:true},
        first:'string',
        last:'string',
    });
    };

    exports.down = function(db) {
    return db.dropTable('users');
    };
  ```

