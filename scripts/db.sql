-- Connect to postgres in terminal and run the following commands

create database todoapp;

\c todoapp;

create table tasks
(
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    isDone BOOLEAN NOT NULL DEFAULT FALSE,
    assignee TEXT NOT NULL,
    deadline DATE NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT now()
);

insert into tasks
    (content, assignee, deadline)
values('Create a proposal for next meet', 'Harun', '2022-09-12');
insert into tasks
    (content, assignee, deadline)
values('Create a presentation for morning briefing', 'Rudy', '2022-08-30');
insert into tasks
    (content, assignee, deadline)
values('Check the journal', 'Siska', '2022-10-12');