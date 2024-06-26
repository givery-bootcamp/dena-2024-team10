import {
	type LoaderFunctionArgs,
	json,
	type MetaFunction,
	type SerializeFrom,
} from "@remix-run/node";
import {
	Link,
	isRouteErrorResponse,
	useFetcher,
	useLoaderData,
	useRouteError,
} from "@remix-run/react";
import classNames from "classnames";
import formatDate from "utils/formatDate";
import apiClient from "~/apiClient/apiClient";
import Observable from "~/components/observable";
import { useInfinitieLoading } from "~/hooks/infinitieLoading";

export const meta: MetaFunction = () => {
	return [
		{ title: "New Remix App" },
		{ name: "description", content: "Welcome to Remix!" },
	];
};

export const loader = async ({ request, params }: LoaderFunctionArgs) => {
	try {
		const url = new URL(request.url);
		const posts = await apiClient.getPosts({
			queries: {
				offset: Number.parseInt(url.searchParams.get("offset") ?? "0"),
				limit: Number.parseInt(url.searchParams.get("limit") ?? "20"),
			},
			headers: {
				Cookie: request.headers.get("Cookie"),
			},
		});
		return json({ posts });
	} catch (error) {
		// console.error(error);
		if (error instanceof Error) {
			throw new Response(`name: ${error.name}, message: ${error.message}`, {
				status: 500,
			});
		}
		throw new Response("エラーが発生しました", { status: 500 });
	}
};

type Post = SerializeFrom<typeof loader>["posts"][0];

export default function Index() {
	const {
		data: posts,
		loadNext,
		state,
	} = useInfinitieLoading<typeof loader, Post>((data) => data.posts);

	return (
		<main className={classNames("mx-auto", "w-1/2")}>
			<div className={classNames("flex", "py-4")}>
				<h1
					className={classNames("text-3xl", "font-bold", "underline", "flex-1")}
				>
					投稿一覧
				</h1>
				<Link
					to="posts/new"
					className={classNames(
						"p-2",
						"rounded-md",
						"bg-blue-500",
						"text-white",
						"hover:bg-blue-200",
						"text-sm",
					)}
				>
					新しい投稿を作成
				</Link>
			</div>
			<ul>
				{posts.map((post) => (
					<li
						key={post.id}
						className={classNames("border", "flex", "h-16", "px-4", "py-2")}
					>
						<Link
							to={`/posts/${post.id}`}
							className={classNames(
								"text-blue-500",
								"font-bold",
								"underline",
								"flex-1",
							)}
						>
							{post.title}
						</Link>
						<p className={classNames("text-sm", "mx-1", "self-end")}>
							{post.username}
						</p>
						<p className={classNames("text-sm", "mx-1", "self-end")}>
							更新日時: {formatDate(post.updated_at)}
						</p>
					</li>
				))}
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

export function ErrorBoundary() {
	const error = useRouteError();
	console.log(error);
	console.error(isRouteErrorResponse(error));
	return (
		<main>
			<h1
				className={classNames(
					"text-3xl",
					"font-bold",
					"underline",
					"text-red-500",
					"text-center",
					"mt-8",
				)}
			>
				一覧取得できなかったよーー <br />
				{isRouteErrorResponse(error) ? error.data : "エラー"} <br />
				<Link to="/">トップに戻る</Link>
			</h1>
		</main>
	);
}
