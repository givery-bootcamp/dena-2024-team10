import { Form } from "@remix-run/react";
import classNames from "classnames";

export default function SignIn() {
	return (
		<main className={classNames("mx-auto", "w-1/3")}>
			<Form>
				<div className={classNames("my-5")}>
					<label htmlFor="name" className={classNames("my-2", "block")}>
						ユーザー名
					</label>
					<input
						id="name"
						type="text"
						className={classNames(
							"border-2",
							"rounded-md",
							"border-gray-300",
							"w-full",
						)}
					/>
				</div>
				<div className={classNames("my-5")}>
					<label htmlFor="pass" className={classNames("my-2", "block")}>
						パスワード
					</label>
					<input
						id="pass"
						type="password"
						className={classNames(
							"border-2",
							"rounded-md",
							"border-gray-300",
							"w-full",
						)}
					/>
				</div>
				<input
					type="submit"
					value="サインイン"
					className={classNames(
						"text-white",
						"bg-blue-600",
						"p-2",
						"rounded-md",
						"hover:bg-indigo-200",
					)}
				/>
			</Form>
		</main>
	);
}
