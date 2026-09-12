import { Database } from "bun:sqlite";

let dbInstance: Database | null = null;

/**
 * Initializes and returns the embedded SQLite database instance.
 *
 * TODO(contributor): [Issue #3]
 *   - Configure path via DATABASE_URL or default to "./portalis.db"
 *   - Enforce mandatory SQLite Pragmas:
 *     - PRAGMA journal_mode = WAL;
 *     - PRAGMA foreign_keys = ON;
 *     - PRAGMA busy_timeout = 5000;
 *     - PRAGMA synchronous = NORMAL;
 */
export function openDb(): Database {
  if (!dbInstance) {
    const dbPath = process.env.DATABASE_URL || "./portalis.db";
    dbInstance = new Database(dbPath);

    // Enforce WAL mode and foreign key integrity
    dbInstance.exec("PRAGMA journal_mode = WAL;");
    dbInstance.exec("PRAGMA foreign_keys = ON;");
    dbInstance.exec("PRAGMA busy_timeout = 5000;");
    dbInstance.exec("PRAGMA synchronous = NORMAL;");
  }
  return dbInstance;
}

export function closeDb(): void {
  if (dbInstance) {
    dbInstance.close();
    dbInstance = null;
  }
}
