import { createApiClient } from "./output.generated";

const apiClient = createApiClient(
	process.env.API_BASE_URL ?? "http://localhost:4010",
);

export default apiClient;
