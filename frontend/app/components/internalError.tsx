import { Link } from "@remix-run/react";
import classNames from "classnames";

export default function InternalError({
	title,
	to,
	toMessage,
}: { title: string; to: string; toMessage: string }) {
	return (
		<main className={classNames("w-1/2", "mx-auto")}>
			<h1 className={classNames("text-4xl", "my-4")}>{title}</h1>
			<p>しばらく待ってからやり直してください</p>
			<div className={classNames("text-end", "mr-5")}>
				<Link to={to} className={classNames("text-blue-600", "underline")}>
					{toMessage}
				</Link>
			</div>
		</main>
	);
}
