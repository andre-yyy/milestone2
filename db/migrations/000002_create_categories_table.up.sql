create table categories (
   id serial,
   name varchar(255) not null unique,
   constraint pk_categories primary key(id)
);

insert into categories (name) values
('Wheelchairs'),
('Walking Aids'),
('Bed');