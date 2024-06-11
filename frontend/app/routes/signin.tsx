import type { ActionFunctionArgs } from "@remix-run/node";
import { Form, Link } from "@remix-run/react";
import classNames from "classnames";

export async function action({ request }: ActionFunctionArgs) {
	const formData = await request.formData();
	const name = formData.get("name");
	const pass = formData.get("pass");
	console.log(name, pass);
	return new Response("サインインしました", {
		status: 200,
	});
}

export default function SignIn() {
	return (
		<main className={classNames("mx-auto", "w-1/3")}>
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
			<Link
				to="/signup"
				className={classNames(
					"text-sm",
					"text-blue-800",
					"underline",
					"text-center",
					"block",
					"mt-3",
				)}
			>
				アカウントをお持ちでない方はこちら
			</Link>
		</main>
	);
}
