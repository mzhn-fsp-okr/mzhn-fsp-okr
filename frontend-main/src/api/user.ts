export interface User {
  id: string;
  lastName?: string;
  firstName?: string;
  middleName?: string;
  email: string;
  roles: string[];
}

const roles: Record<string, string> = {
  admin: "Администратор",
};

export function getAvatar(user: User) {
  return `https://api.dicebear.com/9.x/micah/svg?baseColor=f9c9b6&earringColor[]&earrings[]&earringsProbability=0&ears=detached,attached&facialHair=beard&facialHairProbability=5&hair=dannyPhantom,fonze,full,pixie,turban,mrT&backgroundColor=b6e3f4,c0aede,d1d4f9,ffd5dc,ffdfbf&seed=${user.email}`;
}

export function getFullName(user: User) {
  const parts = [user.firstName, user.middleName, user.lastName];
  if (parts.every((v) => !v)) return null;
  return parts.filter((v) => v != null).join(" ");
}

export function getRoleFriendly(user: User) {
  if (user.roles.length == 0) return "Пользователь";
  const roleKey = user.roles[0];
  if (Object.hasOwn(roles, roleKey)) return roles[roleKey];
  return "Неизвестно";
}
