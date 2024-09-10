-- Create the `tasks` table
CREATE TABLE IF NOT EXISTS tasks (
    id CHAR(36) NOT NULL PRIMARY KEY,       -- UUID for unique identification (using CHAR(36) for UUID format)
    title VARCHAR(255) NOT NULL,            -- Task title
    description TEXT,                       -- Task description
    status INT NOT NULL,                    -- Status (assuming an integer type for Status enum)
    deleted BOOLEAN NOT NULL                -- Deleted status
);