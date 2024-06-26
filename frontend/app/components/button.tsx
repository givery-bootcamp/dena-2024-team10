import classNames from "classnames";

const Button: React.FC<{
	type: "none" | "danger";
	children: React.ReactNode;
	onClick: () => void;
}> = ({ type, children, onClick }) => {
	return (
		<button
			onClick={onClick}
			type="button"
			className={classNames("py-2", "px-4", "rounded-md", {
				"bg-red-500": type === "danger",
				"text-white": type === "danger",
				"hover:bg-red-200": type === "danger",
				"text-black": type === "none",
				"hover:bg-gray-200": type === "none",
			})}
		>
			{children}
		</button>
	);
};

export default Button;
