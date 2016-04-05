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
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
var core_1 = require('angular2/core');
var common_1 = require('angular2/common');
var position_1 = require('../position');
var tooltip_options_class_1 = require('./tooltip-options.class');
var TooltipContainer = (function () {
    function TooltipContainer(element, options) {
        this.element = element;
        Object.assign(this, options);
        this.classMap = { 'in': false };
        this.classMap[options.placement] = true;
    }
    TooltipContainer.prototype.ngAfterViewChecked = function () {
        if (this.hostEl !== null) {
            var p = position_1.positionService
                .positionElements(this.hostEl.nativeElement, this.element.nativeElement.children[0], this.placement, this.appendToBody);
            this.top = p.top + 'px';
            this.left = p.left + 'px';
            this.classMap['in'] = true;
        }
    };
    TooltipContainer.prototype.position = function (hostEl) {
        this.display = 'block';
        this.top = '-1000px';
        this.left = '-1000px';
        this.hostEl = hostEl;
    };
    TooltipContainer = __decorate([
        core_1.Component({
            selector: 'tooltip-container',
            directives: [common_1.NgClass, common_1.NgStyle],
            template: "<div class=\"tooltip\" role=\"tooltip\"\n     [ngStyle]=\"{top: top, left: left, display: display}\"\n     [ngClass]=\"classMap\">\n      <div class=\"tooltip-arrow\"></div>\n      <div class=\"tooltip-inner\">\n        {{content}}\n      </div>\n    </div>"
        }),
        __param(1, core_1.Inject(tooltip_options_class_1.TooltipOptions)), 
        __metadata('design:paramtypes', [core_1.ElementRef, tooltip_options_class_1.TooltipOptions])
    ], TooltipContainer);
    return TooltipContainer;
}());
exports.TooltipContainer = TooltipContainer;
//# sourceMappingURL=tooltip-container.component.js.map