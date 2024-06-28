import type { ActionFunctionArgs, LoaderFunctionArgs } from "@remix-run/node";
import {
	Form,
	Link,
	redirect,
	useLoaderData,
	useParams,
} from "@remix-run/react";
import classNames from "classnames";
import formatDate from "utils/formatDate";
import apiClient, { API_BASE_URL } from "~/apiClient/apiClient";
import { useDialog } from "~/components/dialog";

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
		const user = await apiClient.getSignedInUser({
			headers: {
				Cookie: request.headers.get("Cookie"),
			},
		});

		return { ...post, isMyPost: user.id === post.user_id };
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

export default function PostsDetails() {
	const data = useLoaderData<typeof loader>();
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
				<h1 className={classNames("text-3xl", "my-3")}>{data.title}</h1>
			</div>
			<div className={classNames("flex", "gap-3")}>
				{TimeTopic("作成日時", data.created_at)}
				{TimeTopic("更新日時", data.updated_at)}
			</div>
			<hr className={classNames("my-3")} />
			<pre>{data.body}</pre>
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
				{data.username}
			</p>
			{data.isMyPost && (
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
			)}
		</main>
	);
}
