CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role ENUM('participant', 'organizer', 'admin') NOT NULL DEFAULT 'participant',
    avatar_url TEXT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    INDEX idx_users_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS events (
    id CHAR(36) PRIMARY KEY,
    slug VARCHAR(100) NOT NULL UNIQUE,
    title VARCHAR(150) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50) NOT NULL,
    status ENUM('open', 'ongoing', 'completed') NOT NULL DEFAULT 'open',
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    location VARCHAR(255) NOT NULL,
    image_url TEXT NULL,
    prize_pool VARCHAR(100) NOT NULL,
    max_teams INT NOT NULL DEFAULT 32,
    current_teams INT NOT NULL DEFAULT 0,
    registration_fee DECIMAL(12, 2) NOT NULL DEFAULT 0.00,
    early_bird_price DECIMAL(12, 2) NULL,
    early_bird_quota INT NULL,
    registration_deadline DATETIME(6) NOT NULL,
    tm_date DATETIME(6) NULL,
    organizer_id CHAR(36) NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    CONSTRAINT fk_events_organizer
        FOREIGN KEY (organizer_id) REFERENCES users(id)
        ON DELETE CASCADE,
    INDEX idx_events_category_status (category, status),
    INDEX idx_events_organizer_id (organizer_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS prizes (
    id CHAR(36) PRIMARY KEY,
    event_id CHAR(36) NOT NULL,
    `rank` VARCHAR(50) NOT NULL,
    reward VARCHAR(100) NOT NULL,
    CONSTRAINT fk_prizes_event
        FOREIGN KEY (event_id) REFERENCES events(id)
        ON DELETE CASCADE,
    INDEX idx_prizes_event_id (event_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS rundown_items (
    id CHAR(36) PRIMARY KEY,
    event_id CHAR(36) NOT NULL,
    label VARCHAR(150) NOT NULL,
    start_time DATETIME(6) NOT NULL,
    end_time DATETIME(6) NULL,
    CONSTRAINT fk_rundown_items_event
        FOREIGN KEY (event_id) REFERENCES events(id)
        ON DELETE CASCADE,
    INDEX idx_rundown_items_event_id (event_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS teams (
    id CHAR(36) PRIMARY KEY,
    event_id CHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL,
    logo_url TEXT NULL,
    captain_id CHAR(36) NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    CONSTRAINT fk_teams_event
        FOREIGN KEY (event_id) REFERENCES events(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_teams_captain
        FOREIGN KEY (captain_id) REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT unique_team_name_per_event UNIQUE (event_id, name),
    INDEX idx_teams_event_id (event_id),
    INDEX idx_teams_captain_id (captain_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS team_members (
    id CHAR(36) PRIMARY KEY,
    team_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    role VARCHAR(30) NOT NULL DEFAULT 'member',
    joined_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    CONSTRAINT fk_team_members_team
        FOREIGN KEY (team_id) REFERENCES teams(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_team_members_user
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT unique_user_per_team UNIQUE (team_id, user_id),
    INDEX idx_team_members_team_id (team_id),
    INDEX idx_team_members_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS rounds (
    id CHAR(36) PRIMARY KEY,
    event_id CHAR(36) NOT NULL,
    title VARCHAR(100) NOT NULL,
    round_number INT NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    CONSTRAINT fk_rounds_event
        FOREIGN KEY (event_id) REFERENCES events(id)
        ON DELETE CASCADE,
    CONSTRAINT unique_round_number_per_event UNIQUE (event_id, round_number),
    INDEX idx_rounds_event_id (event_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS matches (
    id CHAR(36) PRIMARY KEY,
    event_id CHAR(36) NOT NULL,
    round_id CHAR(36) NOT NULL,
    team1_id CHAR(36) NULL,
    team2_id CHAR(36) NULL,
    team1_score INT NOT NULL DEFAULT 0,
    team2_score INT NOT NULL DEFAULT 0,
    winner_id CHAR(36) NULL,
    status ENUM('scheduled', 'live', 'completed') NOT NULL DEFAULT 'scheduled',
    start_time DATETIME(6) NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    CONSTRAINT fk_matches_event
        FOREIGN KEY (event_id) REFERENCES events(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_matches_round
        FOREIGN KEY (round_id) REFERENCES rounds(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_matches_team1
        FOREIGN KEY (team1_id) REFERENCES teams(id)
        ON DELETE SET NULL,
    CONSTRAINT fk_matches_team2
        FOREIGN KEY (team2_id) REFERENCES teams(id)
        ON DELETE SET NULL,
    CONSTRAINT fk_matches_winner
        FOREIGN KEY (winner_id) REFERENCES teams(id)
        ON DELETE SET NULL,
    INDEX idx_matches_event_id (event_id),
    INDEX idx_matches_round_id (round_id),
    INDEX idx_matches_team1_id (team1_id),
    INDEX idx_matches_team2_id (team2_id),
    INDEX idx_matches_winner_id (winner_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS submissions (
    id CHAR(36) PRIMARY KEY,
    event_id CHAR(36) NOT NULL,
    team_id CHAR(36) NOT NULL,
    submitted_by CHAR(36) NOT NULL,
    file_url TEXT NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_size INT NOT NULL,
    status ENUM('pending', 'reviewed', 'approved', 'rejected') NOT NULL DEFAULT 'pending',
    feedback TEXT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    CONSTRAINT fk_submissions_event
        FOREIGN KEY (event_id) REFERENCES events(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_submissions_team
        FOREIGN KEY (team_id) REFERENCES teams(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_submissions_submitter
        FOREIGN KEY (submitted_by) REFERENCES users(id)
        ON DELETE CASCADE,
    INDEX idx_submissions_event_id (event_id),
    INDEX idx_submissions_team_id (team_id),
    INDEX idx_submissions_submitted_by (submitted_by)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS chat_messages (
    id CHAR(36) PRIMARY KEY,
    channel_name VARCHAR(100) NOT NULL,
    sender_id CHAR(36) NOT NULL,
    content TEXT NOT NULL,
    `timestamp` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    CONSTRAINT fk_chat_messages_sender
        FOREIGN KEY (sender_id) REFERENCES users(id)
        ON DELETE CASCADE,
    INDEX idx_chat_channel_timestamp (channel_name, `timestamp` DESC),
    INDEX idx_chat_messages_sender_id (sender_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
