import type { ActionFunctionArgs } from "@remix-run/node";
import { Form } from "@remix-run/react";
import classNames from "classnames";
import SubmitButton from "~/components/submitButton";

export async function action({ request }: ActionFunctionArgs) {
	try {
		const formData = await request.formData();
		const comment = formData.get("comment") as string;

		const body = { comment };
		// schemas.createComment_Body.parse(body);
		// const res = await apiClient.createComment(
		//     { comment },
		//     { headers: { cookie: request.headers.get("cookie") } },
		// );
	} catch (e) {
		console.error(e);
	}
}
