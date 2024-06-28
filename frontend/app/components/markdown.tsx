import {
	headingsPlugin,
	listsPlugin,
	quotePlugin,
	thematicBreakPlugin,
	markdownShortcutPlugin,
	linkPlugin,
	linkDialogPlugin,
	codeBlockPlugin,
	sandpackPlugin,
	codeMirrorPlugin,
	MDXEditor,
	toolbarPlugin,
	UndoRedo,
	BoldItalicUnderlineToggles,
	BlockTypeSelect,
	CodeToggle,
	CreateLink,
	InsertCodeBlock,
	ListsToggle,
	ChangeCodeMirrorLanguage,
	ShowSandpackInfo,
	InsertSandpack,
	ConditionalContents,
	type SandpackConfig,
} from "@mdxeditor/editor";
import "@mdxeditor/editor/style.css";
import { ClientOnly } from "remix-utils/client-only";

export default function Markdown({
	markdown,
	onChange,
}: { markdown: string; onChange?: (markdown: string) => void }) {
	const readOnly = onChange === undefined;
	const defaultSnippetContent = `
        export default function App() {
        return (
            <div className="App">
            <h1>Hello CodeSandbox</h1>
            <h2>Start editing to see some magic happen!</h2>
            </div>
        );
        }
        `.trim();
	const simpleSandpackConfig: SandpackConfig = {
		defaultPreset: "react",
		presets: [
			{
				label: "React",
				name: "react",
				meta: "live react",
				sandpackTemplate: "react",
				sandpackTheme: "light",
				snippetFileName: "/App.js",
				snippetLanguage: "jsx",
				initialSnippetContent: defaultSnippetContent,
			},
		],
	};

	return (
		<ClientOnly fallback={<p>Loading...</p>}>
			{() => (
				<MDXEditor
					plugins={[
						headingsPlugin(),
						listsPlugin(),
						quotePlugin(),
						thematicBreakPlugin(),
						markdownShortcutPlugin(),
						linkPlugin(),
						linkDialogPlugin(),
						codeBlockPlugin({ defaultCodeBlockLanguage: "js" }),
						sandpackPlugin({ sandpackConfig: simpleSandpackConfig }),
						codeMirrorPlugin({
							codeBlockLanguages: { js: "JavaScript", css: "CSS" },
						}),
						toolbarPlugin({
							toolbarContents: () =>
								!readOnly && (
									<>
										{" "}
										<UndoRedo />
										<BoldItalicUnderlineToggles />
										<BlockTypeSelect />
										<CodeToggle />
										<CreateLink />
										<ListsToggle />
										<ConditionalContents
											options={[
												{
													when: (editor) => editor?.editorType === "codeblock",
													contents: () => <ChangeCodeMirrorLanguage />,
												},
												{
													when: (editor) => editor?.editorType === "sandpack",
													contents: () => <ShowSandpackInfo />,
												},
												{
													fallback: () => (
														<>
															<InsertCodeBlock />
															<InsertSandpack />
														</>
													),
												},
											]}
										/>
									</>
								),
						}),
					]}
					contentEditableClassName="prose"
					markdown={markdown}
					onChange={onChange}
					readOnly={readOnly}
				/>
			)}
		</ClientOnly>
	);
}
