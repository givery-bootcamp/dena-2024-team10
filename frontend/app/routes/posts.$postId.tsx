import { json, type LoaderFunctionArgs } from "@remix-run/node";
import classNames from "classnames";
import apiClient from "~/apiClient/apiClient";

export async function loader({ params }: LoaderFunctionArgs) {
	try {
		// const detailes = await apiClient.getPostDetails(id: params.postId);
		// return json({ detailes });
		return json({
			id: params.postId,
			title: "title",
			body: "body",
			user_id: 1,
			username: "username",
			created_at: "2022-01-01T00:00:00.000Z",
			updated_at: "2022-01-01T00:00:00.000Z",
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

export default function PostsDetails() {
	return (
		<main className={classNames("mx-auto", "w-1/2")}>
			<h1 className="text-3xl font-bold underline">投稿詳細</h1>
			<p>投稿詳細ページです</p>
		</main>
	);
}
