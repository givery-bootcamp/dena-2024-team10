import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { createRemixStub } from "@remix-run/testing";
import { render, screen, waitFor } from "@testing-library/react";
import PostsDetails from "~/routes/posts.$postId";
import { expect, test } from "vitest";

import Edit from "~/routes/posts.$postId_.edit";

test("投稿詳細ページが表示できる", async () => {
	const RemixStub = createRemixStub([
		{
			path: "/",
			Component: Edit,
			loader() {
				return json({
					id: 1,
					title: "投稿タイトル",
					body: "投稿内容",
					user_id: 1,
					username: "taro",
					created_at: "2021-01-01T00:00:00Z",
					updated_at: "2021-01-01T00:00:00Z",
				});
			},
		},
	]);

	render(<RemixStub />);

	await waitFor(() => screen.findByText("投稿を編集"));
	const input = await waitFor(() => screen.findByLabelText("タイトル"));
	expect((input as HTMLInputElement).value).toBe("投稿タイトル");
	const area = await waitFor(() => screen.findByLabelText("内容"));
	expect((area as HTMLInputElement).value).toBe("投稿内容");
});
