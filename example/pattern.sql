drop table if exists t_user;
create table t_user
(
    id         serial      not null,
    name       text        not null,
    created_at timestamptz not null default now(),

    constraint t_user_pkey primary key (id)
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
    id         serial      not null,
    amount     int         not null default 0,
    created_at timestamptz not null default now(),
    constraint t_order_pkey primary key (id)
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
    constraint t_user_order_pkey primary key (user_id, order_id)
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
     , json_agg(
        json_build_object('order_id', o.id
            , 'order_amount', o.amount
            , 'order_created_at', o.created_at)
    )         as orders

from t_user u
         inner join t_user_order uo on u.id = uo.user_id
         left join t_order o on uo.order_id = o.id
group by u.id
    )
;


select json_agg(
               json_build_object('user_id', v.user_id,
                                 'user_name', v.user_name,
                                 'orders', v.orders)
           )
from user_order_view v;
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