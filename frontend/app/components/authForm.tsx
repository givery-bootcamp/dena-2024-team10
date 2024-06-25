import { Form } from "@remix-run/react";
import classNames from "classnames";

const AuthForm = ({ authType }: { authType: string }) => {
	return (
		<Form method="post">
			<div className={classNames("my-5")}>
				<label htmlFor="name" className={classNames("my-2", "block")}>
					ユーザー名
				</label>
				<input
					id="name"
					type="text"
					name="name"
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
					name="pass"
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
				value={authType === "signin" ? "サインイン" : "サインアップ"}
				className={classNames(
					"text-white",
					authType === "signin" ? "bg-blue-600" : "bg-green-600",
					"p-2",
					"rounded-md",
					authType === "signin" ? "hover:bg-indigo-200" : "hover:bg-green-200",
				)}
			/>
		</Form>
	);
};

export default AuthForm;
