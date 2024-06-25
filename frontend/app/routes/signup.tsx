import type { ActionFunctionArgs, SerializeFrom } from "@remix-run/node";
import { Link, json, redirect, useActionData } from "@remix-run/react";
import classNames from "classnames";
import { ZodError } from "zod";
import { API_BASE_URL } from "~/apiClient/apiClient";
import { schemas } from "~/apiClient/output.generated";
import InternalError from "~/components/InternalError";
import AuthForm from "~/components/authForm";

export async function action({ request }: ActionFunctionArgs) {
	try {
		const formData = await request.formData();
		const username = formData.get("name") as string;
		const password = formData.get("pass") as string;

		const body = schemas.signupUser_Body.parse({ username, password });
		const res = await fetch(`${API_BASE_URL}/signup`, {
			method: "POST",
			body: JSON.stringify(body),
			headers: {
				"content-type": "application/json",
			},
		});
		if (res.status === 400) throw new Error("Duplicate username");
		if (!res.ok) throw new Error("Failed to sign up");

		return redirect("/", {
			headers: res.headers,
		});
	} catch (e) {
		console.error(e);
		if (e instanceof ZodError) {
			return json({
				errors: e.errors.map((error) => {
					return { path: error.path, message: error.message };
				}),
			});
		}
		if (e instanceof Error) {
			if (e.message === "Duplicate username") {
				return json({
					errors: [
						{
							path: "username",
							message: "このユーザー名はすでに使用されています",
						},
					],
				});
			}
			throw new Response(e.toString(), { status: 500 });
		}
	}
}

export type AuthErrorType = SerializeFrom<typeof action> | undefined;

export default function SignIn() {
	const actionData = useActionData<typeof action>();
	return (
		<main className={classNames("mx-auto", "w-1/3")}>
			<AuthForm authType="signup" actionData={actionData} />
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

export function ErrorBoundary() {
	return (
		<InternalError
			title="サインアップに失敗しました"
			to="/signup"
			toMessage="サインアップ画面へ戻る"
		/>
	);
}
