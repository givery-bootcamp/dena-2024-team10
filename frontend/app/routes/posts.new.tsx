import {
	type ActionFunctionArgs,
	redirect,
	type SerializeFrom,
} from "@remix-run/node";
import { Form, json, useActionData } from "@remix-run/react";
import classNames from "classnames";
import { ZodError } from "zod";
import apiClient from "~/apiClient/apiClient";
import { schemas } from "~/apiClient/output.generated";
import PostForm from "~/components/postForm";
import SubmitButton from "~/components/submitButton";

export async function action({ request }: ActionFunctionArgs) {
	try {
		const formData = await request.formData();
		const title = formData.get("title") as string;
		const content = formData.get("content") as string;

		const body = { title, body: content };
		schemas.createPost_Body.parse(body);
		const res = await apiClient.createPost(
			{ title, body: content },
			{ headers: { cookie: request.headers.get("cookie") } },
		);

		return redirect(`/posts/${res.id}`);
		// return redirect('/posts');
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

export type PostErrorType = SerializeFrom<typeof action>;

export default function () {
	const actionData = useActionData<typeof action>();

	return (
		<main className={classNames("w-1/2", "mx-auto")}>
			<h1 className={classNames("text-4xl", "my-4")}>新しい投稿を作成する</h1>
			<PostForm actionData={actionData} />
		</main>
	);
}
