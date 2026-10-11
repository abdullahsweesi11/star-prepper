CREATE TABLE IF NOT EXISTS roles (
    role_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role_name varchar(255) NOT NULL UNIQUE
);

INSERT INTO roles (role_name)
VALUES
('admin'),
('user'),
('guest')
ON CONFLICT (role_name) DO NOTHING;

CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email varchar(255) NOT NULL,
    hashed_password varchar(255) NOT NULL,
    role_id UUID REFERENCES roles(role_id) NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS stories (
    story_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title varchar(255) NOT NULL,
    user_id UUID REFERENCES users(user_id) NOT NULL,
    default_situation text NOT NULL,
    default_result text NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS narratives (
    narrative_id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    label varchar(255) NOT NULL,
    story_id UUID REFERENCES stories(story_id) NOT NULL,
    situation_override text,
    task text NOT NULL,
    "action" text NOT NULL,
    result_override text,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS competencies (
    competency_id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    competency_name VARCHAR(255) NOT NULL,
    user_id UUID REFERENCES users(user_id),
    custom boolean NOT NULL DEFAULT TRUE,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS narratives_competencies (
    narrative_id UUID REFERENCES narratives(narrative_id) NOT NULL,
    competency_id UUID REFERENCES competencies(competency_id) NOT NULL,
    PRIMARY KEY(narrative_id, competency_id)
);


CREATE OR REPLACE FUNCTION mark_updated()
RETURNS trigger AS $$
    BEGIN
        NEW.updated_at := CURRENT_TIMESTAMP;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_mark_updated_stories ON stories;
CREATE TRIGGER trg_mark_updated_stories BEFORE UPDATE ON stories
    FOR EACH ROW
    EXECUTE FUNCTION mark_updated();

DROP TRIGGER IF EXISTS trg_mark_updated_narratives ON narratives;
CREATE TRIGGER trg_mark_updated_narratives BEFORE UPDATE ON narratives
    FOR EACH ROW
    EXECUTE FUNCTION mark_updated();

DROP TRIGGER IF EXISTS trg_mark_updated_competencies ON competencies;
CREATE TRIGGER trg_mark_updated_competencies BEFORE UPDATE ON competencies
    FOR EACH ROW
    EXECUTE FUNCTION mark_updated();


CREATE OR REPLACE FUNCTION set_default_role()
RETURNS TRIGGER AS $$
    BEGIN
        IF NEW.role_id IS NULL THEN
            SELECT role_id INTO NEW.role_id 
            FROM roles 
            WHERE role_name = 'user' 
            LIMIT 1;
            
            IF NEW.role_id IS NULL THEN
                RAISE EXCEPTION 'The default role "user" does not exist';
            END IF;
        END IF;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS tgr_set_default_role ON users;
CREATE TRIGGER tgr_set_default_role BEFORE INSERT ON users
    FOR EACH ROW
    EXECUTE FUNCTION set_default_role();