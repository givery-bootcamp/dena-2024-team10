import {
	type ActionFunctionArgs,
	type LoaderFunctionArgs,
	type SerializeFrom,
	json,
} from "@remix-run/node";
import {
	Form,
	Link,
	redirect,
	useFetcher,
	useLoaderData,
	useParams,
} from "@remix-run/react";
import classNames from "classnames";
import formatDate from "utils/formatDate";
import "@mdxeditor/editor/style.css";
import Markdown from "~/components/markdown";
import apiClient, { API_BASE_URL } from "~/apiClient/apiClient";
import Button from "~/components/button";
import { useDialog } from "~/components/dialog";
import SubmitButton from "~/components/submitButton";
import { useState } from "react";

export async function loader({ params, request }: LoaderFunctionArgs) {
	const LIMIT = 100;
	try {
		const url = new URL(request.url);
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
			queries: {
				offset: Number.parseInt(url.searchParams.get("offset") ?? "0"),
				limit: LIMIT,
			},
			params: {
				postId,
			},
			headers: {
				Cookie: request.headers.get("Cookie"),
			},
		});
		const user = await apiClient.getSignedInUser({
			headers: {
				Cookie: request.headers.get("Cookie"),
			},
		});

		return json({
			post: { ...post, isMyPost: user.id === post.user_id },
			comments: comments,
			user: user,
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
		await fetch(`${API_BASE_URL}/posts/${postId}`, {
			method: "DELETE",
			headers: {
				Cookie: request.headers.get("Cookie") as string,
			},
		});
		return redirect("/");
	} catch (e) {
		console.error(e);
		throw new Response(JSON.stringify(e), { status: 500 });
	}
}

type Comment = SerializeFrom<typeof loader>["comments"][0];

export default function PostsDetails() {
	const fetcher = useFetcher();
	const { post, comments, user } = useLoaderData<typeof loader>();

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
	const [comment, setComment] = useState("");

	return (
		<main className={classNames("mx-auto", "w-1/2")}>
			<div className={classNames("flex")}>
				<Link
					to="/"
					className={classNames(
						"-ml-24",
						"mr-8",
						"mt-4",
						"text-gray-300",
						"hover:text-gray-500",
					)}
				>
					＜一覧へ
				</Link>
				<h1 className={classNames("text-3xl", "my-3")}>{post.title}</h1>
			</div>
			<div className={classNames("flex", "gap-3")}>
				{TimeTopic("作成日時", post.created_at)}
				{TimeTopic("更新日時", post.updated_at)}
			</div>
			<hr className={classNames("my-3")} />
			<Markdown markdown={post.body} />
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
			{post.isMyPost && (
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
				</div>
			)}
			<hr className={classNames("mb-12", "mt-4")} />
			<fetcher.Form
				method="post"
				action={`/posts/${params.postId}/comments`}
				onSubmit={() => setComment("")}
				className={classNames("flex", "my-4")}
			>
				<input
					type="text"
					id="comment"
					name="comment"
					placeholder="コメントを入力..."
					value={comment}
					onChange={(e) => setComment(e.target.value)}
					className={classNames(
						"w-full",
						"border-b",
						"outline-none",
						"focus:border-b-blue-500",
						"border-b-2",
						"mr-4",
					)}
				/>
				<SubmitButton
					color="primary"
					text="投稿"
					isDisabled={comment.length === 0}
				/>
			</fetcher.Form>
			<ul>
				{comments.map((comment, index) => {
					const isLastChild = index === comments.length - 1;
					const isMyComment = comment.user_id === user.id;
					return (
						<li key={comment.id} className={classNames("p-2")}>
							<Comment
								comment={comment}
								confirm={confirm}
								isLastChild={isLastChild}
								isMyComment={isMyComment}
							/>
						</li>
					);
				})}
			</ul>
			{dialog}
		</main>
	);
}

function Comment({
	comment,
	confirm,
	isLastChild,
	isMyComment,
}: {
	comment: Comment;
	confirm: () => Promise<boolean>;
	isLastChild: boolean;
	isMyComment: boolean;
}) {
	const fetcher = useFetcher();
	const params = useParams();
	const [isEditingComment, setIsEditingComment] = useState(false);
	return (
		<>
			{isEditingComment ? (
				<fetcher.Form
					method="put"
					action={`/posts/${params.postId}/comments/${comment.id}`}
					onSubmit={() => setIsEditingComment(false)}
				>
					<textarea
						name="comment"
						id="comment"
						defaultValue={comment.body}
						className={classNames(
							"border",
							"border-gray-400",
							"w-full",
							"mb-2",
						)}
					/>
					<div className={classNames("flex", "justify-end", "gap-4")}>
						<SubmitButton color="primary" text="更新" />
						<Button type="none" onClick={() => setIsEditingComment(false)}>
							Cancel
						</Button>
					</div>
				</fetcher.Form>
			) : (
				<>
					<p>{comment.body}</p>
					<div className={classNames("flex", "my-2")}>
						{isMyComment && (
							<div className={classNames("gap-4", "flex")}>
								<button
									type="button"
									onClick={() => setIsEditingComment(true)}
									className={classNames(
										"text-blue-500",
										"underline",
										"cursor-pointer",
									)}
								>
									編集
								</button>
								<button
									onClick={async () => {
										if (await confirm())
											fetcher.submit(
												{},
												{
													method: "delete",
													action: `/posts/${params.postId}/comments/${comment.id}`,
												},
											);
									}}
									type="button"
									className={classNames(
										"text-blue-500",
										"underline",
										"cursor-pointer",
									)}
								>
									削除
								</button>
							</div>
						)}
						<div className={classNames("ml-auto", "w-fit", "flex", "text-sm")}>
							<p>{comment.username}</p>
							<p className={classNames("ml-1", "opacity-20")}>
								{formatDate(comment.created_at)}
							</p>
						</div>
					</div>
					{!isLastChild && <hr className={classNames("mt-4")} />}
				</>
			)}
		</>
	);
}
