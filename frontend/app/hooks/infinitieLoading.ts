import type { SerializeFrom } from "@remix-run/node";
import { useFetcher, useLoaderData } from "@remix-run/react";
import { useEffect, useState } from "react";

/**
 * 無限読み込みを実現する
 * loaderがqueryParameterでoffset, limitを受け取る事が前提
 * @param extractArrayData loaderから帰ってるくるデータから、対象となるItemの配列を取り出す関数
 * @param limit 一回の読み込みで取ってくる要素のサイズ
 * @returns data: 取得したすべての要素 loadNext: 次の要素群を読み込むための関数 isLoading: 読み込み中かどうか。ローディング表示に使える
 */
export function useInfinitieLoading<T, U>(
	extractArrayData: (data: SerializeFrom<T>) => U[],
	limit = 20,
) {
	// remixのデフォルトの読み込みを行う
	const initialData = useLoaderData<T>();
	// 追加読み込みをするためのfetcher
	const fetcher = useFetcher<T>();
	// 取得したすべての要素を保持する配列
	const [list, setList] = useState<U[]>(extractArrayData(initialData));

	const [hasNoMoreItems, setHasNoMoreItems] = useState(list.length < limit);

	// fetcherが新しい要素を読み込んだら配列に追加
	// biome-ignore lint/correctness/useExhaustiveDependencies: <explanation>
	useEffect(() => {
		if (!fetcher.data) return;
		const arr = extractArrayData(fetcher.data);
		if (arr.length < limit) setHasNoMoreItems(true);
		setList((l) => [...l, ...arr]);
	}, [fetcher.data, setHasNoMoreItems]);

	// 他のリクエストが走ってなければ、新しい要素をとりに行く
	const loadNext = () => {
		if (fetcher.state !== "idle" || hasNoMoreItems) return;
		fetcher.submit({ offset: list.length, limit });
	};

	return {
		data: list,
		loadNext,
		state: (fetcher.state !== "idle"
			? "loading"
			: hasNoMoreItems
				? "end"
				: "idle") as "loading" | "end" | "idle",
	};
}
