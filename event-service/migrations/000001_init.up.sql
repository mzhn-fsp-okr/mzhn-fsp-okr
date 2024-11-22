create extension if not exists "uuid-ossp";

create table if not exists sport_types (
  id uuid primary key default uuid_generate_v4(),
  sport_type varchar not null
);

create table if not exists sport_subtypes (
  id uuid primary key default uuid_generate_v4(),
  sport_subtype varchar not null,
  sport_type_id uuid not null references sport_types(id) on delete cascade
);

create table if not exists events (
  id uuid primary key default uuid_generate_v4(),
  ekp_id varchar not null,
  sport_subtype_id uuid not null references sport_subtypes(id) on delete cascade,
  name varchar not null,
  description varchar,
  location varchar not null,
  participants int not null,
  created_at timestamp default now(),
  updated_at timestamp
);

create table if not exists event_dates (
  event_id uuid primary key references events(id) on delete cascade,
  date_from timestamp not null,
  date_to timestamp not null
);

create table if not exists event_participants_requirements (
  id serial primary key,
  event_id uuid references events(id) on delete cascade,
  gender boolean not null,
  min_age int,
  max_age int
);