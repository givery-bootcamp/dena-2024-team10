import type {
	ErrorResponse,
	LinksFunction,
	LoaderFunctionArgs,
} from "@remix-run/node";
import {
	Link,
	Links,
	Meta,
	Outlet,
	Scripts,
	ScrollRestoration,
	isRouteErrorResponse,
	redirect,
	useLoaderData,
	useRouteError,
} from "@remix-run/react";
import { AxiosError } from "axios";
import classNames from "classnames";
import stylesheet from "~/tailwind.css?url";
import apiClient from "./apiClient/apiClient";
import Header from "./components/header";

export const links: LinksFunction = () => [
	{ rel: "stylesheet", href: stylesheet },
];

export async function loader({
	request,
}: LoaderFunctionArgs): Promise<{ id?: number; username?: string }> {
	try {
		const user = await apiClient.getSignedInUser({
			headers: {
				Cookie: request.headers.get("Cookie"),
			},
		});
		return user;
	} catch (e) {
		if (e instanceof AxiosError) {
			if (e.response?.status === 401) {
				const ALLOW_UNAUTHORIZED_PATHS = ["/signin", "/signup"];
				const host = request.headers.get("Host");
				const path = host && request.url.split(host)[1];
				if (path && ALLOW_UNAUTHORIZED_PATHS.includes(path)) return {};
				throw redirect("/signin");
			}
		}
		return {};
	}
}

export function Layout({ children }: { children: React.ReactNode }) {
	const user = useLoaderData<typeof loader>();

	return (
		<html lang="en">
			<head>
				<meta charSet="utf-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<Meta />
				<Links />
			</head>
			<body>
				<Header isSignedIn={!!user?.id} username={user?.username} />
				{children}
				<ScrollRestoration />
				<Scripts />
			</body>
		</html>
	);
}

export function ErrorBoundary() {
	const error = useRouteError();
	type MyRouteError = {
		type: "MyRouteError";
		status: number;
		data: string;
		string: string;
	};
	const errorResponseToMyRouteError = (error: ErrorResponse): MyRouteError => {
		const { status, data } = error;
		return {
			type: "MyRouteError",
			status,
			data,
			string: JSON.stringify(error),
		};
	};
	let title = "Oops!";
	let message = "Something went wrong!";
	const toMessage = "Go to Top Page";
	const toLink = "/";

	if (isRouteErrorResponse(error)) {
		if (error.status === 404) {
			title = "Page Not Found";
			message = "Sorry, we couldn't find the page you're looking for.";
		}
		console.error(JSON.stringify(errorResponseToMyRouteError(error)));
	} else {
		console.error(JSON.stringify(error));
	}
	return (
		<html lang="ja">
			<head>
				<title>{title}</title>
				<Meta />
				<Links />
			</head>
			<body>
				<h1
					className={classNames(
						"text-3xl",
						"font-bold",
						"underline",
						"text-red-500",
						"text-center",
						"mt-8",
					)}
				>
					{title}
				</h1>
				<p className={classNames("text-center", "mt-4")}>{message}</p>
				<p
					className={classNames(
						"text-center",
						"mt-4",
						"text-blue-600",
						"underline",
						"hover:text-blue-300",
					)}
				>
					<Link to={toLink}>{toMessage}</Link>
				</p>
				<Scripts />
			</body>
		</html>
	);
}

export default function App() {
	return <Outlet />;
}
