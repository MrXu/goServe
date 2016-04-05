System.register(['angular2/core', '../service/EventEmitterService'], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
        var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
        if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
        else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
        return c > 3 && r && Object.defineProperty(target, key, r), r;
    };
    var __metadata = (this && this.__metadata) || function (k, v) {
        if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
    };
    var core_1, EventEmitterService_1;
    var HomePage;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (EventEmitterService_1_1) {
                EventEmitterService_1 = EventEmitterService_1_1;
            }],
        execute: function() {
            HomePage = (function () {
                function HomePage() {
                }
                /*
                    ui states
                */
                /*
                    class functions
                */
                HomePage.prototype.openHowItWorks = function () {
                    EventEmitterService_1.EmitterService.get("Open_How_It_Works").emit("open");
                };
                HomePage = __decorate([
                    core_1.Component({
                        selector: 'home-page',
                        templateUrl: 'app/components/templates/HomePage.html'
                    }), 
                    __metadata('design:paramtypes', [])
                ], HomePage);
                return HomePage;
            }());
            exports_1("HomePage", HomePage);
        }
    }
});
//# sourceMappingURL=HomePage.js.map