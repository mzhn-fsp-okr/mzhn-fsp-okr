import db from "../db/kysely";
import { NewUser } from "../db/schema/users-table";

class UserService {
  async save(user: NewUser) {
    console.log("saving user", user);
    const u = await this.find(user.username);

    if (u) {
      return await this.update(user);
    }

    const res = await db.insertInto("users").values(user).executeTakeFirst();
    return res;
  }

  async find(username: string) {
    const res = await db
      .selectFrom("users")
      .selectAll()
      .where("username", "=", username)
      .executeTakeFirst();
    console.log("finding user", username, res);
    return res;
  }

  async update(user: NewUser) {
    const res = await db
      .updateTable("users")
      .set(user)
      .where("chat_id", "=", user.chat_id)
      .executeTakeFirst();
    return res;
  }
}

export const userService = new UserService();
