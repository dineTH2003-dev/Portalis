import { app } from "./app";
import { loadEnv } from "./core/env";
import { openDb } from "./core/db/db";

const env = loadEnv();

// Initialize database connection
openDb();

console.log("==================================================");
console.log(`🚀 Portalis Control Plane starting on port ${env.port}`);
console.log("==================================================");

export default {
  port: env.port,
  fetch: app.fetch,
};
