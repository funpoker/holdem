CREATE TABLE player (
        id INTEGER NOT NULL,
        username VARCHAR(20) NOT NULL,
        password VARCHAR(64) NOT NULL,
        avatar_url VARCHAR(256),
        chip INTEGER,
        created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        status BOOLEAN NOT NULL,
        PRIMARY KEY (id),
        UNIQUE (username),
        CHECK (status IN (0, 1))
)

CREATE TABLE room (
        id INTEGER NOT NULL,
        name VARCHAR(20) NOT NULL,
        background_url VARCHAR(256),
        created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        status BOOLEAN NOT NULL,
        PRIMARY KEY (id),
        UNIQUE (name),
        CHECK (status IN (0, 1))
)

CREATE TABLE game (
        id INTEGER NOT NULL,
        room_id INTEGER NOT NULL,
        period INTEGER NOT NULL,
        round INTEGER NOT NULL,
        flop VARCHAR(16),
        turn VARCHAR(16),
        river VARCHAR(16),
        begin_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        end_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        status BOOLEAN NOT NULL,
        PRIMARY KEY (id),
        CHECK (status IN (0, 1))
)

CREATE TABLE relation_game_player (
        id INTEGER NOT NULL,
        game_id INTEGER NOT NULL,
        player_id INTEGER NOT NULL,
        player_role INTEGER NOT NULL,
        player_position VARCHAR(16),
        hands VARCHAR(16),
        created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        status BOOLEAN NOT NULL,
        PRIMARY KEY (id),
        CHECK (status IN (0, 1))
)

CREATE TABLE game_record (
        id INTEGER NOT NULL,
        game_id INTEGER NOT NULL,
        player_id INTEGER NOT NULL,
        round INTEGER NOT NULL,
        phase INTEGER NOT NULL,
        amount INTEGER NOT NULL,
        created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        PRIMARY KEY (id)
)
