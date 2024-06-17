import { Link } from "@remix-run/react";
import classNames from "classnames";
import apiClient from "~/apiClient/apiClient";

export default function Header({
	isSignedIn,
	username,
}: { isSignedIn: boolean; username: string | undefined }) {
	const signout = async () => {
		try {
			await apiClient.signoutUser(undefined);
		} catch (e) {
			console.error(e);
		}
	};

	return (
		<header
			className={classNames("bg-blue-800", "p-2", "flex", "items-center")}
		>
			<h1 className={classNames("text-lg", "text-white", "flex-1")}>
				サンプルアプリケーション
			</h1>
			{isSignedIn ? (
				<div className={classNames("flex", "items-center")}>
					<p className={classNames("text-white", "mx-4")}>{username}</p>
					<Link
						to="/signin"
						className={classNames("p-2", "rounded-md", "bg-white")}
						onClick={signout}
					>
						サインアウト
					</Link>
				</div>
			) : (
				<div>
					<Link
						to="/signin"
						className={classNames("p-2", "rounded-md", "bg-white")}
					>
						サインイン
					</Link>
				</div>
			)}
		</header>
	);
}
