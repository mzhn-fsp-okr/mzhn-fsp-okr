create extension if not exists "uuid-ossp";

create table if not exists integrations(
  user_id uuid primary key,
  telegram_username varchar,
  wanna_mail boolean not null default true
);

create table if not exists verification (
  user_id uuid primary key,
  token varchar not null,
  created_at timestamp default now()
);