-- Connect to postgres in terminal and run the following commands

create database todoapp;

\c todoapp;

create table tasks (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    isDone BOOLEAN NOT NULL DEFAULT FALSE,
    created TIMESTAMP NOT NULL DEFAULT now()
);

insert into tasks(content) values('Wash the dog');
insert into tasks(content) values('Take out the trash');
insert into tasks(content) values('Go for a walk after dinner');
