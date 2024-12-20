create table orders (
   id serial,
   user_id uuid not null,
   product_id int not null,
   start_date timestamptz,
   rent_days int,
   end_date timestamptz,
   total_price decimal(10, 2) not null,
   destination_id varchar(100) not null,
   status varchar(100) default 'pending',
   created_at timestamptz default current_timestamp,
   updated_at timestamptz default current_timestamp,
   constraint pk_orders primary key(id),
   constraint fk_orders_user foreign key(user_id) references users(id) on delete cascade,
   constraint fk_orders_product foreign key(product_id) references products(id) on delete cascade
);