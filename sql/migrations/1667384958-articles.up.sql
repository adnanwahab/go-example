create table machines (
  id integer primary key,
  name text not null,
) strict;

create table customers (
  id integer primary key,
  name text not null,
) strict;

create table customer_machines (
  customer_id INTEGER,
  machine_id INTEGER,
  FOREIGN KEY(customer_id) REFERENCES customers(id),
  FOREIGN KEY(machine_id) REFERENCES machines(id)
) strict;

create table metrics (
  id integer primary key,
  name text not null,
) strict;

create table machine_metric (
   metric_id INTEGER,
   machine_id INTEGER,
   FOREIGN KEY(metric_id) REFERENCES metrics(id),
   FOREIGN KEY(machine_id) REFERENCES machines(id)
) strict;