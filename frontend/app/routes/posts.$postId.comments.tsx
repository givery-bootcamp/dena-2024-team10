import type { ActionFunctionArgs } from "@remix-run/node";
import { json, redirect } from "@remix-run/react";
import { ZodError } from "zod";
import apiClient from "~/apiClient/apiClient";
import { z } from "zod";

export async function action({ request, params }: ActionFunctionArgs) {
	const postId = Number.parseInt(params.postId as string);
	try {
		const formData = await request.formData();
		const comment = formData.get("comment") as string;

		const body = { body: comment };
		createComment_Body.parse(body);
		await apiClient.createComment(body, {
			headers: { cookie: request.headers.get("cookie") },
			params: { postId },
		});
		return redirect(`/posts/${postId}`);
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

const createComment_Body = z.object({ body: z.string() }).passthrough();
