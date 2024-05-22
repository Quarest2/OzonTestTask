CREATE TABLE IF NOT EXISTS public.posts
(
    id              SERIAL PRIMARY KEY,
    author          TEXT,
    CommentsAllowed BOOLEAN
);

CREATE TABLE IF NOT EXISTS public.posts
(
    id              SERIAL PRIMARY KEY,
    author          TEXT,
    parentPostID    SERIAL,
    ParentCommentID SERIAL
);