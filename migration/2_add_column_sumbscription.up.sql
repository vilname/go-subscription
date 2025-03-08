alter table subscription
    add card_uuid varchar;

comment on column subscription.cardUuid is 'Uuid карты из вебцитруса';