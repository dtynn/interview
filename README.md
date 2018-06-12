#### interview

```
查找资料根据下列需求利用任意web框架+ 任意数据库 构建一个api项目

Web框架尽量选择简单的 不要找那种一站式的框架 太臃肿  如 express diango rails

实现世界杯小组赛模拟排期功能

1 提供一个录入球队的api  名称 编号
2 通过各种方式插入16支球队的数据到数据库 数据可以参考 http://data.2018.163.com/schedule.html#/group

3 实现排期功能
	1  四个小队一组 每组的球队都和另外的球队进行一次比赛  小组名称可以用A-D
	2 每天进行2次比赛 比赛时间设置在当天的23:00 和 02:00
4 写个api 返回 赛程列表
每个列表项包含
小组名称 参数球队1  参赛球队2 比赛时间
```



##### install

```
// source code
go get -v github.com/dtynn/interview
go install -v github.com/dtynn/interview/...

// init database
psql -c "CREATE DATABASE interview;"
psql -f $GOPATH/src/github.com/dtynn/interview/proto/database.sql interview


// start server
interview --dsn="postgres://user:pass@localhost:5432/interview?sslmode=disable&connect_timeout=10&search_path=interview"
```



##### api

1.  add team

   ```
   curl -XPOST http://127.0.0.1:10080/teams/add --data-binary '{"teams": [{"id": 1, "name": "巴西"}]}'
   
   {
     "code": 0,
     "message": ""
   }
   ```

2. schedule

   ```
   curl -XPOST http://127.0.0.1:10080/schedule
   
   {
     "groups": [
       {
         "id": 1,
         "team": [
           {
             "id": 1,
             "name": "1"
           },
           {
             "id": 1,
             "name": "巴西"
           },
           {
             "id": 2,
             "name": "2"
           },
           {
             "id": 3,
             "name": "3"
           }
         ],
         "matches": [
           {
             "day": 1,
             "time": "23:00",
             "home": {
               "id": 1,
               "name": "1"
             },
             "away": {
               "id": 1,
               "name": "巴西"
             }
           },
           {
             "day": 1,
             "time": "02:00",
             "home": {
               "id": 1,
               "name": "1"
             },
             "away": {
               "id": 2,
               "name": "2"
             }
           },
           {
             "day": 2,
             "time": "23:00",
             "home": {
               "id": 1,
               "name": "1"
             },
             "away": {
               "id": 3,
               "name": "3"
             }
           },
           {
             "day": 2,
             "time": "02:00",
             "home": {
               "id": 1,
               "name": "巴西"
             },
             "away": {
               "id": 2,
               "name": "2"
             }
           },
           {
             "day": 3,
             "time": "23:00",
             "home": {
               "id": 1,
               "name": "巴西"
             },
             "away": {
               "id": 3,
               "name": "3"
             }
           },
           {
             "day": 3,
             "time": "02:00",
             "home": {
               "id": 2,
               "name": "2"
             },
             "away": {
               "id": 3,
               "name": "3"
             }
           }
         ]
       },
       {
         "id": 2,
         "team": [
           {
             "id": 4,
             "name": "4"
           },
           {
             "id": 5,
             "name": "5"
           },
           {
             "id": 6,
             "name": "6"
           },
           {
             "id": 7,
             "name": "7"
           }
         ],
         "matches": [
           {
             "day": 1,
             "time": "23:00",
             "home": {
               "id": 4,
               "name": "4"
             },
             "away": {
               "id": 5,
               "name": "5"
             }
           },
           {
             "day": 1,
             "time": "02:00",
             "home": {
               "id": 4,
               "name": "4"
             },
             "away": {
               "id": 6,
               "name": "6"
             }
           },
           {
             "day": 2,
             "time": "23:00",
             "home": {
               "id": 4,
               "name": "4"
             },
             "away": {
               "id": 7,
               "name": "7"
             }
           },
           {
             "day": 2,
             "time": "02:00",
             "home": {
               "id": 5,
               "name": "5"
             },
             "away": {
               "id": 6,
               "name": "6"
             }
           },
           {
             "day": 3,
             "time": "23:00",
             "home": {
               "id": 5,
               "name": "5"
             },
             "away": {
               "id": 7,
               "name": "7"
             }
           },
           {
             "day": 3,
             "time": "02:00",
             "home": {
               "id": 6,
               "name": "6"
             },
             "away": {
               "id": 7,
               "name": "7"
             }
           }
         ]
       },
       {
         "id": 3,
         "team": [
           {
             "id": 8,
             "name": "8"
           },
           {
             "id": 9,
             "name": "9"
           },
           {
             "id": 10,
             "name": "10"
           },
           {
             "id": 11,
             "name": "11"
           }
         ],
         "matches": [
           {
             "day": 1,
             "time": "23:00",
             "home": {
               "id": 8,
               "name": "8"
             },
             "away": {
               "id": 9,
               "name": "9"
             }
           },
           {
             "day": 1,
             "time": "02:00",
             "home": {
               "id": 8,
               "name": "8"
             },
             "away": {
               "id": 10,
               "name": "10"
             }
           },
           {
             "day": 2,
             "time": "23:00",
             "home": {
               "id": 8,
               "name": "8"
             },
             "away": {
               "id": 11,
               "name": "11"
             }
           },
           {
             "day": 2,
             "time": "02:00",
             "home": {
               "id": 9,
               "name": "9"
             },
             "away": {
               "id": 10,
               "name": "10"
             }
           },
           {
             "day": 3,
             "time": "23:00",
             "home": {
               "id": 9,
               "name": "9"
             },
             "away": {
               "id": 11,
               "name": "11"
             }
           },
           {
             "day": 3,
             "time": "02:00",
             "home": {
               "id": 10,
               "name": "10"
             },
             "away": {
               "id": 11,
               "name": "11"
             }
           }
         ]
       },
       {
         "id": 4,
         "team": [
           {
             "id": 12,
             "name": "12"
           },
           {
             "id": 13,
             "name": "13"
           },
           {
             "id": 14,
             "name": "14"
           },
           {
             "id": 15,
             "name": "15"
           }
         ],
         "matches": [
           {
             "day": 1,
             "time": "23:00",
             "home": {
               "id": 12,
               "name": "12"
             },
             "away": {
               "id": 13,
               "name": "13"
             }
           },
           {
             "day": 1,
             "time": "02:00",
             "home": {
               "id": 12,
               "name": "12"
             },
             "away": {
               "id": 14,
               "name": "14"
             }
           },
           {
             "day": 2,
             "time": "23:00",
             "home": {
               "id": 12,
               "name": "12"
             },
             "away": {
               "id": 15,
               "name": "15"
             }
           },
           {
             "day": 2,
             "time": "02:00",
             "home": {
               "id": 13,
               "name": "13"
             },
             "away": {
               "id": 14,
               "name": "14"
             }
           },
           {
             "day": 3,
             "time": "23:00",
             "home": {
               "id": 13,
               "name": "13"
             },
             "away": {
               "id": 15,
               "name": "15"
             }
           },
           {
             "day": 3,
             "time": "02:00",
             "home": {
               "id": 14,
               "name": "14"
             },
             "away": {
               "id": 15,
               "name": "15"
             }
           }
         ]
       }
     ]
   }
   ```

   