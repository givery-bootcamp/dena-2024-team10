import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { createRemixStub } from "@remix-run/testing";
import { render, screen, waitFor } from "@testing-library/react";
import PostsDetails from "~/routes/posts.$postId";
import { expect, test } from "vitest";

test("投稿詳細ページが表示できる", async () => {
	const RemixStub = createRemixStub([
		{
			path: "/",
			Component: PostsDetails,
			loader() {
				return json({
					id: 1,
					title: "投稿タイトル",
					body: "投稿内容",
					user_id: 1,
					username: "taro",
					created_at: "2021-01-01T00:00:00Z",
					updated_at: "2021-01-01T00:00:00Z",
					isMyPost: true,
				});
			},
		},
	]);

	render(<RemixStub />);

	await waitFor(() => screen.findByText("投稿タイトル"));
	await waitFor(() => screen.findByText("投稿内容"));
	await waitFor(() => screen.findByText("taro"));
});
test("自分の投稿なら削除ボタンがある", async () => {
	const RemixStub = createRemixStub([
		{
			path: "/",
			Component: PostsDetails,
			loader() {
				return json({
					id: 1,
					title: "投稿タイトル",
					body: "投稿内容",
					user_id: 1,
					username: "taro",
					created_at: "2021-01-01T00:00:00Z",
					updated_at: "2021-01-01T00:00:00Z",
					isMyPost: true,
				});
			},
		},
	]);

	render(<RemixStub />);

	await waitFor(() => screen.findByText("削除"));
});

test("自分の投稿ではないなら削除ボタンがない", async () => {
	const RemixStub = createRemixStub([
		{
			path: "/",
			Component: PostsDetails,
			loader() {
				return json({
					id: 1,
					title: "投稿タイトル",
					body: "投稿内容",
					user_id: 1,
					username: "taro",
					created_at: "2021-01-01T00:00:00Z",
					updated_at: "2021-01-01T00:00:00Z",
					isMyPost: false,
				});
			},
		},
	]);

	render(<RemixStub />);

	expect(screen.queryByText("削除")).toBeNull();
});
