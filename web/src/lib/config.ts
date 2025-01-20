export const PROD = import.meta.env.MODE === "production";
export const SERVER_URL = PROD ? "" : "http://localhost:8080";
