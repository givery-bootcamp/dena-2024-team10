import { redirect, type ActionFunctionArgs } from "@remix-run/node";
import { API_BASE_URL } from "~/apiClient/apiClient";

export async function action({ request }: ActionFunctionArgs) {
	try {
		const res = await fetch(`${API_BASE_URL}/signout`, {
			method: "POST",
			headers: {
				Cookie: request.headers.get("Cookie") as string,
			},
		});

		return redirect("/signin", {
			headers: {
				"Set-Cookie": res.headers.get("Set-Cookie") as string,
			},
		});
	} catch (e) {
		return new Response((e as Error).message, {
			status: 500,
		});
	}
}

export function loader() {
	return new Response("Not found", {
		status: 404,
	});
}
