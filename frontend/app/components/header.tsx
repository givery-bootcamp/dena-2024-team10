import { Link } from "@remix-run/react";
import classNames from "classnames";
import apiClient from "~/apiClient/apiClient";

export default function Header({
	isSignedIn,
	username,
}: { isSignedIn: boolean; username: string | undefined }) {
	return (
		<header
			className={classNames("bg-blue-800", "p-2", "flex", "items-center")}
		>
			<h1 className={classNames("text-lg", "text-white", "flex-1")}>
				サンプルアプリケーション
			</h1>
			{isSignedIn ? (
				<>
					<p className={classNames("text-white", "mx-4")}>{username}</p>
					<form method="post" action="/signout">
						<input
							type="submit"
							className={classNames("p-2", "rounded-md", "bg-white")}
							value="サインアウト"
						/>
					</form>
				</>
			) : (
				<>
					<Link
						to="/signin"
						className={classNames("p-2", "rounded-md", "bg-white")}
					>
						サインイン
					</Link>
				</>
			)}
		</header>
	);
}
