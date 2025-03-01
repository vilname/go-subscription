create table subscription
(
    id      serial not null
        constraint subscription_pk
            primary key,
    email   varchar,
    hash    varchar,
    created         timestamp(6) with time zone,
    updated         timestamp(6) with time zone
);