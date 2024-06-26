import classNames from "classnames";
import { useRef, useState } from "react";
import Button from "./button";

export default function Dialog({
	children,
	isOpen,
	confirm,
}: {
	children: React.ReactNode;
	isOpen: boolean;
	confirm: (result: boolean) => void;
}) {
	return (
		// biome-ignore lint/a11y/useKeyWithClickEvents: <explanation>
		<dialog
			open={isOpen}
			onClick={() => confirm(false)}
			className={classNames(
				"fixed",
				"inset-0",
				"w-full",
				"h-full",
				"bg-black",
				"bg-opacity-30",
				"inset-0",
				"m-auto",
			)}
		>
			{/* biome-ignore lint/a11y/useKeyWithClickEvents: <explanation> */}
			<div
				onClick={(e) => e.stopPropagation()}
				className={classNames(
					"bg-white",
					"w-1/2",
					"fixed",
					"top-1/2",
					"left-1/2",
					"-translate-x-1/2",
					"-translate-y-1/2",
					"p-4",
					"rounded",
					"drop-shadow-lg",
				)}
				style={{}}
			>
				<div>{children}</div>
				<hr className={classNames("my-4")} />
				<div className={classNames("flex", "justify-end", "gap-4")}>
					<Button type="danger" onClick={() => confirm(true)}>
						OK
					</Button>
					<Button type="none" onClick={() => confirm(false)}>
						Cancel
					</Button>
				</div>
			</div>
		</dialog>
	);
}

export function useDialog(children: React.ReactNode) {
	const p = useRef(Promise.withResolvers<boolean>());
	const [isOpen, setIsOpen] = useState(false);
	const handleConfirm = (result: boolean) => {
		p.current.resolve(result);
	};

	const dialog = (
		<Dialog isOpen={isOpen} confirm={handleConfirm}>
			{children}
		</Dialog>
	);

	const confirm = async () => {
		setIsOpen(true);
		const result = await p.current.promise;
		p.current = Promise.withResolvers<boolean>();
		setIsOpen(false);
		return result;
	};

	return {
		dialog,
		confirm,
	};
}
