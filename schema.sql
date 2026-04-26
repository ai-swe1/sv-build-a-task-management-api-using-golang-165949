-- SQLite schema for task management
-- Table: tasks
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    due_date TEXT,
    completed INTEGER NOT NULL DEFAULT 0
);
