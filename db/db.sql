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

insert into users values (1, 'francisco');
insert into questions values (1, 'question', 'answer', 1);