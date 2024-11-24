import { formatDate } from "./lib";

export const formatNewSubMessage = (event) => {
  return `🎉 Вы успешно подписались на событие!

📌 "${event.name}"
📅 *Дата проведения:* ${formatDate(event.dates.from)} - ${formatDate(event.dates.to)}
📍 *Место проведения:* ${event.location}

Мы отправим вам напоминание о начале события! ⏰`;
};

export const formatSportSubMessage = (sport) => {
  return `🎉 Вы успешно начали отслеживать вид спорта *"${sport.name}"*!`;
};
