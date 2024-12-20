create table users (
   id uuid default gen_random_uuid(),
   full_name varchar(255) not null,
   role varchar(100) default 'user',
   email varchar(255) not null unique,
   password varchar(255) not null,
   address varchar(255),
   city_id varchar(100),
   deposit decimal(10, 2) default 0.0,
   created_at timestamptz default current_timestamp,
   updated_at timestamptz default current_timestamp,
   constraint pk_users primary key(id)
);

insert into users (full_name, role, email, password) values
('User One', 'admin', 'user1@mail.com', '$2a$10$wByAwQ.vqytrn.W4NIHfq.IHAPUpTlm2zEUFcV04sTrKhDjLYzLAu'),
('User Two', 'user', 'user2@mail.com', '$2a$10$v9ZqNrVjLKuQA4txP4WFeeNFEgp3cFBSQbL1GYsfB6xkIMhbwlEVm');