-- Here, you can execute any CQL commands, e.g.

CREATE KEYSPACE oauth WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};


-- CREATE TABLE oauth.tweet (
--     id UUID,
--     timeline text, 
--     text text,
--     PRIMARY KEY(id));


-- CREATE INDEX IF NOT EXISTS timeline_index
-- ON oauth.tweet (timeline);

CREATE TABLE oauth.access_tokens ( 
    access_token varchar, 
    user_id bigint, 
    client_id bigint, 
    expires bigint,
    PRIMARY KEY(access_token));

CREATE INDEX IF NOT EXISTS access_token_index
ON oauth.access_tokens (access_token);

CREATE INDEX IF NOT EXISTS userid_index
ON oauth.access_tokens (user_id);