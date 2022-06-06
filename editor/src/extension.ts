import * as vscode from 'vscode';
import { FluxEditorProvider } from './fluxEditorProvider';

export function activate(context: vscode.ExtensionContext) {
	
	let openWithCommand = vscode.commands.registerTextEditorCommand('flux-editor.openWithEditor', (textEditor: vscode.TextEditor, edit: vscode.TextEditorEdit) => {
		vscode.commands.executeCommand(
			"vscode.openWith",
			textEditor.document.uri,
			FluxEditorProvider.viewType
		  );
	});
	
	context.subscriptions.push(openWithCommand);
	context.subscriptions.push(FluxEditorProvider.register(context));
}

export function deactivate() {}
