import { Kysely, SqliteDialect } from "kysely";
import database from ".";
import { Database } from "./schema";

const dialect = new SqliteDialect({ database: database });
const db = new Kysely<Database>({ dialect });

export default db;
