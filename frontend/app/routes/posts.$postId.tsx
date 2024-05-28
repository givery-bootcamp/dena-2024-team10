import { json, type LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import classNames from "classnames";
import formatDate from "utils/formatDate";
import apiClient from "~/apiClient/apiClient";

export async function loader({ params }: LoaderFunctionArgs) {
	try {
		// const detailes = await apiClient.getPostDetails(id: params.postId);
		// return json({ detailes });
		return json({
			id: params.postId,
			title: "title",
			body: "body\n\nboooddddyyyyyy",
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
	const data = useLoaderData<typeof loader>();
	const TimeTopic = (topic: string, time: string) => (
		<div className={classNames("flex", "text-sm")}>
			<p className={classNames("opacity-20", "mr-2")}>{topic}</p>
			<p>{formatDate(time)}</p>
		</div>
	);
	return (
		<main className={classNames("mx-auto", "w-1/2")}>
			<h1 className={classNames("text-3xl", "my-3")}>{data.title}</h1>
			<div className={classNames("flex", "gap-3")}>
				{TimeTopic("作成日時", data.created_at)}
				{TimeTopic("更新日時", data.updated_at)}
			</div>
			<hr className={classNames("my-3")} />
			{data.body.split("\n").map((line, index) => (
				// biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
				<p key={index} className={classNames("my-1")}>
					{/* 空行対応できてない */}
					{line}
				</p>
			))}
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
		</main>
	);
}
