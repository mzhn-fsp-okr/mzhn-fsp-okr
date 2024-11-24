import {
  daysUntil,
  formatDate,
  formatDays,
  formatParticipantRequirements,
} from "./lib";

export const formatNewEventMessage = (event) => {
  const daysLeft = daysUntil(event.dates.from);
  return `Новое событие для фанатов ${event.sportSubtype.sportType.name} уже в календаре! 🔥 ${event.name}
  
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
${formatParticipantRequirements(event.participantRequirements)}`;
};
