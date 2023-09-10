DROP VIEW IF EXISTS improve_requests_previews;
DROP VIEW IF EXISTS improve_requests_latest_revisions;
DROP VIEW IF EXISTS improve_requests_revisions_list;

--bun:split

DROP INDEX IF EXISTS improve_requests_source;
DROP INDEX IF EXISTS improve_requests_user;
DROP INDEX IF EXISTS improve_requests_fts;
DROP INDEX IF EXISTS improve_requests_last_rev;
DROP INDEX IF EXISTS improve_suggestions_source;
DROP INDEX IF EXISTS improve_suggestions_request;
DROP INDEX IF EXISTS improve_suggestions_user;

--bun:split

DROP TRIGGER IF EXISTS format_searchable_content ON improve_requests_revisions;

--bun:split

DROP FUNCTION IF EXISTS format_searchable_content;

--bun:split

DROP TABLE IF EXISTS improve_suggestions;
DROP TABLE IF EXISTS improve_requests_revisions;
DROP TABLE IF EXISTS improve_requests;

--bun:split

DROP EXTENSION IF EXISTS unaccent;
DROP EXTENSION IF EXISTS pg_trgm;
