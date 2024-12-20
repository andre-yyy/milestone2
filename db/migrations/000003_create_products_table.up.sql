create table products (
   id serial,
   name varchar(255) not null,
   stock int default 0,
   rental_cost decimal(10, 2) not null,
   daily decimal(10, 2),
   weight int default 1,
   category_id int not null,
   image_url text default 'https://www.hacktiv8.com/_next/image?url=%2Flogo.png&w=1920&q=75',
   description text,
   constraint pk_products primary key(id),
   constraint fk_products_category foreign key(category_id) references categories(id)
);

insert into products (name, stock, rental_cost, daily, weight, category_id, description) values
('Wheelchair One', 10, 10, 2.5, 500, 1, 'Wheelchair One Description'),
('Wheelchair Two', 6, 15, 2.5, 750, 1, 'Wheelchair Two Description'),
('Wheelchair Three', 4, 20, 3, 1000, 1, 'Wheelchair Three Description'),
('Walking Aid One', 15, 5, 1, 350, 2, 'Walking Aid One Description'),
('Walking Aid Two', 8, 8, 1.5, 500, 2, 'Walking Aid Two Description'),
('Walking Aid Three', 3, 12, 2, 750, 2, 'Walking Aid Three description'),
('Bed One', 3, 50, 2.5, 3500, 3, 'Bed One description'),
('Bed Two', 2, 60, 3, 3800, 3, 'Bed Two description'),
('Bed Three', 1, 75, 3.5, 4000, 3, 'Bed Three description');
