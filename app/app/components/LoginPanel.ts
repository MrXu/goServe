import {Component} from "angular2/core";
import {AuthServices} from "../service/AuthServices";

@Component({
	selector: 'login-panel',
	templateUrl: 'app/components/templates/LoginPanel.html',
	providers:[AuthServices]
})

export class LoginPanel {
	isPanelVisible = false;
	
	// states
	email: String;
	password: String;

	constructor(private _authService:AuthServices) {
		this.email = "";
		this.password = "";
	}

	showPanel() {
		this.isPanelVisible = true;
	}

	hidePanel(){
		this.isPanelVisible = false;
	}

	loginWithEmail(){
		if(!this.email || !this.password){
			return;
		}
		this._authService.loginWithEmail(this.email, this.password);
	}

}