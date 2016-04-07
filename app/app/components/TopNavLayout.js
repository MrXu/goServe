System.register(["angular2/core", 'angular2/router', 'angular2/common', './HomePage', './JoinPage', './LoginPanel', '../service/EventEmitterService'], function(exports_1, context_1) {
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
    var core_1, router_1, common_1, HomePage_1, JoinPage_1, LoginPanel_1, EventEmitterService_1;
    var TopNavLayout;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (router_1_1) {
                router_1 = router_1_1;
            },
            function (common_1_1) {
                common_1 = common_1_1;
            },
            function (HomePage_1_1) {
                HomePage_1 = HomePage_1_1;
            },
            function (JoinPage_1_1) {
                JoinPage_1 = JoinPage_1_1;
            },
            function (LoginPanel_1_1) {
                LoginPanel_1 = LoginPanel_1_1;
            },
            function (EventEmitterService_1_1) {
                EventEmitterService_1 = EventEmitterService_1_1;
            }],
        execute: function() {
            TopNavLayout = (function () {
                function TopNavLayout() {
                    var _this = this;
                    // child components
                    // ui states
                    this.openHowItWorks = false;
                    // event emitter
                    this.emitter = EventEmitterService_1.EmitterService.get("Open_How_It_Works");
                    this.emitter.subscribe(function (msg) {
                        console.log(msg);
                        if (msg === "open") {
                            _this.openHowItWorks = true;
                        }
                    });
                }
                // close how it works
                TopNavLayout.prototype.closeHowItWorks = function () {
                    this.openHowItWorks = false;
                };
                TopNavLayout = __decorate([
                    core_1.Component({
                        selector: 'top-nav-layout',
                        templateUrl: 'app/components/templates/TopNavLayout.html',
                        styleUrls: ['css/top-nav-layout.css'],
                        directives: [router_1.ROUTER_DIRECTIVES, common_1.NgClass, LoginPanel_1.LoginPanel],
                        providers: [
                            router_1.ROUTER_PROVIDERS
                        ]
                    }),
                    router_1.RouteConfig([
                        { path: '/home', name: 'HomeUrl', component: HomePage_1.HomePage },
                        { path: '/join', name: 'JoinUsUrl', component: JoinPage_1.JoinPage }
                    ]), 
                    __metadata('design:paramtypes', [])
                ], TopNavLayout);
                return TopNavLayout;
            }());
            exports_1("TopNavLayout", TopNavLayout);
        }
    }
});
//# sourceMappingURL=TopNavLayout.js.map