const formatMessage = (event) => {
  const formatDate = (dateString) => {
    const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: 'long', day: 'numeric' };
    return new Date(dateString).toLocaleDateString('ru-RU', options);
  };

  const formatParticipantRequirements = (requirements) => {
    return requirements.map((req, index) => {
      let reqText = ``;
      const details = [];

      if (req.gender !== undefined) {
        const genderText = req.gender === true ? '👨‍🦱 Мужчины' : req.gender === false ? '👱‍♀️ Женщины' : '👨‍🦱👱‍♀️ Мужчины и женщины';
        details.push(`- ${genderText}`);
      }

      if (req.minAge !== null && req.minAge !== undefined) {
        details.push(` > Минимальный возраст: ${req.minAge} лет`);
      }

      if (req.maxAge !== null && req.maxAge !== undefined) {
        details.push(` > Максимальный возраст: ${req.maxAge} лет`);
      }

      if (details.length === 0) {
        reqText += `\n- Нет дополнительных требований.`;
      } else {
        reqText += `\n${details.join('\n')}`;
      }

      return reqText.trim();
    }).join('\n');
  };

  const formattedMessage = `🌟*${event.name}*🌟
  
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

  return formattedMessage
}

export { formatMessage }