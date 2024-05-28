import { json, type MetaFunction } from "@remix-run/node";
import {
	Link,
	isRouteErrorResponse,
	useLoaderData,
	useRouteError,
} from "@remix-run/react";
import classNames from "classnames";
import dayjs from "dayjs";
import { createApiClient } from "~/apiClient";

export const meta: MetaFunction = () => {
	return [
		{ title: "New Remix App" },
		{ name: "description", content: "Welcome to Remix!" },
	];
};

const api = createApiClient("http://localhost:4010/");
export const loader = async () => {
	try {
		const posts = await api.getPosts();
		return json({ posts });
	} catch (error) {
		console.error(error);
		if (error instanceof Error) {
			throw new Response(`name: ${error.name}, message: ${error.message}`, {
				status: 500,
			});
		}
		throw new Response("エラーが発生しました", { status: 500 });
	}
};

const formatDate = (date: string) => {
	return dayjs(date).format("YYYY/MM/DD HH:mm");
};

export default function Index() {
	const data = useLoaderData<typeof loader>();
	console.log(data.posts[0]);

	return (
		<main className={classNames("mx-auto", "w-1/2")}>
			<h1 className="text-3xl font-bold underline">投稿一覧</h1>
			<ul>
				{data.posts.map((post) => (
					<li
						key={post.id}
						className={classNames("border", "flex", "h-16", "px-4", "py-2")}
					>
						<Link
							to={"/id"}
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
