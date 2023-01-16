CREATE TABLE doctors
(
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE specializations
(
    id serial not null unique,
    name varchar(255) not null
);

CREATE TABLE patients
(
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255) not null,
    birthdate date not null
);

CREATE TABLE visits
(
    id serial not null unique,
    docId int references doctors(id) on delete cascade not null,
    patientId int references patients(id) on delete cascade not null
);

CREATE TABLE doc_spec
(
    id serial not null unique,
    docId int references doctors(id) on delete cascade not null,
    specId int references specializations(id) on delete cascade not null
);