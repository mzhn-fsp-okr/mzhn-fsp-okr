create table if not exists event_stale (
  event_id uuid primary key references events(id) on delete cascade,
  is_stale boolean not null default false
);