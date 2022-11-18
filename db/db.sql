CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "questions" (
  "id" SERIAL PRIMARY KEY,
  "question" varchar NOT NULL,
  "answer" varchar,
  "user_id" int
);

ALTER TABLE "questions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

insert into users (name) values ('francisco');
insert into questions (question, answer, user_id) values ('question', 'answer', 1);

commit;