CREATE EXTENSION IF NOT EXISTS unaccent;
CREATE EXTENSION IF NOT EXISTS pg_trgm;

--bun:split

CREATE TABLE IF NOT EXISTS improve_requests (
    id uuid PRIMARY KEY NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,

    up_votes BIGINT,
    down_votes BIGINT
);

CREATE TABLE IF NOT EXISTS improve_requests_revisions (
    id uuid PRIMARY KEY NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,

    source_id uuid NOT NULL,
    user_id uuid NOT NULL,
    title VARCHAR(256) NOT NULL,
    content TEXT NOT NULL,
    text_searchable_index_col tsvector,

    CONSTRAINT title_filled CHECK ( title <> '' ),
    CONSTRAINT content_filled CHECK ( content <> '' ),
    CONSTRAINT content_length CHECK ( char_length(content) <= 4096 )
);

CREATE TABLE IF NOT EXISTS improve_suggestions (
    id uuid PRIMARY KEY NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,

    user_id uuid NOT NULL,
    source_id uuid NOT NULL,
    request_id uuid NOT NULL,
    validated boolean,

    title VARCHAR(256) NOT NULL,
    content TEXT NOT NULL,

    up_votes BIGINT,
    down_votes BIGINT,

    CONSTRAINT title_filled CHECK ( title <> '' ),
    CONSTRAINT content_filled CHECK ( content <> '' ),
    CONSTRAINT content_length CHECK ( char_length(content) <= 4096 )
);

--bun:split

/* unaccent cannot be used in a standard constraint, as it is not immutable */
CREATE FUNCTION format_searchable_content()
    RETURNS TRIGGER AS $format_searchable_content$
BEGIN
    NEW.text_searchable_index_col :=
        setweight(to_tsvector('french',  unaccent(NEW.title)), 'A') ||
        setweight(to_tsvector('french', unaccent(NEW.content)), 'B');
    RETURN NEW;
END;
$format_searchable_content$ LANGUAGE plpgsql;

--bun:split

CREATE TRIGGER format_searchable_content
    BEFORE INSERT ON improve_requests_revisions
    FOR EACH ROW
EXECUTE FUNCTION format_searchable_content();

--bun:split

CREATE INDEX IF NOT EXISTS improve_requests_source ON improve_requests_revisions (source_id);
CREATE INDEX IF NOT EXISTS improve_requests_user ON improve_requests_revisions (user_id);
CREATE INDEX IF NOT EXISTS improve_requests_fts ON improve_requests_revisions USING GIN (text_searchable_index_col);
CREATE INDEX IF NOT EXISTS improve_requests_last_rev ON improve_requests_revisions (source_id, created_at DESC NULLS LAST);

CREATE INDEX IF NOT EXISTS improve_suggestions_source ON improve_suggestions (source_id);
CREATE INDEX IF NOT EXISTS improve_suggestions_request ON improve_suggestions (request_id);
CREATE INDEX IF NOT EXISTS improve_suggestions_user ON improve_suggestions (user_id);

--bun:split

CREATE VIEW improve_requests_revisions_list AS
    SELECT
        improve_requests_revisions.id,
        improve_requests_revisions.created_at,
        improve_requests_revisions.updated_at,
        improve_requests_revisions.source_id,
        suggestions.total AS suggestions_count,
        accepted_suggestions.total AS accepted_suggestions_count
    FROM improve_requests_revisions
    LEFT JOIN LATERAL (
        SELECT COUNT(*) AS total FROM improve_suggestions
            WHERE improve_suggestions.request_id = improve_requests_revisions.id
    ) AS suggestions ON TRUE
    LEFT JOIN LATERAL (
        SELECT COUNT(*) AS total FROM improve_suggestions
            WHERE improve_suggestions.request_id = improve_requests_revisions.id AND improve_suggestions.validated = TRUE
    ) AS accepted_suggestions ON TRUE;

CREATE VIEW improve_requests_latest_revisions AS
    SELECT DISTINCT ON (source_id) *
    FROM improve_requests_revisions
    ORDER BY source_id, created_at DESC NULLS LAST;

CREATE VIEW improve_requests_previews AS
SELECT
    improve_requests.id,
    improve_requests.created_at,
    improve_requests.updated_at,
    improve_requests.up_votes,
    improve_requests.down_votes,
    improve_requests_latest_revisions.title AS title,
    improve_requests_latest_revisions.content AS content,
    improve_requests_latest_revisions.user_id AS user_id,
    improve_requests_latest_revisions.text_searchable_index_col AS text_searchable_index_col,
    suggestions.total AS suggestions_count,
    accepted_suggestions.total AS accepted_suggestions_count,
    revisions.total AS revisions_count
FROM improve_requests
    LEFT JOIN improve_requests_latest_revisions ON improve_requests_latest_revisions.source_id = improve_requests.id
    LEFT JOIN LATERAL (
        SELECT COUNT(*) AS total FROM improve_suggestions
            WHERE improve_suggestions.source_id = improve_requests.id
    ) AS suggestions ON TRUE
    LEFT JOIN LATERAL (
        SELECT COUNT(*) AS total FROM improve_suggestions
            WHERE improve_suggestions.source_id = improve_requests.id AND improve_suggestions.validated = TRUE
    ) AS accepted_suggestions ON TRUE
    LEFT JOIN LATERAL (
        SELECT COUNT(improve_requests_revisions.id) AS total
        FROM improve_requests_revisions
        WHERE improve_requests_revisions.source_id = improve_requests.id
    ) AS revisions ON TRUE;
