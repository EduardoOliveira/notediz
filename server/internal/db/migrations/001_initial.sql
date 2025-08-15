-- +goose Up
-- +goose StatementBegin
CREATE TABLE note (
	id char(36) PRIMARY KEY,
	kind char(20) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
);

CREATE TABLE file (
	id char(36) PRIMARY KEY,
	note_id char(36) NOT NULL,
	name text NOT NULL,
	path text NOT NULL,
	size bigint NOT NULL,
	mime_type text NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	FOREIGN KEY (note_id) REFERENCES note(id) ON DELETE CASCADE
);

CREATE TABLE text (
	id char(36) PRIMARY KEY,
	note_id char(36) NOT NULL,
	content text NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	FOREIGN KEY (note_id) REFERENCES note(id) ON DELETE CASCADE
);

CREATE TABLE image (
	id char(36) PRIMARY KEY,
	note_id char(36) NOT NULL,
	url text NOT NULL,
	alt text,
	width integer NOT NULL,
	height integer NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	FOREIGN KEY (note_id) REFERENCES note(id) ON DELETE CASCADE
);

CREATE TABLE audio (
	id char(36) PRIMARY KEY,
	note_id char(36) NOT NULL,
	url text NOT NULL,
	title text NOT NULL,
	duration integer NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	FOREIGN KEY (note_id) REFERENCES note(id) ON DELETE CASCADE
);

CREATE TABLE video (
	id char(36) PRIMARY KEY,
	note_id char(36) NOT NULL,
	url text NOT NULL,
	title text NOT NULL,
	duration integer NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	FOREIGN KEY (note_id) REFERENCES note(id) ON DELETE CASCADE
);

CREATE TABLE bookmark (
	id char(36) PRIMARY KEY,
	note_id char(36) NOT NULL,
	url text NOT NULL,
	title text NOT NULL,
	description text NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	FOREIGN KEY (note_id) REFERENCES note(id) ON DELETE CASCADE
);

CREATE TABLE todo (
	id char(36) PRIMARY KEY,
	note_id char(36) NOT NULL,
	content text NOT NULL,
	completed boolean NOT NULL DEFAULT false,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	FOREIGN KEY (note_id) REFERENCES note(id) ON DELETE CASCADE
);


CREATE TABLE tag (
	id char(36) PRIMARY KEY,
	name text NOT NULL,
	color char(7) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
);

CREATE TABLE note_tag (
	note_id char(36) NOT NULL,
	tag_id char(36) NOT NULL,
	PRIMARY KEY (note_id, tag_id),
	FOREIGN KEY (note_id) REFERENCES note(id) ON DELETE CASCADE,
	FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE
);

CREATE TABLE connection (
	id char(36) PRIMARY KEY,
	src_note_id char(36) NOT NULL,
	dst_note_id char(36) NOT NULL,
	type text NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	UNIQUE (src, dst, type),
	FOREIGN KEY (src_note_id) REFERENCES note(id) ON DELETE CASCADE,
	FOREIGN KEY (dst_note_id) REFERENCES note(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
