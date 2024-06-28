import {
	type ActionFunctionArgs,
	json,
	type LoaderFunctionArgs,
	type SerializeFrom,
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
import { useState } from "react";
import formatDate from "utils/formatDate";
import apiClient from "~/apiClient/apiClient";
import Button from "~/components/button";
import { useDialog } from "~/components/dialog";
import Observable from "~/components/observable";
import SubmitButton from "~/components/submitButton";
import { useInfinitieLoading } from "~/hooks/infinitieLoading";

export async function loader({ params, request }: LoaderFunctionArgs) {
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
				limit: Number.parseInt(url.searchParams.get("limit") ?? "20"),
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
			post: post,
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

type Comment = SerializeFrom<typeof loader>["comments"][0];

export default function PostsDetails() {
	const fetcher = useFetcher();
	const { post, user } = useLoaderData<typeof loader>();
	const {
		data: comments,
		loadNext,
		state,
	} = useInfinitieLoading<typeof loader, Comment>((data) => data.comments, 3);

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
								isLastChild={isLastChild}
								isMyComment={isMyComment}
							/>
						</li>
					);
				})}
			</ul>
			{state === "loading" && (
				<div className={classNames("text-center", "m-4")}>読み込み中</div>
			)}
			{state === "end" && (
				<div className={classNames("text-center", "m-4")}>
					これ以上投稿がありません
				</div>
			)}
			<Observable
				callback={() => {
					loadNext();
				}}
			/>
		</main>
	);
}

function Comment({
	comment,
	isLastChild,
	isMyComment,
}: {
	comment: Comment;
	isLastChild: boolean;
	isMyComment: boolean;
}) {
	const fetcher = useFetcher();
	const params = useParams();
	const [isEditingComment, setIsEditingComment] = useState(false);
	const { dialog, confirm } = useDialog(
		<h1 className={classNames("text-lg", "font-bold")}>削除しますか？</h1>,
	);
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
					{dialog}
				</>
			)}
		</>
	);
}
