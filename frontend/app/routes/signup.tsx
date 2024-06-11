import type { ActionFunctionArgs } from "@remix-run/node";
import { Link } from "@remix-run/react";
import classNames from "classnames";
import AuthForm from "~/components/authForm";

export async function action({ request }: ActionFunctionArgs) {
	const formData = await request.formData();
	const name = formData.get("name");
	const pass = formData.get("pass");
	console.log(name, pass);
	return new Response("サインアップしました", {
		status: 200,
	});
}

export default function SignIn() {
	return (
		<main className={classNames("mx-auto", "w-1/3")}>
			<AuthForm authType="signup" />
			<Link
				to="/signin"
				className={classNames(
					"text-sm",
					"text-blue-800",
					"underline",
					"text-center",
					"block",
					"mt-3",
				)}
			>
				アカウントをすでにお持ちの方はこちら
			</Link>
		</main>
	);
}
