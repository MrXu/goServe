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
var common_1 = require('angular2/common');
var typeahead_utils_1 = require('./typeahead-utils');
var typeahead_options_class_1 = require('./typeahead-options.class');
var position_1 = require('../position');
var ng2_bootstrap_config_1 = require('../ng2-bootstrap-config');
var TEMPLATE = (_a = {},
    _a[ng2_bootstrap_config_1.Ng2BootstrapTheme.BS4] = "\n  <div class=\"dropdown-menu\"\n      [ngStyle]=\"{top: top, left: left, display: display}\"\n      style=\"display: block\">\n      <a href=\"#\"\n         *ngFor=\"#match of matches\"\n         class=\"dropdown-item\"\n         (click)=\"selectMatch(match, $event)\"\n         (mouseenter)=\"selectActive(match)\"\n         [class.active]=\"isActive(match)\"\n         [innerHtml]=\"hightlight(match, query)\"></a>\n  </div>\n  ",
    _a[ng2_bootstrap_config_1.Ng2BootstrapTheme.BS3] = "\n  <ul class=\"dropdown-menu\"\n      [ngStyle]=\"{top: top, left: left, display: display}\"\n      style=\"display: block\">\n    <li *ngFor=\"#match of matches\"\n        [class.active]=\"isActive(match)\"\n        (mouseenter)=\"selectActive(match)\">\n        <a href=\"#\" (click)=\"selectMatch(match, $event)\" tabindex=\"-1\" [innerHtml]=\"hightlight(match, query)\"></a>\n    </li>\n  </ul>\n  ",
    _a
);
var TypeaheadContainer = (function () {
    function TypeaheadContainer(element, options) {
        this.element = element;
        this._matches = [];
        Object.assign(this, options);
    }
    Object.defineProperty(TypeaheadContainer.prototype, "matches", {
        get: function () {
            return this._matches;
        },
        set: function (value) {
            this._matches = value;
            if (this._matches.length > 0) {
                this._active = this._matches[0];
            }
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(TypeaheadContainer.prototype, "field", {
        set: function (value) {
            this._field = value;
        },
        enumerable: true,
        configurable: true
    });
    TypeaheadContainer.prototype.position = function (hostEl) {
        this.display = 'block';
        this.top = '0px';
        this.left = '0px';
        var p = position_1.positionService
            .positionElements(hostEl.nativeElement, this.element.nativeElement.children[0], this.placement, false);
        this.top = p.top + 'px';
        this.left = p.left + 'px';
    };
    TypeaheadContainer.prototype.selectActiveMatch = function () {
        this.selectMatch(this._active);
    };
    TypeaheadContainer.prototype.prevActiveMatch = function () {
        var index = this.matches.indexOf(this._active);
        this._active = this.matches[index - 1 < 0 ? this.matches.length - 1 : index - 1];
    };
    TypeaheadContainer.prototype.nextActiveMatch = function () {
        var index = this.matches.indexOf(this._active);
        this._active = this.matches[index + 1 > this.matches.length - 1 ? 0 : index + 1];
    };
    TypeaheadContainer.prototype.selectActive = function (value) {
        this._active = value;
    };
    TypeaheadContainer.prototype.isActive = function (value) {
        return this._active === value;
    };
    TypeaheadContainer.prototype.selectMatch = function (value, e) {
        if (e === void 0) { e = null; }
        if (e) {
            e.stopPropagation();
            e.preventDefault();
        }
        this.parent.changeModel(value);
        this.parent.typeaheadOnSelect.emit({
            item: value
        });
        return false;
    };
    TypeaheadContainer.prototype.hightlight = function (item, query) {
        var itemStr = (typeof item === 'object' && this._field ? item[this._field] : item).toString();
        var itemStrHelper = (this.parent.typeaheadLatinize ? typeahead_utils_1.TypeaheadUtils.latinize(itemStr) : itemStr).toLowerCase();
        var startIdx;
        var tokenLen;
        // Replaces the capture string with the same string inside of a "strong" tag
        if (typeof query === 'object') {
            var queryLen = query.length;
            for (var i = 0; i < queryLen; i += 1) {
                // query[i] is already latinized and lower case
                startIdx = itemStrHelper.indexOf(query[i]);
                tokenLen = query[i].length;
                if (startIdx >= 0 && tokenLen > 0) {
                    itemStr = itemStr.substring(0, startIdx) + '<strong>' + itemStr.substring(startIdx, startIdx + tokenLen) + '</strong>' + itemStr.substring(startIdx + tokenLen);
                    itemStrHelper = itemStrHelper.substring(0, startIdx) + '        ' + ' '.repeat(tokenLen) + '         ' + itemStrHelper.substring(startIdx + tokenLen);
                }
            }
        }
        else if (query) {
            // query is already latinized and lower case
            startIdx = itemStrHelper.indexOf(query);
            tokenLen = query.length;
            if (startIdx >= 0 && tokenLen > 0) {
                itemStr = itemStr.substring(0, startIdx) + '<strong>' + itemStr.substring(startIdx, startIdx + tokenLen) + '</strong>' + itemStr.substring(startIdx + tokenLen);
            }
        }
        return itemStr;
    };
    TypeaheadContainer = __decorate([
        core_1.Component({
            selector: 'typeahead-container',
            directives: [common_1.CORE_DIRECTIVES],
            template: TEMPLATE[ng2_bootstrap_config_1.Ng2BootstrapConfig.theme],
            encapsulation: core_1.ViewEncapsulation.None
        }), 
        __metadata('design:paramtypes', [core_1.ElementRef, typeahead_options_class_1.TypeaheadOptions])
    ], TypeaheadContainer);
    return TypeaheadContainer;
}());
exports.TypeaheadContainer = TypeaheadContainer;
var _a;
//# sourceMappingURL=typeahead-container.component.js.map