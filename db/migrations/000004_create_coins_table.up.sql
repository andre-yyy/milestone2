create table coins (
   id serial,
   name varchar(255) not null,
   description text not null,
   price decimal(10, 2) not null,
   constraint pk_coins primary key(id)
);

insert into coins (name, description, price) values
('10 coins', 'increase your deposit by 10', 10),
('25 coins', 'increase your deposit by 25', 25),
('50 coins', 'increase your deposit by 50', 50),
('75 coins', 'increase your deposit by 75', 75);
