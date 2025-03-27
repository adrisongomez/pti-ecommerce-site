export function filterNull<T>(
  v: T | null | undefined,
): v is Exclude<T, null | undefined> {
  return Boolean(v);
}
