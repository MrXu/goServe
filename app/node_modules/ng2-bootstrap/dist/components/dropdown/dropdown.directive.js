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
var dropdown_service_1 = require('./dropdown.service');
var Dropdown = (function () {
    function Dropdown(el, dropdownMenuList) {
        this.el = el;
        this.onToggle = new core_1.EventEmitter();
        this.isOpenChange = new core_1.EventEmitter();
        this.addClass = true;
        // todo: bind to route change event
    }
    Object.defineProperty(Dropdown.prototype, "isOpen", {
        get: function () {
            return this._isOpen;
        },
        set: function (value) {
            this._isOpen = !!value;
            // todo: implement after porting position
            if (this.appendToBody && this.menuEl) {
            }
            // todo: $animate open<->close transitions, as soon as ng2Animate will be ready
            if (this.isOpen) {
                this.focusToggleElement();
                dropdown_service_1.dropdownService.open(this);
            }
            else {
                dropdown_service_1.dropdownService.close(this);
                this.selectedOption = null;
            }
            this.onToggle.emit(this.isOpen);
            this.isOpenChange.emit(this.isOpen);
            // todo: implement call to setIsOpen if set and function
        },
        enumerable: true,
        configurable: true
    });
    Dropdown.prototype.ngOnInit = function () {
        this.autoClose = this.autoClose || dropdown_service_1.NONINPUT;
        if (this.isOpen) {
        }
    };
    Dropdown.prototype.ngOnDestroy = function () {
        if (this.appendToBody && this.menuEl) {
            this.menuEl.nativeElement.remove();
        }
    };
    Object.defineProperty(Dropdown.prototype, "dropDownMenu", {
        set: function (dropdownMenu) {
            // init drop down menu
            this.menuEl = dropdownMenu.el;
            if (this.appendToBody) {
                window.document.body.appendChild(this.menuEl.nativeElement);
            }
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(Dropdown.prototype, "dropDownToggle", {
        set: function (dropdownToggle) {
            // init toggle element
            this.toggleEl = dropdownToggle.el;
        },
        enumerable: true,
        configurable: true
    });
    Dropdown.prototype.toggle = function (open) {
        return this.isOpen = arguments.length ? !!open : !this.isOpen;
    };
    Dropdown.prototype.focusDropdownEntry = function (keyCode) {
        // If append to body is used.
        var hostEl = this.menuEl ?
            this.menuEl.nativeElement :
            this.el.nativeElement.getElementsByTagName('ul')[0];
        if (!hostEl) {
            // todo: throw exception?
            return;
        }
        var elems = hostEl.getElementsByTagName('a');
        if (!elems || !elems.length) {
            // todo: throw exception?
            return;
        }
        // todo: use parseInt to detect isNumber?
        // todo: or implement selectedOption as a get\set pair with parseInt on set
        switch (keyCode) {
            case (40):
                if (typeof this.selectedOption !== 'number') {
                    this.selectedOption = 0;
                    break;
                }
                if (this.selectedOption === elems.length - 1) {
                    break;
                }
                this.selectedOption++;
                break;
            case (38):
                if (typeof this.selectedOption !== 'number') {
                    return;
                }
                if (this.selectedOption === 0) {
                    // todo: return?
                    break;
                }
                this.selectedOption--;
                break;
        }
        elems[this.selectedOption].focus();
    };
    Dropdown.prototype.focusToggleElement = function () {
        if (this.toggleEl) {
            this.toggleEl.nativeElement.focus();
        }
    };
    __decorate([
        core_1.HostBinding('class.open'),
        core_1.Input(), 
        __metadata('design:type', Boolean)
    ], Dropdown.prototype, "isOpen", null);
    __decorate([
        core_1.Input(), 
        __metadata('design:type', String)
    ], Dropdown.prototype, "autoClose", void 0);
    __decorate([
        core_1.Input(), 
        __metadata('design:type', Boolean)
    ], Dropdown.prototype, "keyboardNav", void 0);
    __decorate([
        core_1.Input(), 
        __metadata('design:type', Boolean)
    ], Dropdown.prototype, "appendToBody", void 0);
    __decorate([
        core_1.Output(), 
        __metadata('design:type', core_1.EventEmitter)
    ], Dropdown.prototype, "onToggle", void 0);
    __decorate([
        core_1.Output(), 
        __metadata('design:type', core_1.EventEmitter)
    ], Dropdown.prototype, "isOpenChange", void 0);
    __decorate([
        core_1.HostBinding('class.dropdown'), 
        __metadata('design:type', Object)
    ], Dropdown.prototype, "addClass", void 0);
    Dropdown = __decorate([
        core_1.Directive({ selector: '[dropdown]' }),
        __param(1, core_1.Query('dropdownMenu', { descendants: false })), 
        __metadata('design:paramtypes', [core_1.ElementRef, core_1.QueryList])
    ], Dropdown);
    return Dropdown;
}());
exports.Dropdown = Dropdown;
//# sourceMappingURL=dropdown.directive.js.map