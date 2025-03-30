export function joinClass(...classes: (string | undefined | null)[] | TemplateStringsArray[]): string {
  return classes.filter(Boolean).join(" ");
}
