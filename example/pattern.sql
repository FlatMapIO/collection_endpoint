/*
docker run -d \
    --name mysql \
    -p 3306:3306 \
    --user 1000:1000 \
    -e MYSQL_ROOT_PASSWORD=sa \
    -e MYSQL_PASSWORD=sa \
    -e MYSQL_USER=sa \
    -e MYSQL_DATABASE=test \
    mysql:8
*/


drop table if exists t_user;
create table t_user
(
    id         int      not null auto_increment,
    name       text     not null,
    created_at datetime not null default now(),
    primary key pk_id (id)
)
;

insert into t_user (name, created_at)
values ('test1', '2012-01-01 00:00:00')
     , ('test2', '2012-01-02 00:00:00')
     , ('test3', '2012-01-03 00:00:00')
     , ('test4', '2012-01-04 00:00:00');
;



drop table if exists t_order;;
create table t_order
(
    id         int          not null auto_increment,
    amount     int unsigned not null default 0,
    created_at datetime     not null default now(),
    primary key pk_id (id)
)
;

insert into t_order (amount, created_at)
values (10, '2012-01-01 00:00:00')
     , (20, '2012-01-02 00:00:00')
     , (33, '2012-01-03 00:00:00')
     , (1123, '2012-01-04 00:00:00');
;



drop table if exists t_user_order;
create table t_user_order
(

    user_id  int not null,
    order_id int not null,
    primary key pk_user_order (user_id, order_id)
)
;



insert into t_user_order (user_id, order_id)
values (1, 1)
     , (1, 2)
     , (1, 3)
     , (2, 4);
;

-- json
create view user_order_view as
(
select u.id   as user_id
     , u.name as user_name
     , json_arrayagg(
        json_object('order_id', o.id
            , 'order_amount', o.amount
            , 'order_created_at', o.created_at)
    )         as orders

from t_user u
         inner join t_user_order uo on u.id = uo.user_id
         left join t_order o on uo.order_id = o.id
group by u.id
    )
;


with r1 as (select *
            from user_order_view v)
select count(*) as count
     , json_arrayagg(json_object(
        'user_id', r1.user_id
    , 'user_name', r1.user_name
    , 'orders', r1.orders
    ))          as data
from r1
group by null
;

-- with meta
select count(*)
from (select user_id from t_user_order where t_user_order.user_id > 1) as tuoui;


with r1 as (select user_id
                 , order_id
            from t_user_order
            where user_id > 0)
   , r2 as (select * from r1 order by user_id limit 2)
select json_object('count', count(*))
from r1
;


select json_arrayagg(json_object('id', t.id, 'name', t.name, 'created_at', t.created_at)) as data
from (select id, name, created_at from sample where 1 = 1 order by name asc, created_at asc limit 10) as t
group by 1
;

-- types

create table types_test
(
    id         int      not null auto_increment,
    name       text     not null,
    created_at datetime not null default now(),
    primary key pk_id (id)
)
;
insert into types_test (name)
values ('admin')
     , ('guest')
;

select *
from types_test
where 1 = 1
  and created_at between '2022-08-01' and '2022-08-19'
;
select *
from types_test
where 1 = 1
  and id between 1 and 2
order by id desc
       , name
       , created_at desc