import classNames from "classnames";
import SubmitButton from "~/components/submitButton";

export default function () {
	return (
		<main className={classNames("w-1/2", "mx-auto")}>
			<h1 className={classNames("text-4xl", "my-4")}>新しい投稿を作成する</h1>
			<form method="post" className={classNames("p-2")}>
				<label htmlFor="title" className={classNames("block")}>
					タイトル
				</label>
				<input
					type="text"
					id="title"
					name="title"
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
				<textarea
					name="content"
					id="content"
					className={classNames(
						"block",
						"border",
						"border-gray-400",
						"w-full",
						"mb-4",
					)}
				/>
				<SubmitButton color="primary" text="投稿" />
			</form>
		</main>
	);
}
