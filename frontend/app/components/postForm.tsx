import { Form, useNavigate } from "@remix-run/react";
import classNames from "classnames";
import type { PostErrorType } from "~/routes/posts.new";
import "@mdxeditor/editor/style.css";
import Button from "./button";
import SubmitButton from "./submitButton";
import { useState } from "react";
import Markdown from "./markdown";

export default function PostForm({
	title,
	content,
	actionData,
	submitText,
}: {
	title?: string;
	content?: string;
	actionData?: PostErrorType;
	submitText: string;
}) {
	const navigate = useNavigate();
	const [contentState, setContentState] = useState(content);
	return (
		<Form method="post" className={classNames("p-2")}>
			<label htmlFor="title" className={classNames("block")}>
				タイトル
			</label>
			<input
				type="text"
				id="title"
				name="title"
				defaultValue={title}
				className={classNames(
					"block",
					"border",
					"border-gray-400",
					"w-full",
					"mb-2",
				)}
			/>
			<label htmlFor="content" className={classNames("block", "w-full")}>
				内容
			</label>
			<input
				type="text"
				id="content"
				name="content"
				value={contentState}
				hidden
			/>
			<Markdown markdown={content ?? ""} onChange={(e) => setContentState(e)} />
			<div className={classNames("flex", "justify-end", "gap-4")}>
				<SubmitButton color="primary" text={submitText} />
				<Button type="none" onClick={() => navigate(-1)}>
					Cancel
				</Button>
			</div>
			{actionData?.errors && (
				<ul>
					{actionData.errors.map((error, i) => (
						// biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
						<li key={i} className={classNames("text-red-500")}>
							{error.path} : {error.message}
						</li>
					))}
				</ul>
			)}
		</Form>
	);
}
