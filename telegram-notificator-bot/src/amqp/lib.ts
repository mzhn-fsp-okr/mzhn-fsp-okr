export const formatDate = (dateString) => {
  const options: Intl.DateTimeFormatOptions = {
    year: "numeric",
    month: "long",
    day: "numeric",
  };
  return new Date(dateString).toLocaleDateString("ru-RU", options);
};

export const daysUntil = (targetDate) => {
  const currentDate = new Date().getTime();
  const target = new Date(targetDate).getTime();
  const differenceInMilliseconds = target - currentDate;
  const daysRemaining = Math.ceil(
    differenceInMilliseconds / (1000 * 60 * 60 * 24)
  );
  return daysRemaining >= 0 ? daysRemaining : 0;
};

const pluralRules = new Intl.PluralRules("ru-RU");

export const formatDays = (number) => {
  if (number === 0) {
    return `уже *сегодня*`;
  }

  const forms = {
    one: "день",
    few: "дня",
    many: "дней",
    other: "дней",
  };

  const category = pluralRules.select(number);

  return `через ${number} ${forms[category]}`;
};

export const formatParticipantRequirements = (requirements) => {
  return requirements
    .map((req, index) => {
      let reqText = ``;
      const details = [];

      if (req.gender !== undefined) {
        const genderText =
          req.gender === true
            ? "👨‍🦱 Мужчины"
            : req.gender === false
              ? "👱‍♀️ Женщины"
              : "👨‍🦱👱‍♀️ Мужчины и женщины";
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
        reqText += `\n${details.join("\n")}`;
      }

      return reqText.trim();
    })
    .join("\n");
};
