create extension if not exists "uuid-ossp";

create table if not exists integrations(
  user_id uuid primary key,
  telegram_username varchar,
  wanna_mail boolean default true
);