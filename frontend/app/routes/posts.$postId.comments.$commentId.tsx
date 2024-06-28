import { type ActionFunctionArgs, json } from "@remix-run/node";
import { ZodError } from "zod";
import { z } from "zod";
import apiClient, { API_BASE_URL } from "~/apiClient/apiClient";

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
			const res = await fetch(
				`${API_BASE_URL}/posts/${postId}/comments/${commentId}`,
				{
					method: "DELETE",
					headers: {
						Cookie: request.headers.get("Cookie") as string,
					},
				},
			);
			if (!res.ok) throw new Error("Failed to delete comment");
			return new Response("Comment deleted", { status: 200 });
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
