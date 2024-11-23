type Locale = Record<string, string>;

const locale = {
  example_message: "Hello",
} satisfies Locale;
type AvailableLocaleKeys = keyof typeof locale;

/** Get localized string by translation key */
export const _ = (
  key: AvailableLocaleKeys,
  replacements?: Record<string, any>
) => {
  if (key in locale) {
    let val = locale[key];
    if (!replacements) return val;
    for (const [key, value] of Object.entries(replacements)) {
      val = val.replace(`%${key}%`, value);
    }
    return val;
  }
  throw new Error(`Unknown locale key: ${key}`);
};
