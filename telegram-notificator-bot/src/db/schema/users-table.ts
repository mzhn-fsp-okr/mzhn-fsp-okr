import { Insertable, Selectable, Updateable } from "kysely";

export interface UsersTable {
  chat_id: string;
  username: string;
}

export type User = Selectable<UsersTable>;
export type NewUser = Insertable<UsersTable>;
export type UserUpdate = Updateable<UsersTable>;
