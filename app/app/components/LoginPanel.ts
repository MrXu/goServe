import {Component} from "angular2/core"

@Component({
	selector: 'login-panel',
	templateUrl: 'app/components/templates/LoginPanel.html'
})

export class LoginPanel {
	isPanelVisible = false;

	constructor() {

	}

	showPanel() {
		this.isPanelVisible = true;
	}

	hidePanel(){
		this.isPanelVisible = false;
	}

}