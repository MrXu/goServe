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
var tooltip_options_class_1 = require('./tooltip-options.class');
var tooltip_container_component_1 = require('./tooltip-container.component');
var Tooltip = (function () {
    function Tooltip(element, loader) {
        this.element = element;
        this.loader = loader;
        this.placement = 'top';
        this.visible = false;
    }
    Tooltip.prototype.ngOnInit = function () {
    };
    // todo: filter triggers
    // params: event, target
    Tooltip.prototype.show = function () {
        var _this = this;
        if (this.visible) {
            return;
        }
        this.visible = true;
        var options = new tooltip_options_class_1.TooltipOptions({
            content: this.content,
            placement: this.placement
        });
        var binding = core_1.Injector.resolve([
            new core_1.Provider(tooltip_options_class_1.TooltipOptions, { useValue: options })
        ]);
        this.tooltip = this.loader
            .loadNextToLocation(tooltip_container_component_1.TooltipContainer, this.element, binding)
            .then(function (componentRef) {
            componentRef.instance.position(_this.element);
            return componentRef;
        });
    };
    // params event, target
    Tooltip.prototype.hide = function () {
        if (!this.visible) {
            return;
        }
        this.visible = false;
        this.tooltip.then(function (componentRef) {
            componentRef.dispose();
            return componentRef;
        });
    };
    __decorate([
        core_1.Input('tooltip'), 
        __metadata('design:type', String)
    ], Tooltip.prototype, "content", void 0);
    __decorate([
        core_1.Input('tooltipPlacement'), 
        __metadata('design:type', String)
    ], Tooltip.prototype, "placement", void 0);
    __decorate([
        core_1.Input('tooltipIsOpen'), 
        __metadata('design:type', Boolean)
    ], Tooltip.prototype, "isOpen", void 0);
    __decorate([
        core_1.Input('tooltipEnable'), 
        __metadata('design:type', Boolean)
    ], Tooltip.prototype, "enable", void 0);
    __decorate([
        core_1.Input('tooltipAppendToBody'), 
        __metadata('design:type', Boolean)
    ], Tooltip.prototype, "appendToBody", void 0);
    __decorate([
        core_1.HostListener('focusin', ['$event', '$target']),
        core_1.HostListener('mouseenter', ['$event', '$target']), 
        __metadata('design:type', Function), 
        __metadata('design:paramtypes', []), 
        __metadata('design:returntype', void 0)
    ], Tooltip.prototype, "show", null);
    __decorate([
        core_1.HostListener('focusout', ['$event', '$target']),
        core_1.HostListener('mouseleave', ['$event', '$target']), 
        __metadata('design:type', Function), 
        __metadata('design:paramtypes', []), 
        __metadata('design:returntype', void 0)
    ], Tooltip.prototype, "hide", null);
    Tooltip = __decorate([
        core_1.Directive({ selector: '[tooltip]' }), 
        __metadata('design:paramtypes', [core_1.ElementRef, core_1.DynamicComponentLoader])
    ], Tooltip);
    return Tooltip;
}());
exports.Tooltip = Tooltip;
//# sourceMappingURL=tooltip.directive.js.map