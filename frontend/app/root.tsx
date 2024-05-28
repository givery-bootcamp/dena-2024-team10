import {
	Links,
	Meta,
	Outlet,
	Scripts,
	ScrollRestoration,
	useRouteError,
} from "@remix-run/react";
import type { LinksFunction } from "@remix-run/node";
import stylesheet from "~/tailwind.css?url";
import classNames from "classnames";

export const links: LinksFunction = () => [
	{ rel: "stylesheet", href: stylesheet },
];

export function Layout({ children }: { children: React.ReactNode }) {
	return (
		<html lang="en">
			<head>
				<meta charSet="utf-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<Meta />
				<Links />
			</head>
			<body>
				{children}
				<ScrollRestoration />
				<Scripts />
			</body>
		</html>
	);
}

export function ErrorBoundary() {
	const error = useRouteError();
	console.error(error);
	return (
		<html lang="ja">
			<head>
				<title>Oh no!</title>
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
					Oh no!
				</h1>
				<Scripts />
			</body>
		</html>
	);
}

export default function App() {
	return <Outlet />;
}
