DROP SCHEMA IF EXISTS nyla_treatment;
CREATE SCHEMA nyla_treatment;
USE nyla_treatment;

CREATE TABLE medicine(
  id integer unsigned not null auto_increment,
  name varchar(50),
  type varchar(255),
  time_period_hr smallint unsigned,
  PRIMARY KEY(id)
);

CREATE TABLE treatment_time(
  id integer unsigned not null auto_increment,
  recorded_time_taken datetime,
  medicine_id integer unsigned not null,
  constraint fk_treatment_time_medicine FOREIGN KEY(medicine_id) REFERENCES medicine(id),
  next_time_taken datetime,
  PRIMARY KEY(id)
);

INSERT INTO medicine(name, type, time_period_hr) VALUES(
  'Prednisolone',
  'Anti-Inflammatory Steroid',
  12
);

INSERT INTO medicine(name, type, time_period_hr) VALUES(
  'Animax',
  'Anti-infection Ointment',
  12
);

INSERT INTO medicine(name, type, time_period_hr) VALUES(
  'Clavamox',
  'Antibiotic',
  12
);
