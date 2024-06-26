import {
	type ActionFunctionArgs,
	type LoaderFunctionArgs,
	type SerializeFrom,
	redirect,
} from "@remix-run/node";
import { Form, json, useActionData, useLoaderData } from "@remix-run/react";
import classNames from "classnames";
import { ZodError } from "zod";
import apiClient from "~/apiClient/apiClient";
import { schemas } from "~/apiClient/output.generated";
import PostForm from "~/components/postForm";
import SubmitButton from "~/components/submitButton";

export async function loader({ request, params }: LoaderFunctionArgs) {
	const postId = Number.parseInt(params.postId as string);
	return json({
		id: postId,
		title: "title",
		body: "body\n\nboooddddyyyyyy",
		user_id: 1,
		username: "username",
		created_at: "2022-01-01T00:00:00.000Z",
		updated_at: "2022-01-01T00:00:00.000Z",
	});
}

export async function action({ request, params }: ActionFunctionArgs) {
	const postId = Number.parseInt(params.postId as string);
	try {
		const formData = await request.formData();
		const title = formData.get("title") as string;
		const content = formData.get("content") as string;

		const body = { title, body: content };
		schemas.createPost_Body.parse(body);
		const res = await apiClient.updatePost(
			{ title, body: content },
			{
				headers: { cookie: request.headers.get("cookie") },
				params: { postId },
			},
		);

		return redirect(`/posts/${res.id}`);
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
	const data = useLoaderData<typeof loader>();

	return (
		<main className={classNames("w-1/2", "mx-auto")}>
			<h1 className={classNames("text-4xl", "my-4")}>投稿を編集</h1>
			<PostForm
				title={data.title}
				content={data.body}
				actionData={actionData}
				submitText="更新"
			/>
		</main>
	);
}
