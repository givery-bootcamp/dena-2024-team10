import { json, type MetaFunction } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";
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
	const posts = await api.getPosts();
	return json({ posts });

	// return json({
	// 	posts: [
	// 		{
	// 			title: "投稿1",
	// 			id: 123,
	// 			username: "user1",
	// 			updated_at: "2022-03-01T13:00:00.000Z",
	// 		},
	// 		{
	// 			title: "投稿2",
	// 			id: 456,
	// 			username: "user2",
	// 			updated_at: "2022-03-01T13:00:00.000Z",
	// 		},
	// 	],
	// });
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

// export function loader() {
//   // return json({count: counter})
// }
