import {Component} from 'angular2/core';
import { RouteConfig, ROUTER_DIRECTIVES, ROUTER_PROVIDERS } from 'angular2/router';
import {HomePage} from './HomePage'
import {JoinPage} from './JoinPage'


@Component({
    selector: 'slide-layout',
    templateUrl: 'app/components/templates/SlideLayout.html',
	directives: [ROUTER_DIRECTIVES],
	providers: [
	  ROUTER_PROVIDERS
	]
})

@RouteConfig([
	{ path: '/home', name: 'HomeUrl', component:HomePage},
	{ path:'/join', name:'JoinUsUrl', component:JoinPage}
])

export class SlideLayout { 

	SideMenuVisible = false;

}