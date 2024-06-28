import classNames from "classnames";

type Props = {
	color: "primary" | "secondaly" | "none";
	text: string;
	isDisabled?: boolean;
};

const SubmitButton: React.FC<Props> = ({ color, text, isDisabled }) => {
	const Style = {
		primary: classNames("bg-blue-500", "text-white", "hover:bg-blue-200"),
		secondaly: classNames("bg-green-500", "text-white", "hover:bg-green-200"),
		none: classNames("text-gray-800", "hover:text-gray-400"),
	};
	return (
		<input
			type="submit"
			value={text}
			disabled={isDisabled}
			className={classNames(
				"p-2",
				"rounded-md",
				"cursor-pointer",
				"disabled:opacity-30",
				"disabled:cursor-not-allowed",
				Style[color],
			)}
		/>
	);
};

export default SubmitButton;
