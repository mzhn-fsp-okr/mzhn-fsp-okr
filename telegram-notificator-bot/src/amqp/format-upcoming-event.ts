import {
  daysUntil,
  formatDate,
  formatDays,
  formatParticipantRequirements,
} from "./lib";

export const formatMessage = (event) => {
  const daysLeft = daysUntil(event.dates.from);
  const formattedMessage = `🔔 Напоминание о предстоящем событии!

📅 *${event.name}* ${formatDays(daysLeft)}

📖 *Описание:*
${event.description}

🏅 *Вид спорта:* 
${event.sportSubtype.sportType.name} - ${event.sportSubtype.name}

📅 *Дата проведения:* 
${formatDate(event.dates.from)} - ${formatDate(event.dates.to)}

📍 *Место проведения:* 
${event.location}

👥 *Количество участников:* ${event.participants}

✅ *Участвовать могут:*
${formatParticipantRequirements(event.participantRequirements)}

`;

  return formattedMessage;
};
