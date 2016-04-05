"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var core_1 = require('angular2/core');
var carousel_component_1 = require('./carousel.component');
var Slide = (function () {
    function Slide(carousel) {
        this.carousel = carousel;
        this.addClass = true;
    }
    Slide.prototype.ngOnInit = function () {
        this.carousel.addSlide(this);
    };
    Slide.prototype.ngOnDestroy = function () {
        this.carousel.removeSlide(this);
    };
    __decorate([
        core_1.Input(), 
        __metadata('design:type', Number)
    ], Slide.prototype, "index", void 0);
    __decorate([
        core_1.Input(), 
        __metadata('design:type', Number)
    ], Slide.prototype, "direction", void 0);
    __decorate([
        core_1.HostBinding('class.active'),
        core_1.Input(), 
        __metadata('design:type', Boolean)
    ], Slide.prototype, "active", void 0);
    __decorate([
        core_1.HostBinding('class.item'),
        core_1.HostBinding('class.carousel-item'), 
        __metadata('design:type', Boolean)
    ], Slide.prototype, "addClass", void 0);
    Slide = __decorate([
        core_1.Component({
            selector: 'slide',
            template: "\n    <div [class.active]=\"active\" class=\"item text-center\">\n      <ng-content></ng-content>\n    </div>\n  "
        }), 
        __metadata('design:paramtypes', [carousel_component_1.Carousel])
    ], Slide);
    return Slide;
}());
exports.Slide = Slide;
//# sourceMappingURL=slide.component.js.map