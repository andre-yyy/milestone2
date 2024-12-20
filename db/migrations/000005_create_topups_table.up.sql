create table topups (
   id serial,
   external_id varchar(255) not null,
   user_id uuid not null,
   status varchar(255) default 'UNPAID',
   created_at timestamptz default current_timestamp,
   constraint pk_topups primary key(id),
   constraint fk_topups_user foreign key(user_id) references users(id)
);