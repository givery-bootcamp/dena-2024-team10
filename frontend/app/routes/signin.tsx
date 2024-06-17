import type { ActionFunctionArgs } from "@remix-run/node";
import { Link } from "@remix-run/react";
import classNames from "classnames";
import { ZodError } from "zod";
import apiClient, { API_BASE_URL } from "~/apiClient/apiClient";
import { schemas } from "~/apiClient/output.generated";
import AuthForm from "~/components/authForm";

export async function action({ request }: ActionFunctionArgs) {
	try {
		const formData = await request.formData();
		const username = formData.get("name") as string;
		const password = formData.get("pass") as string;

		const body = schemas.signupUser_Body.parse({ username, password });
		const res = await fetch(`${API_BASE_URL}/signin`, {
			method: "POST",
			body: JSON.stringify(body),
			headers: {
				"content-type": "application/json",
			},
		});

		return new Response(await res.text(), {
			status: res.status,
			headers: res.headers,
		});
	} catch (e) {
		console.error(e);
		if (e instanceof ZodError) return new Response(e.message, { status: 400 });
		if (e instanceof Error) return new Response(e.toString(), { status: 500 });
	}
}

export default function SignIn() {
	return (
		<main className={classNames("mx-auto", "w-1/3")}>
			<AuthForm authType="signin" />
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
