import { Creds } from "../generated";

const CREDS_KEY = "admin-creds";

export function getCreds(): Creds | null {
  const data = localStorage.getItem(CREDS_KEY);
  if (!data) {
    return null;
  }
  const parseData = JSON.parse(data);
  if (isCreds(parseData)) {
    return parseData;
  }
  return null;
}

export function isCreds(data: unknown): data is Creds {
  const keys = Object.keys(data as Record<string, unknown>);
  return keys.includes("accessToken") && keys.includes("refreshToken");
}

export function writeCreds(data: Creds) {
  localStorage.setItem(CREDS_KEY, JSON.stringify(data));
}

export function removeCreds() {
  localStorage.removeItem(CREDS_KEY);
}
