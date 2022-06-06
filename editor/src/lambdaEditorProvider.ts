import * as vscode from 'vscode';

export class FluxEditorProvider implements vscode.CustomTextEditorProvider {

    public static register(context: vscode.ExtensionContext): vscode.Disposable {
        const provider = new FluxEditorProvider(context);
        const providerRegistration = vscode.window.registerCustomEditorProvider(FluxEditorProvider.viewType, provider);
        return providerRegistration;
    }
    
    public static readonly viewType = 'flux-editor.editor';

    constructor(private readonly context: vscode.ExtensionContext) { }

    resolveCustomTextEditor(document: vscode.TextDocument, webviewPanel: vscode.WebviewPanel, token: vscode.CancellationToken): void | Thenable<void> {
       
        webviewPanel.webview.options = {
            enableScripts: true,
        };
        webviewPanel.webview.html = this.getHtmlForWebview(webviewPanel.webview, document);

        const changeDocumentSubscription = vscode.workspace.onDidChangeTextDocument(e => {
			if (e.document.uri.toString() === document.uri.toString()) {
				this.updateWebview(document, webviewPanel);
			}
		});
        
        webviewPanel.webview.onDidReceiveMessage(e => {
			switch (e.type) {
				case 'change':
					this.updateTextDocument(document, e.text);
					return;
			}
		});


		webviewPanel.onDidDispose(() => {
			changeDocumentSubscription.dispose();
		});

        this.updateWebview(document, webviewPanel);
    }

    private updateWebview(document: vscode.TextDocument, webviewPanel: vscode.WebviewPanel) {
        webviewPanel.webview.postMessage({
            command: 'update',
            text: document.getText()
        });
    }

    private updateTextDocument(document: vscode.TextDocument, text: string) {
        const edit = new vscode.WorkspaceEdit();
        edit.replace(document.uri, new vscode.Range(0, 0, document.lineCount, 0), text);
        vscode.workspace.applyEdit(edit);
    }

    private getHtmlForWebview(webview: vscode.Webview, document: vscode.TextDocument): string {
    
        return `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Flux Editor</title>
        </head>
        <body>
            <h1>Flux Editor</h1>
        </body>
        </html>`;
    }
}