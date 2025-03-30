export function joinClass(...classes: string[] | TemplateStringsArray[]): string {
  return classes.filter(Boolean).join(" ");
}
