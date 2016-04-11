import {Injectable} from 'angular2/core';
import {Http, Response, Headers} from 'angular2/http';
import {UrlFactory} from '../factory/UrlFactory';


@Injectable()
export class AuthServices {
	
	constructor(private http: Http) {}

	loginWithEmail(email, password){
		var load = {
			email:		email,
			password:	password
		}
		var credential = JSON.stringify(load);
		var url = UrlFactory.getInstance().getLoginUrl();
		var headers = new Headers();
		headers.append('Content-Type', 'application/x-www-form-urlencoded');
		console.log(url);
		this.http.post(url, credential, {headers:headers})
			 .map(res => res.json())
			 .subscribe(
				res => console.log(res),
				error => console.log(error),
				() => console.log("Auth completed")
			 );	

	}

}