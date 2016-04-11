export class UrlFactory{
	static instance: UrlFactory;
	static isCreating: boolean = false;

	// base url for api
	baseUrl: String = "http://localhost:8080";

	constructor(){
		if(!UrlFactory.isCreating) {
			throw new Error("UrlFactory is a singleton");
		}
	}

	static getInstance(){
		if(UrlFactory.instance == null) {
			UrlFactory.isCreating = true;
			UrlFactory.instance = new UrlFactory();
			UrlFactory.isCreating = false;
		}
		return UrlFactory.instance;
	}

	// auth api
	getLoginUrl(){
		return this.baseUrl + "/login";
	}

}