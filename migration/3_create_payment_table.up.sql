create table payment
(
    id             serial
        constraint payment_pk
            primary key,
    sum            integer not null,
    subscriptionId integer
        constraint payment_subscription_id_fk
            references subscription,
    created         timestamp(6) with time zone,
    updated         timestamp(6) with time zone
);