System.register(['angular2/core', 'angular2/router', './HomePage', './JoinPage'], function(exports_1, context_1) {
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
    var core_1, router_1, HomePage_1, JoinPage_1;
    var SlideLayout;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (router_1_1) {
                router_1 = router_1_1;
            },
            function (HomePage_1_1) {
                HomePage_1 = HomePage_1_1;
            },
            function (JoinPage_1_1) {
                JoinPage_1 = JoinPage_1_1;
            }],
        execute: function() {
            SlideLayout = (function () {
                function SlideLayout() {
                    this.SideMenuVisible = false;
                }
                SlideLayout = __decorate([
                    core_1.Component({
                        selector: 'slide-layout',
                        templateUrl: 'app/components/templates/SlideLayout.html',
                        directives: [router_1.ROUTER_DIRECTIVES],
                        providers: [
                            router_1.ROUTER_PROVIDERS
                        ]
                    }),
                    router_1.RouteConfig([
                        { path: '/home', name: 'HomeUrl', component: HomePage_1.HomePage },
                        { path: '/join', name: 'JoinUsUrl', component: JoinPage_1.JoinPage }
                    ]), 
                    __metadata('design:paramtypes', [])
                ], SlideLayout);
                return SlideLayout;
            }());
            exports_1("SlideLayout", SlideLayout);
        }
    }
});
//# sourceMappingURL=SlideLayout.js.map