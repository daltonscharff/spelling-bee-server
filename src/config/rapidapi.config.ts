import { registerAs } from "@nestjs/config";

export default registerAs('rapidapi', () => ({
    host: process.env.RAPID_API_HOST,
    key: process.env.RAPID_API_KEY
}));