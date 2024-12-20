create table activities (
   id serial,
   user_id uuid not null,
   description text not null,
   created_at timestamptz default current_timestamp,
   constraint pk_activities primary key(id),
   constraint fk_activities_user foreign key(user_id) references users(id)
);