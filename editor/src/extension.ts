import * as vscode from 'vscode';
import { LambdaEditorProvider } from './lambdaEditorProvider';

export function activate(context: vscode.ExtensionContext) {
	
	let openWithCommand = vscode.commands.registerTextEditorCommand('lambda-editor.openWithEditor', (textEditor: vscode.TextEditor, edit: vscode.TextEditorEdit) => {
		vscode.commands.executeCommand(
			"vscode.openWith",
			textEditor.document.uri,
			LambdaEditorProvider.viewType
		  );
	});
	
	context.subscriptions.push(openWithCommand);
	context.subscriptions.push(LambdaEditorProvider.register(context));
}

export function deactivate() {}
