import { json, type ActionFunctionArgs } from "@remix-run/node";
import { ZodError } from "zod";
import apiClient from "~/apiClient/apiClient";
import { z } from "zod";

export async function action({ request, params }: ActionFunctionArgs) {
	const postId = Number.parseInt(params.postId as string);
	const commentId = Number.parseInt(params.commentId as string);
	try {
		if (request.method === "PUT") {
			const formData = await request.formData();
			const comment = formData.get("comment") as string;

			const body = { body: comment };
			updateComment_Body.parse(body);
			return await apiClient.updateComment(body, {
				headers: { cookie: request.headers.get("cookie") },
				params: { postId, commentId },
			});
		}
		if (request.method === "DELETE") {
			return await apiClient.deleteComment(undefined, {
				headers: { cookie: request.headers.get("cookie") },
				params: { postId, commentId },
			});
		}
	} catch (e) {
		console.error(e);
		if (e instanceof ZodError) {
			return json({
				errors: e.errors.map((error) => {
					return { path: error.path, message: error.message };
				}),
			});
		}
		if (e instanceof Error) {
			throw new Response("Oh no! Something went wrong!", {
				status: 500,
			});
		}
	}
}

export function loader() {
	return new Response("Not found", {
		status: 404,
	});
}

const updateComment_Body = z.object({ body: z.string() }).passthrough();
