export function replaceEmpty<
  T extends {
    [k: string]: any;
  },
>(
  obj: T
): {
  [k: string]: any;
} {
  return Object.fromEntries(
    Object.entries(obj).filter(
      ([key, value]) => value !== "" && value !== undefined
    )
  );
}

export function replaceArrayEmpty<
  T extends {
    [k: string]: any;
  },
>(
  obj: T
): {
  [k: string]: any;
} {
  return Object.fromEntries(
    Object.entries(obj).filter(
      ([key, value]) => !Array.isArray(value) || value.length > 0
    )
  );
}
