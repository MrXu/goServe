import {Component} from 'angular2/core';
import {EmitterService} from '../service/EventEmitterService'

@Component({
	selector: 'home-page',
	templateUrl: 'app/components/templates/HomePage.html'
})

export class HomePage {

	/*
		ui states
	*/
	

	/*
		class functions
	*/
	openHowItWorks(){
		EmitterService.get("Open_How_It_Works").emit("open");
	}

}