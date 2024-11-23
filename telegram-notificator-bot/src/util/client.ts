import axios from "axios";
import { Agent } from "https";

const env = process.env;
const httpsAgent = new Agent();
const client = axios.create({
  baseURL: env.API_URL,
  headers: { "Content-Type": "application/json", scheme: "https" },
  httpsAgent,
});

export default client;
