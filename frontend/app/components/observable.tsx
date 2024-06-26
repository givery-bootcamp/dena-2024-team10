import { useEffect, useRef } from "react";

// このコンポーネントが画面内に現れるとcallbackが呼ばれる
// infinitie-loadingで使う想定
export default function ({ callback }: { callback: () => void }) {
	const ref = useRef<HTMLDivElement>(null);
	const obs = useRef<IntersectionObserver | null>(null);

	useEffect(() => {
		obs.current = new IntersectionObserver((entries) => {
			if (entries[0].intersectionRatio <= 0) return;
			callback();
		});
		if (ref.current) obs.current.observe(ref.current);
		return () => {
			if (ref.current) obs.current?.unobserve(ref.current);
		};
	}, [callback]);

	return <div ref={ref} />;
}
