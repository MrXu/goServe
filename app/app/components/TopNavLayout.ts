import {Component} from "angular2/core";
import { RouteConfig, ROUTER_DIRECTIVES, ROUTER_PROVIDERS } from 'angular2/router';
import {NgClass} from 'angular2/common'
import {HomePage} from './HomePage';
import {JoinPage} from './JoinPage';
import {EmitterService} from '../service/EventEmitterService'


@Component({
    selector: 'top-nav-layout',
    templateUrl: 'app/components/templates/TopNavLayout.html',
    styleUrls:['css/top-nav-layout.css'],
	directives: [ROUTER_DIRECTIVES,NgClass],
	providers: [
	  ROUTER_PROVIDERS
	]
})

@RouteConfig([
	{ path: '/home', name: 'HomeUrl', component:HomePage},
	{ path:'/join', name:'JoinUsUrl', component:JoinPage}
])

export class TopNavLayout { 

	// ui states
	openHowItWorks = false

	// event emitter
	emitter = EmitterService.get("Open_How_It_Works");

	constructor() {
	    this.emitter.subscribe(msg => {
			console.log(msg);
			if(msg==="open") {
				this.openHowItWorks = true;
			}
		});
	}

	// close how it works
	closeHowItWorks(){
		this.openHowItWorks = false;
	}


}
