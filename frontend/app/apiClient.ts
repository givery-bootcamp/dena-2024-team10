import { makeApi, Zodios, type ZodiosOptions } from "@zodios/core";
import { z } from "zod";

const Post = z
	.object({
		id: z.number().int(),
		title: z.string(),
		body: z.string(),
		user_id: z.number().int(),
		username: z.string(),
		created_at: z.string().datetime({ offset: true }),
		updated_at: z.string().datetime({ offset: true }),
	})
	.passthrough();

export const schemas = {
	Post,
};

const endpoints = makeApi([
	{
		method: "get",
		path: "/posts",
		alias: "getPosts",
		description: `Retrieve a list of posts sorted by id in descending order.`,
		requestFormat: "json",
		response: z.array(Post),
	},
]);

export const api = new Zodios(endpoints);

export function createApiClient(baseUrl: string, options?: ZodiosOptions) {
	return new Zodios(baseUrl, endpoints, options);
}
