CREATE TABLE users (
    id uuid NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    image_name VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT users_pk PRIMARY KEY (id),
    CONSTRAINT email_unique UNIQUE (email)
);

CREATE TABLE tags (
    id uuid NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    CONSTRAINT tags_pk PRIMARY KEY (id)
);

CREATE TABLE roles (
    id uuid NOT NULL ,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    CONSTRAINT roles_pk PRIMARY KEY (id)
);

CREATE TABLE topics (
    id uuid NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    image_name VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    user_id uuid,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT topics_pk PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

   
CREATE TABLE user_roles (
    user_id uuid,
    role_id uuid,
    CONSTRAINT user_roles_pk PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE questions (
    id uuid NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    user_id uuid,
    topic_id uuid,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT questions_pk PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (topic_id) REFERENCES topics(id)
);

CREATE TABLE answers (
    id uuid NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    user_id uuid,
    question_id uuid,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT answers_pk PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (question_id) REFERENCES questions(id)
);

CREATE TABLE question_tags (
    question_id uuid,
    tag_id uuid,
    CONSTRAINT question_tags_pk PRIMARY KEY (question_id, tag_id),
    FOREIGN KEY (question_id) REFERENCES questions(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);

CREATE TABLE user_votes_question (
    user_id uuid,
    question_id uuid,
    vote_type TEXT NOT NULL CHECK (vote_type IN ('upvote', 'downvote')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_votes_question_pk PRIMARY KEY (user_id, question_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (question_id) REFERENCES questions(id)
);

CREATE TABLE user_votes_answer (
    user_id uuid,
    answer_id uuid,
    vote_type TEXT NOT NULL CHECK (vote_type IN ('upvote', 'downvote')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_votes_answer_pk PRIMARY KEY (user_id, answer_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (answer_id) REFERENCES answers(id)
);

-- CREATE INDEX idx_questions_user_id ON questions(user_id);
-- CREATE INDEX idx_questions_topic_id ON questions(topic_id);
-- CREATE INDEX idx_answers_user_id ON answers(user_id);
-- CREATE INDEX idx_answers_question_id ON answers(question_id);
-- CREATE INDEX idx_topics_user_id ON topics(user_id);
-- CREATE INDEX idx_question_tags_question_id ON question_tags(question_id);
-- CREATE INDEX idx_question_tags_tag_id ON question_tags(tag_id);
-- CREATE INDEX idx_user_votes_question_user_id ON user_votes_question(user_id);
-- CREATE INDEX idx_user_votes_question_question_id ON user_votes_question(question_id);
-- CREATE INDEX idx_user_votes_answer_user_id ON user_votes_answer(user_id);
-- CREATE INDEX idx_user_votes_answer_answer_id ON user_votes_answer(answer_id);
-- CREATE INDEX idx_user_roles_user_id ON user_roles(user_id);
-- CREATE INDEX idx_user_roles_role_id ON user_roles(role_id);