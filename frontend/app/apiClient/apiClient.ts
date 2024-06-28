import { createApiClient } from "./output.generated";

// export const API_BASE_URL = process.env.API_BASE_URL ?? "http://localhost:4010";
export const API_BASE_URL = process.env.API_BASE_URL ?? "http://localhost:9000";

const apiClient = createApiClient(API_BASE_URL, {});

export default apiClient;
