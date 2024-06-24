import { Form } from "@remix-run/react";
import classNames from "classnames";
import SubmitButton from "./submitButton";
import type { PostErrorType } from "~/routes/posts.new";
import { ClientOnly } from "remix-utils/client-only";
import {
	headingsPlugin,
	listsPlugin,
	quotePlugin,
	thematicBreakPlugin,
	markdownShortcutPlugin,
	MDXEditor,
	type MDXEditorMethods,
	type MDXEditorProps,
} from "@mdxeditor/editor";

export default function PostForm({
	title,
	content,
	actionData,
}: { title?: string; content?: string; actionData?: PostErrorType }) {
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
			<ClientOnly fallback={<p>Loading...</p>}>
				{() => (
					<MDXEditor
						plugins={[
							headingsPlugin(),
							listsPlugin(),
							quotePlugin(),
							thematicBreakPlugin(),
							markdownShortcutPlugin(),
						]}
						markdown="Hello world"
					/>
				)}
			</ClientOnly>
			<SubmitButton color="primary" text="投稿" />
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
