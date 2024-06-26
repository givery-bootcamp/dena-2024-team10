import {
	type ActionFunctionArgs,
	json,
	type LoaderFunctionArgs,
} from "@remix-run/node";
import {
	Form,
	Link,
	redirect,
	useLoaderData,
	useParams,
} from "@remix-run/react";
import classNames from "classnames";
import { useState } from "react";
import formatDate from "utils/formatDate";
import apiClient from "~/apiClient/apiClient";
import Dialog, { useDialog } from "~/components/dialog";
import SubmitButton from "~/components/submitButton";

export async function loader({ params, request }: LoaderFunctionArgs) {
	try {
		const postId = Number.parseInt(params.postId as string);
		const post = await apiClient.getPost({
			params: {
				postId,
			},
			headers: {
				Cookie: request.headers.get("Cookie"),
			},
		});
		const comments = await apiClient.getComments({
			params: {
				postId,
			},
			headers: {
				Cookie: request.headers.get("Cookie"),
			},
		});

		return json({
			post: post,
			comments: comments,
		});
	} catch (error) {
		console.error(error);
		if (error instanceof Error) {
			throw new Response(`name: ${error.name}, message: ${error.message}`, {
				status: 500,
			});
		}
		throw new Response("エラーが発生しました", { status: 500 });
	}
}

export async function action({ params, request }: ActionFunctionArgs) {
	const postId = Number.parseInt(params.postId as string);
	try {
		await apiClient.deletePost(undefined, {
			params: {
				postId,
			},
			headers: {
				Cookie: request.headers.get("Cookie") as string,
			},
		});
		return redirect("/");
	} catch (e) {
		return new Response((e as Error).message, {
			status: 400,
		});
	}
}

export default function PostsDetails() {
	const { post, comments } = useLoaderData<typeof loader>();
	const params = useParams();
	const { dialog, confirm } = useDialog(
		<h1 className={classNames("text-lg", "font-bold")}>削除しますか？</h1>,
	);

	const TimeTopic = (topic: string, time: string) => (
		<div className={classNames("flex", "text-sm")}>
			<p className={classNames("opacity-20", "mr-2")}>{topic}</p>
			<p>{formatDate(time)}</p>
		</div>
	);
	return (
		<main className={classNames("mx-auto", "w-1/2")}>
			<h1 className={classNames("text-3xl", "my-3")}>{post.title}</h1>
			<div className={classNames("flex", "gap-3")}>
				{TimeTopic("作成日時", post.created_at)}
				{TimeTopic("更新日時", post.updated_at)}
			</div>
			<hr className={classNames("my-3")} />
			<pre>{post.body}</pre>
			<p
				className={classNames(
					"bg-blue-200",
					"p-2",
					"ml-auto",
					"w-fit",
					"text-xs",
					"rounded-md",
				)}
			>
				{post.username}
			</p>
			<div className={classNames("flex", "gap-4")}>
				<Form
					method="delete"
					onSubmit={async (e) => {
						e.preventDefault();
						if (await confirm()) (e.target as HTMLFormElement).submit();
					}}
				>
					<input
						type="submit"
						value="削除"
						className={classNames(
							"text-blue-500",
							"underline",
							"cursor-pointer",
						)}
					/>
				</Form>
				<Link
					to={`/posts/${params.postId}/edit`}
					className={classNames("text-blue-500", "underline")}
				>
					編集
				</Link>
				{dialog}
			</div>
			<hr className={classNames("my-4")} />
			<Form
				method="post"
				action={`/posts/${params.postId}/comments`}
				className={classNames("flex", "my-2")}
			>
				<input
					type="text"
					id="comment"
					name="comment"
					placeholder="コメントを入力..."
					className={classNames("w-full", "border-b", "")}
				/>
				<SubmitButton color="primary" text="投稿" />
			</Form>
			<ul>
				{comments.map((comment, index) => (
					<li key={comment.id} className={classNames("p-2")}>
						<p>{comment.body}</p>
						<div className={classNames("ml-auto", "w-fit", "flex", "text-sm")}>
							<p>{comment.username}</p>
							<p className={classNames("ml-1", "opacity-20")}>
								{formatDate(comment.created_at)}
							</p>
						</div>
						{index !== comments.length - 1 && (
							<hr className={classNames("mt-4")} />
						)}
					</li>
				))}
			</ul>
		</main>
	);
}
